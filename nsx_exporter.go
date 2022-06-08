package main

import (
	"net/http"
	"nsx_exporter/collector"
	"os"
        "fmt"

	"github.com/go-kit/kit/log/level"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/promlog"
	"github.com/prometheus/common/promlog/flag"
	"github.com/prometheus/common/version"
	nsxt "github.com/vmware/go-vmware-nsxt"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

type nsxtOpts struct {
	host     string
	username string
	password string
	insecure bool
}

func newNSXTClient(opts nsxtOpts) (*nsxt.APIClient, error) {
	cfg := nsxt.Configuration{
		BasePath:           "/api/v1",
		Host:               opts.host,
		Scheme:             "https",
		UserAgent:          "nsx_exporter/1.0",
		ClientAuthCertFile: "",
		RemoteAuth:         false,
		UserName:           opts.username,
		Password:           opts.password,
		Insecure:           opts.insecure,
	}
        fmt.Println("%T", cfg)
	return nsxt.NewAPIClient(&cfg)
}

func main() {
	var (
		listenAddress = kingpin.Flag("listen", "Address to listen on for web interface and telemetry.").OverrideDefaultFromEnvar("NSX_LISTEN").Default(":9744").String()
		metricsPath   = kingpin.Flag("path", "Path under which to expose metrics.").Default("/metrics").String()
		opts          = nsxtOpts{}
	)
	kingpin.Flag("host", "URI of NSX-T manager").OverrideDefaultFromEnvar("NSX_HOST").Default("localhost").StringVar(&opts.host)
	kingpin.Flag("username", "The username to connect to the NSX-T manager").OverrideDefaultFromEnvar("NSX_USERNAME").StringVar(&opts.username)
	kingpin.Flag("password", "The password for the NSX-T manager user").OverrideDefaultFromEnvar("NSX_PASSWORD").StringVar(&opts.password)
	kingpin.Flag("insecure", "Disable TLS host verification").Default("true").BoolVar(&opts.insecure)

	promlogConfig := &promlog.Config{}
	flag.AddFlags(kingpin.CommandLine, promlogConfig)
	kingpin.HelpFlag.Short('h')
	kingpin.Parse()
	logger := promlog.New(promlogConfig)

	level.Info(logger).Log("msg", "Starting nsx_exporter", "version", version.Info())
	level.Info(logger).Log("msg", "Build context", "context", version.BuildContext())

	nsxtClient, err := newNSXTClient(opts)
	if err != nil {
		level.Error(logger).Log("msg", "Error creating nsx-t client", "err", err)
		os.Exit(1)
	}

	collector := collector.NewNSXTCollector(nsxtClient, logger)
	prometheus.MustRegister(collector)
	prometheus.MustRegister(version.NewCollector("nsx_exporter"))

	level.Info(logger).Log("msg", "Listening on address", "address", *listenAddress)
	http.Handle(*metricsPath, promhttp.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
		<head><title>NSX-T Exporter</title></head>
		<body>
		<h1>NSX Exporter</h1>
		<p>URL=<a href="` + *metricsPath + `">`+ *metricsPath +`</a></p>
		</body>
		</html>`))
	})
	if err := http.ListenAndServe(*listenAddress, nil); err != nil {
		level.Error(logger).Log("msg", "Error starting HTTP server", "err", err)
		os.Exit(1)
	}
}

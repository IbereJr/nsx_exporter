# NSX Exporter for Prometheus
Simple server that scrapes NSX stats and exports them via HTTP for Prometheus consumption.

## Getting Started

To run it:

```bash
./nsx_exporter [flags]
```

Help on flags:

```bash
./nsx_exporter --help
```

## Usage

Specify host URI for the NSX API using the `--host` flag. 
Add the credentials as well by using `--username` and `--password` flags:
```bash
./nsxt_exporter --host localhost --username user --password password
```

or with environment variables: (dont forget the prefix "NSX_")
```bash
NSX_HOST="localhost"  NSX_LISTEN=":9999" NSX_USERNAME="user" NSX_PASSWORD="pass" ./nsxt_exporter
```


Certificate validation is disabled by default, but
you can enable it using the `--nsxt.insecure=false` flag:
```bash
./nsxt_exporter --host localhost --username user --password password --insecure=false
```

### Docker

To run the nsx exporter as a Docker container, run:

```bash
docker run -p 9744:9744 cloudnativeid/nsx-exporter-linux-amd64:latest --nsxt.host localhost --nsxt.username user --nsxt.password password
```

### Building

```bash
make build
```

### Testing

```bash
make test
```


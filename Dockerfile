FROM quay.io/prometheus/busybox:latest
LABEL maintainer="ibere.tizio@tivit.com"

ARG ARCH="amd64"
ARG OS="linux"
COPY .build/${OS}-${ARCH}/nsxt_exporter /bin/nsxt_exporter

EXPOSE      9744
USER        nobody
ENTRYPOINT  [ "/bin/nsx_exporter" ]

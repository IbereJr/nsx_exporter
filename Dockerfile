#Build
FROM golang:1.16-alpine AS build
WORKDIR /src
ENV CGO_ENABLED=0
COPY go.* ./
RUN go mod download
COPY . ./
ARG TARGETOS
ARG TARGETARCH
RUN GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o /out/nsx_exporter .

#Deploy
FROM quay.io/prometheus/busybox:latest
LABEL maintainer="ibere.tizio@tivit.com"

COPY --from=build /out/nsx_exporter /bin/nsx_exporter

EXPOSE      9744
USER        nobody
ENTRYPOINT  [ "/bin/nsx_exporter" ]

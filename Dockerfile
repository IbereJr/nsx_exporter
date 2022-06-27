#Build
FROM golang:1.16-buster AS build
WORKDIR /app

#COPY go.mod ./
#COPY go.sum ./
COPY client ./
COPY collector ./

ENV GOROOT /app

#RUN go mod download
#RUN go mod tidy
#RUN go get -u -v -f all
RUN go env -w GO111MODULE=off

COPY *.go ./

RUN go build -o /nsx-exporter ./nsx_exporter.go

#Deploy
FROM quay.io/prometheus/busybox:latest
LABEL maintainer="ibere.tizio@tivit.com"

COPY --from=build /nsx-exporter /bin/nsx_exporter

EXPOSE      9744
USER        nobody
ENTRYPOINT  [ "/bin/nsx_exporter" ]

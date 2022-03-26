FROM golang

WORKDIR /go/src/github.com/pmokeev/covid-statistic

ENV GOPATH=/

COPY ./ ./
RUN apt-get update
RUN go mod download
RUN go build -o covid-statistic ./cmd/main.go

EXPOSE 8000
CMD ["./covid-statistic"]

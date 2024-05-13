FROM golang:1.20.3-alpine3.16 AS build-env

WORKDIR /src
COPY . /src
WORKDIR /src/twse-service

RUN cd /src && go build -o proxy ./twse-service/*.go

FROM ubuntu:latest

RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

COPY --from=build-env /src/twse-service/data/mis.twse.com.tw.cer /usr/local/share/ca-certificates/mis.twse.com.tw.crt

RUN update-ca-certificates

WORKDIR /app
ENV STOCK_LIST_PATH="/app/data/stock_list.csv"
COPY --from=build-env /src/proxy /app/proxy
COPY --from=build-env /src/twse-service/data /app/data

CMD [ "/app/proxy" ]
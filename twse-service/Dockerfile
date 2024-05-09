FROM golang:1.20.3-alpine3.16 AS build-env

WORKDIR /src
COPY . /src
WORKDIR /src/twse-service

RUN cd /src && go build -o proxy ./*.go

FROM alpine:3.16

WORKDIR /app
ENV STOCK_LIST_PATH="/app/data/stock_list.csv"
COPY --from=build-env /src/proxy /app/proxy
COPY --from=build-env /src/data /app/data

CMD [ "/app/proxy" ]
FROM golang:1.22.2-alpine

RUN mkdir /app

COPY brokerApp /app

CMD ["/app/brokerApp"]
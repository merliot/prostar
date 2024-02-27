# syntax=docker/dockerfile:1

FROM ghcr.io/merliot/device/device-base:latest

WORKDIR /app
COPY . .
RUN go work use .

RUN go build -tags prime -o /ps30m ./cmd
RUN /ps30m -uf2

EXPOSE 8000

ENV PORT_PRIME=8000
CMD ["/ps30m"]

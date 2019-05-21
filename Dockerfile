FROM golang:1.12-stretch AS build-env
WORKDIR /src
ENV GOFLAGS=-mod=vendor
ADD . .
RUN go build -v -o tfproviderlint cmd/tfproviderlint/tfproviderlint.go

FROM golang:1.12-stretch
WORKDIR /src
COPY --from=build-env /src/tfproviderlint /app/tfproviderlint
ENTRYPOINT ["/app/tfproviderlint"]
CMD ["./..."]

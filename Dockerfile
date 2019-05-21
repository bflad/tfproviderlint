FROM golang:1.12-stretch AS build-env
WORKDIR /src
ADD go.mod go.sum ./
RUN go mod download
ADD . .
RUN go build -v -o tfproviderlint cmd/tfproviderlint/tfproviderlint.go

FROM golang:1.12-stretch
WORKDIR /src
COPY --from=build-env /src/tfproviderlint /app/tfproviderlint
ENTRYPOINT ["/app/tfproviderlint"]
CMD ["-v", "./..."]

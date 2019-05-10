FROM golang:1.12-stretch AS build-env
WORKDIR /src
ADD go.mod go.sum ./
RUN go mod download
ADD . .
RUN go build -v -o linter cmd/linter/linter.go

FROM golang:1.12-stretch
WORKDIR /src
COPY --from=build-env /src/linter /app/linter
ENTRYPOINT ["/app/linter"]
CMD ["-v", "./..."]

FROM golang:1.17-alpine
WORKDIR /src
COPY tfproviderlint /usr/bin/tfproviderlint
ENTRYPOINT ["/usr/bin/tfproviderlint"]
CMD ["./..."]

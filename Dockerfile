FROM golang:1.21-bookworm
WORKDIR /src
COPY tfproviderlint /usr/bin/tfproviderlint
ENTRYPOINT ["/usr/bin/tfproviderlint"]
CMD ["./..."]

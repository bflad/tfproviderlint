FROM golang:1.22-bookworm
WORKDIR /src
COPY tfproviderlint /usr/bin/tfproviderlint
ENTRYPOINT ["/usr/bin/tfproviderlint"]
CMD ["./..."]

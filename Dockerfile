FROM golang:1.13-stretch
WORKDIR /src
COPY tfproviderlint /usr/bin/tfproviderlint
ENTRYPOINT ["/usr/bin/tfproviderlint"]
CMD ["./..."]

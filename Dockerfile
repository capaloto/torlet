
FROM alpine:latest

# Install a Tor node on the container
RUN apk add --no-cache tor && \
    sed "1s/^/SocksPort 127.0.0.1:9050\n/" /etc/tor/torrc.sample | \
    sed "1s/^/ControlPort 127.0.0.1:9051\n/" | \
    sed "1s/^/RunAsDaemon 1\n/" | \
    sed "1s/^/HashedControlPassword 16:FDDFE480B1DE4CDD6034CB80648F0AA1693B039F21808AB4866E0BE39A\n/" > /etc/tor/torrc

EXPOSE 9001
EXPOSE 9030
EXPOSE 9050-9080
EXPOSE 9150-9180

VOLUME "/var/lib/tor"
USER tor
CMD ["/usr/bin/tor"]

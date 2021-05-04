# TeamSpeak 3 Prometheus exporter

Simple exporter using the [HTTP query API](https://community.teamspeak.com/t/webquery-discussion-help-3-12-0-onwards/7184).

It is based on the similar [ts3exporter](https://github.com/hikhvar/ts3exporter)
from [hikhvar](ttps://github.com/hikhvar/ts3exporter).

## Docker usage

```yaml
services:
  ts3exporter:
    image: bboehmke/ts3exporter
    ports:
      - "8080:8080/tcp"
    environment:
      TS3_API_URL: "http://ts3server:10080"
      TS3_API_KEY: "ABCDEFGHIKLM"
    restart: unless-stopped
```

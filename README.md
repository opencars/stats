# Statsd

[goreport]: https://goreportcard.com/report/github.com/opencars/statsd
[goreport-img]: https://goreportcard.com/badge/github.com/opencars/statsd
[version]: https://img.shields.io/github/v/tag/opencars/statsd?sort=semver

[![Go Report][goreport-img]][goreport]
[![Version][version]][version]

## Overview

Responsible for collecting events from all micro-services over the stack.

## Event API

### Example

```JSON
{
  "kind": "authorization",
  "data": {
    "enabled": false,
    "error": "auth.token.revoked",
    "id": "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
    "ip": "172.18.0.1",
    "name": "xxx-xxx",
    "status": "succeed",
    "timestamp": "2020-03-14T00:43:20"
  }
}
```

## Development

Run PostgreSQL database

```sh
docker-compose up -Vd postgres
```

Migrate the database

```sh
migrate -source file://migrations -database postgres://postgres:password@127.0.0.1:5432/stats\?sslmode=disable up
```

Build the project

```sh
make
```

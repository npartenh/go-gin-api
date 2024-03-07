# go testing

[gin](https://github.com/gin-gonic/gin)

## Quick Start

```
go get -u github.com/gin-gonic/gin
go run .
```

## docker

```
docker build -t npartenh/go-gin-api:1.0.0 .
docker run -p 8080:8080 npartenh/go-gin-api:1.0.0 -d
```

## Confirm things are working

```
curl http://127.0.0.1:8080/ | jq
```

should produce: {"status": "ok"}

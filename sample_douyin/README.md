# Sample Douyin

## Introduction
A simple Douyin service built with `Kitex` and `Hertz` which is divided into seven microservices.

| Service Name | Usage                    | Framework   | protocol | Path         | IDL                 |
|--------------|--------------------------|-------------|----------|--------------|---------------------|
| api          | HTTP interface           | kitex/hertz | http     | cmd/api      | idl/api.thrift      |
| user         | user data management     | kitex/gorm  | thrift   | cmd/user     | idl/user.thrift     |
| video        | video data management    | kitex/gorm  | thrift   | cmd/video    | idl/video.thrift    |
| favorite     | favorite data management | kitex/gorm  | thrift   | cmd/favorite | idl/favorite.thrift |
| comment      | comment data management  | kitex/gorm  | thrift   | cmd/comment  | idl/comment.thrift  |
| message      | message data management  | kitex/gorm  | thrift   | cmd/message  | idl/message.thrift  |
| relation     | relation data management | kitex/gorm  | thrift   | cmd/relation | idl/relation.thrift |

### Architecture

![Architecture](./images/architecture.jpg)

### Call Relations

![Call Relations](./images/call-relation.jpg)

### Basic Features

- Hertz
  - Use `thrift` IDL to define HTTP interface
  - Use `hz` to generate server/client code
  - Use `Hertz` binding and validate
  - Use `obs-opentelemetry` and `jarger` for `tracing`, `metrics`, `logging`
  - Middleware
    - Use `requestid`, `jwt`, `recovery`, `pprof`, `gzip`
- Kitex
  - User `thrift` IDL to define `RPC` interface
  - Use `kitex` to generate code
  - Use `obs-opentelemetry` and `jarger` for `tracing`, `metrics`, `logging`
  - Use `registry-etcd` for service discovery and register

### Catalog Introduce

| catalog       | introduce               |
|---------------|-------------------------|
| handler       | HTTP handler            |
| service       | business logic          |
| rpc           | RPC call logic          |
| dal           | DB operation            |
| pack          | data pack               |
| apimodel      | struct definition       |
| cache         | Redis operation         |
| videoHandler  | Video stream processing |
| pkg/mw        | RPC middleware          |
| pkg/consts    | constants               |
| pkg/errno     | customized error number |
| pkg/configs   | SQL and Tracing configs |

## Code Generation

| catalog               | command                              |
|-----------------------|--------------------------------------|
| hertz_api_model       | make hertz_gen_model                 |
| hertz_api_client      | make hertz_gen_client                |
| hertz_api_new         | cd cmd/api && make hertz_new_api     |
| hertz_api_update      | cd cmd/api && make hertz_update_api  |
| kitex_user_client     | make kitex_gen_user                  |
| kitex_video_client    | make kitex_gen_video                 |
| kitex_favorite_client | make kitex_gen_favorite              |
| kitex_comment_client  | make kitex_gen_comment               |
| kitex_message_client  | make kitex_gen_message               |
| kitex_relation_client | make kitex_gen_relation              |
| kitex_user_server     | cd cmd/user && make kitex_gen_server |
| kitex_video_server    | cd cmd/video && make kitex_gen_server |
| kitex_favorite_server | cd cmd/favorite && make kitex_gen_server |
| kitex_comment_server  | cd cmd/comment && make kitex_gen_server |
| kitex_message_server  | cd cmd/message && make kitex_gen_server |
| kitex_relation_server | cd cmd/relation && make kitex_gen_server |

## Quick Start

### Setup Basic Dependence

```shell
docker-compose up -d
```

### Change config
Edit the following code block in pkg/consts/consts.go and replace it with your own Alibaba Cloud OSS AKID & ASK etc.
```go
Endpoint = "oss-c**************cs.com"
AKID     = "LTAI****************92kxo"
ASK      = "SmEa**************LuS9N3K9"
Bucket   = "douy******************67"
CDNURL   = "http://*************.cn/"
```

### Run User RPC Server

```shell
cd cmd/user
sh build.sh
sh output/bootstrap.sh
```

### Run Video RPC Server

```shell
cd cmd/video
sh build.sh
sh output/bootstrap.sh
```

### Run Relation RPC Server

```shell
cd cmd/relation
sh build.sh
sh output/bootstrap.sh
```

### Run Favorite RPC Server

```shell
cd cmd/favorite
sh build.sh
sh output/bootstrap.sh
```

### Run Comment RPC Server

```shell
cd cmd/comment
sh build.sh
sh output/bootstrap.sh
```
### Run Message RPC Server

```shell
cd cmd/message
sh build.sh
sh output/bootstrap.sh
```
### Run API Server

```shell
cd cmd/api
sh build.sh
sh output/bootstrap.sh
```

### Jaeger

Visit `http://127.0.0.1:16686/` on browser

### Grafana

Visit `http://127.0.0.1:3000/` on browser

### APK

[:memo:Sample Douyin APP Package Instructions](https://bytedance.feishu.cn/docs/doccnM9KkBAdyDhg8qaeGlIz7S7#quPkfu)

## [:point_right: MORE](https://a6i0rzkzjm.feishu.cn/docx/Xa8sdTIGJopWrNxYgeVcrmxPnKe#doxcnPDVQrEZq14hwckf1K1Taqg)


module Distributed-cloud-storage

go 1.22.2

replace Distributed-cloud-storage/handler => ./handler

replace Distributed-cloud-storage/meta => ./meta

replace Distributed-cloud-storage/db/conn => ./db/conn

require (
	// github.com/garyburd/redigo v1.6.4
	github.com/go-sql-driver/mysql v1.8.1
	github.com/gomodule/redigo v1.9.2
)

// github.com/aws/aws-sdk-go-v2 v1.26.1
// github.com/aws/aws-sdk-go-v2/config v1.27.11
// github.com/aws/aws-sdk-go-v2/service/s3 v1.53.1
require gopkg.in/amz.v3 v3.0.0-20201001071545-24fc1eceb27b

require (
	github.com/aliyun/aliyun-oss-go-sdk v3.0.2+incompatible
	github.com/gin-gonic/gin v1.9.1
	github.com/micro/cli v0.2.0
	github.com/streadway/amqp v1.1.0
)

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/bytedance/sonic v1.9.1 // indirect
	github.com/chenzhuoyu/base64x v0.0.0-20221115062448-fe3a3abad311 // indirect
	github.com/gabriel-vasile/mimetype v1.4.2 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.14.0 // indirect
	github.com/goccy/go-json v0.10.2 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/cpuid/v2 v2.2.4 // indirect
	github.com/leodido/go-urn v1.2.4 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pelletier/go-toml/v2 v2.0.8 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/ugorji/go/codec v1.2.11 // indirect
	golang.org/x/arch v0.3.0 // indirect
	golang.org/x/crypto v0.9.0 // indirect
	golang.org/x/net v0.10.0 // indirect
	golang.org/x/sys v0.8.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	golang.org/x/time v0.5.0 // indirect
	google.golang.org/protobuf v1.30.0 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

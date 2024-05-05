module Distributed-cloud-storage

go 1.22.2

replace Distributed-cloud-storage/handler => ./handler

replace Distributed-cloud-storage/meta => ./meta

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
	filippo.io/edwards25519 v1.1.0 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
)

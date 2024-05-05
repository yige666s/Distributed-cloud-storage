module Distributed-cloud-storage

go 1.22.2

replace Distributed-cloud-storage/handler => ./handler

replace Distributed-cloud-storage/meta => ./meta

require (
	// github.com/garyburd/redigo v1.6.4
	github.com/go-sql-driver/mysql v1.8.1
	github.com/gomodule/redigo v1.9.2
)

require filippo.io/edwards25519 v1.1.0 // indirect

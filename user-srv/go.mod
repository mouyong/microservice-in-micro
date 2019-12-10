module microservice-in-micro/user-srv

go 1.13

require (
	github.com/go-sql-driver/mysql v1.4.1
	github.com/golang/protobuf v1.3.2
	github.com/micro/cli v0.2.0
	github.com/micro/go-micro v1.18.0
	microservice-in-micro v0.0.0-00010101000000-000000000000
)

replace microservice-in-micro => ../

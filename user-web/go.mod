module microservice-in-micro/user-web

go 1.13

require (
	github.com/micro/cli v0.2.0
	github.com/micro/go-micro v1.18.0
	microservice-in-micro/auth v0.0.0-00010101000000-000000000000
	microservice-in-micro/user-srv v0.0.0-00010101000000-000000000000
)

replace (
	microservice-in-micro => ../
	microservice-in-micro/auth => ../auth
	microservice-in-micro/user-srv => ../user-srv
)

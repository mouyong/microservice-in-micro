module microservice-in-micro/auth

go 1.13

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-redis/redis v6.15.6+incompatible
	github.com/golang/protobuf v1.3.2
	github.com/micro/go-micro v1.18.0
	microservice-in-micro v0.0.0-00010101000000-000000000000
)

replace microservice-in-micro => ../

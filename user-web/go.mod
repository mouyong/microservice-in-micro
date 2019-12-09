module user-web

go 1.13

require (
	github.com/micro/cli v0.2.0
	github.com/micro/go-micro v1.18.0
	user-srv v0.0.0-00010101000000-000000000000
)

replace user-srv => ../user-srv/

module gateway

go 1.22.4

require (
	golang.org/x/net v0.24.0 // indirect
	golang.org/x/sys v0.19.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240318140521-94a12d6c2237 // indirect
	google.golang.org/grpc v1.64.0
	google.golang.org/protobuf v1.34.2 // indirect
)

require (
	authproto v1.0.0
	counterproto v1.0.0
	github.com/golang-jwt/jwt/v5 v5.2.1
	github.com/labstack/echo/v4 v4.12.0
	loggerclient v1.0.0
)

require (
	github.com/labstack/gommon v0.4.2 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/rabbitmq/amqp091-go v1.10.0 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.2 // indirect
	golang.org/x/crypto v0.22.0 // indirect
)

replace authproto v1.0.0 => ../authproto

replace counterproto v1.0.0 => ../counterproto

replace loggerclient v1.0.0 => ../logger-client

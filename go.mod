module gomicro_example

go 1.12

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-redis/redis v6.15.2+incompatible
	github.com/go-sql-driver/mysql v1.4.1
	github.com/golang/protobuf v1.3.1
	github.com/google/uuid v1.1.1
	github.com/gorilla/sessions v1.1.3
	github.com/micro/cli v0.2.0
	github.com/micro/go-config v1.1.0
	github.com/micro/go-micro v1.5.0
	github.com/nats-io/nats-server/v2 v2.0.0 // indirect
	github.com/prometheus/common v0.4.1
	go.uber.org/zap v1.9.1
	google.golang.org/grpc v1.21.1
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
)

replace github.com/testcontainers/testcontainer-go => github.com/testcontainers/testcontainers-go v0.0.2

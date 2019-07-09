module gomicro_example

go 1.12

require (
	github.com/SAP/go-hdb v0.14.1 // indirect
	github.com/afex/hystrix-go v0.0.0-20180502004556-fa1af6a1f4f5
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-redis/redis v6.15.2+incompatible
	github.com/go-sql-driver/mysql v1.4.1
	github.com/golang/protobuf v1.3.1
	github.com/google/uuid v1.1.1
	github.com/gorilla/sessions v1.1.3
	github.com/hashicorp/consul v1.4.4 // indirect
	github.com/hashicorp/go-gcp-common v0.5.0 // indirect
	github.com/hashicorp/go-memdb v1.0.0 // indirect
	github.com/hashicorp/go-plugin v1.0.0 // indirect
	github.com/hashicorp/vault v1.1.0 // indirect
	github.com/hashicorp/vault-plugin-auth-alicloud v0.0.0-20190320211238-36e70c54375f // indirect
	github.com/hashicorp/vault-plugin-auth-azure v0.0.0-20190320211138-f34b96803f04 // indirect
	github.com/hashicorp/vault-plugin-auth-centrify v0.0.0-20190320211357-44eb061bdfd8 // indirect
	github.com/hashicorp/vault-plugin-auth-kubernetes v0.0.0-20190328163920-79931ee7aad5 // indirect
	github.com/hashicorp/vault-plugin-secrets-ad v0.0.0-20190327182327-ed2c3d4c3d95 // indirect
	github.com/hashicorp/vault-plugin-secrets-alicloud v0.0.0-20190320213517-3307bdf683cb // indirect
	github.com/hashicorp/vault-plugin-secrets-azure v0.0.0-20190320211922-2dc8a8a5e490 // indirect
	github.com/hashicorp/vault-plugin-secrets-gcp v0.0.0-20190320211452-71903323ecb4 // indirect
	github.com/hashicorp/vault-plugin-secrets-gcpkms v0.0.0-20190320213325-9e326a9e802d // indirect
	github.com/influxdata/influxdb v1.7.5 // indirect
	github.com/micro/cli v0.2.0
	github.com/micro/go-config v1.1.0
	github.com/micro/go-micro v1.5.0
	github.com/micro/go-plugins v1.1.0
	github.com/nats-io/nats-server/v2 v2.0.0 // indirect
	github.com/prometheus/common v0.4.1
	github.com/ugorji/go/codec v0.0.0-20190320090025-2dc34c0b8780 // indirect
	go.uber.org/zap v1.9.1
	google.golang.org/grpc v1.21.1
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
	layeh.com/radius v0.0.0-20190322222518-890bc1058917 // indirect
)

replace github.com/testcontainers/testcontainer-go => github.com/testcontainers/testcontainers-go v0.0.2

replace github.com/golang/lint v0.0.0-20190313153728-d0100b6bd8b3 => golang.org/x/lint v0.0.0-20190313153728-d0100b6bd8b3

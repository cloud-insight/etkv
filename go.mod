module github.com/ucloud/etkv

go 1.12

require (
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/go-systemd v0.0.0-20190719114852-fd7a80b32e1f // indirect
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/groupcache v0.0.0-20191025150517-4a4ac3fbac33 // indirect
	github.com/gorilla/websocket v1.4.1 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.1.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.11.3
	github.com/pkg/errors v0.8.1
	github.com/prometheus/client_golang v1.2.1 // indirect
	github.com/soheilhy/cmux v0.1.4
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/tikv/client-go v0.0.0-20190822125924-d9c03d0f448b
	go.etcd.io/bbolt v1.3.3
	go.etcd.io/etcd v3.3.17+incompatible
	go.uber.org/zap v1.11.0
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9
	golang.org/x/time v0.0.0-20191024005414-555d28b269f0 // indirect
	google.golang.org/genproto v0.0.0-20191009194640-548a555dbc03 // indirect
	google.golang.org/grpc v1.24.0
	gopkg.in/yaml.v2 v2.2.4 // indirect
)

replace go.etcd.io/etcd => github.com/etcd-io/etcd v0.0.0-20191023171146-3cf2f69b5738

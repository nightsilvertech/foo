module github.com/nightsilvertech/foo

go 1.17

require (
	contrib.go.opencensus.io/exporter/zipkin v0.1.2
	github.com/afex/hystrix-go v0.0.0-20180502004556-fa1af6a1f4f5
	github.com/go-kit/kit v0.12.0
	github.com/go-kit/log v0.2.0
	github.com/go-sql-driver/mysql v1.6.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.6.0
	github.com/nightsilvertech/bar v0.0.0-20211031131457-723f23e3bc64
	github.com/openzipkin/zipkin-go v0.3.0
	github.com/nightsilvertech/utl v0.0.0-20211031131457-723f23e3bc64
	github.com/satori/go.uuid v1.2.0
	github.com/soheilhy/cmux v0.1.5
	go.opencensus.io v0.23.0
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	google.golang.org/grpc v1.41.0
	google.golang.org/protobuf v1.27.1
)

replace (
	github.com/nightsilvertech/bar v0.0.0-20211031131457-723f23e3bc64 => /home/stoic/Go/src/github.com/nightsilvertech/bar
	github.com/nightsilvertech/utl v0.0.0-20211031131457-723f23e3bc64 => /home/stoic/Go/src/github.com/nightsilvertech/utl
)

require (
	github.com/go-logfmt/logfmt v0.5.1 // indirect
	github.com/golang-jwt/jwt/v4 v4.0.0 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/gopherjs/gopherjs v0.0.0-20181017120253-0766667cb4d1 // indirect
	github.com/jtolds/gls v4.20.0+incompatible // indirect
	github.com/smartystreets/assertions v1.2.0 // indirect
	golang.org/x/net v0.0.0-20210917221730-978cfadd31cf // indirect
	golang.org/x/sys v0.0.0-20210917161153-d61c044b1678 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20211029142109-e255c875f7c7 // indirect
)

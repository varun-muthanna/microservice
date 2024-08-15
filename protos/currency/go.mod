module github.com/varun-muthanna/products-api/currency

replace github.com/varun-muthanna/products-api/currency v0.0.0 => ./protos/currency

require (
	github.com/hashicorp/go-hclog v1.6.3
	github.com/varun-muthanna/products-api/currency v0.0.0
)

go 1.22.1

require (
	google.golang.org/grpc v1.65.0
	google.golang.org/protobuf v1.34.2
)

require (
	golang.org/x/net v0.25.0 // indirect
	golang.org/x/sys v0.20.0 // indirect
	golang.org/x/text v0.15.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240528184218-531527333157 // indirect
)

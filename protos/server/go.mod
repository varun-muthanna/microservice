module github.com/varun-muthanna/products-api/server

replace github.com/varun-muthanna/products-api/currency v0.0.0 => ../currency

require (
	github.com/hashicorp/go-hclog v1.6.3
	github.com/varun-muthanna/products-api/currency v0.0.0
)

require (
	github.com/fatih/color v1.13.0 // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	golang.org/x/net v0.25.0 // indirect
	golang.org/x/sys v0.20.0 // indirect
	golang.org/x/text v0.15.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240528184218-531527333157 // indirect
	google.golang.org/grpc v1.65.0 // indirect
	google.golang.org/protobuf v1.34.2 // indirect
)

go 1.22.1

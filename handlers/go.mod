module github.com/varun-muthanna/handlers

replace github.com/varun-muthanna/data v0.0.0 => ../data

replace github.com/varun-muthanna/products-api/currency v0.0.0 => ../protos/currency

require (
	github.com/gorilla/mux v1.8.1
	github.com/varun-muthanna/data v0.0.0
	github.com/varun-muthanna/products-api/currency v0.0.0
)

require (
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator v9.31.0+incompatible // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	golang.org/x/net v0.25.0 // indirect
	golang.org/x/sys v0.20.0 // indirect
	golang.org/x/text v0.15.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240528184218-531527333157 // indirect
	google.golang.org/grpc v1.65.0 // indirect
	google.golang.org/protobuf v1.34.2 // indirect
	gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
)

go 1.22.1

module github.com/varunmuthanna/main

replace github.com/varun-muthanna/handlers v0.0.0 => ../handlers

replace github.com/varun-muthanna/data v0.0.0 => ../data

require github.com/varun-muthanna/handlers v0.0.0

require github.com/varun-muthanna/data v0.0.0 // indirect

go 1.22.1

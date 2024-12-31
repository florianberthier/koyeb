run:
	go run main.go

start-nomad:
	nomad agent -dev -bind 0.0.0.0 -config nomad.hcl

test:
	go test -v ./...
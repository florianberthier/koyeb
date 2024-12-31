run:
	go run main.go

setup:
	go mod tidy

test:
	go test -v ./...

nomad-start:
	nomad agent -dev -bind 0.0.0.0 -config nomad.hcl

nomad-stop:
	nomad agent -dev -stop

nomad-jobs:
	nomad job status

nomad-node-status:
	nomad node status -verbose

docker-containers:
	docker container ls -a

docker-nginx:
	docker pull nginx:latest
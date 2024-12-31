all: nomad-start wait-for-nomad run

nomad-start:
	@echo "Starting Nomad..."
	nomad agent -dev -bind 0.0.0.0 -config nomad.hcl &

wait-for-nomad:
	@echo "Waiting for Nomad to be ready..."
	@while ! nc -z 127.0.0.1 4646; do echo "Waiting for Nomad..."; sleep 1; done
	@echo "Nomad is ready!"

run:
	go run main.go

setup:
	go mod tidy

test:
	go test -v ./...

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
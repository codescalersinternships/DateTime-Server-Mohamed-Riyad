format:
	go fmt ./...

lint:
	golangci-lint run

build:
	go build -o httpservergin ./httpServerGin/httpServerGin.go
	go build -o httpserver ./httpServer/httpServer.go
	chmod +x httpserver
	chmod +x httpservergin


docker-build:
	docker-compose -f docker-compose.yml build

run:
	./httpservergin &
	./httpserver &

docker-run:
	docker-compose -f docker-compose.yml up

clean:
	rm -f httpservergin httpserver

test:
	go test ./...


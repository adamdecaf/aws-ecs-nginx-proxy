build:
	go tool vet .
	echo "== OSX build"
	GOOS=darwin GOARCH=386 go build .
	mv aws-ecs-nginx-proxy aws-ecs-nginx-proxy-osx

	echo "== Linux build"
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o aws-ecs-nginx-proxy-linux .
	docker build -t aws-ecs-nginx-proxy-linux:latest .

test: build
	go test -v ./...

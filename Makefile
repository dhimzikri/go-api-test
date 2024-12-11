hello:
	echo "Hello"

build:
	go build -o go/go_api_cimol main.go

run:
	go run main.go

compile:
	echo "Compiling for every OS and Platform"
	GOOS=freebsd GOARCH=386 go build -o bin/main-freebsd-386 main.go
	GOOS=linux GOARCH=386 go build -o bin/main-linux-386 main.go
	GOOS=windows GOARCH=386 go build -o bin/main-windows-386 main.go

server.swagger:
	swagger generate server --name go_api_cimol --spec swagger.yaml

all:
	compile hello
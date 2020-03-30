fmt:
	@echo ">> Running Gofmt.."
	gofmt -l -s -w .

get:
	@echo ">> Getting any missing dependencies.."
	go get ./...

install:
	go install github.com/nbw/rps

build:
	go build main.go

run :
	go run main.go

test:
	go test -v ./...




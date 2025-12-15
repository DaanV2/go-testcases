assembly:
	go tool compile -S x.go

test:
	go test ./... --cover -coverprofile=reports/coverage.out --covermode atomic --coverpkg=./...

show-coverage-report: test
	go tool cover -html=reports/coverage.out

lint:
	golangci-lint run -v --fix

install-lint:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
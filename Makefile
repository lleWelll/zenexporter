build:
	go build ./cmd/exporter/main.go

test:
	go build ./cmd/exporter/main.go
	./main --out=.

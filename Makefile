build:
	rm ./main
	go build ./cmd/exporter/main.go

test:
	make build
	./main --out=./tmp/

dw:
	make build
	cp ./main ~/bin/zenexport

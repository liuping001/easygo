.PHONY: build
build:
	go build  -gcflags "all=-N -l" -o to_redis main.go
.PHONY: clean
clean:
	rm -f to_redis
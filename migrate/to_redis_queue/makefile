.PHONY: build
build:
	go build  -gcflags "all=-N -l" -o to_redis_queue main.go
.PHONY: clean
clean:
	rm -f to_redis_queue
wyllow:
	@go build ./cmd/wyllow/
run:
	@go run ./cmd/wyllow/
test:
	@go test ./...
clean:
	@rm ./wyllow
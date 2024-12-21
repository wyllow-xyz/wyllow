wyllow:
	@make templ_gen && go build ./cmd/wyllow/
run:
	@make templ_gen && go run ./cmd/wyllow/
dev:
	@air -c .air.toml
test:
	@make templ_gen && go test ./...
clean:
	@rm ./wyllow

templ_check:
	@templ fmt -fail .
templ_gen:
	@templ fmt . && templ generate
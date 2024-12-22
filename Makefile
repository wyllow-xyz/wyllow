wyllow:
	@make templ_gen && go build -o ./tmp/wyllow ./cmd/wyllow/
run:
	@make templ_gen && go run ./cmd/wyllow/
dev:
	@air -c .air.toml
test:
	@make templ_gen && go test ./...
clean:
	@if [ -f ./tmp/wyllow ]; then rm ./tmp/wyllow; fi

templ_check:
	@templ fmt -fail .
templ_gen:
	@templ fmt . && templ generate
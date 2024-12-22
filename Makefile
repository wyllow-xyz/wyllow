wyllow:
	@make gen && go build -o ./tmp/wyllow ./cmd/wyllow/
run:
	@make && ./tmp/wyllow
dev:
	@air -c .air.toml
test:
	@make gen && go test ./...
gen:
	@make css_gen && make templ_gen
fmt:
	@templ fmt -fail .
clean:
	@if [ -f ./tmp/wyllow ]; then rm ./tmp/wyllow; fi

templ_gen:
	@templ fmt . && templ generate
css_gen:
	@pnpm run build:css
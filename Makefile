.PHONY: build run css clean

build: css
	go build -o bin/server cmd/server/main.go

run: css
	go run cmd/server/main.go

css:
	tailwindcss -i ./web/static/css/input.css -o ./web/static/css/output.css --minify

clean:
	rm -rf bin/
	rm -f web/static/css/output.css

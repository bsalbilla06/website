.PHONY: build css clean run

build: css

run:
	python3 -m http.server 8000

css:
	tailwindcss -i ./static/css/input.css -o ./static/css/output.css --minify

clean:
	rm -f static/css/output.css

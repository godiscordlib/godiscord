all: clean bundle push

push:
	./git.sh

bundle:
	./build.sh
clean:
	rm -f www/static/css/styles.min.css

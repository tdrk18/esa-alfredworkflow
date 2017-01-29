all: build alfredworkflow

build:
	go build -o bin/esa-alfredworkflow

alfredworkflow:
	awp --version `cat VERSION`

clean:
	rm -f bin/esa-alfredworkflow

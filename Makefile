

install: build
	go install .


build:
	go build .


race:
	go build -race .
	go install -race .
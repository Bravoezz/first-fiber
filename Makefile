run:
	@go run main.go

build:
	mkdir bin
	@go build -o bin/app.exe main.go

start: 
	./bin/app.exe
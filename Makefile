build:
	go build -o ./art-go

circles: build
	./art-go circles
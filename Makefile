all:
	go fmt
	GOOS=linux go build -o main main.go
debug:
	go fmt
	GOOS=linux go build -gcflags='-N -l' -o main main.go
clean:
	rm -rf main

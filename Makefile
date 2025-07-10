build:
	go build -o ./bin/git-cd main.go

install: build
	cp ./bin/got-cd /usr/local/bin/git-cd

run:
	go run main.go

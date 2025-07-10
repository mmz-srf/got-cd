build:
	go build -o ./bin/got-cd main.go

install:
	cp ./bin/got-cd /usr/local/bin/git-cd

run:
	go run main.go

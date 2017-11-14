dev:
	gofmt -w .
	dbhost=localhost dbport=27017 dbname=pokemon go run main.go

build:
	go build github.com/kurozakizz/go-api-starter

run:
	go build github.com/kurozakizz/go-api-starter
	dbhost=localhost dbport=27017 dbname=pokemon ./

# dbuild:
# 	docker build -t api-start .

# dstart:
# 	docker run -d -p 5000:5000 --name api api-starter

# dstop:
# 	docker stop api && docker rm api

# dclean:
# 	docker rmi $(docker images -q --filter "dangling=true")


dev:
	gofmt -w .
	API_VERION=1 DB_HOST=localhost DB_PORT=27017 DB_NAME=pokemon go run main.go

# build:
# 	go build github.com/kurozakizz/go-api-starter

# run:
# 	go build github.com/kurozakizz/go-api-starter
# 	API_VERION=1 DB_HOST=localhost DB_PORT=27017 DB_NAME=pokemon ./

# dbuild:
# 	docker build -t api-start .

# dstart:
# 	docker run -d -p 5000:5000 --name api api-starter

# dstop:
# 	docker stop api && docker rm api

# dclean:
# 	docker rmi $(docker images -q --filter "dangling=true")


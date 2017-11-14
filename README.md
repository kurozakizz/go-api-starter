# go-api-starter

docker build -t api-start .
docker run -it --rm -p 8080:8080 --name hello-api api-starter

docker run --rm -w /go/src/api-starter -v $(pwd):/go/src/api-starter golang:1.9.0-alpine3.6 go test -coverprofile=a.out api-starter/fav

docker run --rm -w /go/src/api-starter -v $(pwd):/go/src/api-starter golang:1.9.0-alpine3.6 go test -coverprofile=b.out api-starter/somewhere && sed '1d' b.out >> a.out && sed -i.bak 's/api-starter/./g' a.out && go tool cover -html=a.out

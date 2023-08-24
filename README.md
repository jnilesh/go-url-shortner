# go-url-shortner
a simple go gin app to shorten urls using redis

for redis container
$ docker run -d -p 6379:6379 --name my-redis-container redis:latest


for live reload
go install github.com/codegangsta/gin@latest
$ C:/Users/"user"/go/bin/gin --appPort 8080 --all --immediate run main.go
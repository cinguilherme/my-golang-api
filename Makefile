run-local:
	go run main.go

build-image:
	docker build . -t my-golang-api

run-dockerized:
	docker run -p 8080:8080 my-golang-api

git-all:
	git add . && git commit -m "all changes WIP" && git push

run:
	docker-compose down
	docker-compose up

test:
	docker-compose run app go test -v ./...

pull:
	git pull

wip:
	make git-all

bomb:
	make git-all && git push heroku master
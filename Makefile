build-image:
	docker build . -t my-golang-api

run-dockerized:
	docker run -p 8080:8080 my-goland-api

git-all:
	git add . && git commit -m "all changes" && git push

run:
	docker-compose down
	docker-compose up

test:
	docker-compose run app go test -v ./...

bomb:
	make git-all && git push heroku master
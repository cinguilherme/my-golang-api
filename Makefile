build-image:
	docker build . -t my-golang-api

run-dockerized:
	docker run -p 8080:8080 my-goland-api

git-all:
	git add . && git commit -m "all changes"
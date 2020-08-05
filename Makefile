run:
		go run main.go

fmt:
		go fmt ./...

deploy:
		git push heroku develop:master

push:
		git push origin develop
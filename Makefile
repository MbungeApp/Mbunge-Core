rund:
	docker-compose build && docker-compose up -d
run:
		go run main.go

fmt:
		go fmt ./...

deploy:
		git push heroku develop:master

swagger:
		../../bin/swag init

push:
		git push origin develop
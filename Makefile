rund:
	docker-compose build && docker-compose up -d
run:
		go run main.go

fmt:
		go fmt ./...

deploy:
		sudo git pull origin master && sudo systemctl restart api.service && sudo systemctl status api.service && sudo systemctl restart nginx.service

swagger:
		../../bin/swag init

push:
		git push origin develop

linesOfCode:
		find . -name '*.go' | xargs wc -l



.PHONY: build run stop log compup compdwn

compup:
	@docker compose -f ./deploy/docker-compose.yaml up -d

compdwn: stop

build:
	@docker build -t lattots/hasukoira -f ./build/Dockerfile .

stop:
	@docker stop hasukoira-container
	@docker rm hasukoira-container
	@docker image rm lattots/hasukoira

log:
	@docker logs hasukoira-container

push: build
	@docker push lattots/hasukoira:latest
	@scp ./deploy/.env otso@piikki.stadi.ninja:hasukoira/deploy/.env

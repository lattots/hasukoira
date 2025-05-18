.PHONY: build run stop log

build:
	@docker build -t lattots/hasukoira -f ./build/Dockerfile .

run: build
	@docker run -d --name hasukoira-container lattots/hasukoira

stop:
	@docker stop hasukoira-container
	@docker rm hasukoira-container
	@docker image rm lattots/hasukoira

log:
	@docker logs hasukoira-container

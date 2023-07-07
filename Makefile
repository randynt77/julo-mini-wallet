## clean : Remove all unused docker data from local
clean:
	@docker container prune
	@docker image prune
	@docker volume prune

## run : build and run using docker
run:
	@docker-compose up --build

## stop : stop all running service on docker
stop:
	@echo ">>>> Stopping all service on docker..."make
	@docker-compose -f docker-compose.yaml down
.PHONY: up
up:
	docker-compose up -d mongodb
	sleep 10
	docker-compose up -d mongo-express
	docker ps

.PHONY: down
down:
	docker-compose down

.PHONY: clean
clean:
	@rm -rf ./data
	echo "database data has been clean"
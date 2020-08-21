.PHONY: up
up:
	docker-compose up -d
	docker ps

.PHONY: down
down:
	docker-compose down

.PHONY: clean
clean:
	@rm -rf ./data
	echo "database data has been clean"
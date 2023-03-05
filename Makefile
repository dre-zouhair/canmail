init: redis-build redis-run mongo-init mailer-build mailer-run

redis-build:
	docker build -t redis-mailer  -f ./docker/redis.dockerfile .

redis-run:
	docker-compose -f ./docker/redis-compose.yml up -d

mongo-init:
	docker-compose -f ./docker/mongo-compose.yml up -d

mailer-build:
	docker build -t mailer  -f ./docker/mailer.dockerfile .

mailer-run:
	docker-compose -f ./docker/mailer-compose.yml up -d

.PHONY: init
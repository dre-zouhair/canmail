init: redis-build redis-run mongo-init smtp-build smtp-run mailer-build mailer-run

redis-build:
	docker build -t redis-mailer  -f ./docker/redis.dockerfile .

redis-run:
	docker-compose -f ./docker/redis-compose.yml up -d

mongo-init:
	docker-compose -f ./docker/mongo-compose.yml up -d

mailer-build:
	docker build -t mailer-image -f ./docker/mailer.dockerfile .

mailer-run:
	docker-compose -f ./docker/mailer-compose.yml up -d

smtp-build:
	docker build -t mailer-smtp-image  -f ./docker/mail-hog.dockerfile .

smtp-run:
	docker-compose -f ./docker/mail-hog-compose.yml up -d

.PHONY: init
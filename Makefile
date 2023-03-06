init: init-dependency init-app
init-redis: redis-build redis-run

init-dependency: mongodb-build mongodb-run smtp-build smtp-run

init-app: mailer-build mailer-run

redis-build:
	docker build -t redis-mailer  -f ./docker/redis.dockerfile .

redis-run:
	docker-compose -f ./docker/redis-compose.yml up -d

mongodb-build:
	docker build -t mongodb-mailer-image -f ./docker/mongodb.dockerfile .

mongodb-run:
	docker-compose -f ./docker/mongodb-compose.yml up -d

mailer-build:
	docker build -t mailer-image -f ./docker/mailer.dockerfile .

mailer-run:
	docker-compose -f ./docker/mailer-compose.yml up -d

smtp-build:
	docker build -t mailer-smtp-image  -f ./docker/mail-hog.dockerfile .

smtp-run:
	docker-compose -f ./docker/mail-hog-compose.yml up -d

.PHONY: init init-dependency init-app
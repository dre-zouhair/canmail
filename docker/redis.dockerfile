FROM alpine:latest AS build

RUN apk add --no-cache build-base git

RUN git clone --branch 6.2.8 --depth 1 https://github.com/redis/redis.git /redis

WORKDIR /redis

RUN make

# Final stage
FROM alpine:latest

COPY --from=build /redis/src/redis-server /usr/bin/redis-server

# Set the Redis password
ENV REDIS_PASSWORD password

EXPOSE 6379

CMD ["redis-server", "--requirepass", "${REDIS_PASSWORD}"]
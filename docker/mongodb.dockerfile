FROM mongo

ENV MONGO_INITDB_ROOT_USERNAME=${MONGO_USERNAME}
ENV MONGO_INITDB_ROOT_PASSWORD=${MONGO_PASSWORD}

COPY ../docker/init.js /docker-entrypoint-initdb.d/
FROM mysql:8.0

ENV MYSQL_ROOT_PASSWORD=password

ENV MYSQL_DATABASE=mydatabase

COPY ./schema.sql /docker-entrypoint-initdb.d/


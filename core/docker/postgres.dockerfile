FROM postgres

LABEL AUTHOR="RONY"

WORKDIR /my_map_points/code

ENV POSTGRES_PASSWORD="123456"

EXPOSE 5432

ADD ./core/config /docker-entrypoint-initdb.d

FROM mysql:8.0.13

ARG PUID=1000
ARG PGID=1000

RUN usermod -u $PUID mysql; groupmod -g $PGID mysql;

RUN chown -R mysql:mysql /var/lib/mysql /var/run/mysqld

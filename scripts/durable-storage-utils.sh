#!/bin/bash

init_psql_db_cluster() {
    if [ ! "$(ls -A $PGDATA)" ]; then
        echo -e "Initialising PostgreSQL DB Cluster..."
        rm -rf "${PG_HOME}"/logfile
        "${PG_BASEDIR}"/bin/pg_ctl -U "${PG_ROOTUSER}" -D "${PG_DATADIR}" initdb -o "-E=UTF8 --no-locale --auth=trust"
    fi
}

start_psql_db_cluster() {
    mkdir -p "${PG_HOME}"
    mkdir -p "${PGDATA}"
    export PG_PID=$(ps -ax | grep -v " grep " | grep "${PG_BASEDIR}"/bin/postgres | awk '{ print $1 }')
    if [ -z "${PG_PID}" ]; then
        init_psql_db_cluster
        echo -e "Starting up PostgreSQL DB Server..."
        "${PG_BASEDIR}"/bin/pg_ctl -U "${PG_ROOTUSER}" -D "${PGDATA}" -o "-p ${PG_TCP_PORT} -k ${PG_HOME}" -l "${PG_HOME}"/logfile start
        export PG_PID=$!
    fi
}

create_doko_db_user() {
    DB_USER_NAME="${1}"
    if psql -h localhost -p "${PG_TCP_PORT}" "${PG_ROOTUSER}" -t -c '\du' | cut -d \| -f 1 | grep -qw "${DB_USER_NAME}"; then
        echo "PostgreSQL DB User ${DB_USER_NAME} already exists."
    else
        createuser -e -h localhost -p "${PG_TCP_PORT}" -d -r -S "${DB_USER_NAME}"
    fi
}

create_doko_db() {
    DB_NAME="${1}"
    DB_OWNER_NAME="${2:-$PG_ROOTUSER}"
    if psql -h localhost -p "${PG_TCP_PORT}" "${PG_ROOTUSER}" -lqt | cut -d \| -f 1 | grep -qw "${DB_NAME}"; then
        echo "PostgreSQL doko DB ${DB_NAME} already exists. Continuing..."
    else
        createdb -e -h localhost -p "${PG_TCP_PORT}" -O "${DB_OWNER_NAME}" "${DB_NAME}"
    fi
}

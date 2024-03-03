#!/bin/bash

export POSTGRESQL_URL='postgres://postgres:password@localhost:5432/postgres?sslmode=disable&search_path=cumulus'

migrate -database ${POSTGRESQL_URL} -path db/migrations up
#!/bin/bash
psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER"  <<-EOSQL
    create database blog;
EOSQL
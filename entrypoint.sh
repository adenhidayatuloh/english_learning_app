#!/bin/bash

echo "Running database migrations..."
psql $DATABASE_URL -f /docker-entrypoint-initdb.d/database_backup1.sql

echo "Starting application..."
./app

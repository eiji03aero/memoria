#!/bin/bash

echo "Starting database migration"

# env vars are to be set through render-task-definition action
encoded_password=$(echo ${DB_PASSWORD} | jq -Rr @uri)
migrate --path db/migrations --database "postgresql://${DB_USER}:${encoded_password}@${DB_HOST}:5432/${DB_NAME}" -verbose up

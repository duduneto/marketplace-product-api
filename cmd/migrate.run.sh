#!/bin/bash

source ../.env

migrate -path db/migrations -database postgresql://$DB_USER:$DB_PASSWORD@$DB_HOST:$DB_PORT/$DB_DATABASE
 up

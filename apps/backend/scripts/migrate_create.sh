#!/bin/bash
docker compose exec backend migrate create -ext sql -dir ./internal/postgres/migrations $1

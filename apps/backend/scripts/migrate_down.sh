#!/bin/bash
docker compose exec backend migrate -path ./internal/postgres/migrations -database postgres://user:password@db:5432/nectlife?sslmode=disable down $1

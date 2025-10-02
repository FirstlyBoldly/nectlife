.PHONY: migrate_create migrate_up migrate_down migrate_force sqlc_generate seed gen_docs

migrate_create:
	bash scripts/migrate_create.sh $(NAME)

migrate_up:
	bash scripts/migrate_up.sh $(LEVELS)

migrate_down:
	bash scripts/migrate_down.sh $(LEVELS)

migrate_force:
	bash scripts/migrate_force.sh $(V)

sqlc_generate:
	bash scripts/sqlc_generate.sh

seed:
	bash scripts/seed.sh

gen_docs:
	bash scripts/generate_openapi_docs.sh

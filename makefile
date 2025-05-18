# Makefile

# Configuration
GOOSE=goose
DB_DRIVER=postgres
DB_DSN=user=nwt_user password=networth dbname=networth host=localhost sslmode=disable
MIGRATIONS_DIR=./migrations

# Run goose up
up:
	$(GOOSE) -dir $(MIGRATIONS_DIR) $(DB_DRIVER) "$(DB_DSN)" up

# Run goose down
down:
	$(GOOSE) -dir $(MIGRATIONS_DIR) $(DB_DRIVER) "$(DB_DSN)" down


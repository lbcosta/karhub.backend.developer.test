.PHONY: seed

COLOR_RESET=\033[0m
COLOR_GREEN=\033[0;32m
COLOR_BRIGHT_BLUE=\033[1;34m

deploy:
	@printf "$(COLOR_BRIGHT_BLUE)Deploying application...$(COLOR_RESET)\n"
	@docker-compose up --build -d --wait
	@printf "$(COLOR_GREEN)Application deployed.$(COLOR_RESET)\n\n"

seed:
	@printf "$(COLOR_BRIGHT_BLUE)Seeding database...$(COLOR_RESET)\n"
	@docker cp ./scripts/beers.sql beer_db:/tmp/beers.sql
	@docker exec -it beer_db psql -U karhub -d karhub -f /tmp/beers.sql
	@printf "$(COLOR_GREEN)Database seeded.$(COLOR_RESET)\n\n"

test:
	@printf "$(COLOR_BRIGHT_BLUE)Running tests...$(COLOR_RESET)\n"
	@docker build -t beer_api-test -f Dockerfile.test .
	@docker run --rm --env-file .env beer_api-test
	@printf "$(COLOR_GREEN)Tests passed.$(COLOR_RESET)\n\n"

cleanup:
	@printf "$(COLOR_BRIGHT_BLUE)Stopping containers...$(COLOR_RESET)\n"
	@docker-compose down
	@printf "$(COLOR_GREEN)Containers stopped.$(COLOR_RESET)\n\n"
MIGRATE = for f in supabase/migrations/*.sql; do \
  PGPASSWORD=postgres psql -h localhost -p 54322 -U postgres -d postgres -f $$f; \
done

migrate:
	@echo "⏳ Running all migrations against local dev DB..."
	@$(MIGRATE)

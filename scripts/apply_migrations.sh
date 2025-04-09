#!/bin/bash
set -e

# 環境変数（必要に応じて .env から読み込んでもOK）
DB_HOST=${DB_HOST:-localhost}
DB_PORT=${DB_PORT:-54322}
DB_NAME=${DB_NAME:-postgres}
DB_USER=${DB_USER:-postgres}
DB_PASS=${DB_PASS:-postgres}

echo "📦 Running migrations on $DB_HOST:$DB_PORT/$DB_NAME"

# 各マイグレーションSQLファイルを順に適用
for file in supabase/migrations/*.sql; do
  echo "🚀 Applying migration: $file"
  PGPASSWORD=$DB_PASS psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME -f "$file"
done

echo "✅ All migrations applied successfully."

.PHONY: build migrate seed run test setup clean dev install-air help

# =========================
# ⚙️ Default Configuration
# =========================
DB_HOST ?= localhost
DB_PORT ?= 3306
DB_USER ?= root
DB_PASS ?= @Joseph13101
DB_NAME ?= virtual_lab

# =========================
# 🏗️ Build Targets
# =========================
build:
	@echo "🔧 Building application..."
	go build -o bin/server ./cmd/server
	go build -o bin/migrate ./cmd/migrate
	@echo "✅ Build completed."

# =========================
# 🗄️ Database Migrations
# =========================
migrate: build
	@echo "📦 Running database migrations..."
	./bin/migrate
	@echo "✅ Migrations applied successfully."

# =========================
# 🌱 Seed Database
# =========================
seed:
	@echo "🌱 Seeding database..."
	mysql -h$(DB_HOST) -P$(DB_PORT) -u$(DB_USER) -p$(DB_PASS) $(DB_NAME) < scripts/seed.sql
	@echo "✅ Database seeded successfully."

# =========================
# 🚀 Run Application
# =========================
run: build
	@echo "🚀 Starting Virtual Lab API..."
	./bin/server

# =========================
# 🧪 Test Backend
# =========================
test:
	@echo "🧪 Running backend tests..."
	go run scripts/test_backend.go

# =========================
# ⚡ Setup: Migrate + Seed
# =========================
setup: migrate seed
	@echo "✅ Environment setup completed."

# =========================
# 🧹 Clean Build Artifacts
# =========================
clean:
	@echo "🧹 Cleaning up build artifacts..."
	@if exist bin (rmdir /S /Q bin) || true
	@if exist tmp (rmdir /S /Q tmp) || true
	@echo "✅ Cleanup done."

# =========================
# 🔥 Development Mode (Air)
# =========================
dev:
	@echo "🔥 Starting development server (Air hot reload)..."
	air

# =========================
# 🛠️ Install Air (One-time)
# =========================
install-air:
	@echo "⬇️ Installing Air hot reload tool..."
	go install github.com/air-verse/air@latest
	@echo "✅ Air installed successfully."

# =========================
# 🧭 Help Menu
# =========================
help:
	@echo ""
	@echo "🧭 Available Make Commands:"
	@echo "  make build        → Build the application binaries"
	@echo "  make migrate      → Run database migrations"
	@echo "  make seed         → Seed database with sample data"
	@echo "  make run          → Run the backend server"
	@echo "  make dev          → Start dev mode with Air hot reload"
	@echo "  make setup        → Migrate + seed database"
	@echo "  make clean        → Remove build artifacts"
	@echo "  make test         → Run backend test script"
	@echo "  make install-air  → Install Air tool for hot reload"
	@echo ""

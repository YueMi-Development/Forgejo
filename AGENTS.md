# Custom Forgejo Build For YueMi Use Case

> **Note**: This is a custom build of Forgejo for internal use by the YueMi project. It is not the upstream Forgejo project (https://codeberg.org/forgejo/forgejo) and is not subject to Forgejo's AI-generated code policy. This build may include modifications, additions, and integrations specific to YueMi's requirements.

---

# Forgejo Development Guide

## Project Overview

Forgejo is a self-hosted Git forge (software development platform) written in Go. It was forked from Gitea in late 2022 with the goal of keeping the project community-owned and independent. The project provides Git hosting, issue tracking, pull requests, wikis, CI/CD (Actions), package registries, and more.

- **Website**: https://forgejo.org
- **Repository**: https://codeberg.org/forgejo/forgejo
- **License**: GPL-3.0-or-later (versions 9.0+), MIT (versions before 9.0)

## Technology Stack

### Backend
- **Language**: Go 1.26.0+
- **Web Framework**: chi (go-chi/chi/v5)
- **ORM**: XORM (code.forgejo.org/xorm/xorm)
- **Database**: SQLite (default), MySQL, PostgreSQL
- **Template Engine**: Go templates (html/template)
- **CLI Framework**: urfave/cli/v3

### Frontend
- **Languages**: JavaScript, TypeScript, Vue 3, CSS
- **Build Tool**: Webpack 5
- **CSS Framework**: Tailwind CSS + Fomantic UI (custom build)
- **Testing**: Vitest, Playwright
- **Node.js**: >= 20.0.0

### Infrastructure
- **Container**: Docker (Alpine Linux-based)
- **Package Registry**: npm, go modules (code.forgejo.org)

## Build Commands

### Prerequisites
- Go 1.26.0+
- Node.js 20.0.0+
- npm
- Git with LFS support

### Full Build
```bash
make build        # Build everything (frontend + backend)
make backend      # Build Go backend only
make frontend     # Build webpack assets only
```

### Development
```bash
make watch        # Watch and rebuild on changes
make watch-backend  # Watch Go files only
make watch-frontend # Watch JS/CSS files only
```

### Dependency Management
```bash
make deps              # Install all dependencies
make deps-frontend     # npm install
make deps-backend      # go mod download
```

### Asset Generation
```bash
make generate        # Run all code generators
make generate-swagger  # Generate OpenAPI spec from annotations
make svg             # Generate SVG sprites
make fomantic        # Build Fomantic UI CSS/JS
```

## Testing Commands

### Backend Tests
```bash
make test-backend                    # Run all Go tests
make test-backend#TestSpecificName   # Run specific test by name
make test-sqlite                     # Integration tests with SQLite
make test-sqlite#TestName           # Specific integration test
make test-mysql                      # MySQL integration tests
make test-pgsql                      # PostgreSQL integration tests
```

### Frontend Tests
```bash
make test-frontend              # Vitest unit tests
make test-frontend-coverage    # With coverage report
```

### E2E Tests
```bash
make playwright                      # Install Playwright browsers
make test-e2e-sqlite                # Run e2e tests (SQLite)
make test-e2e-sqlite#TestName       # Specific e2e test
PLAYWRIGHT_PROJECT=firefox make test-e2e-sqlite  # Firefox browser
```

### Coverage
```bash
make coverage-reset
make coverage-run           # Collect coverage for all packages
make coverage-show-html     # View HTML coverage report
```

## Linting Commands

```bash
make lint              # Lint everything
make lint-fix         # Fix lint issues automatically
make lint-backend      # Go linting
make lint-go           # golangci-lint
make lint-go-fix      # Fix Go lint issues
make lint-frontend    # JS/CSS linting
make lint-js          # eslint
make lint-css         # stylelint
make lint-md          # Markdown linting
make pr-go            # CI checks for pull requests
```

## Code Style Guidelines

### Go Code
- Follow Go formatting standards (gofmt)
- Use `gofumpt` for additional formatting rules
- Import paths use `forgejo.org` module name
- Use `context.Context` for request-scoped values
- Return errors rather than logging them in internal packages
- Use testify for assertions: `require` for fatal, `assert` for non-fatal

### Import Aliasing Rules (from .golangci.yml)
- Models use pattern: `forgejo.org/models/<name>` -> `name_model`
- Services use pattern: `forgejo.org/services/<name>` -> `name_service`
- Example: `forgejo.org/models/user` -> `user_model`
- Example: `forgejo.org/services/repository` -> `repository_service`

### Forbidden Patterns
- Do not use `encoding/json` - use `forgejo.org/modules/json`
- Do not use `gopkg.in/ini.v1` - use the internal config system
- Do not use `github.com/go-git/go-git` - use `forgejo.org/modules/git`
- Do not use internal Git packages - use public AddXxx functions

### Frontend Code
- JavaScript follows ESLint configuration
- CSS uses Tailwind CSS utilities + custom styles
- Vue components use Composition API
- TypeScript is preferred for new code

### Template Guidelines
- Go templates in `templates/` directory
- Use `{{` with no space before variable names
- HTML attributes use standard Go template syntax

## Configuration

### Development Configuration
```bash
# Create a development config
cp custom/conf/app.example.ini custom/conf/app.ini

# Or use environment-to-ini tool
./environment-to-ini > custom/conf/app.ini
```

### Key Environment Variables
- `FORGEJO_CUSTOM`: Path to custom directory
- `FORGEJO_WORK_DIR`: Working directory
- `GITEA_URL`: Base URL for e2e tests (default: http://localhost:3003)
- `PROJECT_ROOT`: Root directory for tests

### Test Configuration Templates
- `tests/sqlite.ini.tmpl` - SQLite test config
- `tests/mysql.ini.tmpl` - MySQL test config
- `tests/pgsql.ini.tmpl` - PostgreSQL test config

## Database Migrations

Migrations are located in `models/forgejo_migrations/` and `models/gitea_migrations/`.

### Rules for Migrations
- Migrations must NOT import application models
- Migrations must NOT import application services
- Only allowed imports: `forgejo.org/models/db`, `forgejo.org/models/gitea_migrations/base`, `forgejo.org/models/gitea_migrations/test`

## API Development

### Swagger/OpenAPI
- API spec: `public/assets/forgejo/api.v1.yml`
- Generated from Go code annotations using go-swagger
- Run `make generate-swagger` to regenerate

### Forgejo API (OpenAPI 3.0)
- Spec: `public/assets/forgejo/api.v1.yml`
- Generated with: `make generate-forgejo-api`
- Uses chi-server code generation

## Release Process

```bash
make release                    # Create release binaries
make release-linux             # Linux binaries only
make release-sources          # Source tarball
make reproduce-build#version   # Reproducible build
```

Version is determined by git tags or `VERSION` file.

## Docker

### Build Image
```bash
docker build -t forgejo .
```

### Multi-arch Build
Uses buildx for cross-platform builds:
```bash
docker buildx build --platform linux/amd64,linux/arm64 -t forgejo .
```

### Localization
- Translations are managed via Weblate
- Do NOT submit translation PRs
- See https://forgejo.org/docs/next/contributor/localization/

## Key Dependencies (Notable)

- `code.gitea.io/sdk/gitea` - Gitea SDK (for API compatibility)
- `code.forgejo.org/xorm/xorm` - Fork of XORM ORM
- `github.com/go-chi/chi/v5` - HTTP router
- `github.com/mattn/go-sqlite3` - SQLite driver
- `github.com/go-git/go-gogit` - Git operations
- `github.com/stretchr/testify` - Testing assertions
- `github.com/prometheus/client_golang` - Metrics

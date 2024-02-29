.DEFAULT_GOAL := local-dev-all

.PHONY: clean-dev-env
clean-dev-env:
	@echo "Removing Python virtual env and Node module directories..."
	rm -rf .venv/
	rm -rf node_modules/

.PHONY: go-fmt
go-fmt:
	@echo "Go fmt...\n"

.PHONY: go-staticcheck
go-staticcheck:
	@echo "Go staticcheck...\n"
	go install honnef.co/go/tools/cmd/staticcheck@latest
	@echo

.PHONY: go-test
go-test:
	@echo "Go test...\n"

.PHONY: go-vet
go-vet:
	@echo "Go vet...\n"

.PHONY: local-dev-all
local-dev-all: prettier-format markdown-lint yaml-lint go-fmt go-test go-vet go-staticcheck

.PHONY: markdown-lint
markdown-lint:
	@echo "Linting Markdown...\n"
	npm run markdown-lint

.PHONY: prettier-format
prettier-format:
	@echo "Formatting with Prettier...\n"
	npm run format
	@echo

.PHONY: setup-dev-env
setup-dev-env: clean-dev-env
	poetry install --with dev
	npm i

.PHONY: yaml-lint
yaml-lint:
	@echo "Linting YAML...\n"
	poetry run yamllint -c .github/linters/yamllint.yml .
	@echo

define go_fmt
	gofmt -d -s -w $(1)
	@echo
endef

define go_staticcheck
	staticcheck -go 1.21 $(1)/...
	@echo
endef

define go_test
	go test $(1)/...
	@echo
endef

define go_vet
	go vet $(1)/...
	@echo
endef

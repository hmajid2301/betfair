.PHONY: build
build:
	@go build ./...

.PHONY: test
test:
	@go test -run Test ./...

.PHONY: coverage
coverage:
	@go test -run Test -cover ./...

.PHONY: lint
lint:
	@go vet ./...

.PHONY: codegen
codegen:
	@go run scripts/gen_betting_enums.go

.PHONY: docs-server
docs-server:
	@echo "Documentation:"
	@bash scripts/run-docs-server.sh

.PHONY: find_todo
find_todo:
	@grep --color=always --include=\*.go -PnRe '(//|/*).*TODO' ./ || true

.PHONY: find_fixme
find_fixme:
	@grep --color=always --include=\*.go -PnRe '(//|/*).*FIXME' ./ || true

.PHONY: find_xxx
find_xxx:
	@grep --color=always --include=\*.go -PnRe '(//|/*).*XXX' ./ || true

.PHONY: clean
clean:
	# @rm -f file

.PHONY: count
count:
	@echo "Lines of code:"
	@find . -type f -name "*.go" | xargs wc -l

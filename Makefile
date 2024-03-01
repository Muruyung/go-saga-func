install-test-coverage:
	@go get github.com/vladopajic/go-test-coverage
	@go install github.com/vladopajic/go-test-coverage

install-pre-commit:
ifdef pip
	@pip install pre-commit || sudo apt install pre-commit
else
	@pip3 install pre-commit || sudo apt install pre-commit
endif
	@make install-pre-commit-config

install-pre-commit-config:
	@pre-commit autoupdate || sudo pre-commit autoupdate
	@pre-commit install --hook-type pre-commit --hook-type pre-push --hook-type post-commit || sudo pre-commit install --hook-type pre-commit --hook-type pre-push --hook-type post-commit

unit-test-coverage:
	@go test --tags=!integration ./... -coverprofile=coverage.cov
	@go tool cover -func coverage.cov
	@make go-test-coverage || echo "ok"

go-test-coverage:
	@go-test-coverage -config .testcoverage.yml

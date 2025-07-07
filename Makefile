.PHONY: setup-git-hooks go-test

setup-git-hooks:
	./scripts/git-setup.sh

go-test:
	GO111MODULE=off go test -v $(ARGS)
	# eg make go-test ARGS="./1-unit-test-introduction/arrays"

.PHONY: install
install:
	go get github.com/stretchr/testify/require

.PHONY: test
test:
	go test ./hashmap

.PHONY: bench
bench:
	go test ./hashmap -bench=. -benchmem

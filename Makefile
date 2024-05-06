GO_MODULE = github.com/KyloReneo/googleTrendsBot

.PHONY: run
run:
	go run main.go

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: git
git:
	git add .
	git commit -m "$m"
	git push
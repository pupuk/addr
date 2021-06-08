.PHONY: all check gen

all: gen check

gen:
	go generate

check:
	gofmt -s -w ../
	go vet ../

help:
	@echo "make all - 自动构建项目代码，然后再格式化代码"
	@echo "make gen - 自动构建项目代码"
	@echo "make check - 运行 Go 工具 'fmt' and 'vet'"
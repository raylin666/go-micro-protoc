GOHOSTOS:=$(shell go env GOHOSTOS)

ifeq ($(GOHOSTOS), windows)
	#the `find.exe` is different from `find` in bash/shell.
	#to see https://docs.microsoft.com/en-us/windows-server/administration/windows-commands/find.
	#changed to use git-bash.exe to run find cli or other cli friendly, caused of every developer has a Git.
	Git_Bash= $(subst cmd\,bin\bash.exe,$(dir $(shell where git)))
	API_PROTO_FILES=$(shell $(Git_Bash) -c "find api -name *.proto")
else
	API_PROTO_FILES=$(shell find api -name *.proto)
endif

.PHONY: init
# 初始化安装脚本
init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-errors/v2@latest
	go install github.com/envoyproxy/protoc-gen-validate@latest
	go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest

.PHONY: api
# 生成API相关 Proto 文件
api:
	protoc --proto_path=./api \
		   --proto_path=./third_party \
		   --go_out=paths=source_relative:./api \
		   --go-http_out=paths=source_relative:./api \
		   --go-grpc_out=paths=source_relative:./api \
		   --validate_out=paths=source_relative,lang=go:./api \
		   --go-errors_out=paths=source_relative:./api \
		   --openapi_out=fq_schema_naming=true,default_response=false:. \
		   --openapiv2_out ./api \
		   --openapiv2_opt logtostderr=true \
		   --openapiv2_opt json_names_for_fields=false \
		   $(API_PROTO_FILES)

# 帮助命令
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help

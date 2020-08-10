# Makefileの絶対パス
ROOT := $(abspath $(dir $(lastword $(MAKEFILE_LIST))))

# 各種ツールのパス
# PROTOC := $(ROOT)/bin/protoc
PROTOC := /usr/local/bin/protoc
SKAFFOLD := $(ROOT)/bin/skaffold

# skaffold の設定ファイルパス
SKAFFOLD_YAML := skaffold.yaml

.PHONY: dev
dev: ## 自動リロードを有効にしてローカルでサーバーを起動する
	$(SKAFFOLD) dev --filename $(SKAFFOLD_YAML)

.PHONY: build
build: ## コンテナイメージのビルドを行う
	$(SKAFFOLD) build

.PHONY: test
test: ## バックエンドテストを実行する
	go test --race --coverprofile=./tmp/coverage.go.out --covermode=atomic ./...

.PHONY: download_backend
download_backend: ## バックエンドの開発で使うパッケージをダウンロードする
	go mod download

.PHONY: git_config
git_config: commit_template ## gitの設定を更新する

.PHONY: commit_template
commit_template: ## コミットメッセージのテンプレートを設定する
	git config commit.template $(ROOT)/commit_template.txt

.PHONY: go_generate
go_generate: ## go generate ./...
	go generate ./...

.PHONY: codegen_gql
codegen_gql: ## GraphQLスキーマ => Goコード生成
	cd ./backend/ && go run github.com/99designs/gqlgen --config=gqlgen.yml

.PHONY: codegen_pb
codegen_pb: ## protobuf => Goコード生成
	$(PROTOC) \
		--go_out=plugins=grpc,paths=source_relative:backend/rpcs \
		--proto_path=backend/protobuf \
		$(wildcard backend/protobuf/football/v1/*.proto)

## 順番は
## - go generate ./...
## - protobuf => Goコード生成
## - GraphQLスキーマ => Goコード生成
.PHONY: codegen
codegen: go_generate codegen_pb codegen_gql ## 自動生成処理を一括で実行する

.PHONY: clean_local
clean_local: clean_branch clean_docker ## master マージ済みのローカルブランチと不要なコンテナ等を削除する

.PHONY: clean_branch
clean_branch: ## master マージ済みのローカルブランチを削除
	git branch --merged | \
		grep -vE '^\*|master$$' | \
		xargs -I % git branch -d %
	yarn run delete-branch

.PHONY: clean_docker
clean_docker: ## 使っていないdockerコンテナ等を削除する
	docker system prune -fa --volumes

.PHONY: list
list: ## コマンド一覧を表示
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "} {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

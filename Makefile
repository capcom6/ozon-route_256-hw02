# Copyright 2022 Aleksandr Soloshenko
# 
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
# 
#     http://www.apache.org/licenses/LICENSE-2.0
# 
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

cr=gitlab-registry.ozon.dev/capcom6/homework-2/

.PHONY: run-server
run-server:
	go run cmd/server/main.go

.PHONY: run-bot
run-bot:
	go run cmd/bot/main.go

.PHONY: air-bot
air-bot:
	air -c bot.air.toml

.PHONY: air-server
air-server:
	air -c server.air.toml

.PHONY: protobuf
grpc: api/mail.proto
	protoc --go_out=pkg --go_opt=paths=source_relative --go-grpc_out=pkg --go-grpc_opt=paths=source_relative api/mail.proto \
	&& protoc -I . --grpc-gateway_out ./pkg \
    --grpc-gateway_opt logtostderr=true \
    --grpc-gateway_opt paths=source_relative \
    --grpc-gateway_opt generate_unbound_methods=true \
    api/mail.proto

.PHONY: up
up:
	docker-compose -f deployments/docker-compose-dev.yml up

.PHONY: migrate
migrate:
	goose --dir=assets/migrations up

.PHONY: docker
docker: docker-bot docker-server

.PHONY: docker-bot
docker-bot:
	docker build -f build/package/Dockerfile.bot -t "$(cr)bot:latest" .

.PHONY: docker-server
docker-server:
	docker build -f build/package/Dockerfile.server -t "$(cr)server:latest" .

.PHONY: push
push: push-bot push-server

.PHONY: docker-bot
push-bot: docker-bot
	docker push "$(cr)bot:latest"

.PHONY: docker-server
push-server: docker-server
	docker push "$(cr)server:latest"

.PHONY: test
test:
	go test -race -v -mod=readonly ./...

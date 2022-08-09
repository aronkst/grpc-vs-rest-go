run:
	docker compose up

start:
	docker compose up -d

stop:
	docker compose stop

logs:
	docker compose logs -f

test-rest:
	docker exec grpc-vs-rest-go-rest ./rest-client -count $(count)

test-grpc:
	docker exec grpc-vs-rest-go-grpc ./grpc-client -count $(count)

generate-proto:
	protoc --go_out=grpc --go-grpc_out=grpc grpc/proto/message.proto

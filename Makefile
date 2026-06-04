APP_NAME=prompt-analyzer
IMAGE_NAME=prompt-analyzer
CONTAINER_NAME=prompt-analyzer-api
CONFIG_PATH=configs/server/config.yaml
PORT=8000

.PHONY: run test test-v test-cover build clean docker-build docker-run docker-run-detached docker-stop docker-logs

run:
	CONFIG_PATH=$(CONFIG_PATH) go run .

test:
	go test ./...

test-v:
	go test ./... -v

test-cover:
	go test ./... -cover

build:
	go build -o bin/$(APP_NAME) .

clean:
	rm -rf bin
	rm -f coverage.out

docker-build:
	docker build -t $(IMAGE_NAME) .

docker-run:
	docker run --rm \
		-p $(PORT):$(PORT) \
		-e CONFIG_PATH=$(CONFIG_PATH) \
		$(IMAGE_NAME)

docker-run-detached:
	docker run -d \
		--name $(CONTAINER_NAME) \
		--restart unless-stopped \
		-p $(PORT):$(PORT) \
		-e CONFIG_PATH=$(CONFIG_PATH) \
		$(IMAGE_NAME)

docker-stop:
	docker stop $(CONTAINER_NAME) || true
	docker rm $(CONTAINER_NAME) || true

docker-logs:
	docker logs -f $(CONTAINER_NAME)
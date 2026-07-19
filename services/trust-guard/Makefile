IMAGE_NAME := bearbreach/trust-guard
IMAGE_TAG := week8

docker-build:
	docker build -t $(IMAGE_NAME):$(IMAGE_TAG) .

docker-run:
	docker run --rm \
		--name trust-guard \
		-p 8000:8000 \
		$(IMAGE_NAME):$(IMAGE_TAG)

docker-stop:
	docker stop trust-guard
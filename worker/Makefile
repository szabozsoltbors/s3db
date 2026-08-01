# Makefile

## Color definition
WHITE=\033[0m
RED=\033[1;31m
PURPLE=\033[1;35m
GREEN=\033[1;32m
YELLOW=\033[1;33m

PROFILE ?= dev
BUILD_DIR := build
ENV_FILE := env/$(PROFILE).env

ifneq (,$(wildcard $(ENV_FILE)))
include $(ENV_FILE)
export
else
$(error Environment file $(ENV_FILE) not found)
endif

.PHONY: run
run:
	@echo "\n${YELLOW}Run the project${WHITE}" && \
	LOCAL=1 S3_BUCKET_NAME=$(S3_BUCKET_NAME) go run .

.PHONY: clean
clean:
	@echo "\n${YELLOW}Cleaning build artifacts${WHITE}" && \
	rm -f -r ./$(BUILD_DIR)

.PHONY: build
build: clean
	@echo "\n${YELLOW}Build the project${WHITE}" && \
	go build -o $(BUILD_DIR)/bootstrap

zip: build
	cd $(BUILD_DIR) && zip function.zip bootstrap

.PHONY: deploy
deploy: zip
	@echo "\n${YELLOW}Deploy new AWS Lambda version${WHITE}" && \
	aws lambda update-function-code \
		--no-cli-pager \
		--function-name $(LAMBDA_NAME) \
		--zip-file fileb://$(BUILD_DIR)/$(PACKAGE_NAME).zip && \
	echo "${GREEN}Lambda code updated for $(LAMBDA_NAME).${WHITE}"
	aws lambda wait function-updated \
		--function-name $(LAMBDA_NAME)

.PHONY: invoke
invoke:
	@echo "\n${YELLOW}Invoke AWS Lambda function${WHITE}" && \
	aws lambda invoke \
		--function-name $(LAMBDA_NAME) \
		--payload '{}' \
		--log-type Tail \
		response.json && \
		echo "\n${GREEN}Lambda function invoked: ${WHITE}\n $(shell cat response.json)\n"

.PHONY: test
test:
	@echo "\n${YELLOW}Run tests${WHITE}" && \
	go test -v ./...

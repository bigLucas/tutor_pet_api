.PHONY: build
build:
	go run scripts/build-template/main.go
	sam build

.PHONY: run-local
run-local: build
	sudo sam local start-api

.PHONY: deploy
deploy: build
	sam deploy
	# rm -rf .aws-sam template.yaml

.PHONY: delete
delete:
	sam delete

.PHONY: unit-test
unit-test:
	go test -tags=unit -v ./...

.PHONY: integraiton-test
integraiton-test:
	go test -tags=integraiton

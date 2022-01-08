.PHONY: build
build:
	go run scripts/build-template/main.go
	sam build

.PHONY: run-local
run-local: build
	sudo docker rm -f dynamodb_local
	sudo docker network rm lambda_local
	sudo docker network create lambda_local
	sudo docker run --name dynamodb_local --network lambda_local -d -p 8000:8000 amazon/dynamodb-local
	aws dynamodb create-table --table-name pet-table --attribute-definitions AttributeName=ID,AttributeType=S --key-schema AttributeName=ID,KeyType=HASH --billing-mode PAY_PER_REQUEST --endpoint-url http://localhost:8000
	sudo sam local start-api --parameter-overrides "ParameterKey=APIEnv,ParameterValue=local" --docker-network lambda_local

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

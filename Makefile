test:
	go test --tags=integration --count=1 ./internal/servicedb/.

run:
	@echo "Remember to source environment variables\n\n"
	@go run cmd/app/main.go

# Used for generating tools and documentation around the openapi spec
openapi-docs:
	swagger-codegen generate -i openapi-doc.yaml -o ./openapi -l go

resources:
	docker-compose up -d
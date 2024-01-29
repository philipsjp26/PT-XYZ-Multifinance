
start:
	@echo "Starting apps..."
	@go run main.go serve

install:
	@echo "Installing ..."
	@go mod tidy && go mod vendor
	
clean:
	@echo "Cleaning"
	@rm -rf vendor

build: 
	@echo "building to binary . . ."
	@go build -o binary
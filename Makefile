build:
	@go build -o bin/expense-tracker 
run: build
	@./bin/expense-tracker
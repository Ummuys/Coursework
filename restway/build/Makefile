APP_PATH=../app/
BUILD_FILES=main.go
APP_NAME=cwapp

all: clear build run

run:
	@echo "Start the exec\n\n"
	@./$(APP_PATH)$(APP_NAME)

build:
	@echo "Building go. . ."
	@go mod tidy
	@go build -o $(APP_PATH)$(APP_NAME) $(APP_PATH)$(BUILD_FILES)
	@echo "Exec is builded\n"


clear:
	@echo "Remove old files\n"
	@rm -rf $(APP_PATH)$(APP_NAME)

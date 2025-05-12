APP_PATH=../app/
BUILD_FILES=main.go
APP_NAME=cwapp

PROTO_PATH_HEALTH = ../proto/health
PROTO_PATH_FILE_HEALTH = ../proto/health/health.proto
GEN_PATH_HEALTH = ../proto/health/gen


PROTO_PATH_STUDENTS = ../proto/students
PROTO_PATH_FILE_STUDENTS = ../proto/students/students.proto
GEN_PATH_STUDENTS = ../proto/students/gen


all: clear buildAllContracts build run
	
run:
	@echo "Start the exec\n\n"
	@./$(APP_PATH)$(APP_NAME)

build:
	@echo "Building go. . ."
	@go build -o $(APP_PATH)$(APP_NAME) $(APP_PATH)$(BUILD_FILES)
	@echo "Exec is builded\n"

deps:
	@echo "Updating package . . ."
	@go get google.golang.org/grpc@latest
	@go get google.golang.org/protobuf@latest
	@echo "Package updated\n"

	@echo "Install packages grpc/protoc . . ."
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	@echo "Package installed\n"

buildAllContracts: checkFolders deps buildContractStudents buildContractHealth
	@echo "All contracts are generated\n"
	@go mod tidy

buildContractHealth: 
	@echo "Building health contract. . ."
	@protoc -I $(PROTO_PATH_HEALTH) $(PROTO_PATH_FILE_HEALTH) \
	--go_out=$(GEN_PATH_HEALTH) \
	--go_opt=paths=source_relative \
	--go-grpc_out=$(GEN_PATH_HEALTH) \
	--go-grpc_opt=paths=source_relative
	@echo "Health contract generated\n"

buildContractStudents: 
	@echo "Building students contract. . ."
	@protoc -I $(PROTO_PATH_STUDENTS) $(PROTO_PATH_FILE_STUDENTS) \
	--go_out=$(GEN_PATH_STUDENTS) \
	--go_opt=paths=source_relative \
	--go-grpc_out=$(GEN_PATH_STUDENTS) \
	--go-grpc_opt=paths=source_relative
	@echo "Student contract generated\n"


checkFolders:
	@echo "Checking folders. . ."
	@mkdir -p $(GEN_PATH_HEALTH)
	@mkdir -p $(GEN_PATH_STUDENTS)
	@echo "Folders is exists\n"

clear:
	@echo "Remove old files\n"
	@rm -rf $(GEN_PATH_HEALTH)
	@rm -rf $(GEN_PATH_STUDENTS)
	@rm -rf $(APP_PATH)$(APP_NAME)

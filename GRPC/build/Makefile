APP_PATH=../app/
BUILD_FILES=main.go
APP_NAME=cwapp

PROTO_PATH_HEALTH = ../proto/health
PROTO_PATH_FILE_HEALTH = ../proto/health/health.proto
GEN_PATH_HEALTH = ../proto/health/gen


PROTO_PATH_STUDENTS = ../proto/students
PROTO_PATH_FILE_STUDENTS = ../proto/students/students.proto
GEN_PATH_STUDENTS = ../proto/students/gen


all: clear buildAllContracts
	

build: 
	go build -o $(APP_PATH)$(APP_NAME) $(APP_PATH)$(BUILD_FILES)

deps:
	go get google.golang.org/grpc@latest
	go mod tidy

buildAllContracts: deps buildContractStudents buildContractHealth
	@echo "All contracts are generated"

buildContractHealth: 
	@protoc -I $(PROTO_PATH_HEALTH) $(PROTO_PATH_FILE_HEALTH) \
	--go_out=$(GEN_PATH_HEALTH) \
	--go_opt=paths=source_relative \
	--go-grpc_out=$(GEN_PATH_HEALTH) \
	--go-grpc_opt=paths=source_relative
	@echo "Health contract generated"

buildContractStudents: 
	@protoc -I $(PROTO_PATH_STUDENTS) $(PROTO_PATH_FILE_STUDENTS) \
	--go_out=$(GEN_PATH_STUDENTS) \
	--go_opt=paths=source_relative \
	--go-grpc_out=$(GEN_PATH_STUDENTS) \
	--go-grpc_opt=paths=source_relative
	@echo "Student contract generated"



clear:
	rm -rf $(APP_PATH)$(APP_NAME)

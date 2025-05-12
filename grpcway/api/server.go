package server

import (
	"net"

	handHealth "github.com/ummuys/courseproject/grpcway/handlers/server/health"
	handStudents "github.com/ummuys/courseproject/grpcway/handlers/server/students"
	pbHealth "github.com/ummuys/courseproject/grpcway/proto/health/gen"
	pbStudents "github.com/ummuys/courseproject/grpcway/proto/students/gen"
	repos "github.com/ummuys/courseproject/grpcway/repository"

	"google.golang.org/grpc"
)

func StartListen(serverStatus chan struct{}) {
	// Получаем запросы по адресу http://localhost:8080
	list, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	// Инициализируем базу данных
	db := repos.NewMapDB()
	db.SetFields(1000)
	db.ToProto()

	// Инициализируем новый сервер
	grpcServer := grpc.NewServer()

	// Регистрируем путь для "прогрева" канала
	pbHealth.RegisterHealthServer(grpcServer, &handHealth.Health{})

	// Регистрируем путь для отравки словаря
	pbStudents.RegisterStudentsServer(grpcServer, &handStudents.Students{DB: db})

	// Отправляем сигнал, что сервер успешно запущен
	serverStatus <- struct{}{}

	// Если произошла ошибка на стороне сервера, то экстренно закрываем программу
	if err := grpcServer.Serve(list); err != nil {
		panic(err)
	}
}

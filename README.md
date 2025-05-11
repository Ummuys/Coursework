# 📃Coursework— Курсовая работа

## Тема курсовой работы: Сравнение gRPC и REST API для обработки больших данных

## Основные задачи:

- Изучение модели вызовов (RPC vs HTTP/JSON)
- Сравнение формата сериализации (Protocol Buffers vs JSON)
- Общий вывод о двух архитектурах

## 💫 В чем идея проекта?

Получение новых знаний в области разработки backend'a, а также ознакомление с самими архитектурами и их особенностями.

## 🚀 Быстрый старт

### 📦 Требования

- Go 1.24.2
- Make (но можно и без него)

---

### 🔧 Установка

```bash
git clone https://github.com/Ummuys/Coursework.git
cd Coursework
go mod tidy

REST-way:
<----------->
cd rest

<\> Если у вас есть make:
	cd build
	make

<\> Если у вас нет make:
	mkdir ../proto/students/gen/
	mkdir ../proto/health/gen/

	protoc -I ../proto/students ../proto/students/students.proto \
	--go_out=../proto/students/gen \
	--go_opt=paths=source_relative \
	--go-grpc_out=../proto/students/gen \
	--go-grpc_opt=paths=source_relative

	protoc -I ../proto/health ../proto/health/health.proto \
	--go_out=../proto/health/gen \
	--go_opt=paths=source_relative \
	--go-grpc_out=../proto/health/gen \
	--go-grpc_opt=paths=source_relative

	cd ../app
	go run main.go
<----------->

GRPC-way:
<----------->
cd grpc

<\> Если у вас есть make:
	cd build
	make

<\> Если у вас нет make:
	cd app
	go run main.go
<----------->




```

## 🧩 Структура проекта

```
coursework/
├── grpc/
│   ├── api/
│   ├── app/
│   ├── build/
│   ├── handlers/
│   ├── proto/
│   ├── repository/
│   ├── tools/
│   ├── go.mod
│   └── go.sum
│
├── rest/
|   ├── api/
│   ├── app/
│   ├── repository/
│   ├── tools/
│   ├── build/
│   ├── handlers/
│   ├── go.mod
│   └── go.sum
|
├── LICENSE
└── README.md
 
```

## ⚙️ Основной функционал

* Создание рандомных структур для будущего тестирования, нужно будет указать их количество.
* Проверка скорости доставки от мелких до крупных пакетов данных с помощью gRPC и REST api.

## ⚡️Автор: Евгений Егоров/Ummuys

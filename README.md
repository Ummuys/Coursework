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

- Go 1.20+
- Make (опционально, если используешь `makefile`)

---

### 🔧 Установка

```bash
git clone https://github.com/Ummuys/Coursework.git
cd Coursework
go mod tidy

REST-way:
<----------->
cd REST
cd app
go run main.go
<----------->

GRPC-way:
<----------->
cd GRPC
cd build
make (если у вас нет Make, то сгенерируйте контракт сами)
cd ../app
go run main.go
<----------->




```

## 🧩 Структура проекта

```
COURSEWORK/
├── GRPC/
│   ├── api/
│   ├── app/
│   ├── build/
│   ├── handlers/
│   ├── proto/
│   ├── repository/
│   ├── server/
│   ├── tools/
│   ├── go.mod
│   └── go.sum
│
├── REST/
|   ├── api/
│   ├── app/
│   ├── repository/
│   ├── tools/
│   ├── handlers/
│   ├── go.mod
│   └── go.sum
├── LICENSE
└── README.md
 
```

## ⚙️ Основной функционал

* Создание рандомных структур для будущего тестирования, нужно будет указать их количество.
* Проверка скорости доставки от мелких до крупных пакетов данных с помощью gRPC и REST api.

## ⚡️Автор: Евгений Егоров/Ummuys

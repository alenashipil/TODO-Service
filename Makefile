
APP_NAME = todo-app
PORT = 8080
MAIN_DIR = ./cmd/todo-app

help:  
	@echo "Используйте:"
	@echo "  make build    - Собрать приложение"
	@echo "  make run      - Запустить локально"
	@echo "  make docker   - Собрать и запустить Docker"
	@echo "  make clean    - Очистить проект"

build:  
	go build -o bin/todo-app ./cmd/todo-app

run:  
	./bin/todo-app

docker:  
	docker build -t todo-app .
	docker run -p 8080:8080 todo-app

clean:  
	rm -rf bin/
	docker stop todo-app || true
	docker rm todo-app || true
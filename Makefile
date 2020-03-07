PROJECT_NAME="github.com/nobuenhombre/suikat"

help: Makefile
	@echo "Выберите опицию сборки "$(BINARY_NAME)":"
	@sed -n 's/^##//p' $< | column -s ':' |  sed -e 's/^/ /'

## all: Удалить старые сборки, скачать необходимые пакеты, протестировать, скомпилировать
all: clean deps test build

## test: Запустить тесты
test:
	go test -v ./...

## coverage: Получить информацию о покрытии тестами кода
cover:
	go test -coverprofile=cover.out ./...
	go tool cover -html=cover.out -o cover.html
	rm -f cover.out

## clean: Удалить старые сборки
clean:
	go clean
	rm -f cover.out

## deps: Инициализация модулей, скачать все необходимые програме модули
deps:
	rm -f go.mod
	rm -f go.sum
	go mod init $(PROJECT_NAME)
	go get -u ./...

## lint: Проверка кода линтерами
lint: lint-standart lint-bugs lint-complexity lint-format lint-performance lint-style lint-unused

## lint-standart: Проверка кода стандартным набором линтереров
lint-standart:
	golangci-lint run ./...

## lint-bugs: Проверка кода линтерами bugs
lint-bugs:
	golangci-lint run -p=bugs ./...

## lint-complexity: Проверка кода линтерами complexity
lint-complexity:
	golangci-lint run -p=complexity ./...

## lint-format: Проверка кода линтерами format
lint-format:
	golangci-lint run -p=format ./...

## lint-performance: Проверка кода линтерами performance
lint-performance:
	golangci-lint run -p=performance ./...

## lint-style: Проверка кода линтерами style
lint-style:
	golangci-lint run -p=style ./...

## lint-unused: Проверка кода линтерами unused
lint-unused:
	golangci-lint run -p=unused ./...

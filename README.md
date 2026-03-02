1. Входная точка проект -- [./cmd/main.go](https://github.com/CALLlA-74/cashing/blob/master/cmd/main.go) Для запуска вызвать:
```
go run ./cmd/main.go
```
(сервер будет слушать порт 8080. на http://localhost:8080/ будет доступна html-страничка для задания параметров размена)

2. Реализация алгоритма описана в модуле [./pkg/changing_money](https://github.com/CALLlA-74/cashing/tree/master/pkg/changing_money)
3. Для получения результата с затраченным временем на выполнение лучше вызывать через метод ChangeMoney() в [./internal/domain/usecases.go](https://github.com/CALLlA-74/cashing/blob/master/internal/domain/usecases.go)

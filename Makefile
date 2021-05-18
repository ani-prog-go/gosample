.PHONY:
#.SILENT:

build:
	go build -o ./bin/main cmd/app/main.go

run: build
	./bin/main



# make run - скомпилирует и запустит
# make build- скомпилиррует
# .SILENT  не будет выводить на экран то что build
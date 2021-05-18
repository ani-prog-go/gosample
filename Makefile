.PHONY:
#.SILENT:

build:
	go build -o ./bin/main cmd/app/main.go

run: build
	./bin/main

docs: 
	go doc -all ./internal > ./doc/internal.md
	go doc -all ./cmd/app > ./doc/app.md

# make run - скомпилирует и запустит
# make build- скомпилиррует
# .SILENT  не будет выводить на экран то что build
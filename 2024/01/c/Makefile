.PHONY: run

ALL: build run

build: main

main: main.c
	@echo "Building main.c"
	gcc -O3 -o main main.c

run: build
	./main
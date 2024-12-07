.PHONY: run real

ALL: build run

build: main

main: main.c
	@echo "Building main.c"
	cc -O3 -o main main.c

run: build
	./main testinput

real: build
	./main input
ifeq ($(OS), Windows_NT)
	EXE := webserver.exe
	DELCMD := del
else
	EXE := webserver
	DELCMD := rm
endif

all: build

.PHONY: build
build:
	go build -o $(EXE) cmd/webserver/main.go

.PHONY: clean
clean:
	$(DELCMD) $(EXE)

.DEFAULT_GOAL = build

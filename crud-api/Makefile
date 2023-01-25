ifeq ($(OS), Windows_NT)
	EXE := crudserver.exe
	DELCMD := del
else
	EXE := crudserver
	DELCMD := rm
endif

MAINPATH := cmd/crudserver/main.go

all: build

.PHONY: build
build:
	go build -o $(EXE) $(MAINPATH)

.PHONY: run
run:
	go run $(MAINPATH)

.PHONY: clean
clean:
	$(DELCMD) $(EXE)

.DEFAULT_GOAL = build

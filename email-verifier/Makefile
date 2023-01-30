ifeq ($(OS), Windows_NT)
	EXE := verifier.exe
	DELCMD := del
else
	EXE := verifier
	DELCMD := rm
endif

MAINPATH := cmd/verifier/main.go

all: build

.PHONY: build
build:
	go build -o $(EXE) $(MAINPATH)

.PHONY: run
run:
	go run $(MAINPATH)

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.PHONY: clean
clean:
	$(DELCMD) $(EXE)

.DEFAULT_GOAL = build

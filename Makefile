GO           ?= go
BIN           = rc-town
SRC           = $(shell find . -type f -name '*.go')
.DEFAULT_GOAL = build

build: $(BIN)

run: $(BIN)
	./$<

$(BIN): $(SRC) go.mod go.sum
	$(GO) build

clean:
	$(RM) $(BIN)

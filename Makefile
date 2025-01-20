BINARY_LINUX=upgrader_linux
BINARY_WIN=upgrader_win.exe
BUILDS_DIR=dist
SRC_DIR=./cmd/app/main.go

all: clean test build

build_linux:
		env GOOS="linux" GOARCH="amd64" go build -o $(BUILDS_DIR)/$(BINARY_LINUX) $(SRC_DIR)

build_win:
		env GOOS="windows" GOARCH="amd64" go build -o $(BUILDS_DIR)/$(BINARY_WIN) $(SRC_DIR)

run:
		go run $(SRC_DIR)

test:
		go test ./...

clean:
		rm -f $(BUILDS_DIR)/$(BINARY_WIN) $(BUILDS_DIR)/$(BINARY_LINUX)

build: build_linux build_win


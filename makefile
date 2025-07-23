EXE=syno

.PHONY: all build run clean rebuild

all: build run

build :
		go build -o $(EXE)

run:
		./$(EXE)

clean:
		rm -f $(EXE)

rebuild: clean all


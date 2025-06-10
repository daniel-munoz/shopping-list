all: shopping-list

shopping-list: $(shell find . -name '*.go')
	go build -o shopping-list

clean:
	rm shopping-list

run: shopping-list
	./shopping-list

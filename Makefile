all: shopping-list

shopping-list: $(shell find src -name '*.go')
	(cd src; go build -o ../shopping-list)

clean:
	rm shopping-list

run: shopping-list
	./shopping-list

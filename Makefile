.PHONY: clean all

all: clean
	mkdir -p build
	gox -output="build/{{.Dir}}-{{.OS}}-{{.Arch}}" \
		-os="linux darwin windows" -arch="386 amd64"

clean:
	rm -rf build


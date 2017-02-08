.PHONY: build build-release

build:
	rm -rf build/
	gox -os="darwin" -arch="amd64" -output="build/savetofile"

build-release:
	rm -rf release/
	gox -os="linux darwin" -arch="amd64" -output="release/savetofile_{{.OS}}_{{.Arch}}"

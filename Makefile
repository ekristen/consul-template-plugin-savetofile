.PHONY: hashit

build:
	gox -os="darwin" -arch="amd64" -output="build/savetofile"

release:
	rm -rf release/
	gox -os="linux darwin" -arch="amd64" -output="release/savetofile_{{.OS}}_{{.Arch}}"
	./resources/hashes.sh

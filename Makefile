build-linux-x64:
	mkdir -p out/linux-x64
	GOOS=linux GOARCH=amd64 go build -o out/linux-x64/presigner ./main.go

build-osx-x64:
	mkdir -p out/osx-x64
	GOOS=darwin GOARCH=amd64 go build -o out/osx-x64/presigner ./main.go

release: build-linux-x64 build-osx-x64
	tar czvf out/presigner-${RELEASE_VERSION}-linux-x64.tar.gz --directory out/linux-x64/ presigner
	tar czvf out/presigner-${RELEASE_VERSION}-osx-x64.tar.gz --directory out/osx-x64/ presigner

clean:
	rm -rf out/

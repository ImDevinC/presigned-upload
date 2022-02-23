build-linux-x64:
	GOOS=linux GOARCH=amd64 go build -o presigner ./main.go

release: build-linux-x64
	tar czvf presigner-${RELEASE_VERSION}.tar.gz presigner

clean:
	rm -rf presigner
	rm -rf presigner*.tar.gz

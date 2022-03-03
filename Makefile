.DEFAULT_GOAL := build

clean:
	rm -fr release

install:
	go install -v

build: install
	CGO_ENABLED=0 DEBUG=false GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -a -v -o rattus

release: clean install
	mkdir -p release
	CGO_ENABLED=0 
	DEBUG=false
	GOOS=linux GOARCH=386 go build -ldflags="-s -w" -a -v -o release/rattus-linux-i386
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -a -v -o release/rattus-linux-amd64
	GOOS=freebsd GOARCH=386 go build -ldflags="-s -w" -a -v -o release/rattus-freebsd-i386
	GOOS=freebsd GOARCH=amd64 go build -ldflags="-s -w" -a -v -o release/rattus-freebsd-amd64
	GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -a -v -o release/rattus-darwin-amd64
	GOOS=windows GOARCH=386 go build -ldflags="-s -w" -a -v -o release/rattus-windows-i386.exe
	GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -a -v -o release/rattus-windows-amd64.exe
	upx -9 --best --ultra-brute release/rattus-linux-i386
	upx -9 --best --ultra-brute release/rattus-linux-amd64
	upx -9 --best --ultra-brute release/rattus-darwin-amd64
	upx -9 --best --ultra-brute release/rattus-windows-i386.exe
	upx -9 --best --ultra-brute release/rattus-windows-amd64.exe

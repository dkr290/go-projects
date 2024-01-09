compile:
	# compile it for Linux
	GOOS=linux GOARCH=amd64 go build -o ./bin/httphit_linux_amd64 ./cmd/httphit
	# compile it for macOS
	GOOS=darwin GOARCH=amd64 go build -o ./bin/httphit_darwin_amd64 ./cmd/httphit
	# compile it for Apple M1
	GOOS=darwin GOARCH=arm64 go build -o ./bin/httphit_darwin_arm64 ./cmd/httphit
	# compile it for Windows
	GOOS=windows GOARCH=amd64 go build -o ./bin/httphit_win_amd64.exe ./cmd/httphit
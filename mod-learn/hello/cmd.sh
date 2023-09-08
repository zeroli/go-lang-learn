go mod init example.com/hello
edit hello.go
go mod edit -replace example.com/greetings=../greetings
go mod tidy
go run .

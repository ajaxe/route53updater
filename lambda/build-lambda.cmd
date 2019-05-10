set GOOS=linux
go build -o route53updater main.go
build-lambda-zip.exe -o route53updater.zip route53updater
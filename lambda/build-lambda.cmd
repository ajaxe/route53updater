set output_dir=..\dist\lambda
del /q %output_dir%\*
set GOOS=linux
set GOARCH=amd64
set CGO_ENABLED=0
go build -o %output_dir%\route53updater main.go
build-lambda-zip.exe -o %output_dir%\route53updater.zip %output_dir%\route53updater
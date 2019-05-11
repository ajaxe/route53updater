set output_dir=..\dist\lambda
del /q %output_dir%\*
set GOOS=linux
go build -o %output_dir%\route53updater main.go
build-lambda-zip.exe -o %output_dir%\route53updater.zip %output_dir%\route53updater
# AWS Lambda To Update Route53

This folder contains AWS Lambda in Go which updates Route53 on behalf of the calling client.

## Creating a Deployment Package on Windows

To download the tool, run the following command:

```cmd
go.exe get -u github.com/aws/aws-lambda-go/cmd/build-lambda-zip
```

Run commands in powershell:

```powershell
$env:GOOS = "linux"
go build -o route53updater main.go
build-lambda-zip.exe -o ..\dist\lambda\route53updater.zip route53updater
```

Lambda package is create in `dist` folder which is used by the serverless deployment.
# aws-go-lambda
AWS Lambda using Golang

#### delete existing main.zip
`rm -rf build/main.zip` 
#### compile new project
`GOOS=linux GOARCH=amd64 go build -o build/main cmd/main.go`

#### zip result, upload this main.zip to your lambda project
`zip -jrm build/main.zip build/main`

Upload main.zip to your lambda project on AWS, change your main handler from 'hello' to 'main'

#### Using dynamodb
You can use [this branch code](https://github.com/aazzou/aws-go-lambda/tree/aws/dynamodb) for using dynamodb

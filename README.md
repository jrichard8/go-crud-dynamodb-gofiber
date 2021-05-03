# go-crud-dynamodb-gofiber
Simple of CRUD REST API backed by DynamoDB wrote in Golang based on AWS SDK V2

## TECH STACK
* GOLANG 1.16.3
* [GOFIBER v2](https://github.com/gofiber/fiber)
* [AWS SDK GO V2](https://github.com/aws/aws-sdk-go-v2)

## Getting Started
### Pre-requirements
* In your machine install [AWS-CLI](https://docs.aws.amazon.com/cli/latest/userguide/install-cliv2.html) and [CONFIGURE](https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-configure.html) to save configuration of the your database
* You also need [DOCKER](https://docs.docker.com/get-docker/)

### Install
If you already have golang installed you can install by running the command:
```sh
go get -u ./...
```

### Start DynamoDB local
docker run -p 8000:8000 amazon/dynamodb-local

### Create Table
```shell
go run scripts/createTable.go
```
Verify that table is created:
```shell
go run scripts/listTable.go
```
You should see:
```shell
0 Books
```

### Init server
To start the project run the command:
```sh
go run api/main.go
```
You can see in your terminal this log:
`service running on port  :3000`

## Usage

### Base Route
```text
    GET - http://localhost:3000/
```

#### Post
This route create on item in your database
```text
    POST - http://localhost:3000/api/books

    {
        "Title": "WhyJavaSucks"
        "Author": "me"
    }
```
```shell
curl -d "Title=WhyJavaSucks&Author=Me" -X POST -H "Accept: application/json" http://localhost:3000/api/books
```

### Get_All
This route return all data in your database
```text
    GET - http://localhost:3000/api/books
```

#### Put
This route update on item in your database
```text
    PUT - http://localhost:3000/api/books/
    {   
        "Title": "WhyJavaSucks"
        "Author": "me"
        "Rate": "0.8"
    }
```
or
```shell
curl -d "Title=WhyJavaSucks&Author=Me&Rating=0.8" -X PUT -H "Accept: application/json" http://localhost:3000/api/books
```

#### Delete
This route remove on item in your database
```text
    DELETE - http://localhost:3000/api/books/
    {   
        "Title": "WhyJavaSucks"
        "Author": "me"
    }
```
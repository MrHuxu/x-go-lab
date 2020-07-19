# Regression in Memory

## Environment

1. Nodejs
2. Go with go module enabled
3. MySQL, a database named `regression` must have been created
4. Sqlite3

## Usage

1. Clone this project;

2. Run `npm install` in `client` directory;

3. Run `go mod download` in `server` directory;

4. List and create people on web page by using MySQL:

        # in server directory
        go build -o main *.go
        ./main

        # in client directory
        npm start

        # the web page is bound to localhost:3000

5. Run ut by using Sqlite3:

        // in server directory
        go test *.go -v
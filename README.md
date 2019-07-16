# Mobile Log Server

## Project prerequisites

- GOlang
- Node
- Angular
- MySQL Database

## Project dependency getting

In main project folder: `go get -u ./`

In UI project folder: `npm install`

## Project setup

Define port, host and publicHost in *`ui/package.json`*

```
"serve": {
  ...
  "port": 4201,
  "host": "0.0.0.0",
  "publicHost": "someservice:4201"
}
```

Define API URL in *`ui/src/environments/environment.ts`*

```
server: {
  protocol: 'http://',
  host: 'someHost',
  port: ':3334'
}
```

Define server config in *`/config`* directory
By default there are two types of config:

- *config/production.json* for production
- *config/default.json* for any other type of server running

 Config example:

```
{
    "Port": ":2222",
    "DB": {
        "Username": "root",
        "Password": "password",
        "Port": ":3306",
        "Name": "logging_server"
    }
  }
```

## Running Server & UI

In root directory `go run server.go` to run API.

In **ui** directory `ng serve --open` to run client.

## Models

_Log Create Request/Response example:_

Request returns **sessionId** which is needed to identify message source.

| Request        | Response           |
| ------------- |:-------------:|
| { "appId": "com.example.sample", "deviceId": "some-model"}      | { "sessionId": 1563178048, "status:": 200 } |

_Log Send Request/Response example:_

For log sending it is necessary to attach **sessionId** to message in JSON form.

| Request        | Response           |
| ------------- |:-------------:|
| { "sessionId": 1563177895, "message": "Some message"}      | { "message": "Log succesfully added", "status:": 200 } |

_Log Record Schema in DB:_

```

ID - int

CreatedAt - date

UpdatedAt - date

DeletedAt - date

sessionId - unsigned int

appId - string

deviceID - string

sessionStart - string

sessionLastActive - string

sessionPath - string
```

#

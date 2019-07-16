# Mobile Log Server

## Project prerequisites

- Node
- Angular
- GOlang

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
  host: 'somehost',
  port: ':3334'
}
```

Define username, password, DB port and DB name in *`repository/db.go`*

```
db, err = gorm.Open("mysql", "*username*:*password*@tcp(:*port*)/*database_name*?charset=utf8&parseTime=True&loc=Local")
```

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

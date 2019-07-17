# Mobile Log Server

## Project prerequisites

- GOlang
- Node
- Angular
- MySQL Server

## Project dependency getting

In main project folder: `go get -u ./`

In UI project folder: `npm install` and `npx -p devextreme-cli devextreme add devextreme-angular` to install Angular DevExtreme.

## Project setup

### Client setup

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
### Database setup

1. Download MySQL Server, create database for the server.
1. Run server, create repository for the server side then.

### Server setup

Define server config in *`/config`* directory
By default there are two types of config:

- *config/production.json* for production
- *config/default.json* for any other type of server running

For **"DB"** object use username and password of a MySQL user for which you created database on the previous step.

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

# License

```
/*
 * ----------------------------------------------------------------------------
 * "THE BEER-WARE LICENSE" (Revision 42):
 * Pilowl<valdis.abrams@gmail.com> wrote this file. As long as you retain this notice you
 * can do whatever you want with this stuff. If we meet some day, and you think
 * this stuff is worth it, you can buy me a beer in return Poul-Henning Kamp
 * ----------------------------------------------------------------------------
 */
```
 #

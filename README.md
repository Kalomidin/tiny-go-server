Development Environment:
- Go version **go1.18.2 darwin/amd64**

Environment Setup: [Reference](https://go.dev/)

Execution:
```
go run server.go
```

Default Host & Port:
* Host/IP: 127.0.0.1 (localhost)
* Port: 1423

Endpoint List:
- **GET - /echo** : This route sends the Body of the Request back to the User.
- **POST - /reverse** : This route reverses the message that the user sent in the Body of the Request.
- **POST - /skip_odd** : This route skips characters at odd indices from the message in the Body of the Request.

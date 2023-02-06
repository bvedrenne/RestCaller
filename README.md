# RestCaller
Little executable that allow you to call URL with HTTP command. See https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods

To run the script, just call it like above replacing **CMD** by the HTTP command and **URL** by the url.
```sh
go run restlauncher.go CMD --url URL
```

The supported command are:
- GET
- POST
- HEAD
- PUT
- DELETE
- OPTIONS
- PATCH

Supported options are:
- **H** to set request header information. It supports multiple header options.
```
-H accept-type=application/json -H accept-type=text/html
```
- **C** (or **content**, only for POST, PUT, DELETE and PATCH) to set the request body.
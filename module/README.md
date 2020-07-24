- module: `example.com/hello` define in `go.mod`
- package:
  + `main` in root dir
  + `foo` in foo dir

```sh
$ go mod init example.com/hello
$ go install example.com/hello
$ go test example.com/hello/foo
```

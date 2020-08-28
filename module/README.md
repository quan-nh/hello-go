- module: `example.com/hello` define in `go.mod`
- package:
  + `main` in root dir (special package)
  + `foo` in foo dir (multiple files in a dir have the same package)

```sh
$ go mod init example.com/hello
$ go install example.com/hello
$ go test example.com/hello/foo
```

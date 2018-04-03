```
export GOPATH=$(pwd)
export CGO_LDFLAGS="`mecab-config --libs`"
export CGO_CFLAGS="-I`mecab-config --inc-dir`"
```

```
dep ensure
go run main.go
```

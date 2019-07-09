# protobuf-bug

Steps to reproduce

```bash
# we assume that you have protoc installed already
protoc --version
# libprotoc 3.7.1
protoc  --go_out=./ simple/*.proto
dep ensure
go run main.go
# Data in source: []byte{}
# Nil check result in original object: false
# Data after Unmarshal: []byte(nil)
# Nil check result in unmarshalled object: true
```

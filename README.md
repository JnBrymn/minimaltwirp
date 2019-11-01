# build

```
go mod init minimaltwirp
script/bootstrap
script/generate-twirp
```

# run

```
go run cmd/main.go
```

# query from CURL

```
curl -XPOST http://localhost:1234/twirp/github.minimaltwirp.rpc.Search/Search -H "Content-Type: application/json"    --data '{"user_query":"some query string", "repository_id": 4, "unrequired_param": "what is this?"}'
```



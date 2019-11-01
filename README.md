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

# query from Ruby

```
$ gem install twirp
$ irb

load 'ruby/rpc/search/service_twirp.rb'
client = Github::Minimaltwirp::Rpc::SearchClient.new 'http://localhost:1234/twirp'
resp = client.search(user_query:"thing", repository_id: 123)
```



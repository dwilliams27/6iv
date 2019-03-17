<div style="text-align:center"><img src ="https://www.pcgamesn.com/wp-content/uploads/legacy/civilization-2.jpg" /></div>

```bash
go get github.com/graphql-go/graphql
go get github.com/graphql-go/handler
export GOROOT=/usr/local/go
export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
go run cmd/http_handler.go
```

Example working GraphQL Queries:

> query {getNameQuery}
> mutation {incrementScountQuery(increaseAmount: 12)}
> mutation {putNameQuery(name : "peter")}

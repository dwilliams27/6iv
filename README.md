<div style="text-align:center"><img src ="https://www.pcgamesn.com/wp-content/uploads/legacy/civilization-2.jpg" /></div>

Running Go server:

```bash
go get github.com/graphql-go/graphql
go get github.com/graphql-go/handler
export GOROOT=/usr/local/go
export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
go run cmd/http_handler.go
```

Example working GraphQL Queries:

End Point: 127.0.0.1:8080/graphql
Headers: ( Content-Type -> application/graphql )
Body:

> query {getNameQuery}. 
> mutation {incrementScountQuery(increaseAmount: 12)}. 
> mutation {putNameQuery(name : "peter")}. 


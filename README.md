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

> query {getNameQuery}    
> query {getPlayersQuery}   
> mutation {incrementScoreQuery(increaseAmount: 12)}  
> mutation {putNameQuery(name : "peter")}  
> mutation {incrementScoreQuery(increaseAmount: 12), putNameQuery(name : "peter")}    
> mutation {addPlayerQuery(newPlayerName: "Peter2", newPlayerGoldCount: 30)}    
>mutation {t1: addPlayerQuery(newPlayerName: "Peter2", newPlayerGoldCount: 30), t2 : addPlayerQuery(newPlayerName: "Peter3", newPlayerGoldCount: 30), t3 : addPlayerQuery(newPlayerName: "Peter4", newPlayerGoldCount: 30), t4 : addPlayerQuery(newPlayerName: "Peter5", newPlayerGoldCount: 30), t5 : addPlayerQuery(newPlayerName: "Peter12", newPlayerGoldCount: 30), t6 : addPlayerQuery(newPlayerName: "Peter6", newPlayerGoldCount: 30), t7 : addPlayerQuery(newPlayerName: "Peter7", newPlayerGoldCount: 30), t8 : addPlayerQuery(newPlayerName: "Peter8", newPlayerGoldCount: 30)}.          
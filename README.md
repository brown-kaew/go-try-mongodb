# go-try-mongodb

An example `users` api that lookup user by id in redis before mongoDB
- http://localhost:8080/users/:id

The test data
```js
    { id: 1, name: "Kaew" },
    { id: 2, name: "Arin" },
    { id: 3, name: "John" },
```

## To run this project
```bash
docker-compose up --build app
```


## Notes
### Init module
```bash
go mod init github.com/brown-kaew/go-try-mongodb    
```

### Init server.go main package
```bash
touch server.go
```

### Load dependency after import public package
```bash
go mod tidy    
```

### Run
```bash
go run server.go
```

### Build
```bash
go build -o ./out .
```

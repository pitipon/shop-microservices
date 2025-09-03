

* Run db
```
docker-compose -f docker-compose.db.yml up
```


* Run main.go
```
go run main.go ./env/dev/.env.auth
go run main.go ./env/dev/.env.item
go run main.go ./env/dev/.env.player
```

* migration db
```
go run ./pkg/database/script/migration.go ./env/dev/.env.auth
```
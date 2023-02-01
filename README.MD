####
To run enter:
```
go run cmd/app/driver.go
or
air (hot reload)
```

docker build . --file Dockerfile --tag universal-search
docker run -p 8080:8080 universal-search
docker run -p 8080:8080 --env-file .env universal-search
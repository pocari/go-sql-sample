# setup mysql

```
docker run --rm -e MYSQL_ALLOW_EMPTY_PASSWORD=yes  -p 13306:3306 mysql:5.7
```

# run samp

```
go run main.go

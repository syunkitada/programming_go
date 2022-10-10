# sample1

```
$ curl -XPUT -H "content-type: application/json" localhost:7080/users -d '{"id": "1", "name":"hoge", "age": 10}'
{
 "id": "1",
 "name": "hoge",
 "age": 10
}

$ curl localhost:7080/users/1
{
 "id": "1",
 "name": "hoge",
 "age": 10
}
```

```
$ curl 'http://localhost:7080/apidocs.json' > swagger.json
```

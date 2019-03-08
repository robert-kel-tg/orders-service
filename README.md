# Title
Orders service in GO

```bash
docker build . -t orders-service:prod -f docker/prod/Dockerfile
docker run -it orders-service:prod sh
docker run -it -p 8080:8080 orders-service:prod
```

## Create an Order
```curl
curl -X POST \
  http://127.0.0.1:8080/orders \
  -H 'Content-Type: application/json' \
  -H 'cache-control: no-cache' \
  -d '{
	"ID" : "5678",
	"Name" : "Test1",
	"Desc" : "Aaaaaaaa"
}'
```

# golang-api by afista-pratama for tehcnical-test impactByte
 

# Run
```
go run main.go
```

# Login dummy for testing auth with JWT
```
{
    "username": "admin",
    "password": "admin"
}
```


# routes REST API WEB SERVICE
- login
```
POST http://localhost:8080/login
```
- get All Orders
```
GET http://localhost:8080/api/orders
```
- get Order by ID
```
GET http://localhost:8080/api/orders/:orderID
```
- Create new Order
```
POST http://localhost:8080/api/orders
```
- update Order
```
PUT http://localhost:8080/api/orders/:orderID
```
- delete Order
```
DELETE http://localhost:8080/api/orders/:orderID
```
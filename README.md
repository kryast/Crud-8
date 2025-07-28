CRUD ke-8

POST
curl -X POST http://localhost:8080/employees \
-H "Content-Type: application/json" \
-d '{"name":"Budi", "email":"budi@example.com", "position":"Manager"}'

GET
curl http://localhost:8080/employees
curl http://localhost:8080/employees/1

PUT
curl -X PUT http://localhost:8080/employees/1 \
-H "Content-Type: application/json" \
-d '{"name":"Budi Updated", "email":"budi2@example.com", "position":"CEO"}'

DELETE
curl -X DELETE http://localhost:8080/employees/1
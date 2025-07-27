CRUD ke-8

POST
curl -X POST http://localhost:8080/employees \
-H "Content-Type: application/json" \
-d '{"name":"Budi", "email":"budi@example.com", "position":"Manager"}'

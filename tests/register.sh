curl -X POST http://localhost:8080/api/register \
   -H "Content-Type: application/json" \
   -d '{"firstName": "Joe", "lastName": "Mama", "email": "emacom", "password": "1234", "isHospital": false, "isNGO": false}'

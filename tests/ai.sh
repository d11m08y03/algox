curl -X POST http://localhost:5000/predict \
   -H "Content-Type: application/json" \
   -d '{"date": "2024-07-12", "bloodType": "A-", "hospital": "Wellkin Hospital"}'

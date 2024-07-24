# By Token

for i in {1..11}; do 
    curl -i -X GET http://localhost:8080 -H "API_KEY: abc123";
done
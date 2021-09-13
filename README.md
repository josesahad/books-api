# A single Go app to collect books data

## To add a new book
curl http://localhost:8080/books \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"id": "4","title": "Let it be","author": "Jose Sahad","year": 2020}'

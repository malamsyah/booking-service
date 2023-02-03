# BOOKING SERVICE
## Create
````
POST /booking
````
sample request:
````
curl --location --request POST 'localhost:8080/booking' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "Alam",
    "amount": 4000
}'
````

## Get
````
localhost:8080/booking/{id}
````
sample request:
````
curl --location --request GET 'localhost:8080/booking/46141303-728b-4ed2-bbb2-9f090f37ee98'
````

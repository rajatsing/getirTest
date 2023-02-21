Getir Assignment
======================
How to run the project
----------------------
1. Clone the project
2. Make sure you have installed the latest version of go, which is `go version go1.19.5` at the time of writing assignment.
3. To run the project run `make run`
4. To run the tests run `make test`
5. Then program will be running on `localhost:8080`

----------------------------------------
to get response from the DB like this
```
{
    "code": 0,
    "msg": "Success",
    "records": [
        {
            "Key": "TAKwGc6Jr4i8Z487",
            "CreatedAt": "2017-01-28T01:22:14.398Z",
            "TotalCount": 2800
        },
        {
            "Key": "NAeQ8eX7e5TEg7oH",
            "CreatedAt": "2017-01-27T08:19:14.135Z",
            "TotalCount": 2900
        }
    ]
}
```
curl command is as follows:
```
curl --location 'http://localhost:8080/getir' \
--header 'Content-Type: text/plain' \
--data '{
    "startDate":"2016-01-26",
    "endDate":"2018-02-02",
    "minCount":2700,
    "maxCount":3000
}'
```
----------------------------------------
For in memory response like this
```
{
    "key": "rajat",
    "value": "getir"
}
```
curl command is as follows:
```
curl --location 'http://localhost:8080/in-memory' \
--header 'Content-Type: text/plain' \
--data '{
"key": "rajat",
"value": "getir"
}'
```


# Running application in docker
```
docker-compose up --build avito-tech
```
When the application starts, the necessary tables are immediately created in the database. 
Table initialization scripts are located in /schema/000001_init.up.sql
# How to use app
To interact with the application, you can use requests in the postman. 
To do this, you need to import the avito-tech.postman_collection.json

You can also send requests from the swagger http://localhost:8080/docs/index.html
# Requests and responses to interact with the application
### 1.The method of accruing funds to the balance. URI: /accrual
* Input example 
```
{
    "user_id":1,
    "amount":100
}
```
user_id - int, amount - float64
* Output example 
```
{
"success": true,
"description": "funds have been successfully credited to the balance of the user with id 1"
}
```
success - request success flag(true/false), description - request result description
### 2.Method of reserving funds from the main balance in a separate account. URI: /create_order
* Input example
```
{
    "order_id":1,
    "user_id":1,
    "service_id":2,
    "amount":50
}
```
order_id - int, service_id - int
* Output example
```
{
    "success": true,
    "description": "order 1 successfully created, funds reserved"
}
```
### 3.Revenue recognition method. URI: /charge
* Input example
```
{
    "order_id":1,
    "user_id":1,
    "service_id":2,
    "amount":50
}
```
* Output example
```
{
    "success": true,
    "description": "funds for the order 1 have been successfully charged"
}
```
### 4.User balance get method. URI: /get_balance
* Input example
```
{
    "user_id":1
}
```
* Output example
```
{
    "success": true,
    "description": "user's balance with id 1 = 50.00"
}
```
### 5.Method for transferring funds from user to user. URI: /transfer
* Input example
```
{
    "sender_id":1,
    "receiver_id":2,
    "amount":25
}
```
sender_id - int, receiver_id - int
* Output example
```
{
    "success": true,
    "description": "funds transfer completed successfully"
}
```
### 6.Method of unlocking money if the service failed to apply. URI: /cancel_order
* Input example
```
{
    "order_id":2
}
```
* Output example
```
{
    "success": true,
    "description": "order 2 cancelled, funds refunded"
}
```
### 7.method to get monthly report. URI: /get_report
* Input example
```
{
    "year":2022,
    "month":10
}
```
year - int, month - int
* Output example
```
{
    "success": true,
    "description": "the requested report has been successfully generated and is available at the link: http://localhost:8080/reports/?report=202210"
}
```
http://localhost:8080/reports/?report=202210 - generated link where you can download the report in CSV format
### 8.Method for getting a list of transactions with comments. URI: /transactions
```
{
    "user_id":1,
    "order_by":"amount, date_time",
    "limit":1,
    "offset":0
}
```
order_by - string, limit - int, offset - int
* Output example
```
[
    {
        "transaction_id": 42,
        "user_id": 1,
        "amount": -50,
        "date": "2022-10-22T17:48:46Z",
        "message": "service payment"
    }
]
```
Response - array of all user transactions with amount, date and comment.

All methods with request examples are presented in the avito-tech.postman_collection.json file in the project root.

The /docs folder contains files for the swagger. 
URI for documentation in swagger format: /docs/index.html

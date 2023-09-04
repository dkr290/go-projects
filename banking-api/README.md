API should retreive all customers

- Url should accept only numeric keys
- API will return customer as JSON object
- In case the customer id does not exists, API should return http status code 404
- In case of un unexpected error, API should return status code 500 (Internal server error) along with the eror message

- default .env file with the env parameters

```
SERVER_ADDRESS="some ip address"
SERVER_PORT="4000"
DB_USER="<some user>"
DB_PASS="<some password>"
DB_PORT="3306"
DB_NAME="banking"
DB_ADDR="localhost"
```
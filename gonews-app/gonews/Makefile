postgres:
  docker run --rm -ti --network host -e POSTGRES_PASSWORD=password POSTGRES_PASSWORD
migrate:
  migrate -source file://migrations \
          -database postgres://postgress:password@localhost/postgres?sslmode=disable up

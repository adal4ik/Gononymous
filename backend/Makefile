docker:
	docker-compose up --build -d
sql:
	docker exec -it gononymous_db_1 psql -U latte frappuccino
run:
	go run cmd/main.go 	
test:
	go test -v ./...
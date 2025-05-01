docker:
	docker-compose up --build -d
sql:
	docker exec -it frappuccino_db_1 psql -U latte frappuccino
run:
	go run cmd/app/main.go 	
test:
	go test -v ./...
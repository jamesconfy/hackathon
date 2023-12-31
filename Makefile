migrate_up:
	migrate -path db/migration -database "postgres://dbkijahx:2Ukxnz9TDvsSX1LYC83B2UqCNw7_ogiv@rajje.db.elephantsql.com/dbkijahx" -verbose up

migrate_down:
	migrate -path db/migration -database """ -verbose down

migrate_force:
	migrate -path db/migration -database "postgres://dbkijahx:2Ukxnz9TDvsSX1LYC83B2UqCNw7_ogiv@rajje.db.elephantsql.com/dbkijahx" force $(version)

run:	
	go build cheque_deposit.go && ./cheque_deposit --m=false

run_migrate:
	go build cheque_deposit.go && ./cheque_deposit --m=true

gotidy:
	go mod tidy

goinit:
	go mod init

swag:
	swag init -g cheque_deposit.go -ot go,yaml 

migrate_init:
	migrate create -ext sql -dir db/migration -seq init_schema

docker_init:
	docker build -t everybody8/cheque_deposit:v$(version) .

docker_push:
	docker push everybody8/cheque_deposit:v$(version)
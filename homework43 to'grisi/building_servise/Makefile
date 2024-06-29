mig-file:
	migrate create -ext sql -dir migrations -seq insert_terminal

# DB_URL='postgres://postgres:root@localhost:5432/go11?sslmode=disable'
mig-up:
	migrate -database 'postgres://postgres:root@localhost:5432/go11?sslmode=disable' -path migrations up

gig-dawn:
	migrate -database 'postgres://postgres:root@localhost:5432/go11?sslmode=disable' -path migrations dawn


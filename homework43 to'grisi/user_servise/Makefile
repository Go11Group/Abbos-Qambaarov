mig-file:
	migrate create -ext sql -dir migrations -seq create_user

mig-up:
	migrate -database 'postgres://postgres:root@localhost:5432/go11?sslmode=disable' -path migrations up

gig-dawn:
	migrate -database 'postgres://postgres:root@localhost:5432/go11?sslmode=disable' -path migrations dawn


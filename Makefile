.PHOY: up
.PHOY: down

up:
	docker-compose stop
	docker-compose down
	#docker volume rm users-service-db
	#docker volume create --name=users-service-db

	docker-compose -f docker-compose.yml up --build  -d
	sleep 5
	docker-compose ps
	docker-compose exec profile-service go run db/dbmigrate.go

down:
	docker-compose stop
	docker-compose down
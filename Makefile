build:
	docker-compose build man-utd

run:
	docker-compose up man-utd

migrate:
	migrate -path ./schema -database 'postgres://postgres:qwerty@0.0.0.0:5440/postgres?sslmode=disable' up
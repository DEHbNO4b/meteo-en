run:
	go build ./cmd/researcher
	./researcher -begin 2022-08-15 -radius 1000

migrate: build_migrate
	./migrator  --migrations-path=./migrations

build_migrate:
	go build -o . ./cmd/migrator

read_meteo:
	go build ./cmd/meteo
	./meteo
read_en:
	go build ./cmd/en
	./en
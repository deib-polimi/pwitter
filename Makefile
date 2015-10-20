all: run

run:
	docker-compose up

clean:
	docker rm -f pwitter_web_1
	docker rm -f pwitter_db_1

fullclean: clean
	docker rmi pwitter_web

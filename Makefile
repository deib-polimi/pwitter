all: run

run: clean
	docker-compose up -d
	echo "Web IP: "`docker inspect --format='{{.NetworkSettings.IPAddress}}' pwitter_web_1`

clean:
	docker rm -f pwitter_web_1; true
	docker rm -f pwitter_db_1; true

fullclean: clean
	docker rmi pwitter_web

client:
	docker build -t pwitter -f Dockerfile.client .

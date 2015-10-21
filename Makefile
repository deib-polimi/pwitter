WEB_OPTS=""

all: run

run: clean
	./templater opts="$(WEB_OPTS)" docker-compose.yml.tmpl > docker-compose.yml
	docker-compose up -d
	echo "Web IP: "`docker inspect --format='{{.NetworkSettings.IPAddress}}' pwitter_web_1`

clean:
	docker rm -f pwitter_web_1; true
	docker rm -f pwitter_db_1; true

fullclean: clean
	docker rmi pwitter_web

client:
	docker build -t pwitter -f Dockerfile.client .

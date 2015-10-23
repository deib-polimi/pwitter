# Pwitter
A web application to be used as a benchmarking tool for IaaS systems.  
It exposes RESTfull endpoints to perform both CPU-intensive and memory-intensive tasks.

### What Does it Do?
__Pwitter__ stores _pweets_, pieces of text that are not limited in length, and indexes them by the polarity of their _sentiment_.  
Thus, _pweets_ are stored together with their [sentiment analysis](https://textblob.readthedocs.org/en/dev/index.html).

## The Endpoints
__Pwitter__ exposes a single endpoint: `/pweets`, that can be requested using method `GET` or `POST`.  
`[GET] /pweets` is meant to be the memory-intensive task, given that at most 20 _pweets_ (each of them could contain a big amount of text) are loaded in memory; while `[POST] / pweets` is meant to be the CPU-intensive task, because the new _pweet_ has to be processed to obtain its sentiment polarity.

For CPU and Memory behavior while stressing those endpoints, go to http://affo.github.io/pwitter/.

### `[POST] /pweets`
Receives a new _pweet_ and stores it together with its sentiment analysis.

Sample request:

```
{
    "user": "Alan Touring", # not compulsory
    "body": "Having a great time with 0s and 1s!"
}
```

Sample response:

```
201

{
    "body": "Having a great time with 0s and 1s!",
    "polarity": 1.0,
    "subjectivity": 0.75,
    "user": "Alan Touring"
}
```

### `[GET] /pweets?lte=<upper_bound>&gte=<lower_bound>`
Returns at most 20 _pweets_ with polarity between `lower_bound` and `upper_bound`.  
If `lower_bound` and/or `upper_bound` are/is not specified, they/it default/s, respectively, to `-1.0` and `1.0`.

Sample requests:

```
[GET] /pweets?gte=-0.5&lte=0.23`
[GET] /pweets?gte=0.42`
[GET] /pweets`
```

Sample response:

```
200

{
    "pweets": [
        {
          "body": "Having a great time with 0s and 1s!",
          "polarity": 1.0,
          "subjectivity": 0.75,
          "user": "Alan Touring"
        }
    ]
}
```

## Installation
The installation of both server and client is performed using [Docker Engine](https://docs.docker.com/installation/ubuntulinux/) and [Docker Compose](https://docs.docker.com/compose/install/).  
Use the `Makefile` to run tasks:

 - `make clean`: destroys web and database containers;
 - `make fullclean`: as above, plus it destroys also the web image, in order to force Docker re-build it on next run;
 - `make WEB_OPTS="<opts>"`, or `make`, or `make run`: runs `clean` task and then launches the application (`docker-compose up`) in daemon mode. You can specify [`opts` for the Gunicorn](http://docs.gunicorn.org/en/19.3/settings.html#settings) web server (except from `-b`, which is fixed). At the end of the processing the IP address of the web container is returned (use it with the client).
 - `make client`: builds the client Docker image.

## Using the client
The client is dockerized. To build it for the first time, run `make client`, in this way, you will obtain a Docker image called `pwitter`.  
From now on you can launch commands as `docker run --rm pwitter -H <web_ip> <command> <opt>`. `web_ip` is the one obtained after `make` call.

For a documentation on available commands, run `docker run --rm pwitter --help`.

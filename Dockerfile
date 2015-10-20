FROM python:2.7

EXPOSE 8080

# we want to be sure to be in a bash shell
RUN rm /bin/sh && ln -s /bin/bash /bin/sh
COPY . /app/
ENV PYTHONPATH /app/
RUN pip install -r requirements.txt
RUN ln -s /app/bin/server /usr/local/bin/pwitter-server

WORKDIR /app

ENTRYPOINT [ "pwitter-server" ]

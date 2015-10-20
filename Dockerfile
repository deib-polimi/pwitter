FROM python:2.7

# we want to be sure to be in a bash shell
RUN rm /bin/sh && ln -s /bin/bash /bin/sh
COPY . /app/
ENV PYTHONPATH /app/
WORKDIR /app
RUN pip install -r requirements.txt
RUN pip install gunicorn
RUN ln -s /app/bin/server /usr/local/bin/pwitter-server


ENTRYPOINT [ "pwitter-server" ]

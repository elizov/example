FROM golang:1.11.2

WORKDIR /go/src/app
COPY . .

RUN adduser app --home /home/app --disabled-password \
    && usermod -u 1000 -G root app \
    && groupmod -g 1000 app \
    && echo 'alias ll="ls -la"' >> /home/app/.bashrc \
    && echo 'alias ..="cd .."' >> /home/app/.bashrc

RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh


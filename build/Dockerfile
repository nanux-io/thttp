# Build:
# docker build -t local/go .

# Build with ARG:
# docker build -t local/go . --build-arg go_version=1.11 uid=1001 user=app

# Run
# `docker run -v "$(pwd)":/go/src/app -it local/go /bin/sh`

ARG go_version=latest
FROM golang:$go_version

ARG uid=1000
ARG user=gopher
ARG GO111MODULE=on

# Time
ARG TZ=Europe/Paris
ARG LANG=fr_FR.UTF-8
ARG LANGUAGE=fr_FR:fr
ARG LC_ALL=fr_FR.UTF-8

# TimeZone and locales
ENV TZ=$TZ
ENV LANG $LANG
ENV LANGUAGE $LANGUAGE
ENV LC_ALL $LANG

USER root

# set the timezone and install utilities
RUN echo $TZ | tee /etc/timezone \
  && apt-get update \
  && apt-get install -y \
    apt-utils \
    locales \
    locales-all \
    tzdata \
    make \
    curl \
    git \
    htop \
    nano \
  && apt-get clean


# Configure the locales
# TODO use env $variables
RUN sed -i -e 's/# en_US.UTF-8 UTF-8/en_US.UTF-8 UTF-8/' /etc/locale.gen && \
    sed -i -e 's/# fr_FR.UTF-8 UTF-8/fr_FR.UTF-8 UTF-8/' /etc/locale.gen && \
    echo 'LANG="fr_FR.UTF-8"' > /etc/default/locale && \
    dpkg-reconfigure --frontend=noninteractive locales && \
    update-locale LANG=fr_FR.UTF-8 LC_ALL=fr_FR.UTF-8

# create user and grant the perm on the GOPATH (/go)
RUN useradd -c 'App user' -m $user -o -u $uid && chown -R $user:$user /go

# Update permissions
RUN chown $user:$user -R /home/$user/
USER $user
WORKDIR /home/go

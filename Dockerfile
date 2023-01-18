
FROM amd64/ubuntu:20.04 as confetti-component
LABEL maintainer="Reindert Vetter"

ARG WWWGROUP

ENV DEBIAN_FRONTEND noninteractive
ENV TZ=UTC
ENV INITRD No
ENV GOVERSION 1.18
ENV GOCACHE /var/www/gocache
ENV GOMODCACHE /var/www/gomodcache

WORKDIR /var/www/html

RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

## Apt-get
RUN apt-get update
RUN apt-get install -y \
       gnupg \
       gosu \
       curl \
       ca-certificates \
       zip \
       unzip \
       git \
       supervisor \
       sqlite3 \
       libcap2-bin \
       libpng-dev \
       python2 \
       wget \
       htop \
    && mkdir -p ~/.gnupg \
    && chmod 600 ~/.gnupg \
    && echo "disable-ipv6" >> ~/.gnupg/dirmngr.conf \
    && apt-key adv --homedir ~/.gnupg --keyserver hkp://keyserver.ubuntu.com:80 --recv-keys E5267A6C \
    && apt-key adv --homedir ~/.gnupg --keyserver hkp://keyserver.ubuntu.com:80 --recv-keys C300EE8C \
    && apt-get update \
    && curl -sL https://deb.nodesource.com/setup_15.x | bash - \
    && apt-get install -y nodejs \
    && curl -sS https://dl.yarnpkg.com/debian/pubkey.gpg | apt-key add - \
    && echo "deb https://dl.yarnpkg.com/debian/ stable main" > /etc/apt/sources.list.d/yarn.list \
    && apt-get update \
    && apt-get install -y yarn \
    && apt-get install -y mysql-client \
    && apt-get -y autoremove \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

RUN apt-get update && apt-get install -y gcc

## Install Go
RUN cd /usr/local && wget --progress=bar:force:noscroll https://golang.org/dl/go$GOVERSION.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go$GOVERSION.linux-amd64.tar.gz

ENV PATH PATH=$PATH:/usr/local/go/bin

RUN groupadd --force -g $WWWGROUP confetti
RUN useradd -ms /bin/bash --no-user-group -g $WWWGROUP -u 1337 confetti
RUN mkdir $GOCACHE $GOMODCACHE
RUN chown -R confetti:confetti $GOCACHE $GOMODCACHE
USER confetti

RUN GO111MODULE=off go get github.com/cespare/reflex

COPY ./ ./

CMD ~/go/bin/reflex -r '(\.go$|\.gohtml$|go\.mod$|\.env$)' -s -- sh -c "go run -race main.go app:serve"

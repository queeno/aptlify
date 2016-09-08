FROM ubuntu:16.04
MAINTAINER Philip Harries / Simon Aquino

RUN echo "deb http://repo.aptly.info/ squeeze main" >> /etc/apt/sources.list.d/aptly.list && \
  echo "deb http://ftp.uk.debian.org/debian/ stretch universe main" >> /etc/apt/sources.list.d/golang.list && \
  apt-key adv --keyserver keys.gnupg.net --recv-keys 9E3E53F19C7DE460 8B48AD6246925553 && \
  apt-get update && \
  apt-get install -y aptly golang-1.7 git && \
  ln -s /usr/lib/go-1.7/bin/go /usr/bin/go && \
  mkdir -p /root/gowork/src/github.com/queeno && \
  ln -s  /root/gowork/src/github.com/queeno/aptlify /aptlify && \
  GOPATH=/root/gowork go get github.com/mattn/gom


ENV GOPATH /root/gowork
ENV PATH /root/gowork/bin:/usr/bin:/usr/local/bin:/bin:/sbin:/usr/sbin

COPY . /root/gowork/src/github.com/queeno/aptlify

RUN cd /root/gowork/src/github.com/queeno/aptlify && \
  rm -rf /root/gowork/src/github.com/queeno/aptlify/vendor && \
  gom install && \
  gom build -o /bin/aptlify main.go

RUN /aptlify/run_tests.sh


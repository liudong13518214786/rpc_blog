# 基础镜像
FROM golang:1.12.4
# 维护者
MAINTAINER Kenny
# 下载所需要的安装包
RUN go get github.com/astaxie/beego && go get github.com/beego/bee
RUN go get github.com/lib/pq
RUN go get -u -v github.com/micro/protoc-gen-micro
RUN go get -u -v github.com/micro/go-micro
RUN go get -u -v github.com/micro/micro
RUN go get -u -v github.com/jmoiron/sqlx
RUN go get -u -v github.com/bradfitz/gomemcache/memcache
RUn go get -u -v github.com/sirupsen/logrus
EXPOSE 8877

#RUN cd /go/src/sxsim
CMD ["bee", "run"]
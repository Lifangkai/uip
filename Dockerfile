# Pull base image from image repository
#FROM 172.16.0.11:5001/centos:7
FROM 172.16.0.11:5001/alpine:v1.0.1
MAINTAINER ygz<ygz@thinkhealthi.com>

#mkdir logs path
RUN mkdir /uip && mkdir /uip/logs && mkdir /uip/bin

# deploy rest uip
COPY ./bin/uip /uip/bin/uip
COPY ./bin/start.sh /uip/bin/start.sh
RUN chmod 755 /uip/bin/start.sh

# Expose rest ports. 根据各自的端口设置
EXPOSE 5523

# Define default command.
CMD ["/uip/bin/start.sh"]


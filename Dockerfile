# Pull base image from image repository
FROM 172.16.0.11:5001/centos:7
MAINTAINER ygz<ygz@thinkhealthi.com>

#mkdir logs path
RUN mkdir /sysService && mkdir /sysService/bin

# deploy rest sysService
COPY ./bin/sys /sysService/bin/sysService
COPY ./bin/start.sh /sysService/bin/start.sh
RUN chmod 755 /sysService/bin/start.sh

# Expose rest ports. 根据各自的端口设置

# Define default command.
ENTRYPOINT ["/bin/sh","-c","/sysService/bin/start.sh"] 

FROM ubuntu:20.04

#RUN sudo apt-get update
RUN apt-get update  
RUN apt-get install -y ca-certificates
ADD main /main
ADD entrypoint.sh /entrypoint.sh
ADD static /static
WORKDIR /

EXPOSE 8091
ENTRYPOINT ["/entrypoint.sh"]


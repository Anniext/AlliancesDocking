FROM ubuntu:latest
LABEL authors="xt"
USER root
RUN apt-get update && apt-get install -y python3
ENV WORKPATH="/app"
RUN mkdir -p ${WORKPATH}
COPY ./AlliancesDocking/ ${WORKPATH}
WORKDIR ${WORKPATH}
CMD ["python tcpServer.py"]
CMD ["./server"]



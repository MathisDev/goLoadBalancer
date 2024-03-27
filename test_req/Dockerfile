FROM debian

WORKDIR /home

COPY test_request.sh .

EXPOSE 80

RUN apt update && apt upgrade -y
RUN apt install curl -y

CMD ["bash","test_request.sh"]



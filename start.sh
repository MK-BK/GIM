# !/bin/bash
docker run -itd --name=GIM.POD \
    -p 6379:6379 \
    -p 6666:3306 \
    --restart unless-stopped \
    gcr.io/google_containers/pause-amd64:3.0

docker run -itd --name=GIM.mysql \
    -e MYSQL_ROOT_PASSWORD=123456 \
    --network container:GIM.POD \
    mysql:5.7.19

docker run -itd --name=GIM.redis \
    --network container:GIM.POD \
    redis \
    --requirepass "123456"
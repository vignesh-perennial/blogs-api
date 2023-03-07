#! /bin/bash

echo "Start up..."

#install Mongo

sudo apt install -y software-properties-common gnupg apt-transport-https ca-certificates

sudo apt install -y mongodb

wget -qO - https://www.mongodb.org/static/pgp/server-5.0.asc | sudo apt-key add -

echo "deb [ arch=amd64,arm64 ] https://repo.mongodb.org/apt/ubuntu focal/mongodb-org/5.0 multiverse" | sudo tee /etc/apt/sources.list.d/mongodb-org-5.0.list

sudo apt update


sudo apt install -y mongodb-org


mongod --version


sudo systemctl status mongod


sudo systemctl start mongod


mongo --eval 'db.runCommand({ connectionStatus: 1 })'


sudo systemctl enable mongod

echo "mongo instalation done...."



# build docker images
echo "build docker images..."

# docker build -t blogs-api:latest .

# docker run -it -d -p 8080:8080 blogs-api:latest
#run
echo "Run docker..."
docker-compose up




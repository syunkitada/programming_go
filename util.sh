#!/bin/sh

COMMAND=$1

prepare_mysql() {
    sudo docker ps | grep mysql || \
        ( \
         ((sudo docker ps --all | grep mysql && sudo docker rm mysql) || echo "mysql not found") && \
         sudo docker run -v "/var/lib/docker-mysql":/var/lib/mysql --net=host --name mysql -e MYSQL_ROOT_PASSWORD=rootpass -d mysql \
        )

    mysql -uroot -prootpass -h127.0.0.1 -e "CREATE USER IF NOT EXISTS 'goapp'@'%' IDENTIFIED BY 'goapppass'; GRANT ALL ON *.* TO 'goapp'@'%'; FLUSH PRIVILEGES;"

    cat << EOS | tee ~/.my.cnf
[client]
host=127.0.0.1
port=3306
user=goapp
password=goapppass
EOS
}

$COMMAND

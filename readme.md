# Demo 使用阿里canal 同步mysql数据到elastic

## go mod 代理
export GOPROXY=https://goproxy.io

## canal
docker pull canal/canal-server:v1.1.1
### run canal
#### mysql
CREATE USER 'canal'@'%' IDENTIFIED BY 'canal';
GRANT ALL PRIVILEGES ON *.* TO 'canal'@'%' WITH GRANT OPTION;
#### canal
sh canal_run.sh
Usage:
  canal_run.sh [CONFIG]
example:
bash canal_run.sh -e canal.instance.master.address=127.0.0.1:3306 \
         -e canal.instance.dbUsername=canal \
         -e canal.instance.dbPassword=canal \
         -e canal.instance.connectionCharset=UTF-8 \
         -e canal.instance.tsdb.enable=true \
         -e canal.instance.gtidon=false \
         -e canal.instance.filter.regex=.*\\..*

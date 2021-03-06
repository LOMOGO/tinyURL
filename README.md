# 🌻tinyURL
gin框架编写的短链接服务

运行方法：

```bash
第一种：
在安装了docker和docker-compose工具之后在docker-compose.yml文件的同级目录下键入如下命令（如果在win或者mac平台下运行，需要将Dockerfile文件中对应go编译参数的部分改变为编译的二进制可以运行在win或者mac平台的编译参数）：
[lomogo@lomogo-kplw0x tinyURL]$ docker-compose up

第二种：
在Dockerfile.link同级目录下打开终端键入如下命令：
$ docker build -f Dockerfile.link -t tinyurl .
$ docker run --name mysql55 -p 13306:3306 -e MYSQL_ROOT_PASSWORD=Syq -e MYSQL_DATABASE=db_tiny_link -v /home/lomogo/docker/mysql:/var/lib/mysql -d mysql:5.5
$ docker run --link=mysql8019:mysql8019 -p 8080:8080 tinyurl
因为本项目依赖mysql，并且mysql需要先于tinyurl项目运行，所以项目跑起来需要分三步走：
第一条命令的意思是：通过Dockerfile.link文件构建镜像
第二条命令意为通过mysql:5.5镜像运行mysql55容器，并且指定了root密码是Syq,并新建了数据库db_tiny_link
第三条命令的含义是：在运行tinyurl容器的时候使用--link的方式与上面的mysql8019容器关联起来。

如果不依赖docker运行本项目，而仅仅是独立运行的话，需要将/config/config.toml中 mysql配置中的host项改为:127.0.0.1
```


👌	日志归档

👌	api文档

👌	docker

👌	配置读取

👌	mysql存储

用到的库：gin、viper、zap、lumberjack、gin-swagger、gorm


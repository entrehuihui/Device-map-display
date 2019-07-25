# Device-map-display
# Equipment map display management system

## 通过特地的接口向服务器发送数据, 服务器对数据进行保存和处理并实时展示在地图页面上.
## The data is sent to the server through a special interface, and the server saves and processes the data and displays it on the map page in real time.

## 可以实现地图地理围栏报警, 设备状态展示
## Can realize map geofence alarm, device status display

## 包含完整的用户系统与管理系统
## Includes complete user system and management system

## 依赖
## rely
    PostgreSQL, Redis, swag

## 包含 docker-compose.yml 文件, 支持在安装有docker的机器上一键启动
## Contains the docker-compose.yml file, which supports one-click launch on a machine with docker installed
    docker-compose.yml 启动步骤:
    1: 编译软件: 
        linux: ./linux.sh
        window: window.cmd
    2: docker-compose up -d

    Docker-compose.yml startup steps:
    1: Compile the software:
        linux: ./linux.sh
        window: window.cmd
    2: docker-compose up -d

## 系统使用websocket与前端页面进行实时数据推送, 在websocket断开是会进行重连
## The system uses websocket and front-end pages for real-time data push. Disconnected in websocket will be reconnected

## API接口: http://127.0.0.1:8800/swagger/index.html
## API interface: http://127.0.0.1:8800/swagger/index.html

## config.toml 配置数据差数
## config.toml Configuration data difference


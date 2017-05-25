
## 应用部署

### Fabric部署

#### 网络环境
1. 无线网络,网关192.168.1.1,无密码
2. 192.168.1.20 部署fabric:vp0,mariadb,慈善应用及rsa管理
3. 192.168.1.21 部署fabric:vp1
4. 192.168.1.22 部署fabric:vp2
5. 192.168.1.23 部署fabric:vp3
6. 20-23用户名a 密码p@ssw0rd

#### 192.168.1.20

- fabric:vp0

    启动: ssh登入,进入/home/a/fabric/vp0,运行sudo docker-compose -f server_vp0.yml up
    停止: ssh登入,进入/home/a/fabric/vp0,运行sudo docker-compose -f server_vp0.yml stop
- mariadb
    
    采用docker运行开机启动, sudo docker start mariadb,sudo docker stop mariadb
- 慈善应用

- rsa管理
    启动: ssh登入,进入/home/a/rsaserver,运行nohup ./main &,浏览器访问 http://192.168.1.20:3000/ping,返回pong则运行成功
    停止: ps -ef | grep main,查找到进程id,kill 进程

#### 192.168.1.21

- fabric:vp1

    启动: ssh登入,进入/home/a/deploy/vp1,运行sudo docker-compose -f server_vp1.yml up
    停止: ssh登入,进入/home/a/deploy/vp1,运行sudo docker-compose -f server_vp1.yml stop

#### 192.168.1.22

- fabric:vp2

    启动: ssh登入,进入/home/a/fabric/vp2,运行sudo docker-compose -f server_vp2.yml up
    停止: ssh登入,进入/home/a/fabric/vp2,运行sudo docker-compose -f server_vp2.yml stop

#### 192.168.1.23

- fabric:vp3

    启动: ssh登入,进入/home/a/fabric/vp3,运行sudo docker-compose -f server_vp3.yml up
    停止: ssh登入,进入/home/a/fabric/vp3,运行sudo docker-compose -f server_vp3.yml stop

### Fabric启动
依次启动vp0,vp1,vp2,vp3

#### 192.168.1.20

- fabric:vp0

    启动: ssh登入,进入/home/a/fabric/vp0,运行sudo docker-compose -f server_vp0.yml up, 启动成功后直接关闭ssh终端
    停止: ssh登入,进入vp0,运行sudo docker-compose -f server_vp0.yml stop

- 慈善应用

- rsa管理
    启动: ssh登入,进入rsaserver,运行./main,浏览器访问 http://192.168.1.20:3000/ping,返回pong则运行成功,直接关闭ssh终端
    停止: ps -ef | grep main,查找到进程id,kill 进程

#### 192.168.1.21

- fabric:vp1

    启动: ssh登入,进入/home/a/deploy/vp1,运行sudo docker-compose -f server_vp1.yml up, 启动成功后直接关闭ssh终端
    停止: ssh登入,进入vp1,运行sudo docker-compose -f server_vp1.yml stop

#### 192.168.1.22

- fabric:vp2

    启动: ssh登入,进入/home/a/fabric/vp2,运行sudo docker-compose -f server_vp2.yml up, 启动成功后直接关闭ssh终端
    停止: ssh登入,进入/home/a/fabric/vp2,运行sudo docker-compose -f server_vp2.yml stop

#### 192.168.1.23

- fabric:vp3

    启动: ssh登入,进入/home/a/fabric/vp3,运行sudo docker-compose -f server_vp3.yml up, 启动成功后直接关闭ssh终端
    停止: ssh登入,进入/home/a/fabric/vp3,运行sudo docker-compose -f server_vp3.yml stop

### Fabric重新部署-清理

#### 192.168.1.20

- fabric:vp0

    运行 sudo docker ps, 会显示vp0_vp0_1,dev-vp0-e...的两个容器实例
    运行 sudo docker stop vp0_vp0_1,dev-vp0-e...,停止容器实例
    运行 sudo docker rm vp0_vp0_1,dev-vp0-e...,删除容器实例
    运行 sudo docker images,会显示dev-vp0-ecd4a...的模板实例
    运行 sudo docker rmi dev-vp0-ecd4a...,删除模板实例

- 清理数据库

#### 192.168.1.21

- fabric:vp1

    运行 sudo docker ps, 会显示vp1_vp1_1,dev-vp1-ecd4ae...的两个容器实例
    运行 sudo docker stop vp1_vp1_1,dev-vp1-ecd4ae...,停止容器实例
    运行 sudo docker rm vp1_vp1_1,dev-vp1-ecd4ae...,删除容器实例
    运行 sudo docker images,会显示dev-vp1-ecd4...的模板实例
    运行 sudo docker rmi dev-vp1-ecd4...,删除模板实例

#### 192.168.1.22

- fabric:vp1

    运行 sudo docker ps, 会显示vp2_vp2_1,dev-vp2-ecd4ae...的两个容器实例
    运行 sudo docker stop vp2_vp2_1,dev-vp2-ecd4ae...,停止容器实例
    运行 sudo docker rm vp2_vp2_1,dev-vp2-ecd4ae...,删除容器实例
    运行 sudo docker images,会显示dev-vp2-ecd4...的模板实例
    运行 sudo docker rmi dev-vp2-ecd4...,删除模板实例

#### 192.168.1.23

- fabric:vp3

- fabric:vp1

    运行 sudo docker ps, 会显示vp3_vp3_1,dev-vp3-ecd4ae...的两个容器实例
    运行 sudo docker stop vp3_vp3_1,dev-vp3-ecd4ae...,停止容器实例
    运行 sudo docker rm vp3_vp3_1,dev-vp3-ecd4ae...,删除容器实例
    运行 sudo docker images,会显示dev-vp3-ecd4...的模板实例
    运行 sudo docker rmi dev-vp3-ecd4...,删除模板实例

### Fabric重新部署-启动

#### 192.168.1.20

- fabric:vp0

    启动: ssh登入,进入/home/a/fabric/vp0,运行sudo docker-compose -f server_vp0.yml up, 启动成功后直接关闭ssh终端
    停止: ssh登入,进入vp0,运行sudo docker-compose -f server_vp0.yml stop

#### 192.168.1.21

- fabric:vp1

    启动: ssh登入,进入/home/a/deploy/vp1,运行sudo docker-compose -f server_vp1.yml up, 启动成功后直接关闭ssh终端
    停止: ssh登入,进入/home/a/deploy/vp1,运行sudo docker-compose -f server_vp1.yml stop

#### 192.168.1.22

- fabric:vp2

    启动: ssh登入,进入/home/a/fabric/vp2,运行sudo docker-compose -f server_vp2.yml up, 启动成功后直接关闭ssh终端
    停止: ssh登入,进入/home/a/fabric/vp2,运行sudo docker-compose -f server_vp2.yml stop

#### 192.168.1.23

- fabric:vp3

    启动: ssh登入,进入/home/a/fabric/vp3,运行sudo docker-compose -f server_vp3.yml up, 启动成功后直接关闭ssh终端
    停止: ssh登入,进入/home/a/fabric/vp3,运行sudo docker-compose -f server_vp3.yml stop

### Fabric重新部署-CC初始化

#### deploy cc

    运行postman或其他工具,执行 post 192.168.31.20:7050/chaincode
    参数为:
        {
        "jsonrpc": "2.0",
        "method": "deploy",
        "params": {
        "type": 1,
        "chaincodeID":{
        "path": "github.com/CebEcloudTime/charitycc"
        },
        "ctorMsg": {
        "function":"deploy",
        "args":[]
        }
        },
        "id": 1
        }

    返回为ccname,
    查看节点高度,如果高度为2则发布成功,此过程包含同步文件及编译,时间会比较长(此时可多次发布,返回ccname不变,不影响业务)

#### 初始化数据

    下载github.com/CebCloudTime/tools/fabric/deploy目录

    在修改 /deploy目录里所有文件中包含chaincode name,把name改为上一步返回的ccname(修改29个文件),同CC发布的ccname应该相同,校验一致可以不修改,如registerBargain01.json中
    "name":"ecd4aed37b1c75ea2d83941c329c9631f2659f292e958d5c39955c1a6d24a8dc21c34faa77f0a6d5872f2a63d8c102c4d583fd69002f31b68bbc9d9800e2670f"

    修改 /deploy目录中deploy.sh, 把如http://192.168.31.20:7050,改为 http://192.168.1.20:7050,文件中修改29处

    执行 ./deploy.sh,初始化数据




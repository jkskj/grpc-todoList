# grpc-todoList
## 配置文件
### api-gateway/config/config.yml
```
server:
  domain: api-gateway
  port: :
  jwtSecret: 
  version: 1.0

user:
  address: 

task:
    address: 

mysql:
  driverName: mysql
  host: 
  port: 
  database: 
  username: 
  password: 
  charset: 

etcd:
  address: 

redis:
  address: 
  password:
  ```
### user/config/config.yml
```
server:
  domain: user
  version: 1.0
  jwtSecret: 
  grpcAddress: 

mysql:
  driverName: mysql
  host: 
  port: 
  database: 
  username: 
  password: 
  charset: 

etcd:
  address: 

redis:
  address: 
  password:
  ```
### task/config/config.yml
```
server:
  domain: task
  version: 1.0
  jwtSecret: 
  grpcAddress: 

mysql:
  driverName: mysql
  host: 
  port: 
  database: 
  username: 
  password: 
  charset: 

etcd:
  address: 

redis:
  address: 
  password:
  ```

# grpc-todoList
## 项目介绍
grpc-todoList是一个基于gRPC的微服务项目demo，
v1文件夹里为之前结构。
## 项目结构
```
grpc-todoList
├── .gitigrone
├── README.md
├── app
│  ├── gateway
│  │  ├── cmd
│  │  │  └── main.go
│  │  ├── internal
│  │  │  └── handler
│  │  │     ├── pkg.go
│  │  │     ├── task.go
│  │  │     └── user.go
│  │  ├── middleware
│  │  │  ├── cors.go
│  │  │  ├── init.go
│  │  │  ├── jwt.go
│  │  │  └── wrapper
│  │  ├── routes
│  │  │  └── router.go
│  │  └── rpc
│  │     ├── init.go
│  │     ├── task.go
│  │     └── user.go
│  ├── task
│  │  ├── cmd
│  │  │  └── main.go
│  │  └── internal
│  │     ├── repository
│  │     │  └── db
│  │     │     ├── dao
│  │     │     │  ├── db_init.go
│  │     │     │  ├── migration.go
│  │     │     │  └── task.go
│  │     │     └── model
│  │     │        └── task.go
│  │     └── service
│  │        └── task.go
│  └── user
│     ├── cmd
│     │  └── main.go
│     └── internal
│        ├── repository
│        │  └── db
│        │     ├── dao
│        │     │  ├── db_init.go
│        │     │  ├── migration.go
│        │     │  └── user.go
│        │     └── model
│        │        └── user.go
│        └── service
│           └── user.go
├── config
│  ├── config.go
│  └── config.yml
├── consts
│  ├── common.go
│  └── user.go
├── doc
│  └── grpc-todoList.json
├── idl
│  └── pb
│     ├── task.proto
│     ├── user.proto
│     ├── task
│     │  ├── task.pb.go
│     │  └── task_grpc.pb.go
│     └── user
│        ├── user.pb.go
│        └── user_grpc.pb.go
└── pkg
   ├── ctl
   │  ├── ctl.go
   │  └── user_info.go
   ├── discovery
   │  ├── instance.go
   │  ├── register.go
   │  └── resolver.go
   ├── e
   │  ├── code.go
   │  └── msg.go
   ├── res
   │  └── response.go
   └── util
      ├── jwt
      │  └── jwt.go
      └── shutdown
         └── gracefully_shutdown.go


```
## 配置文件
### v2
#### config/config.yml
```
server:
  port: :3000
  version: 1.0
  jwtSecret: 38324

mysql:
  driverName: mysql
  host: 127.0.0.1
  port: 3306
  database: grpc_todolist
  username: root
  password: 123456
  charset: utf8mb4

redis:
  user_name: default
  address: 127.0.0.1:6379
  password:

etcd:
  address: 127.0.0.1:2379

services:
  gateway:
    name: gateway
    loadBalance: true
    addr:
      - 127.0.0.1:10001
  user:
    name: user
    loadBalance: false
    addr:
      - 127.0.0.1:10002
  task:
    name: task
    loadBalance: false
    addr:
      - 127.0.0.1:10003

domain:
  user:
    name: user
  task:
    name: task
   ```
### v1
#### api-gateway/config/config.yml
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
#### user/config/config.yml
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
#### task/config/config.yml
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

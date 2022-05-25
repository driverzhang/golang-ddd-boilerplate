
项目结构：

```json
├── Dockerfile
├── Makefile
├── README.md
├── api
│   ├── proto
│   └── swagger
│       ├── docs.go
│       ├── swagger.json
│       └── swagger.yaml
├── cmd
│   └── server
│       └── main.go
├── config
│   └── conf.go
├── config.develop-test.json
├── config.json
├── docs
├── go.mod
├── go.sum
├── internal
│   ├── client
│   │   ├── common
│   │   │   └── role.go
│   │   ├── requestobject // 请求参数的结构体定义
│   │   │   └── role.go
│   │   └── viewobject  // 返回参数的结构体定义
│   │       └── role.go
│   ├── controller
│   │   ├── index.go
│   │   ├── limit.go // 对应领域路由
│   │   ├── role.go  // 对应领域路由
│   │   └── router.go // 公共的路由列表，方便定位
│   ├── domain // 领域层。如果不用 ddd 模型，在里面写 service 就行了
│   │   ├── aggr // 聚合根层
│   │   │   ├── aggr.go
│   │   │   ├── limit.go
│   │   │   └── role.go // 对应领域的聚合根，如果你同一领域小聚合根较多，建议这里直接建立一个文件夹：roleaggr 后缀带上aggr避免引入包重复时需要反复自定义
│   │   ├── entity // 实体层
│   │   │   ├── entity.go
│   │   │   └── role.go // 对应领域的实体，实体较多也可以同上建立文件夹：roleentity
│   │   ├── repo // 存储仓 实现抽象定义层
│   │   │   ├── limit.go
│   │   │   └── role.go
│   │   ├── service // 领域服务层， 跨领域聚合建议在该层操作，如果不用DDD模式，有可以再该层直接定义义务方法
│   │   │   └── role.go
│   │   └── vo // 值对象层
│   │       └── role.go
│   ├── infrastructure // 基础设施实现层
│   │   ├── cache // 缓存
│   │   │   └── cache.go
│   │   ├── database  // 数据库 定义model
│   │   │   └── role_model.go
│   │   ├── index.go
│   │   ├── role_repo_impl.go
│   │   └── rpc // 远程调用
│   │       └── rpc.go
│   └── pkg // internal 内部公共包，不会被外部项目引入
│       └── pkg.go
├── manifest.yaml
├── pre-stop.sh
├── proto.lock
└── yapi-import.json
```
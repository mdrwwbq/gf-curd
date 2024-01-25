# 如何使用
## 安装
```
go install github.com/mdrwwbq/gf-curd@latest
```
## 验证安装是否成功
```
gf-curd --help
```

## 使用命令生成 CRUD 代码

### 1.先生成gf框架代码
```
gf init gf-test-demo
initializing...
initialization done!
you can now run "cd gf-test-demo && gf run main.go" to start your journey, enjoy!
```
### 2.生成CRUD代码
```
gf-curd.exe -t menu
./output/api/v1/menu.go success
./output/i18n/zh-CN//menu.toml success
./output/internal/controller///menu.go success
./output/internal/logic/menu/menu.go success
./output/utility//resp.go success
./output/utility//language.go success
./output/utility//db.go success
./output/utility//consts.go success
```
### 3.把output下面的代码复制到项目中去

### 4.需要做如下的操作
#### 4.1 在`internal/cmd/cmd.go`中添加上述控制器
```
group.Bind(controller.Menu)
```
#### 4.2 在项目根据目录下执行
```
gf gen dao  # 生成menu数据库的model文件如果已经生成可忽略
gf gen service  # 生成menu的service文件
gf run main.go  # 运行项目
```



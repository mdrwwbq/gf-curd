# gf-curd
gf curd tools


#### 目前支持接口
- [x] add（新增）
- [x] edit（修改）
- [x] info（详情）
- [x] list（列表）
- [x] delete（物理删除单个/多个）
- [x] logicDelete（逻辑删除单个/多个）
- [x] enable（启用/禁用单个/多个）

### 如何安装
```
go install github.com/mdrwwbq/gf-curd@latest
```
### 怎么使用
[参见文档](https://github.com/mdrwwbq/how-to-user.md)
```
gf-curd --help
Usage of gfcurd.exe:
  -c    覆盖写入，如：-c true覆盖（默认），-c=false不覆盖 (default true)
  -m string
        模块名称，如：-m basic 非必填
  -o string
        输出路径，如：-o output (default "./output")
  -t string
        具体功能名，如：-t menu或-t userRole
```

### 一些参数约定
- `{{ModuleName}}`原始的模块名,如果未指定`{{ModuleName}}`，在`api`目录中将使用默认的`v1`目录
- `{{TooleName}}`原始的功能名
- `{{UpperToolsName}}`功能名大陀峰
- `{{UpperToolsName}}`功能名小陀峰


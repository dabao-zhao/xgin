# xgin

仅比 gin 多了代码生成、log、config 的功能

## 命令

```shell
# 创建 helloworld 项目
xgin new helloworld 
# 升级 xgin
xgin upgrade 
# 从 ddl 中创建 CURD 相关的方法并保存到 test 文件夹中
xgin curd ddl -s user.sql -d test 
# 从 ddl 中创建 CURD 相关的方法并保存到 test 文件夹中
xgin curd datasource -u="root:123456@tcp(127.0.0.1:3306)/test" -t="*" -d="test" 
# 创建新的 app，需要在项目根目录执行
xgin app hello
```
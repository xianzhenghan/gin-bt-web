## 关于

`gin-api-mono` 是基于 [go-gin-api](https://github.com/xinliangnote/go-gin-api) 的删减版（仅提供 API 接口，去掉了界面），便于大家快速上手，示例中仅提供了对 MySQL 表的增删改查。

运行文档：https://xiaobot.net/post/cfb713f6-75e0-432a-90cf-492b42afebd4

## 初始化 MySQL 信息

```sql
-- 1. MySQL 数据库连接信息，配置到 ./configs/fat_configs.toml 中 --
[mysql.read]
addr = '127.0.0.1:3306'
name = 'gin_api_mono'
pass = '123456789'
user = 'root'

[mysql.write]
addr = '127.0.0.1:3306'
name = 'gin_api_mono'
pass = '123456789'
user = 'root'

-- 2. 创建数据表 --
CREATE TABLE `admin` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `username` varchar(32) NOT NULL DEFAULT '' COMMENT '用户名',
  `mobile` varchar(20) NOT NULL DEFAULT '' COMMENT '手机号',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='管理员表';

-- 3. 初始化数据 --
INSERT INTO `admin` (`id`, `username`, `mobile`) VALUES
(1, '张三', '13888888888'),
(2, '李四', '13888888888'),
(3, '赵五', '13888888888');
```

## 启动

```
$ go run main.go -env fat  

// -env 表示设置哪个环境，主要是区分使用哪个配置文件，默认为 fat
// -env dev 表示为本地开发环境，使用的配置信息为：configs/dev_configs.toml
// -env fat 表示为测试环境，使用的配置信息为：configs/fat_configs.toml
// -env uat 表示为预上线环境，使用的配置信息为：configs/uat_configs.toml
// -env pro 表示为正式环境，使用的配置信息为：configs/pro_configs.toml
```

## 接口文档

- 接口文档：http://127.0.0.1:9999/swagger/index.html
- 心跳检测：http://127.0.0.1:9999/system/health



### 加载的配置文件内容为：

```yaml
# 机器人的昵称，可配置多个
nick_name:
  - "小风"
# bot超级管理员账户
admin: 0
# bot运行地址，若和gocq在同一台机器则只需要填写127.0.0.1即可，否则填写0.0.0.0，gocq配置你的公网地址
host: 127.0.0.1
# bot运行端口
port: 8080
# leafBot日志等级
log_level: info
# bot的管理员用户，类型为数组
super_user:
  - 123456
#    - 123456

# 配置命令的起始字符
command_start:
  - ""
  - /
# 是否启用playwright。热搜插件以及网页截图插件会依赖于该配置
enable_playwright: false
# 插件相关的一些配置
plugins:
  flash_group_id: 0
  # al_api的token，微博热搜依赖于该配置 获取地址：https://admin.alapi.cn/
  al_api_token: ""
  # 自动回复插件在群里是否需要at
  enable_reply_tome: false
  # 配置入群欢迎,类型为一个数组，每一个对象配置群号，后面配置欢迎消息
  welcome:
    -
      welcome: ""
      group_id: 123456
  #          -
  #              welcome: ""
  #              group_id: 123
  # 配置github_token
  github_token: ""
  # 配置自动同意好友的答案，若配置为空字符串则同意所有请求，可设置多个
  auto_pass_friend_request: []

```


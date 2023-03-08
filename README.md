# 特别提醒

将ChatGPT接入微信有封号风险，请勿用于商业用途，本仓库的内容谨慎尝试！！！！！！！！！！



# wechatbot
> 最近chatGPT异常火爆，本项目可以将个人微信化身GPT机器人，
> 项目基于[openwechat](https://github.com/eatmoreapple/openwechat) 开发。


### 目前实现了以下功能
 * 提问增加上下文，更接近官网效果 
 * 机器人群聊@回复
 * 机器人私聊回复
 * 好友添加自动通过
 
# 使用前提
> * 有openai账号，并且创建好api_key，注册事项可以参考[此文章](https://www.codeworld.top/?p=274) 。
> * 微信必须实名认证。

# 注意事项
> * 项目仅供娱乐，滥用可能有微信封禁的风险，请勿用于商业用途。
> * 请注意收发敏感信息，本项目不做信息过滤。

# 获取项目
git clone https://github.com/lihongcheng/chatgptforwechat.git

# 进入项目目录
cd chatgptforwechat

# 启动项目
go run main.go
````

# 配置文件说明
````
{
"api_key": "your api key",
"auto_pass": true
}

api_key：openai api_key
auto_pass:是否自动通过好友添加
````

# 使用示例
### 向机器人发送`我要问下一个问题`，清空会话信息。




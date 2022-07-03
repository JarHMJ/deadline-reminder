# deadline-reminder

倒计时提醒器

配置文件默认为`$HOME/.deadline-reminder.yaml`，也可以通过启动参数配置`--config`进行配置

配置文件参数：

```yaml
feishu:
  # 飞书机器人webhook
  baseUrl: "https://open.feishu.cn/open-apis/bot/v2/hook/xxxxxxxx"
# 倒计时日期
deadline: "2021-06-25"
# 倒计时名称
name: "纪念日"  
```

也支持环境变量的方式覆盖配置文件

|      环境名       |      描述      |
|:--------------:|:------------:|
| FEISHU.BASEURL | 飞书机器人webhook |
|    DEADLINE    |    倒计时日期     |
|      NAME      |    倒计时名称     |

# todo

- [ ] 支持自定义每天提醒时间
- [ ] 支持多个倒计时
- [ ] 支持多种🤖
- [ ] 支持添加一次性提醒，暴露api


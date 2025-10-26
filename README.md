# Pusher-Go 自动化的新闻拉取工具 

Pusher-Go 可以定时自动拉取当日的新闻和股市数据，并通过调用大模型服务对当日的信息进行总结，并发送至用户邮箱

## Configuration 
在```pusher-go/cli/config```中创建一个配置文件```config.local.yaml```, 如下: 

> 为了能正常发送邮件, 需要发送方开通 SMTP 功能
```yaml
model: 
  api_key: "sk-xxx"
  model_name: "deepseek-chat"
  endpoint: "https://api.deepseek.com"

g_news: 
  api_key: "xxx"
  keywords: ["agriculture", "technology", "Artificial Intelligence"]
  endpoint: "https://gnews.io/api/v4/top-headlines?" 
  categories: ["business", "technology", "world"]
  languages: ["en", "zh"] 
  countries: ["hk", "cn", "us"] 
  max_results: 10

email: 
  to: ["xxx@gmail.com"]
  subject: "Daily Report - "
  from: "xxx@gmail.com" 
  password: "xxx"

file: 
  file_name_news: "./data/news/news.txt"
  file_name_stock: "./data/stock/stock.txt" 
  file_name_model_response: "./data/model/model.txt"
```

## Installation 

```bash
git clone https://github.com/chriscco/pusher-go.git
cd pusher-go
```
授予部署脚本运行权限: 
```bash
chmod +x script.sh 
```
通过 crontab 或是其他方式将脚本加入系统的定时任务中, 如： 
```bash
# 在洛杉矶时间晚上 9 点运行脚本
0 19 * * * TZ="America/Los_Angeles" /home/ubuntu/pusher/script.sh >> /tmp/cron_debug.log 2>&1
```
> 如果需要修改脚本或代码, 确保不要出现相对路径

#!/bin/bash 

LOG_FILE="/home/ubuntu/pusher-go/pusher_automation.log"

# 设置时区为 PST（如果失败则继续执行）
if sudo timedatectl >/dev/null 2>&1; then
    sudo timedatectl set-timezone America/Los_Angeles
elif sudo systemsetup -gettimezone >/dev/null 2>&1; then
    sudo systemsetup -settimezone America/Los_Angeles
fi

rm "$LOG_FILE"

# 执行任务日志
echo "Daily task executed at $(date)" >> "$LOG_FILE"

chmod +x /home/ubuntu/pusher-go/script.sh 
/home/ubuntu/pusher-go/script.sh

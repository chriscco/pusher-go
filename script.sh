#!/bin/bash

cd /home/ubuntu/pusher-go/cli 
/usr/local/go/bin/go build -o pusher /home/ubuntu/pusher-go/cli/cmd/main 

# get news 
/home/ubuntu/pusher-go/cli/pusher news 

# call mode 
/home/ubuntu/pusher-go/cli/pusher model 

# send email 
/home/ubuntu/pusher-go/cli/pusher email 

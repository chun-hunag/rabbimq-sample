#!/bin/bash
# 在背景啟動 rabbitmq service (erlang node and rabbitmq application , detached mean in background)
rabbitmq-server -detached
sleep 5s

#停止 rabbitmq application
rabbitmqctl stop_app

# 加入 rabbimtq cluster
rabbitmqctl join_cluster ${MASTER_NODE_HOST_NAME}

# 啟動 rabbitmq application 來 sync 資料
rabbitmqctl start_app

# 停止 rabbitmq service (errlang node and rabbitmq application )
rabbitmqctl stop
sleep 5s

# 將 rabbitmq service 啟動在前景
rabbitmq-server
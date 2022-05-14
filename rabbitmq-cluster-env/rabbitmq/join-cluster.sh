#!/bin/bash
# start rabbitmq service in backgorund (erlang node and rabbitmq application , detached mean in background)
rabbitmq-server -detached

sleep 15s

# Stops the RabbitMQ application, leaving the runtime (Erlang VM) running
rabbitmqctl stop_app

# Instructs the node to become a member of the cluster that the specified node is in
rabbitmqctl join_cluster ${MASTER_NODE_HOST_NAME}

# start rabbitmq application 來 sync 資料
rabbitmqctl start_app
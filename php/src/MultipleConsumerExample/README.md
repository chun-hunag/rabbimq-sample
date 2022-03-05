

# Demo Step

1. open two terminal 
2. execute command below each one
   
   `php QueueConsumer.php {QueueName}`

3. check rabbitmq management GUI, you should see the queue have two consumer
4. publisher messages
    
   `php QueuesPublisher.php {QueueName} 1 2 3 4 5 6 7 8`
5. observer Consumer's log
    * example:
        * consumer - 1:
            ```
            Message received =  1 
            Message received =  3
            Message received =  5
            Message received =  7
            ```
        * consumer - 2:
            ```
            Message received =  2 
            Message received =  4
            Message received =  6
            Message received =  8
            ```
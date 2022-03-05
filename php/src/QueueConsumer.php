<?php

require_once '../vendor/autoload.php';

use PhpAmqpLib\Connection\AMQPStreamConnection;


/**
 * usage:
 * php QueueConsumer.php {QueueName}
 */

$queueName = getQueueNameArgv($argv);

$connection = new AMQPStreamConnection('localhost', 5672, 'guest', 'guest');

$channel = $connection->channel();


$callback = function ($msg) {
    echo "Message received =  $msg->body \n";
};

$channel->basic_consume($queueName, '', false, true, false, false, $callback);

while (true) {
    $channel->wait();
}

$channel->close;

$connection->close();


function getQueueNameArgv(array $argv): string
{
    if (isset($argv[1])) {
        return $argv[1];
    }
    throw new InvalidArgumentException("Missing Command Argument QueueName argv[0]");
}
<?php

require_once '../vendor/autoload.php';

use PhpAmqpLib\Connection\AMQPStreamConnection;

/**
 * usage:
 * php QueueDeclarer.php {QueueName}
 */

$queueName = getQueueNameArgv($argv);

$connection = new AMQPStreamConnection('localhost', 5672, 'guest', 'guest');
$channel = $connection->channel();

$channel->queue_declare($queueName, false, true, false, false);

echo "queue: $queueName was declared.";

$channel->close();
$connection->close();


function getQueueNameArgv(array $argv): string
{
    if (isset($argv[1])) {
        return $argv[1];
    }
    throw new InvalidArgumentException("Missing Command Argument ExchangeName argv[0]");
}
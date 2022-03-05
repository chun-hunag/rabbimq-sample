<?php

require_once '../vendor/autoload.php';

use PhpAmqpLib\Connection\AMQPStreamConnection;
use PhpAmqpLib\Message\AMQPMessage;

/**
 * usage:
 * php QueuePublisher.php {message} {QueueName}
 */

$messageStr = getMessageArgv($argv);
$queueName = getQueueNameArgv($argv);

$connection = new AMQPStreamConnection('localhost', 5672, 'guest', 'guest');

$channel = $connection->channel();

$msg = new AMQPMessage($messageStr);
$channel->basic_publish($msg);
$channel->basic_publish($msg, '', $queueName);
$channel->close();
$connection->close();

echo "Message '$messageStr' has been publish to Queue: '$queueName'";

function getMessageArgv(array $argv): string
{
    if (isset($argv[1])) {
        return $argv[1];
    }
    throw new InvalidArgumentException("Missing Command Argument Message argv[0]");
}

function getQueueNameArgv(array $argv): string
{
    if (isset($argv[2])) {
        return $argv[2];
    }
    throw new InvalidArgumentException("Missing Command Argument QueueName argv[1]");
}
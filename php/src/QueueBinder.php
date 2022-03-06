<?php

require_once '../vendor/autoload.php';

use PhpAmqpLib\Connection\AMQPStreamConnection;

const DIRECT = 'direct';
const FANOUT = 'fanout';
const HEADERS = 'headers';
const TOPIC = 'topic';

/**
 * usage:
 * php QueueBinder.php {QueueName} {ExchangeName} {RoutingKey}
 */

$queueName = getQueueNameArgv($argv);
$exchangeName = getExchangeNameArgv($argv);
$routingKey = getRoutingKeyArgv($argv);

$connection = new AMQPStreamConnection('localhost', 5672, 'guest', 'guest');
$channel = $connection->channel();

$channel->queue_bind($queueName, $exchangeName, $routingKey);
echo "queue: $queueName binding to exchange: $exchangeName with routing key: $routingKey";

$channel->close();
$connection->close();


function getQueueNameArgv(array $argv): string
{
    if (isset($argv[1])) {
        return $argv[1];
    }
    throw new InvalidArgumentException("Missing Command Argument QueueName argv[0]");
}

function getExchangeNameArgv(array $argv): string
{
    if (isset($argv[2])) {
        return $argv[2];
    }
    throw new InvalidArgumentException("Missing Command Argument ExchangeName argv[1]");
}

function getRoutingKeyArgv(array $argv): string
{
    if (isset($argv[3])) {
        return $argv[3];
    }
    throw new InvalidArgumentException("Missing Command Argument RoutingKey argv[2]");
}


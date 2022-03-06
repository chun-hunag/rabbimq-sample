<?php

require_once '../vendor/autoload.php';

use PhpAmqpLib\Connection\AMQPStreamConnection;
use PhpAmqpLib\Message\AMQPMessage;

/**
 * usage:
 * php ExchangePublisher.php {message} {ExchangeName} {RoutingKey}
 */

$messageStr = getMessageArgv($argv);
$exchangeName = getExchangeNameArgv($argv);
$routingKey = getRoutingKeyArgv($argv);

$connection = new AMQPStreamConnection('localhost', 5672, 'guest', 'guest');

$channel = $connection->channel();

$msg = new AMQPMessage($messageStr);
$channel->basic_publish($msg);
$channel->basic_publish($msg, $exchangeName, $routingKey);
$channel->close();
$connection->close();

echo "Message '$messageStr' has been publish to Exchange: '$exchangeName' with routing key: $routingKey";

function getMessageArgv(array $argv): string
{
    if (isset($argv[1])) {
        return $argv[1];
    }
    throw new InvalidArgumentException("Missing Command Argument Message argv[0]");
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
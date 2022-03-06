<?php

require_once '../../vendor/autoload.php';

use PhpAmqpLib\Connection\AMQPStreamConnection;
use PhpAmqpLib\Message\AMQPMessage;
use PhpAmqpLib\Wire\AMQPTable;

/**
 * usage:
 * php QueuePublisher.php {Message} {ExchangeName}
 */

$message = getMessageArgv($argv);
$exchangeName = getExchangeNameArgv($argv);
$routingKey = getRoutingKeyArgv($argv);

$connection = new AMQPStreamConnection('localhost', 5672, 'guest', 'guest');

$channel = $connection->channel();

$headerValues = [
    'item-1' => 'mobile',
    'item-2' => 'television'
];

$header = new AMQPTable($headerValues);
$ampqMessage = new AMQPMessage($message);
$ampqMessage->set('application_headers', $header);
$channel->basic_publish($ampqMessage, $exchangeName, $routingKey);

$channel->close();
$connection->close();

echo "message: $message with header has been publish to exchange: $exchangeName with routingKey: $routingKey";

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
    return '';
}
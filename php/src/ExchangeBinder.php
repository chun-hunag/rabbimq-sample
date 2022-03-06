<?php

require_once '../vendor/autoload.php';

use PhpAmqpLib\Connection\AMQPStreamConnection;

const DIRECT = 'direct';
const FANOUT = 'fanout';
const HEADERS = 'headers';
const TOPIC = 'topic';

/**
 * usage:
 * php ExchangeBinder.php {SourceExchangeName} {DestinationExchangeName} {RoutingKey}
 */

$sourceExchangeName = getSourceExchangeNameArgv($argv);
$destinationExchangeName = getDestinationExchangeNameArgv($argv);
$routingKey = getRoutingKeyArgv($argv);

$connection = new AMQPStreamConnection('localhost', 5672, 'guest', 'guest');
$channel = $connection->channel();

$channel->exchange_bind($destinationExchangeName, $sourceExchangeName, $routingKey);
echo "exchange: $sourceExchangeName binding to exchange: $destinationExchangeName with routing key: $routingKey";

$channel->close();
$connection->close();


function getSourceExchangeNameArgv(array $argv): string
{
    if (isset($argv[1])) {
        return $argv[1];
    }
    throw new InvalidArgumentException("Missing Command Argument SourceExchangeName argv[0]");
}

function getDestinationExchangeNameArgv(array $argv): string
{
    if (isset($argv[2])) {
        return $argv[2];
    }
    throw new InvalidArgumentException("Missing Command Argument DestinationExchangeName argv[1]");
}

function getRoutingKeyArgv(array $argv): string
{
    if (isset($argv[3])) {
        return $argv[3];
    }
    throw new InvalidArgumentException("Missing Command Argument RoutingKey argv[2]");
}
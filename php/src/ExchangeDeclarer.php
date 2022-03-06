<?php

require_once '../vendor/autoload.php';

use PhpAmqpLib\Connection\AMQPStreamConnection;

const DIRECT = 'direct';
const FANOUT = 'fanout';
const HEADERS = 'headers';
const TOPIC = 'topic';

/**
 *  usage:
 *  php ExchangeDeclarer.php {ExchangeName} {ExchangeType}
 */

$exchangeName = getExchangeNameArgv($argv);
$exchangeType = getExchangeTypeArgv($argv);

$connection = new AMQPStreamConnection('localhost', 5672, 'guest', 'guest');
$channel = $connection->channel();

$channel->exchange_declare($exchangeName, $exchangeType, false, false, false);
echo "Exchange name: $exchangeName , type: $exchangeType was declared.";

$channel->close();
$connection->close();


function getExchangeNameArgv(array $argv): string
{
    if (isset($argv[1])) {
        return $argv[1];
    }
    throw new InvalidArgumentException("Missing Command Argument ExchangeName argv[0]");
}

function getExchangeTypeArgv(array $argv): string
{
    if (!isset($argv[2])) {
        throw new InvalidArgumentException("Missing Command Argument ExchangeType argv[1]");
    }

    if (!in_array($argv[2], [DIRECT, FANOUT, HEADERS, TOPIC])) {
        throw new InvalidArgumentException("Wrong ExchangeType $argv[2]");
    }

    return $argv[2];
}
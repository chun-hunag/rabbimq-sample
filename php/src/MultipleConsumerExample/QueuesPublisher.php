<?php

require_once '../../vendor/autoload.php';

use PhpAmqpLib\Connection\AMQPStreamConnection;
use PhpAmqpLib\Message\AMQPMessage;

/**
 * usage:
 * php QueuePublisher.php {QueueName} {message1} {message2} ... etc.
 */

$queueName = getQueueNameArgv($argv);
$messages = getMessagesArgv($argv);


$connection = new AMQPStreamConnection('localhost', 5672, 'guest', 'guest');

$channel = $connection->channel();

foreach ($messages as $message) {
    $AmqpMessage = new AMQPMessage($message);
    $channel->basic_publish($AmqpMessage);
    $channel->basic_publish($AmqpMessage, '', $queueName);
    echo "Message '$message' has been publish to Queue: '$queueName' \n";
}

$channel->close();
$connection->close();


function getQueueNameArgv(array $argv): string
{
    if (isset($argv[1])) {
        return $argv[1];
    }
    throw new InvalidArgumentException("Missing Command Argument QueueName");
}

function getMessagesArgv(array $argv): array
{
    $messages = [];
    for ($i = 1; $i < count($argv); $i++) {
        if (isset($argv[$i])) {
            $messages[] = $argv[$i];
        }
    }

    if (empty($messages)) {
        throw new InvalidArgumentException("Missing Command Argument Message");
    }

    return $messages;
}
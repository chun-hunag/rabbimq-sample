<?php

require_once '../../vendor/autoload.php';

use PhpAmqpLib\Connection\AMQPStreamConnection;
use PhpAmqpLib\Message\AMQPMessage;

/**
 * usage:
 * php QueuePublisher.php {QueueName}
 */

$queueName = getQueueNameArgv($argv);

$connection = new AMQPStreamConnection('localhost', 5672, 'guest', 'guest');

$channel = $connection->channel();

$message = new \stdClass();
$message->to = 'to_address';
$message->from = 'from_address';
$message->subject = 'Subject_of_mail';

$msg = new AMQPMessage(json_encode($message));
$channel->basic_publish($msg, '', $queueName);
$channel->close();
$connection->close();

echo "JSON message has been publish to Queue: '$queueName'";


function getQueueNameArgv(array $argv): string
{
    if (isset($argv[1])) {
        return $argv[1];
    }
    throw new InvalidArgumentException("Missing Command Argument QueueName argv[0]");
}
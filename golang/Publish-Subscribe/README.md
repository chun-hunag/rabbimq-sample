
# Demo

## subscribe
* execute command below on two different terminal:
  * `go run subscriber.go`
  * message you will see after execute:
    ```
        2022/03/19 15:40:33  [*] Waiting for messages. To exit press CTRL+C
    ```
## publish
* execute command below on terminal:
  * `go run publisher.go {message you want to send}
  * message you will see after execute:
    ```azure
        2022/03/19 15:41:32  [x] Sent {message you want to send}
    ```
  * you should see terminal you execute subscriber command show message like below:
    * every message you send should be printed on every terminal you have executed subscriber command
    ```
    2022/03/19 15:41:32 Received a message: 1
    2022/03/19 15:41:32 Done
    2022/03/19 15:41:34 Received a message: 2
    2022/03/19 15:41:34 Done
    2022/03/19 15:41:36 Received a message: 3
    2022/03/19 15:41:36 Done
    ```
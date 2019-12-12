
client <---ws---> middleman <---ws---> server

THe goal is:
* **client (c)**: must be a javascript client that will send/receive messages from the middleman
* **middleman (m)**: must be a Go application that will have a web socket client and a web socket server. It will receive messages from (c) and redirect those to (s). At the same time, it will receive messages from (s) and redirect those to (c) 
* **server (s)**: must be a Go websocket server that will send/receive messages from the middleman
 

gprc-server-stream pipeline:
Client: Connect and send UserId = 123
   |
   v
Server: Subscribe to Redis channel "notifications/123"
   |
   v
Redis Publisher: Publishes "New message <timestamp>" every 5 seconds
   |
   v
Server: Receives Redis messages -> streams them to the client
   |
   v
Client: Receives and logs/display notifications

## Elements of demo
The demo consists of 4 elements:
1) Message sender (go app)
2) Message receiver (go app)
3) RabbitMQ
4) RabbitMQ web UI (http://localhost:15672)

## Scenario
Message sender sends simple message `Hello world` to message receiver.

## Running demo
```bash
docker compose up --build
```

## Other
Additional information can be observed in RabbitMQ web UI. Also, there are different actions that
can be performed in UI e.g. sending new message to a queue.

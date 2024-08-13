# Gque - Golang Queue

Gque is a message queue protocol written in Golang that internally uses a NoSQL database to store broadcast queues and messages. Gque is designed to handle high-throughput message processing with a focus on scalability and reliability.

## Features

- **Message Queuing**: Efficiently queues incoming messages for processing.
- **NoSQL Integration**: Stores and retrieves messages and queues from a NoSQL database.
- **Channel-Based Processing**: Uses Golang channels for asynchronous message handling.
- **Event-Driven Architecture**: Supports the creation of events that can be broadcast to multiple consumers.

## Workflow Overview

1. **Receiving Messages**: 
   - Requests from Gque clients are received by the handler.
   - A new event is created and pushed into the **incoming message channel**.

2. **Processing Incoming Messages**:
   - The **incoming message worker** consumes messages from the incoming message channel.
   - The message data is stored in a NoSQL database.
   - The message is then passed to the **message producer channel**.

3. **Producing Messages**:
   - The **message producer worker** consumes the message from the producer channel.
   - Necessary events are generated.
   - Events are pushed into the appropriate consumer channels.

## Installation

1. Run the following command to pull gque image:

```bash
go run cmd/main.go
```

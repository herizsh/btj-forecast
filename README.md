### Assignment: Simple gRPC-based Push Notification Application

#### Overview
Your task is to develop a simplified push notification system using gRPC. The system will consist of two components:

1. **Server (Go)**: This server will handle sending push notifications about various topics (e.g., weather updates, daily fun facts).
2. **Client (Python)**: The client will request notifications for a specific topic from the server.

#### Requirements

##### Server (Go)
- Implement a gRPC server in Go.
- The server should have a function to send push notifications about different topics.
- Implement at least two different notification types (e.g., weather, fun facts).

##### Client (Python)
- Implement a gRPC client in Python.
- The client should call the server function to request notifications for a specific topic.
- Display the received notifications to the user.

#### Detailed Specifications

1. **gRPC Service Definition**
   - Define a gRPC service with appropriate RPC methods for requesting notifications.
   - Use Protocol Buffers to define the request and response message formats.

2. **Server Implementation (Go)**
   - Implement the service logic to send notifications about the requested topics.
   - Ensure that the server can handle multiple requests simultaneously.

3. **Client Implementation (Python)**
   - Implement functionality to connect to the gRPC server.
   - Create a simple CLI where users can input the topic they want to receive notifications about.
   - Handle and display the notifications received from the server.

#### Deliverables
- Source code for both the Go server and the Python client.
- A README file documenting how to set up and run the server and client.
🏗️ Elevator System: The 5-Day Learning Lab
This repository documents my intensive research and development phase. Before executing the final elevator system, I am spending 5 days mastering the underlying technologies: Go Concurrency, Protobuf Serialization, and C++ Performance Integration.

🧠 The Learning Journey
Day 1: Go Concurrency & Event Handling (Completed ✅)
The Goal: Master the "Handshake." Understand how Go channels facilitate communication without shared memory.

The Experiment: Built a simulation that functions like a chat system but is applied to elevator event logic. It handles multiple incoming streams (User vs. Admin) simultaneously.

The Tech: Used unbuffered channels for synchronized communication and the select statement to multiplex between different event types.

Takeaway: Learned how to prevent system deadlocks when multiple events trigger at the exact same time.

Day 2: Advanced Data Flow (Next Up 🔜)
The Goal: Transition from basic handshakes to high-throughput systems.

The Plan: Experiment with Buffered Channels and Worker Pools.

Challenge: Simulating how a system handles a "rush hour" of 50+ simultaneous requests without losing data.

Day 3 & 4: The gRPC Bridge (Contract First)
The Goal: Getting two different languages to speak the same "dialect."

The Plan: Learn Protobuf (Protocol Buffers) basics.

The Tech: Practice serializing and deserializing messages to ensure the Go "Brain" can send structured data to other services.

Day 5+: C++ & Performance Integration
The Goal: High-speed execution.

The Plan: Slowly introduce C++ and explore ONNX (Open Neural Network Exchange).

Why: Learning how to hand off heavy lifting, like path optimization or sensor-pattern recognition, to a low-level language.

📂 Project Structure
This is a monorepo designed to keep my learning organized:

/go-service: Go concurrency labs and event-driven simulations.

/proto: (Upcoming) Serialization contracts and message definitions.

/cpp-service: (Upcoming) Performance-critical logic experiments.

/docs: My research notes and architectural sketches.

🏃 How to Run the Day 1 Lab
Navigate to the service:

PowerShell
cd go-service
Execute the simulation:

PowerShell
go run .

✍️ Reflection
"You don't build a skyscraper by just stacking bricks; you start by testing the strength of the steel. This sprint is about mastering the 'steel' before I build the 'elevator'."

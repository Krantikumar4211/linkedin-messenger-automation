# LinkedIn Messenger Automation (Golang)

## Overview
This project is a LinkedIn Messenger automation system built using Golang.
The objective of this assignment is to demonstrate clean software architecture,
problem-solving skills, automation logic, and human-like interaction patterns.

This is a technical demonstration project and not a production-ready automation tool.

## Project Structure

```
linkedin-messenger-automation/
│
├── cmd/
│ └── main.go
│
├── internal/
│ ├── auth/
│ │ └── auth.go
│ ├── messaging/
│ │ └── messenger.go
│ ├── automation/
│ │ ├── human_behavior.go
│ │ └── rate_limiter.go
│ ├── storage/
│ │ └── storage.go
│ ├── config/
│ │ └── config.go
│ └── logger/
│ └── logger.go
│
├── data/
│ └── state.json
│
├── .env.example
├── go.mod
└── README.md
```

## Features Implemented

- Modular Golang project structure
- Simulated LinkedIn authentication flow
- Human-like randomized delays
- **Rate limiting to control message frequency**
- Message typing simulation
- Structured logging with timestamps
- State persistence using JSON
- Environment-based configuration using `.env`
- Graceful error handling

---

## Rate Limiting (Important Feature)

To prevent spam-like behavior and mimic human interaction, the system includes a
**rate limiter** that restricts how frequently messages can be sent.

### How it works:
- A minimum time interval is enforced between consecutive messages
- If a message is sent too soon, the system automatically waits
- This ensures natural interaction patterns and avoids aggressive automation

### Example:
If the rate limiter is set to **1 message every 10 seconds**, the program will
wait before sending the next message if the previous one was sent recently.

---

## Configuration

Create a `.env` file using `.env.example`:

```env
LINKEDIN_EMAIL=your_email@example.com
LINKEDIN_PASSWORD=your_password
LOG_LEVEL=INFO
STATE_FILE=./data/state.json


Sensitive data is not committed to version control.

How to Run
1. Install Go

Ensure Go 1.21 or higher is installed.

go version

2. Run the Application

From the project root directory:

go run ./cmd/main.go

State Persistence

The application stores message history in a JSON file located at:

data/state.json


This allows:

Tracking previously sent messages

Preventing duplicate sends

Safe recovery after application restart

Demo Video

A demo video is included showing:

Project structure

Configuration setup

Program execution

Logging output



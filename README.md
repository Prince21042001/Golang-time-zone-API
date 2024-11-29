# Toronto Time API

## Summary

The **Toronto Time API** is a Go-based application that provides the current time in Toronto in JSON format. Each API request is logged into a PostgreSQL database for future reference. This application demonstrates API development, PostgreSQL integration, and time zone handling in a modular and efficient manner.

---

## Features

- **Get Current Time in Toronto**: The `/current-time` endpoint returns the current Toronto time in JSON format.
- **Database Logging**: Logs each API request timestamp into a PostgreSQL database.
- **Time Zone Handling**: Automatically adjusts to Toronto's timezone using Go's `time` package.

---

## Setup and Installation

Follow these instructions to set up and run the application.

### Prerequisites

- Go installed (v1.20+ recommended)
- PostgreSQL installed and running
- `pgAdmin` or a similar database management tool for PostgreSQL (optional)
- Git installed

### Instructions for Setting Up and Running the Application

1. **Clone the Repository**
   ```bash
   git clone https://github.com/your-username/toronto-time-api.git
   cd toronto-time-api

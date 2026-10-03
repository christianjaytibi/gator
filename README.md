# Gator

Gator is a command-line interface (CLI) RSS feed aggregator written in Go. It allows users to register, manage user accounts, follow multiple RSS feeds, and aggregate/fetch posts directly from the command line, backed by a PostgreSQL database.

## Prerequisites

Before running or installing Gator, ensure you have the following installed on your system:
- **Go** (version 1.27 or higher recommended)
- **PostgreSQL** (running locally)

## Installation

You can install the Gator CLI globally using the `go install` command:

```bash
go install github.com/christianjaytibi/gator@latest
```

*(Make sure your Go bin directory, typically `~/go/bin`, is added to your system's `PATH` so you can run `gator` from anywhere.)*

## Configuration

Gator requires a configuration file to connect to your PostgreSQL database. 

1. Create a file named `.gatorconfig.json` in your home directory (`~/.gatorconfig.json`).
2. Add your PostgreSQL database connection string in JSON format:

```json
{
  "db_url": "postgres://username:password@localhost:5432/gator?sslmode=disable"
}
```

## Usage & Commands

Once configured, you can run Gator using the compiled binary (or via `go run .` during development). Here are some of the core commands available:

* **Register a new user:**
  ```bash
  gator register <name>
  ```
* **Log in as an existing user:**
  ```bash
  gator login <name>
  ```
* **Add a new RSS feed to follow:**
  ```bash
  gator addfeed <name> <url>
  ```
* **List all feeds in the database:**
  ```bash
  gator feeds
  ```
* **Follow an existing feed:**
  ```bash
  gator follow <url>
  ```
* **View feeds followed by the current user:**
  ```bash
  gator following
  ```
* **Unfollow a feed:**
  ```bash
  gator unfollow <url>
  ```
* **Start the aggregator to fetch posts periodically:**
  ```bash
  gator agg <time_between_reqs>
  (Example: gator agg 1m)
  ```
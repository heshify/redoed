# Redoed

![Work in Progress](https://img.shields.io/badge/status-WIP-orange) ![Go Version](https://img.shields.io/badge/Go-1.24-blue) ![License](https://img.shields.io/badge/License-MIT-green)
[![Go Report Card](https://goreportcard.com/badge/github.com/heshify/redoed?1)](https://goreportcard.com/report/github.com/heshify/redoed)


> 🚧 **Work in Progress**: This project is in development stage. Follow the progress at heshify.github.io

Redoed is a real-time collaborative Markdown editor built in Go. Multiple users can simultaneously edit a document and see changes in real time. This project serves as an educational exploration of WebSockets, Markdown processing, and real-time synchronization.

## Getting Started

This project requires Go +1.22 and PostgreSQL database.

### Installation

Clone the repository :

```sh
git clone https://github.com/heshify/redoed.git
cd redoed
```

To compile and run the server, run the following command :

```sh
go run main.go
```

The server will start on port 8080.

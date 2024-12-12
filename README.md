## Cli-Chat-Server

A lightweight command-line chat client written in Go. This project provides a simple and efficient way to communicate in real-time via a terminal-based interface.

![image](https://github.com/user-attachments/assets/7ca6bc1e-f093-4cdc-912f-8295e08eb581)


## Features
- Real-time chat functionality.
- Lightweight and fast.
- Simple and intuitive command-line interface.
- Written in Go, ensuring high performance and portability.

## How to run

### Build Binary file
#### Prerequisites
before install or build this project you need to download following package
- [Go Lang](https://go.dev/)

```sh
git clone https://github.com/ThawThuHan/cli-chat-server.git
cd cli-chat-server
go build -o cli-chat-server main.go read_config.go
./cli-chat-server
```

### Run from Pre-Build binary file
#### Windows
Download pre-build binary file [Cli-Chat-Server](https://github.com/ThawThuHan/cli-chat-server/releases/download/v1.0/cli-chat-server.exe) from release and run
```sh
./cli-chat-server.exe
```
OR
```sh
curl -o cli-chat-server.exe https://github.com/ThawThuHan/cli-chat-server/releases/download/v1.0/cli-chat-server.exe
./cli-chat-server.exe
```

## To Change the Server Config
config.yaml file is under AppData/Local/Cli-Chat/config.yaml in windows and under ~/.Cli-Chat/config.yml in linux.
```yaml
name: cli-chat
ip: 0.0.0.0
port: 5000
```

## How to connect the server using Cli-Chat-Client
Explore the following link.
https://github.com/ThawThuHan/cli-chat-client

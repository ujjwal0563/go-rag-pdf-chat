# 🚀 GoRAG - Chat with PDF using Go

> A production-ready Retrieval-Augmented Generation (RAG) application built with Go that allows users to upload PDF documents and ask natural language questions based on their content.

![Go](https://img.shields.io/badge/Go-1.25+-00ADD8?style=for-the-badge&logo=go)
![Gin](https://img.shields.io/badge/Gin-Web_Framework-008ECF?style=for-the-badge)
![Qdrant](https://img.shields.io/badge/Qdrant-Vector_DB-DC244C?style=for-the-badge)
![Docker](https://img.shields.io/badge/Docker-Containerized-2496ED?style=for-the-badge&logo=docker)
![License](https://img.shields.io/badge/License-MIT-green?style=for-the-badge)

---

# 📖 Project Overview

GoRAG is a backend application that enables users to upload PDF documents and interact with them using AI.

Instead of sending the entire document to an LLM, the application retrieves only the most relevant information using Retrieval-Augmented Generation (RAG), making responses more accurate, scalable, and cost-effective.

---

# 🎯 Project Goals

- Learn RAG from scratch
- Build a production-style backend in Go
- Understand Vector Databases
- Learn Embeddings
- Implement Similarity Search
- Integrate Gemini API
- Follow Clean Architecture
- Practice Git & GitHub Workflow
- Dockerize the application
- Deploy a production-ready backend

---

# 🏗️ Architecture

```
                    User
                      │
                      ▼
              Go REST API (Gin)
                      │
          ┌───────────┴───────────┐
          ▼                       ▼
    Upload PDF              Ask Question
          │                       │
          ▼                       ▼
     PDF Processing        Query Embedding
          │                       │
          ▼                       ▼
      Text Chunking      Vector Similarity Search
          │                       │
          └───────────┬───────────┘
                      ▼
              Qdrant Vector Database
                      │
                      ▼
               Prompt Builder
                      │
                      ▼
                 Gemini API
                      │
                      ▼
                 AI Response
```

---

# ✨ Features

## Current

- Project Setup
- Professional Git Workflow
- Clean Project Structure

## Planned

- PDF Upload API
- PDF Text Extraction
- Smart Text Chunking
- Embedding Generation
- Qdrant Integration
- Similarity Search
- Gemini API Integration
- Chat API
- Docker Support
- Multi-document Support
- Conversation Memory
- Authentication
- Deployment

---

# 🛠️ Tech Stack

| Technology | Purpose |
|------------|----------|
| Go | Backend Development |
| Gin | REST API Framework |
| Qdrant | Vector Database |
| Gemini API | Large Language Model |
| Docker | Containerization |
| Git | Version Control |
| GitHub | Repository Hosting |

---

# 📂 Project Structure

```
go-rag-pdf-chat/

├── cmd/
│   └── server/
│       └── main.go
│
├── internal/
│   ├── config/
│   ├── handlers/
│   ├── services/
│   ├── pdf/
│   ├── chunker/
│   ├── embedding/
│   ├── vector/
│   ├── llm/
│   ├── middleware/
│   ├── models/
│   └── utils/
│
├── docs/
│
├── uploads/
│
├── scripts/
│
├── .env.example
├── Dockerfile
├── docker-compose.yml
├── go.mod
├── go.sum
├── README.md
└── LICENSE
```

---

# 🔄 Development Workflow

```
main
 │
 ├── feature/project-setup
 ├── feature/gin-server
 ├── feature/pdf-upload
 ├── feature/pdf-parser
 ├── feature/chunking
 ├── feature/embedding
 ├── feature/qdrant
 ├── feature/retrieval
 ├── feature/chat-api
 ├── feature/docker
 └── feature/documentation
```

Each feature is developed in its own branch and merged into `main` through a Pull Request.

---

# 📚 Learning Objectives

This project is built to understand:

- Go Backend Development
- REST APIs
- RAG Architecture
- Vector Databases
- Embeddings
- Prompt Engineering
- Clean Architecture
- Docker
- Git & GitHub Workflow

---

# 📅 Roadmap

- [x] Project Planning
- [x] Development Environment Setup
- [ ] Git Repository Setup
- [ ] Go Project Initialization
- [ ] Gin Server
- [ ] PDF Upload
- [ ] PDF Parser
- [ ] Chunking
- [ ] Embeddings
- [ ] Qdrant Integration
- [ ] Retrieval Pipeline
- [ ] Prompt Engineering
- [ ] Gemini Integration
- [ ] Chat API
- [ ] Docker Support
- [ ] Deployment

---

# 🚀 Getting Started

Coming Soon...

---

# 📖 Documentation

Detailed documentation will be available inside the `docs/` folder.

---

# 🤝 Contributing

Contributions, suggestions, and feedback are welcome.

Feel free to fork the repository and submit a Pull Request.

---

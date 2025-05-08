# 💬 Real-Time Chat Web

A modern real-time chat application built with **Go**, **Gin**, **WebSocket**, and **Nuxt 3**. Designed for scalability, performance, and seamless user interaction.

---

## 🚀 Tech Stack

### 🔧 Backend — Go + Gin + Sockie
- **Gin**: Lightweight, high-performance HTTP framework.
- **Sockie**: Efficient WebSocket library for real-time communication.

#### 📦 Features
- RESTful API for user, auth, and chat management
- WebSocket-powered real-time messaging
- Scalable performance with Go’s concurrency model

### 🌐 Frontend — Nuxt 3 (Vue 3)
- **Nuxt 3**: Modern framework supporting SSR, SSG, and hybrid rendering
- TypeScript & Vue 3 component architecture

#### 🎨 Features
- Real-time updates via WebSocket
- Integrated REST API & WebSocket support

---

## 📁 Project Structure

```
/backend      → Go Gin server + Sockie for WebSockets
/frontend     → Nuxt 3 frontend (Vue-based SPA/SSR)
```

---

## ⚙️ Backend Setup

```bash
# Run API server
go run . http

# Run database migrations
go run . migrate up

# Rollback migrations
go run . migrate down

# Seed initial data
go run . migrate seed

# Refresh migrations and reseed
go run . migrate refresh
```

> 💡 Make sure you have Go and your database configured before running these.

---

## 🎨 Frontend Setup

```bash
# Install dependencies
yarn

# Start local dev server
yarn dev

# Build for production
yarn build
```

---

## 🌟 Key Highlights

- ⚡ **Real-Time Messaging**: Instant updates using WebSockets
- 🔒 **Authentication & API**: Secure user and chat data management
- 🧹 **Separation of Concerns**: Cleanly split backend & frontend logic
- 🚀 **High Performance**: Built for speed and scale

---

## 📌 Ideal Use Cases
- Live chat applications
- Collaborative dashboards
- Real-time notification systems

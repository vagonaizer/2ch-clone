# 🧵 Go Forum — A Minimal Imageboard in Go

A lightweight imageboard-style forum built with **Golang**, **PostgreSQL**, and **HTMX**.  
Inspired by classic boards like 2ch and 4chan — simple threads, clean layout, and no fluff.

---

## 📸 Screenshots

### 🖼️ 1. Header / Landing Page

The main header and first impression.

![Header](/images/1.jpg)

---

### 💬 2. Thread View with Posts

A thread page with the original post, replies below, and a reply form.

![Thread View](/images/2.jpg)

---

### 🧵 3. Recent Threads

Quick access to the latest active threads.

![Recent Threads](/images/3.jpg)

---

### 🛠️ 4. Admin Panel

Manage threads: pin, lock, delete, and more.

![Admin Panel](/images/4.jpg)

---

### 🏠 5. Homepage

Simple overview of the board with thread previews.

![Homepage](/images/5.jpg)

---

## ⚙️ Stack

- **Go (Golang)** — backend logic
- **PostgreSQL** — persistent thread & post storage
- **HTMX** — minimal JS interactivity
- **HTML Templates** — clean, fast-rendered pages

---

## 🚀 Getting Started

```bash
go mod tidy
go run main.go
# 🍀 Vagonach: 4chan clone

A lightweight imageboard-style forum built with **Golang**, **PostgreSQL**, and **HTMX**.  
Inspired by classic boards like 2channel and 4chan.

![Header](/images/1.jpg)

## 📑 Table of Contents

- [📺 Video Demo](#youtube-video-on-how-things-done)
- [🖼️ Screenshots](#screenshots)
- [⚙️ Stack](#️stack)
- [🔧 Features](#features)
  - [🧭 Public Pages](#-public-pages)
  - [🧵 Boards & Threads](#-boards--threads)
  - [💬 Posts](#-posts)
  - [🔐 Admin Panel](#-admin-panel)
- [📁 Project Structure](#project-structure)

---

## Youtube video on how things done:

Click on the link or the image below

link: https://www.youtube.com/watch?v=QW0xHsxjweg

[![Watch the video](/images/6.png)](https://www.youtube.com/watch?v=QW0xHsxjweg)


---

### 💬 1. Thread View with Posts

A thread page with the original post, replies below, and a reply form.

![Thread View](/images/2.jpg)

---

### 🧵 2. Recent Threads

Quick access to the latest active threads.

![Recent Threads](/images/3.jpg)

---

### 🛠️ 3. Admin Panel

Manage threads: pin, lock, delete, and more.

![Admin Panel](/images/4.jpg)

---

### 🏠 4. Homepage

Simple overview of the board with thread previews.

![Homepage](/images/5.png)

---

## ⚙️ Stack

- **Go (Golang)** — backend logic
- **PostgreSQL** — persistent thread & post storage
- **HTMX** — minimal JS interactivity
- **HTML Templates** — clean, fast-rendered pages

## 🔧 Features

### 🧭 Public Pages

- `/` — Homepage with navigation
- `/about` — About project
- `/boards/:slug` — View board page
- `/threads/:id` — View single thread and posts

---

### 🧵 Boards & Threads

- `GET /boards` — List all boards
- `GET /boards/:slug/threads` — Get threads by board
- `POST /boards/:slug/threads` — Create new thread
- `PATCH /threads/:id/sticky` — Toggle sticky flag
- `PATCH /threads/:id/lock` — Toggle locked flag
- `DELETE /threads/:id` — Delete thread

---

### 💬 Posts

- `GET /threads/:id/posts` — Get posts in thread
- `POST /threads/:id/posts` — Create new post
- `GET /posts/:id` — View single post (API)
- `DELETE /posts/:id` — Delete post

---

### 🔐 Admin Panel

- `GET /admin/login` — Admin login page
- `POST /admin/login` — Login submission
- `GET /admin/logout` — Logout
- `GET /admin` — Admin dashboard (with auth)
- `POST /admin/threads/:id/delete` — Delete thread
- `POST /admin/threads/:id/sticky` — Make thread sticky
- `POST /admin/threads/:id/lock` — Lock thread

---

### Project Structure

```
2ch-clone/
├── cmd/
│   └── vagonach/               # Main entrypoint (main.go)
├── internal/
│   ├── app/                    # App startup logic
│   ├── config/                 # Configuration management
│   ├── db/
│   │   ├── migrations/         # SQL migrations (up/down)
│   │   └── postgres/           # PostgreSQL connection setup
│   ├── entity/                 # Domain models (Board, Thread, Post, etc.)
│   ├── repository/             # Repository interfaces & implementations
│   │   └── postgresql/         # PostgreSQL-specific repos
│   ├── transport/
│   │   └── http/               # HTTP handlers and routing
│   └── usecase/                # Business logic services
├── web/
│   ├── img/                    # Static images (logo, banner, etc.)
│   ├── styles/                 # CSS stylesheets
│   ├── templates/              # HTML templates for pages
│   └── uploads/                # Uploaded files (if used)
├── images/                    # Screenshots for README
├── configs/                   # App configs (WIP or placeholder)
├── go.mod / go.sum            # Go modules
└── README.md
```


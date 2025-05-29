# 🕒 Pomodoro Timer CLI

**Pomodoro Timer** is a simple and lightweight command-line tool written in Go that helps you stay focused using the Pomodoro technique: 25 minutes of work followed by 5 minutes of break.

![Go](https://img.shields.io/badge/Built%20with-Go-blue?style=flat-square)
![Version](https://img.shields.io/github/v/release/msapuan/pomodoro-timer?style=flat-square)

---

## ✨ Features

- ⏱️ Classic Pomodoro timer (25 min work / 5 min break)
- 🔔 Desktop notifications using `notify-send`
- 💬 Motivational quotes at each break
- 🧠 Lightweight, fast, and dependency-free
- 📦 Available as `.deb` installer for Linux (Debian/Ubuntu)

---

## 📦 Installation

### 🔹 Option 1: Install from `.deb` file

1. Download the latest `.deb` from the [Releases page](https://github.com/msapuan/pomodoro-timer/releases)
2. Install with:

```bash
sudo dpkg -i pomodoro-timer_1.0.0.deb
```

3. Run the timer :

```bash
pomodoro-timer
```

### 🔹 Option 2: Build from source

Requires **Go** installed

```bash
git clone https://github.com/msapuan/pomodoro-timer.git
cd pomodoro-timer
go build -o pomodoro-timer
./pomodoro-timer
```

## 🔔 Notifications

This tool uses `notify-send` to trigger desktop notifications.

If you don’t have it installed, run:

```bash
sudo apt install libnotify-bin
```

## 👤 Author

Built with ❤️ by [@msapuan](https://github.com/msapuan)

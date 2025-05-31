
# ⏳ Pomodoro Timer

A simple and effective Pomodoro timer CLI app to help you stay focused and productive.  
Now with sound notifications and motivational quotes! 🎧💬

![Go](https://img.shields.io/badge/Built%20with-Go-blue?style=flat-square)
![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg?style=flat-square)
![Version](https://img.shields.io/github/v/release/msapuan/pomodoro-timer?style=flat-square)

---

## ✨ Features

- ⏱ Pomodoro technique with focus & break cycles
- 🔊 Notification sounds for focus and break
- 💡 Random motivational quotes
- 🐧 Works smoothly on Linux terminal

---

## 🛠 Installation

You can install `pomodoro-timer` in two ways:

### 📦 Option 1: Download `.deb` File Manually

1. Download the latest `.deb` file from [Releases](https://github.com/msapuan/pomodoro-timer/releases).
2. Install it using `dpkg`:

   ```bash
   sudo dpkg -i pomodoro-timer_<version>_amd64.deb
   ```

3. Done! Run it with:

   ```bash
   pomodoro-timer
   ```

> ⚠️ Note: With this method, future updates must be installed manually.

---

### 🔁 Option 2: Install via APT Repository (Recommended)

This method allows you to receive updates via `sudo apt upgrade`.

#### 📋 Step-by-Step:

1. Add the GPG key (optional, skipped if unsigned):
   > _If you don't use signing yet, you can skip this step._

2. Add the custom APT repository:

   ```bash
   echo "deb [trusted=yes] https://msapuan.github.io/pomodoro-timer stable main" | sudo tee /etc/apt/sources.list.d/pomodoro-timer.list
   ```

3. Update your package list:

   ```bash
   sudo apt update
   ```

4. Install the app:

   ```bash
   sudo apt install pomodoro-timer
   ```

✅ Done! Now you can run:

```bash
pomodoro-timer
```

> 🆕 Updates will be available automatically via `sudo apt update && sudo apt upgrade`.

---

## 🚀 Usage

Just run:

```bash
pomodoro-timer
```

You'll hear notification sounds at the beginning and end of each session, and get a motivational quote to keep you going 💪

---

## 📦 How It Works (for developers)

This app is written in Go and packaged as a `.deb` file for easy distribution via a custom APT repo hosted on GitHub Pages.

### Project Structure:

```
pomodoro-timer/
├── cmd/...
├── internal/...
├── assets/sounds/
├── debian/control
├── build/
```

---

## 🤝 Contributing

Pull requests and ideas are always welcome. Let’s build this together!

---

## 🧑‍💻 Author

Made with ❤️ by [@msapuan](https://github.com/msapuan)

---

## 📄 License

This project is licensed under the MIT License.

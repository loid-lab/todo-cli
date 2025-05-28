# 🧋 Todo List TUI (Bubble Tea)

A simple interactive terminal app to manage a to-do list using the [Bubble Tea](https://github.com/charmbracelet/bubbletea) framework in Go.

## ✨ Features

- View to-do items in a terminal interface
- Navigate with arrow keys (`↑/↓`) or `j/k`
- Toggle selection with `space` or `enter`
- Add new items by pressing `a`
- Delete selected items by pressing `d`
- Quit with `q` or `ctrl+c`

## 🖼️ Preview

```text
What should we buy at the market?

> [x] Buy carrots
  [ ] Buy celery
  [ ] Buy kohlrabi

↑/↓: move • space/enter: toggle • a: add • d: delete • q: quit
```

## 🚀 Getting Started

### 1. Clone the repo

```bash
git clone https://github.com/loid-lab/todo-cli.git
cd todo-cli
```

### 2. Install dependencies

```bash
go mod init todo-cli
go get github.com/charmbracelet/bubbletea
```

### 3. Run the app

```bash
go run main.go
```

## 🧠 How it works

- The app keeps a list of items (`choices`)
- You can move the cursor to select items
- `a` enters input mode where you type a new item
- Selected items can be deleted with `d`
- Everything runs in a single terminal interface

## 🛠️ Tech Stack

- [Go](https://golang.org/)
- [Bubble Tea](https://github.com/charmbracelet/bubbletea) – for terminal UI
- [Lip Gloss](https://github.com/charmbracelet/lipgloss) - for styling

## 📄 License

MIT — feel free to use, modify, and distribute.

---

**Made with 💚 using Bubble Tea**
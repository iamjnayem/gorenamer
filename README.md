
# GoRenamer 🐹📁

**GoRenamer** is a simple yet powerful command-line file renaming utility written in [Go](https://golang.org).  
It allows batch renaming of files based on flexible patterns like prefix, suffix, replace, filter by extension, and more.

---

## 🚀 Features

- ✅ Add **prefix** or **suffix** to file names
- 🔁 **Replace** substrings inside file names
- 🔍 Filter files by **extension** or **starting text**
- 🧪 **Dry run** mode to preview changes before renaming
- 🛠️ Lightweight, fast & easy to use

---

## 📦 Installation

### **Option 1: Install via Go**

Make sure you have Go installed (`go version` should return your Go version)

```bash
go install github.com/iamjnayem/gorenamer@latest
```

### **Option 2: Build and Install Locally**

1. Clone the repository:

    ```bash
    git clone https://github.com/iamjnayem/gorenamer.git
    cd gorenamer
    ```

2. Build the binary:

    ```bash
    go build -o gorenamer
    ```

3. Move the binary to a directory in your `$PATH` (e.g., `/usr/local/bin` for system-wide access):

    ```bash
    sudo mv gorenamer /usr/local/bin/
    ```

    This allows you to run `gorenamer` from anywhere in your terminal.

---

## 🧪 Usage

```bash
go run main.go rename [flags]
```

### 🔧 Available Flags:

| Flag          | Description                                           | Example                        |
|---------------|-------------------------------------------------------|--------------------------------|
| `--prefix`    | Add prefix to file names                              | `--prefix new_`                |
| `--suffix`    | Add suffix (before extension)                         | `--suffix _v2`                 |
| `--replace`   | Replace substring (format: old:new)                   | `--replace img:photo`          |
| `--start-with`| Only rename files starting with given text            | `--start-with doc_`            |
| `--ext`       | Only rename files with specific extension             | `--ext .txt`                   |
| `--dry-run`   | Preview changes without actual renaming               | `--dry-run`                    |

---

## 🧑‍💻 Example

Rename all `.txt` files starting with `doc_`, replace `old` with `new`, add a prefix and suffix:

```bash
go run main.go rename --prefix new_ --suffix _v2 --replace old:new --start-with doc_ --ext .txt --dry-run
```

---

## 📁 Directory Structure

```
.
├── cmd/
│   └── rename.go     # rename command logic
├── main.go           # entry point
├── go.mod / go.sum   # Go module files
```

---

## 🧑‍🎓 Author

Made with ❤️ by [@iamjnayem](https://github.com/iamjnayem)

---

## 📜 License

MIT License. Feel free to use and contribute.
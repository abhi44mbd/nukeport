# nukeport

> A fast, safe CLI tool to find and terminate processes running on ports.

Stop wasting time with:
```bash
lsof -i :3000
kill -9 PID
```

Just run:

```bash
nukeport 3000
```

---

## Features

- Find processes running on a port
- Graceful process termination (SIGTERM)
- Force kill support (`--force`)
- Multi-process handling
- Interactive confirmation prompts
- Dry run mode
- Clean colored terminal UI
- Fast and lightweight
- macOS + Linux support

---

## Installation

### Homebrew

```bash
brew install abhi44mbd/tools/nukeport
```

### Build From Source

```bash
git clone https://github.com/abhi44mbd/nukeport.git

cd nukeport

go build -o nukeport
```

---

## Usage

### Kill Process Running On Port

```bash
nukeport 3000
```

### Force Kill

```bash
nukeport 3000 --force
```

or

```bash
nukeport 3000 -f
```

### Skip Confirmation

```bash
nukeport 3000 --yes
```

### Dry Run

```bash
nukeport 3000 --dry-run
```

---

## Example Output

```bash
✓ Process Found

──────────────────────────────

 Name     : node
 PID      : 12345
 Port     : 3000
 Command  : vite dev server

──────────────────────────────

Continue and terminate this process? [y/N]:
```

---

## Flags

| Flag | Description |
|------|-------------|
| `-f, --force` | Force kill using SIGKILL |
| `-y, --yes` | Skip confirmation prompts |
| `--dry-run` | Preview actions without terminating processes |
| `-h, --help` | Show help menu |

---

## Tech Stack

Built with:

- Go
- Cobra CLI
- Lip Gloss
- yacspin

---

## Development

Run locally:

```bash
go run main.go 3000
```

---

## Roadmap

### Current
- [x] Process lookup by port
- [x] Graceful process termination
- [x] Force kill support
- [x] Interactive confirmations
- [x] Dry run mode
- [x] Colored terminal UI

### Planned
- [ ] Multi-port support
- [ ] Smart framework detection
- [ ] JSON output mode
- [ ] Restart command support
- [ ] Homebrew distribution
- [ ] Linux package support
- [ ] Windows support

---

## Contributing

Contributions, issues, and feature requests are welcome.

Feel free to open a PR or issue.

---

## License

MIT License
```bash
342  git clone https://github.com/golang-standards/project-layout
343  go mod init gopackages-layout
344  go mod vendor
345  go get github.com/huandu/xstrings
346  go mod verify
347  go run cmd/gopackages-layout/main.go
```
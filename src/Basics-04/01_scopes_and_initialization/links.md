- https://pkg.go.dev/std
- https://github.com/fatih/color.git github.com/fatih/color

```bash
  185  go mod init rebrain-go
  186  cd src/
  187  git clone https://github.com/fatih/color.git github.com/fatih/color
  188  go get github.com/fatih/color
```

```text
go@linux:~/GolandProjects/rebrain-go$ echo $GOPATH
/home/go/GolandProjects/rebrain-go
go@linux:~/GolandProjects/rebrain-go$ unset GOPATH
```

```text
  212  echo "export GOPATH=~/GolandProjects/rebrain-go" >> ~/.profile
  213  echo $GOPATH
  214  cd $GOPATH/src/Basics-04/09_scopes_and_initialization/
  218  go mod init rebrain-go-09
  219  go mod vendor
  225  go get github.com/fatih/color
  226  go run main.go
  231  go build main.go 
  232  cd $GOPATH
  236  ./src/Basics-04/09_scopes_and_initialization/main 
```

```text
go mod tidy - clean unused
go mod download - download dependencies without building project
go mod vendor - storage of packages in project directory
go build -mod vendor
```

```text
go get github.com/huandu/xstrings
```
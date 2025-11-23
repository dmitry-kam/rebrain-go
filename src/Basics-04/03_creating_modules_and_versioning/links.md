- https://habr.com/ru/articles/421411/
- https://semver.org/lang/ru/

# Begin
A primitive repository was created on the internal GitLab Rebrain.
Connecting it here for use in the program.

# Setting up Go for Private Repositories

## Environment Variables Configuration

```bash
export GOPRIVATE=gitlab.rebrainme.com
```

Not required:
```bash
export GONOPROXY=gitlab.rebrainme.com  
export GONOSUMDB=gitlab.rebrainme.com
```

## Purpose of Each Variable

- **`GOPRIVATE`** - Informs Go that specified domains host private repositories
- **`GONOPROXY`** - Disables public proxy (proxy.golang.org) usage for these domains
- **`GONOSUMDB`** - Prevents checksum verification against public database for these domains

## Why This Configuration is Necessary

- Go's default behavior attempts to download packages through public proxies
- Private repositories return 404 errors when accessed via public infrastructure
- These environment variables bypass public systems for specified private domains

## .netrc File Setup

Create or modify `~/.netrc` with the following content:

```bash
machine gitlab.rebrainme.com
login dmitriy_kam
password <access-token>
```

Create and add in Gitlab Preferences the SSH key
```bash
  ssh-keygen -t ed25519 -C "dd.starikov@gmail.com"
  cat ~/.ssh/id_ed25519.pub
```

```text
git ls-remote https://gitlab.rebrainme.com/golang_users_repos/5807/utils.git

f2e2ab90ea3d026a54b1e8d8e8c7b1c5d13f3b6e	HEAD
f2e2ab90ea3d026a54b1e8d8e8c7b1c5d13f3b6e	refs/heads/main
f2e2ab90ea3d026a54b1e8d8e8c7b1c5d13f3b6e	refs/heads/v3
3ad982e21ced54e9de334912425d41217929e00e	refs/tags/v1.0.0
a5e6d94c3d03b7ec63a3a3f9b1af25737e251eaa	refs/tags/v1.1.0
133e1d05c759446a2a72e7ec787b883f2ec2c6bd	refs/tags/v2.0.0
e8ee5d7fccd6816d55af167f1f62b5b8870b568c	refs/tags/v2.0.0^{}
bb523e9cd371c62826029c9e7188338dd453e549	refs/tags/v3.0.0
f2e2ab90ea3d026a54b1e8d8e8c7b1c5d13f3b6e	refs/tags/v3.0.0^{}
```

## Authentication Process Flow

1. Go identifies domain in GOPRIVATE list
2. Searches ~/.netrc for domain credentials
3. Uses credentials to access private repository
4. Fetches library directly from GitLab

## Security Configuration

```bash
chmod 600 ~/.netrc
```

## Struggle with go.mod and Rebrain repo

```go
module test-rebrain-repo-lib

go 1.22.2

// 1. usual require doesn't work because of repo name contains 3 parts. Go mod cant find it
//require (
//    gitlab.rebrainme.com/golang_users_repos/5807/utils v1.1.0
//    gitlab.rebrainme.com/golang_users_repos/5807/utils/v3 v3.0.0
//)

// 2. hack
//replace (
//    gitlab.rebrainme.com/golang_users_repos/5807/utils => gitlab.rebrainme.com/golang_users_repos/5807/utils.git v1.1.0
//    gitlab.rebrainme.com/golang_users_repos/5807/utils/v3 => gitlab.rebrainme.com/golang_users_repos/5807/utils.git v3.0.0
//)

// 3. via require, but doesn't work because of I called the module just `utils` not `gitlab.rebrainme.com/golang_users_repos/5807/utils`
//go: gitlab.rebrainme.com/golang_users_repos/5807/utils.git@v1.1.0: parsing go.mod:
//	module declares its path as: utils
//	        but was required as: gitlab.rebrainme.com/golang_users_repos/5807/utils.git
//require (
//    gitlab.rebrainme.com/golang_users_repos/5807/utils.git v1.1.0
//    gitlab.rebrainme.com/golang_users_repos/5807/utils.git/v3 v3.0.0
//)

// 4. use require + replace
require (
    utils v1.1.0
    utils/v3 v3.0.0
)

replace (
utils v1.1.0 => gitlab.rebrainme.com/golang_users_repos/5807/utils.git v1.1.0
utils/v3 v3.0.0 => gitlab.rebrainme.com/golang_users_repos/5807/utils.git/v3 v3.0.0
)
```

Run
```bash
  292  go run task.go 
  293  go mod download utils
  294  go mod download utils utils/v3
  295  go run task.go
```

Output:
```text
Using v1.1.0:
Contains 'banana': true
Contains 'grape': false
ContainsInt 30: true
ContainsInt 99: false
Using v3.0.0:
InSlice 'banana': true
InSlice 'grape': false
InSliceInt 30: true
InSliceInt 99: false
```
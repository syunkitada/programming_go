# project の sample

```
$ go version
go version go1.20.3 linux/amd64

$ mkdir [project]
$ cd [project]
$ go mod init "$(git remote -v | grep push | awk '{print $2}' | sed -e 's/git@//g' | sed -e 's/:/\//g' | sed -e 's/.git//g')/project_examples/dnsapi"
$ go mod tidy
```

```
$ watchexec -e go -cr -- go run cmd/hoge-api/main.go -port 1080
```

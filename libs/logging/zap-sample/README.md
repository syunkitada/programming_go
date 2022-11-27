# Zap

- [github](https://github.com/uber-go/zap)

## 出力例

```
{"level":"INFO","time":"2021-12-19T16:55:54.876885229+09:00","caller":"zap-sample/main.go:56","func":"main.main","msg":"Hello World","traceId":"c6veb2hku8rtjm00heog","status":200}
{"level":"INFO","time":"2021-12-19T16:55:54.877026855+09:00","caller":"zap-sample/main.go:62","func":"main.main","msg":"Hello World","second":{"traceId":"c6veb2hku8rtjm00heog","status":200}}
{"level":"INFO","time":"2021-12-19T16:55:54.877053395+09:00","caller":"zap-sample/main.go:70","func":"main.main","msg":"new request, in nested object","req":{"url":"/test","ip":"127.0.0.1","port":8080,"remote":{"ip":"127.0.0.1","port":31200}}}
{"level":"INFO","time":"2021-12-19T16:55:54.877076569+09:00","caller":"zap-sample/main.go:71","func":"main.main","msg":"new request, inline","url":"/test","ip":"127.0.0.1","port":8080,"remote":{"ip":"127.0.0.1","port":31200}}
{"level":"INFO","time":"2021-12-19T16:55:54.877096436+09:00","caller":"zap-sample/main.go:99","func":"main.trace","msg":"Start Trace","traceId":"c6veb2hku8rtjm00heog"}
{"level":"INFO","time":"2021-12-19T16:55:55.877464481+09:00","caller":"zap-sample/main.go:101","func":"main.trace.func1","msg":"End Trace","traceId":"c6veb2hku8rtjm00heog","elapsed":1000}
{"level":"INFO","time":"2021-12-19T16:55:55.877640172+09:00","caller":"zap-sample/main.go:78","func":"main.main","msg":"Hello World traceId=c6veb2hku8rtjm00heog, status=200"}
{"level":"INFO","time":"2021-12-19T16:55:55.877688753+09:00","caller":"zap-sample/main.go:80","func":"main.main","msg":"Hello World","traceId":"c6veb2hku8rtjm00heog","status":200}
{"level":"INFO","time":"2021-12-19T16:55:55.877733137+09:00","caller":"zap-sample/main.go:83","func":"main.main.func1","msg":"Hello Goroutine","traceId":"c6veb2hku8rtjm00heog","status":200}
{"level":"WARN","time":"2021-12-19T16:55:55.877752343+09:00","caller":"zap-sample/main.go:87","func":"main.main","msg":"Hello Warn","traceId":"c6veb2hku8rtjm00heog","status":200}
{"level":"ERROR","time":"2021-12-19T16:55:55.877768914+09:00","caller":"zap-sample/main.go:88","func":"main.main","msg":"Hello Error","traceId":"c6veb2hku8rtjm00heog","status":200,"stacktrace":"main.main\n\t/home/owner/go/1.15.6/src/github.com/syunkitada/go-samples/logging/zap-sample/main.go:88\nruntime.main\n\t/home/owner/.goenv/versions/1.15.6/src/runtime/proc.go:204"}
{"level":"FATAL","time":"2021-12-19T16:55:55.877825861+09:00","caller":"zap-sample/main.go:89","func":"main.main","msg":"Hello Fatal","traceId":"c6veb2hku8rtjm00heog","status":200,"stacktrace":"main.main\n\t/home/owner/go/1.15.6/src/github.com/syunkitada/go-samples/logging/zap-sample/main.go:89\nruntime.main\n\t/home/owner/.goenv/versions/1.15.6/src/runtime/proc.go:204"}
exit status 1
```

## 参考

- [golang の高速な構造化ログライブラリ「zap」の使い方](https://qiita.com/emonuh/items/28dbee9bf2fe51d28153)

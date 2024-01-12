Repro by running:

```sh
TEST_COUNT=<N> make test-race
```

Example failure output:

```
% TEST_COUNT=1 make test-race
go test -v -race -count=1 .
=== RUN   TestShutdownCleanupStream
=== PAUSE TestShutdownCleanupStream
=== RUN   TestShutdownCleanupUnary
    repro_test.go:122:
--- SKIP: TestShutdownCleanupUnary (0.00s)
=== CONT  TestShutdownCleanupStream
    repro_test.go:74: test started
    repro_test.go:97: sending request
    repro_test.go:20: StreamEcho started
    repro_test.go:102: receiving response
    repro_test.go:107: response: text:"foo"
    repro_test.go:110: closing client connection
    repro_test.go:113: stopping server
    repro_test.go:116: server stop finished
    repro_test.go:83: server exited
    repro_test.go:119: test ended
--- PASS: TestShutdownCleanupStream (0.01s)
PASS
panic: Log in goroutine after TestShutdownCleanupStream has completed: StreamEcho error in Recv error: rpc error: code = Canceled desc = context canceled
        panic: Log in goroutine after TestShutdownCleanupStream has completed: StreamEcho ended


goroutine 66 [running]:
testing.(*common).logDepth(0xc000183520, {0xc000026198, 0x11}, 0x3)
        /opt/homebrew/Cellar/go/1.21.6/libexec/src/testing/testing.go:1022 +0x4f8
testing.(*common).log(...)
        /opt/homebrew/Cellar/go/1.21.6/libexec/src/testing/testing.go:1004
testing.(*common).Log(0xc000183520, {0xc0002c33d8, 0x1, 0x1})
        /opt/homebrew/Cellar/go/1.21.6/libexec/src/testing/testing.go:1045 +0x74
panic({0x102a706c0?, 0xc00003a4d0?})
        /opt/homebrew/Cellar/go/1.21.6/libexec/src/runtime/panic.go:920 +0x26c
testing.(*common).logDepth(0xc000183520, {0xc00002a2a0, 0x52}, 0x3)
        /opt/homebrew/Cellar/go/1.21.6/libexec/src/testing/testing.go:1022 +0x4f8
testing.(*common).log(...)
        /opt/homebrew/Cellar/go/1.21.6/libexec/src/testing/testing.go:1004
testing.(*common).Logf(0xc000183520, {0x10293fb60, 0xc}, {0xc0002c3748, 0x2, 0x2})
        /opt/homebrew/Cellar/go/1.21.6/libexec/src/testing/testing.go:1055 +0x84
github.com/sunjayBhatia/grpc-go-issue-repro/shutdown-cleanup_test.(*EchoServer).StreamEcho.func1(...)
        /Users/sunjayb/workspace/grpc-go-issue-repro/shutdown-cleanup/repro_test.go:25
github.com/sunjayBhatia/grpc-go-issue-repro/shutdown-cleanup_test.(*EchoServer).StreamEcho(0xc00022e180, {0x102b3e098, 0xc0003861c0})
        /Users/sunjayb/workspace/grpc-go-issue-repro/shutdown-cleanup/repro_test.go:38 +0x330
github.com/sunjayBhatia/grpc-go-issue-repro/shutdown-cleanup/echo._EchoService_StreamEcho_Handler({0x102aadd00?, 0xc00022e180}, {0x102b3c580?, 0xc0003bc000})
        /Users/sunjayb/workspace/grpc-go-issue-repro/shutdown-cleanup/echo/echo_grpc.pb.go:115 +0xb0
google.golang.org/grpc.(*Server).processStreamingRPC(0xc00029a000, {0x102b3b5a8, 0xc000380210}, {0x102b3e268, 0xc000280b60}, 0xc0003a8000, 0xc00028a210, 0x102eb3760, 0x0)
        /Users/sunjayb/go/pkg/mod/google.golang.org/grpc@v1.60.1/server.go:1666 +0x16f4
google.golang.org/grpc.(*Server).handleStream(0xc00029a000, {0x102b3e268, 0xc000280b60}, 0xc0003a8000)
        /Users/sunjayb/go/pkg/mod/google.golang.org/grpc@v1.60.1/server.go:1787 +0x12f0
google.golang.org/grpc.(*Server).serveStreams.func2.1()
        /Users/sunjayb/go/pkg/mod/google.golang.org/grpc@v1.60.1/server.go:1016 +0xa0
created by google.golang.org/grpc.(*Server).serveStreams.func2 in goroutine 43
        /Users/sunjayb/go/pkg/mod/google.golang.org/grpc@v1.60.1/server.go:1027 +0x1d8
FAIL    github.com/sunjayBhatia/grpc-go-issue-repro/shutdown-cleanup    0.515s
FAIL
make: *** [test-race] Error 1
```

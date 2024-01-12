package shutdown_clean_test

import (
	"context"
	"net"
	"sync"
	"testing"

	"github.com/sunjayBhatia/grpc-go-issue-repro/shutdown-cleanup/echo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type EchoServer struct {
	t *testing.T
	echo.UnimplementedEchoServiceServer
}

func (e *EchoServer) StreamEcho(srv echo.EchoService_StreamEchoServer) error {
	e.t.Log("StreamEcho started")
	defer e.t.Log(("StreamEcho ended"))

	done := func(logMsg string, err error) error {
		e.t.Logf("%s error: %v", logMsg, err)
		return err
	}

	ctx := srv.Context()

	for {
		select {
		case <-ctx.Done():
			return done("StreamEcho closing stream", ctx.Err())
		default:
			content, err := srv.Recv()
			if err != nil {
				return done("StreamEcho error in Recv", err)
			}
			err = srv.Send(content)
			if err != nil {
				return done("StreamEcho error in Send", err)
			}
		}
	}
}

func setupServer(t *testing.T) (*grpc.Server, net.Listener) {
	t.Helper()

	s := grpc.NewServer() // Comment this line and uncomment the below to see the failure caused by stream workers.
	// s := grpc.NewServer(grpc.NumStreamWorkers(1))
	e := &EchoServer{t: t}
	echo.RegisterEchoServiceServer(s, e)

	lis, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}

	return s, lis
}

func TestShutdownCleanupStream(t *testing.T) {
	t.Log("test started")
	defer t.Log("test ended")

	s, lis := setupServer(t)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		t.Log("server starting")
		s.Serve(lis)
		t.Log("server exited")

		wg.Done()
	}()

	cc, err := grpc.Dial(lis.Addr().String(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatal(err)
	}
	cli, err := echo.NewEchoServiceClient(cc).StreamEcho(context.TODO())
	if err != nil {
		t.Fatal(err)
	}

	t.Log("sending request")
	err = cli.Send(&echo.Content{Text: "foo"})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("receiving response")
	resp, err := cli.Recv()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("response: %v", resp)

	t.Log("closing client connection")
	cc.Close()

	t.Log("stopping server")
	s.GracefulStop() // Issue can be reproduced with GracefulStop, Stop, or both
	// s.Stop()
	t.Log("server stop finished")

	wg.Wait()
}

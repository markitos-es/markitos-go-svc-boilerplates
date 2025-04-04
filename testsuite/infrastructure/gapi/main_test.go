package gapi_test

import (
	"context"
	"log"
	"net"
	"os"
	"testing"

	"markitos-svc-boilerplates/infrastructure/gapi"
	"markitos-svc-boilerplates/internal/domain"
	"markitos-svc-boilerplates/testsuite/infrastructure/testdb"
	internal_test "markitos-svc-boilerplates/testsuite/internal"

	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener
var grpcServer *grpc.Server
var grpcClient gapi.BoilerplateServiceClient
var ctx context.Context

func TestMain(m *testing.M) {
	setup()

	code := m.Run()

	grpcServer.Stop()
	os.Exit(code)
}

func setup() {
	lis = bufconn.Listen(bufSize)

	grpcServer = grpc.NewServer()

	server := gapi.NewServer(":8080", testdb.GetRepository())
	gapi.RegisterBoilerplateServiceServer(grpcServer, server)

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("['.']:> Error serving gRPC server: %v", err)
		}
	}()

	conn, err := grpc.DialContext(
		context.Background(),
		"bufnet",
		grpc.WithContextDialer(bufDialer),
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("['.']:> Failed to dial bufnet: %v", err)
	}

	grpcClient = gapi.NewBoilerplateServiceClient(conn)
	ctx = context.Background()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func createPersistedRandomBoilerplate() *domain.Boilerplate {
	boilerplate := internal_test.NewRandomBoilerplate()
	testdb.GetRepository().Create(boilerplate)
	return boilerplate
}

func deletePersistedRandomBoilerplate(boilerplateId string) {
	id, err := domain.NewBoilerplateId(boilerplateId)
	if err != nil {
		return
	}
	testdb.GetRepository().Delete(id)
}

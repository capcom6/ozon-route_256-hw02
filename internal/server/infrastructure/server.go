// Copyright 2022 Aleksandr Soloshenko
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package infrastructure

import (
	"context"
	"database/sql"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"gitlab.ozon.dev/capcom6/homework-2/internal/server/config"
	"gitlab.ozon.dev/capcom6/homework-2/internal/server/database"
	"gitlab.ozon.dev/capcom6/homework-2/internal/server/puller"
	"gitlab.ozon.dev/capcom6/homework-2/internal/server/repositories"
	"gitlab.ozon.dev/capcom6/homework-2/internal/server/service"
	pb "gitlab.ozon.dev/capcom6/homework-2/pkg/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Run() error {
	cfg, err := loadConfig()
	if err != nil {
		return err
	}

	log.Println("Config loaded")

	db, err := connectDatabase(cfg.Database)
	if err != nil {
		return err
	}
	defer db.Close()
	log.Println("Database connected")

	mbrepo := repositories.NewMailboxes(db)
	log.Println("Repository created")

	lis, err := net.Listen("tcp", cfg.Server.GRPC)
	if err != nil {
		return err
	}

	s := grpc.NewServer()
	service := service.New(mbrepo, puller.New())

	pb.RegisterMailAggregatorServer(s, service)

	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		// We received an interrupt signal, shut down.
		s.GracefulStop()
		close(idleConnsClosed)
	}()

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	runGateway(ctx, cfg.Server.Gateway, lis.Addr().String())

	log.Printf("Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		return err
	}

	<-idleConnsClosed

	log.Println("Succesfull shutdown")

	return nil
}

func loadConfig() (*config.Config, error) {
	fileName, ok := os.LookupEnv("CONFIG_PATH")
	if !ok {
		fileName = "./configs/server.yml"
	}

	return config.Load(fileName)
}

func connectDatabase(cfg config.Database) (*sql.DB, error) {
	db, err := database.New(database.Config{
		Host:     cfg.Host,
		Port:     cfg.Port,
		Database: cfg.Database,
		User:     cfg.User,
		Password: cfg.Password,
	})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func runGateway(ctx context.Context, listen string, grpcHost string) error {
	if listen == "" {
		return nil
	}

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := pb.RegisterMailAggregatorHandlerFromEndpoint(ctx, mux, grpcHost, opts)
	if err != nil {
		return err
	}

	go func() {
		log.Printf("Gateway listening at %v", listen)
		http.ListenAndServe(listen, mux)
	}()

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return nil
}

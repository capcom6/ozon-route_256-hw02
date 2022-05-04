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
	"database/sql"
	"io"
	"log"
	"net"
	"os"
	"os/signal"

	"gitlab.ozon.dev/capcom6/homework-2/internal/server/config"
	"gitlab.ozon.dev/capcom6/homework-2/internal/server/database"
	"gitlab.ozon.dev/capcom6/homework-2/internal/server/repositories"
	"gitlab.ozon.dev/capcom6/homework-2/internal/server/service"
	pb "gitlab.ozon.dev/capcom6/homework-2/pkg/api"
	"google.golang.org/grpc"
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
	log.Println("Database connected")

	mbrepo := repositories.NewMailboxes(db)
	log.Println("Repository created")

	lis, err := net.Listen("tcp", cfg.HTTP.Listen)
	if err != nil {
		return err
	}

	s := grpc.NewServer()
	service := service.New(mbrepo)

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

	log.Printf("Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		return err
	}

	<-idleConnsClosed

	log.Println("Succesfull shutdown")

	return nil
}

func loadConfig() (*config.Config, error) {
	fileName := os.Getenv("CONFIG_PATH")
	if fileName == "" {
		fileName = "./configs/server.yml"
	}

	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	cfg, err := config.ParseConfig(bytes)

	return cfg, err
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

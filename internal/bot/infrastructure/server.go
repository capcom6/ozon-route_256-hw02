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
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"gitlab.ozon.dev/capcom6/homework-2/internal/bot/config"
	"gitlab.ozon.dev/capcom6/homework-2/internal/bot/handlers"
	"gitlab.ozon.dev/capcom6/homework-2/internal/bot/interpreter"
	pb "gitlab.ozon.dev/capcom6/homework-2/pkg/api"
	"gitlab.ozon.dev/capcom6/homework-2/pkg/telegram"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const timeout = 10 * time.Second

func Run() error {
	cfg, err := loadConfig()
	if err != nil {
		return err
	}

	log.Println("Config loaded")

	backend, err := connectToBackend(cfg.Backend)
	if err != nil {
		return err
	}
	log.Println("Backend connected")

	h := handlers.New(handlers.Config{
		URI:       cfg.HTTP.Path,
		TG:        telegram.New(telegram.Config{Token: cfg.Telegram.Token}),
		Processor: interpreter.New(backend),
	})
	srv := &http.Server{
		Addr:              cfg.HTTP.Listen,
		ReadTimeout:       timeout,
		ReadHeaderTimeout: timeout,
		WriteTimeout:      timeout,
		IdleTimeout:       timeout,
		Handler:           h,
	}

	idleConnsClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		// We received an interrupt signal, shut down.
		if err := srv.Shutdown(context.Background()); err != nil {
			// Error from closing listeners, or context timeout:
			log.Printf("HTTP server Shutdown: %v\n", err)
		}
		close(idleConnsClosed)
	}()

	log.Printf("Listen at %s\n", srv.Addr)
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		return err
	}

	<-idleConnsClosed

	log.Println("Succesfull shutdown")

	return nil
}

func loadConfig() (*config.Config, error) {
	fileName := os.Getenv("CONFIG_PATH")
	if fileName == "" {
		fileName = "./configs/bot.yml"
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

func connectToBackend(cfg config.Backend) (pb.MailAggregatorClient, error) {
	conn, err := grpc.Dial(cfg.Host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return pb.NewMailAggregatorClient(conn), nil
}

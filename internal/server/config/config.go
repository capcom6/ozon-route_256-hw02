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

package config

import (
	"io"
	"os"

	"gopkg.in/yaml.v2"
)

func Load(path string) (*Config, error) {
	if path == "" {
		return FromEnv()
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	cfg, err := Parse(bytes)

	return cfg, err
}

func Parse(fileBytes []byte) (*Config, error) {
	cfg := &Config{}

	if err := yaml.Unmarshal(fileBytes, cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

func FromEnv() (*Config, error) {
	cfg := Config{
		Server: Server{
			GRPC:    os.Getenv("SERVER_GRPC"),
			Gateway: os.Getenv("SERVER_GATEWAY"),
		},
		Database: Database{
			Host:     os.Getenv("DATABASE_HOST"),
			Port:     getIntEnv("DATABASE_PORT", 5432),
			Database: os.Getenv("DATABASE_NAME"),
			User:     os.Getenv("DATABASE_USER"),
			Password: os.Getenv("DATABASE_PASSWORD"),
		},
	}

	return &cfg, nil
}

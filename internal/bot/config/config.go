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
	cfg := &Config{
		Telegram: Telegram{
			Token: os.Getenv("TELEGRAM_TOKEN"),
		},
		HTTP: HTTP{
			Listen: os.Getenv("HTTP_LISTEN"),
			Path:   os.Getenv("HTTP_PATH"),
		},
		Backend: Backend{
			Host: os.Getenv("BACKEND_HOST"),
		},
	}

	return cfg, nil
}

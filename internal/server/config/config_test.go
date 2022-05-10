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
	"os"
	"reflect"
	"testing"
)

var defaultConfig = &Config{
	Server: Server{
		GRPC:    "localhost:8000",
		Gateway: "localhost:8001",
	},
	Database: Database{
		Host:     "db",
		Port:     5432,
		Database: "server",
		User:     "server",
		Password: "server_password",
	},
}

func TestParse(t *testing.T) {
	type args struct {
		fileBytes []byte
	}
	tests := []struct {
		name    string
		args    args
		want    *Config
		wantErr bool
	}{
		{
			name: "Empty config",
			args: args{
				fileBytes: []byte{},
			},
			want:    &Config{},
			wantErr: false,
		},
		{
			name: "Invalid config",
			args: args{
				fileBytes: []byte(`123`),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Valid config",
			args: args{
				fileBytes: []byte(`
server:
    grpc: "localhost:8000"
    gateway: "localhost:8001"
database:
    host: db
    port: 5432
    database: server
    user: server
    password: server_password
`),
			},
			want:    defaultConfig,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.args.fileBytes)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFromEnv(t *testing.T) {
	tests := []struct {
		name    string
		args    map[string]string
		want    *Config
		wantErr bool
	}{
		{
			name: "Empty",
			args: map[string]string{},
			want: &Config{
				Server: Server{},
				Database: Database{
					Host:     "",
					Port:     5432,
					Database: "",
					User:     "",
					Password: "",
				},
			},
			wantErr: false,
		},
		{
			name: "Filled",
			args: map[string]string{
				"SERVER_GRPC":    "localhost:8000",
				"SERVER_GATEWAY": "localhost:8001",

				"DATABASE_HOST":     "db",
				"DATABASE_PORT":     "5432",
				"DATABASE_NAME":     "server",
				"DATABASE_USER":     "server",
				"DATABASE_PASSWORD": "server_password",
			},
			want:    defaultConfig,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Clearenv()
			for k, v := range tt.args {
				os.Setenv(k, v)
			}

			got, err := FromEnv()
			if (err != nil) != tt.wantErr {
				t.Errorf("FromEnv() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromEnv() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLoad(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    *Config
		wantErr bool
	}{
		{
			name: "Non existing config",
			args: args{
				path: "unknown.yml",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Example config",
			args: args{
				path: "../../../configs/server.example.yml",
			},
			want:    defaultConfig,
			wantErr: false,
		},
		{
			name: "Config #1",
			args: args{
				path: "../../../test/data/server/config_01.yml",
			},
			want:    defaultConfig,
			wantErr: false,
		},
		{
			name: "Invalid config",
			args: args{
				path: "../../../test/data/server/config_invalid.yml",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Load(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("Load() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Load() = %v, want %v", got, tt.want)
			}
		})
	}
}

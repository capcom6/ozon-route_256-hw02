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
	"testing"
)

func Test_getIntEnv(t *testing.T) {
	type args struct {
		key string
		def int
	}
	tests := []struct {
		name string
		env  map[string]string
		args args
		want int
	}{
		{
			name: "Default for empty",
			env:  map[string]string{},
			args: args{
				key: "DUMMY",
				def: 5432,
			},
			want: 5432,
		},
		{
			name: "Default for invalid",
			env: map[string]string{
				"DUMMY": "invalid",
			},
			args: args{
				key: "DUMMY",
				def: 5432,
			},
			want: 5432,
		},
		{
			name: "Correct value",
			env: map[string]string{
				"DUMMY": "2345",
			},
			args: args{
				key: "DUMMY",
				def: 5432,
			},
			want: 2345,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Clearenv()
			for k, v := range tt.env {
				os.Setenv(k, v)
			}

			if got := getIntEnv(tt.args.key, tt.args.def); got != tt.want {
				t.Errorf("getIntEnv() = %v, want %v", got, tt.want)
			}
		})
	}
}

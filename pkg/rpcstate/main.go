package rpcstate

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"tenon/pkg/rpcc"

	"go.yaml.in/yaml/v3"
)

type Task struct {
	Id     string     `json:"id"`
	Config TaskConfig `json:"config"`
}

type TaskConfig struct {
	Name        string  `json:"name" yaml:"name"`
	Description *string `json:"description" yaml:"description"`
	Script      string  `json:"script" yaml:"script"`
}

type RPCState struct {
}

func New() *RPCState {
	return &RPCState{}
}
func (self *RPCState) Register(rpcc *rpcc.RPConnection) {
	rpcc.RegisterFn("get_all_task", func() ([]Task, error) {
		tasks := []Task{}
		filepath.WalkDir("tasks", func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if filepath.Ext(path) == ".yaml" {
				data, err := os.ReadFile(path)
				if err != nil {
					return err
				}
				config := TaskConfig{}
				if err := yaml.Unmarshal(data, &config); err != nil {
					return err
				}
				path, err := filepath.Rel("tasks", path)
				if err != nil {
					return err
				}
				tasks = append(tasks, Task{Id: strings.TrimSuffix(path, filepath.Ext(path)), Config: config})
			}
			return nil
		})
		return tasks, nil
	})
}

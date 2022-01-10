package configs

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
)

type ProjectName string

type Config struct {
	ProjectName ProjectName
	RootPath    string
	Raw         []byte
	Debug       bool
	Verbose     bool // if true, prompt enter to exit.
	LogName     string
	API         struct {
		GRPC api `json:"grpc"`
		HTTP api `json:"http"`
	} `json:"api"`
	Web struct {
		Addr string `json:"Addr"`
		Tmpl string `json:"Tmpl"`
	} `json:"web"`
	Err error
}

type api struct {
	Network, Addr, Timeout string
}

func NewConfig(projectName ProjectName) *Config {
	return setRootPath(&Config{ProjectName: projectName}).load()
}

func setRootPath(cfg *Config) *Config {
	cfg.RootPath, cfg.Err = os.Getwd()
	if cfg.Err != nil {
		return cfg
	}
	ps := strings.Split(cfg.RootPath, string(cfg.ProjectName))
	n := 0
	if len(ps) > 1 {
		n = strings.Count(ps[1], string(os.PathSeparator))
	}
	for i := 0; i < n; i++ {
		cfg.RootPath = filepath.Join(
			cfg.RootPath, ".."+string(os.PathSeparator))
	}
	cfg.RootPath = filepath.FromSlash(cfg.RootPath)
	return cfg
}

func (c *Config) load() *Config {
	if c.Err != nil {
		return c
	}
	cfgFile := filepath.Join(c.RootPath, "configs", "configs.json")
	c.Raw, c.Err = os.ReadFile(cfgFile)
	if c.Err != nil {
		if errors.Is(c.Err, os.ErrNotExist) {
			c.Err = errors.WithMessage(c.Err, "ReadFile error: no configs.json")
		}
		return c
	}
	if c.Err = json.Unmarshal(c.Raw, c); c.Err != nil {
		c.Err = errors.WithMessage(c.Err, "Unmarshal configs.json error")
		return c
	}
	return c
}

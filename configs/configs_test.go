package configs

import (
	"fmt"
	"testing"
)

func TestLoad(t *testing.T) {
	cfg := &Config{ProjectName: "hfcms-articles"}
	if err := setRootPath(cfg).load().Err; err != nil {
		t.Error(cfg.Err)
		return
	}
	fmt.Println(cfg)
}

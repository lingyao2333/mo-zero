package service

import (
	"testing"

	"github.com/lingyao2333/mo-zero/core/logx"
)

func TestServiceConf(t *testing.T) {
	c := ServiceConf{
		Name: "foo",
		Log: logx.LogConf{
			Mode: "console",
		},
		Mode: "dev",
	}
	c.MustSetUp()
}

package tools

import (
	"encoding/json"

	"github.com/sirupsen/logrus"
)

type PrettyLog struct {
	Logger *logrus.Logger
}

func (pl *PrettyLog) Log(prefix string, obj interface{}) {
	b, _ := json.MarshalIndent(obj, "", "  ")
	pl.Logger.Infof("%s:\n%s", prefix, string(b))
}

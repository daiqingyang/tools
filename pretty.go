package tools

import (
	"encoding/json"

	"github.com/sirupsen/logrus"
)

type PrettyLog struct {
	Logger logrus.Logger
}

func (pl *PrettyLog) Log(obj interface{}) {
	b, _ := json.MarshalIndent(obj, "", "  ")
	pl.Logger.Infof("object:\n%s", string(b))
}

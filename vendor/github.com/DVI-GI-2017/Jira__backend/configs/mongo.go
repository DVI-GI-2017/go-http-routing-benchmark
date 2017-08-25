package configs

import (
	"fmt"
	"os"
)

type Mongo struct {
	Server string `json:"server"`
	Host   string `json:"host"`
	Port   int    `json:"port"`
	DB     string `json:"db"`
	Drop   bool   `json:"drop"`
}

func (m *Mongo) URL() (url string) {
	url = os.Getenv("MONGO_PROD")

	if url == "" {
		url = fmt.Sprintf("%s://%s:%d", m.Server, m.Host, m.Port)
	}

	return
}

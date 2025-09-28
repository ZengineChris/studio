package infrastructure

import (
	"database/sql"
	"sync"

	"github.com/zenginechris/studio/config"
)

var once sync.Once

type single = struct {
	client sql.Conn
}

var instance *single


func GetClient(config.Database) *sql.Conn {
	once.Do(func(){
		

	})

	return &instance.client
}

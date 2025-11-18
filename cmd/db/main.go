package main

import (
	"loadept.com/pg-mcp/internal/config"
)

func init() {
	config.LoadEnvs()
}

func main() {
	pg, err := config.NewDBPostgres()
	if err != nil {
		panic(err)
	}
	defer pg.Close()
}

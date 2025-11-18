package main

import (
	"flag"

	"loadept.com/pg-mcp/internal/config"
)

func init() {
	config.LoadEnvs()
}

func main() {
	uri := flag.String("u", "", "PostgreSQL connection URI")
	flag.Parse()

	pg, err := config.NewDBPostgres(*uri)
	if err != nil {
		panic(err)
	}
	defer pg.Close()
}

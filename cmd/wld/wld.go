package main

import (
	"os"

	"github.com/warlock-org/server/pkg/log"

	"github.com/spf13/cobra"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var flags struct {
	Port int
	Sock string
	db   *gorm.DB
}

var cmd = cobra.Command{
	Use:     "wld",
	Short:   "warlock daemon server",
	Example: "wld",
	Run:     run,
}

func init() {
	var (
		p      = cmd.PersistentFlags()
		dbfile string
		err    error
	)

	p.IntVarP(&flags.Port, "port", "p", 10010, "server port")
	p.StringVarP(&flags.Sock, "sock", "", "/var/run/wld.sock", "sock file path")
	p.StringVarP(&dbfile, "db", "", "/var/lib/wld/wld.db", "sqlite db file")

	if flags.db, err = gorm.Open(sqlite.Open(dbfile), &gorm.Config{
		Logger: log.GormLogger(),
	}); nil != err {
		log.Error(err)
		os.Exit(1)
	}
}

func main() {
	cmd.Execute()
}

func run(cmd *cobra.Command, _ []string) {
	log.Debugf("wld flags: %+v", flags)
}

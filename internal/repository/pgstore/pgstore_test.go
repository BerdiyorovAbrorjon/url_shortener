package pgstore

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/BerdiyorovAbrorjon/url-shortener/config"
	"github.com/BerdiyorovAbrorjon/url-shortener/pkg/database/postgres"
)

var testPgStore *PgStore

func TestMain(m *testing.M) {

	cfg, err := config.NewConfig("../../..")
	if err != nil {
		log.Fatal("test - PgStore - config.NewConfig: %w", err)
	}

	dbPool, err := postgres.NewPool(cfg.DbSource)
	if err != nil {
		log.Fatal(fmt.Errorf("test - PgStore - postgres.NewPool: %w", err))
	}

	testPgStore = NewPgStore(dbPool)

	os.Exit(m.Run())
}

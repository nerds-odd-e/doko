package services

import (
	"os"
	"testing"

	"github.com/go-pg/pg/v10"
	"tdd.com/v1/utils"
)

var db *pg.DB
var tx *pg.Tx

func TestMain(m *testing.M) {
	db = utils.GetTestPGDB()
	exitCode := m.Run()
	cleanup()
	os.Exit(exitCode)
}

func cleanup() {
	if utils.Teardown != nil {
		defer utils.Teardown()
	}
}

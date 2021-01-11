package main

import (
	"log"
	"net/http"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/vkuznet/dbs2go/dbs"
)

// TestDataTypesSQL API
func TestDataTypesSQL(t *testing.T) {

	// initialize DB for testing
	db := initDB(false)
	defer db.Close()

	// fetch this record from DB, here we can either use nil writer
	// or use StdoutWriter instance (defined in test/main.go)
	params := make(dbs.Record)
	var w http.ResponseWriter
	w = StdoutWriter("")
	log.Println("Test DataTypes API")
	var api dbs.API
	_, err := api.DataTypes(params, w)
	if err != nil {
		t.Errorf("Fail to look-up data tiers %v\n", err)
	}
}
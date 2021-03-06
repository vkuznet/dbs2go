package dbs

import (
	"errors"
	"net/http"
)

// BlockOrigin DBS API
func (API) BlockOrigin(params Record, w http.ResponseWriter) (int64, error) {
	// variables we'll use in where clause
	var args []interface{}
	var conds []string

	// parse given parameters
	site := getValues(params, "origin_site_name")
	if len(site) > 1 {
		msg := "Unsupported list of sites"
		return 0, errors.New(msg)
	} else if len(site) == 1 {
		conds, args = AddParam("origin_site_name", "B.ORIGIN_SITE_NAME", params, conds, args)
	}
	block := getValues(params, "block_name")
	if len(block) > 1 {
		msg := "Unsupported list of block"
		return 0, errors.New(msg)
	} else if len(block) == 1 {
		conds, args = AddParam("block_name", "B.BLOCK_NAME", params, conds, args)
	}
	dataset := getValues(params, "dataset")
	if len(dataset) > 1 {
		msg := "Unsupported list of dataset"
		return 0, errors.New(msg)
	} else if len(dataset) == 1 {
		conds, args = AddParam("dataset", "DS.DATASET", params, conds, args)
	}

	// get SQL statement from static area
	stm := getSQL("blockorigin")
	stm = WhereClause(stm, conds)

	// use generic query API to fetch the results from DB
	return executeAll(w, stm, args...)
}

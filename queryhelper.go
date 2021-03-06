package main

import (
	"github.com/couchbase/gocb"
)

func GetViewQuery(dname, vname string, parameters ViewQueryParameters) *gocb.ViewQuery {
	viewquery := gocb.NewViewQuery(dname, vname)

	if parameters.Limit > 0 {
		viewquery = viewquery.Limit(parameters.Limit)
	}

	if parameters.Stale == false {
		viewquery = viewquery.Stale(gocb.Before)
	} else if parameters.Stale == true {
		viewquery = viewquery.Stale(gocb.None)
	}

	if parameters.UpdateAfter == true {
		viewquery = viewquery.Stale(gocb.After)
	}

	if parameters.Skip > 0 {
		viewquery = viewquery.Skip(parameters.Skip)
	}

	return viewquery
}

func processResults(viewresults gocb.ViewResults) error {
	var val interface{}
	for {
		success := viewresults.Next(&val)
		if success == false {
			err := viewresults.Close()
			return err
		}
	}
}

func GetN1QLQuery(statement string, scanconsistency string) *gocb.N1qlQuery {
	n1qlQuery := gocb.NewN1qlQuery(statement)
	n1qlQuery.AdHoc(false)
	if scanconsistency == "not_bounded" {
		n1qlQuery.Consistency(gocb.NotBounded)
	} else if scanconsistency == "request_plus" {
		n1qlQuery.Consistency(gocb.RequestPlus)
	} else if scanconsistency == "statement_plus" {
		n1qlQuery.Consistency(gocb.StatementPlus)
	}
	return n1qlQuery
}

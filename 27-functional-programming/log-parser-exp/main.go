// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"fmt"
	"os"
)

func main() {
	defer recoverErr()

	_, err := newReport().
		from(os.Stdin).
		to(os.Stdout).
		retrieveFrom(textReader).
		filterBy(orgDomainsFilter).
		// filterBy(notUsing(domainExtFilter("org", "io"))).
		// groupBy(pageGrouper).
		groupBy(domainGrouper).
		writeTo(textWriter).
		run()

	if err != nil {
		fmt.Println("> Err:", err)
	}
}

func recoverErr() {
	val := recover()

	if val == nil {
		return
	}

	if err, ok := val.(string); ok {
		fmt.Println("> Error occurred:", err)
	}
}

/*
newReport -> report.New().
Result    -> report.Line

notUsing = report.Not

_, err := stats.Parse().
	From(os.Stdin).
	To(os.Stdout).
	RetrieveFrom(stats.TextReader).
	FilterBy(notUsing(stats.DomainExtFilter("org", "io"))).
	GroupBy(stats.DomainGrouper).
	WriteTo(stats.TextWriter).
	Run()
*/
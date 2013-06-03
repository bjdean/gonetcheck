package gonetcheck

/*********************************************************************
 * gonetcheck - Go package to check general network health
 *
 * func: SetUrlList
 *
 * Copyright 2013 Bradley Dean
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http:www.gnu.org/licenses/>.
 */

import (
	"net/url"
)

// A set of default test URLs used to check
// remote network access
var testUrls = []string{
	"http://www.google.com/",
	"http://www.bing.com/",
	"http://www.microsoft.com/",
	"http://www.apple.com/",
	"http://yahoo.com/",
	"http://www.unimelb.edu.au/",
	"http://abc.net.au/",
	"http://www.monash.edu.au/",
	"http://bbb.co.uk/",
}

// Change the test URL list
func GetUrlList() []string {
	return testUrls
}

// Update the test URL list
// Valide that the array of strings contains parseable URLs
func SetUrlList(newUrlList []string) []error {
	var errList []error
	for _, urlString := range newUrlList {
		_, err := url.Parse(urlString)
		if err != nil {
			errList = append(errList, err)
		}
	}
	if len(errList) > 0 {
		return errList
	} else {
		testUrls = newUrlList
		return nil
	}
}

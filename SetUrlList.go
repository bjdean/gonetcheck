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
var test_urls = []string{
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
	return test_urls
}

// Update the test URL list
// Valide that the array of strings contains parseable URLs
func SetUrlList(new_url_list []string) []error {
	var err_list []error
	for _, url_string := range new_url_list {
		_, err := url.Parse(url_string)
		if err != nil {
			err_list = append(err_list, err)
		}
	}
	if len(err_list) > 0 {
		return err_list
	} else {
		test_urls = new_url_list
		return nil
	}
}

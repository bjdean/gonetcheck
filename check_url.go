package gonetcheck

/*********************************************************************
 * gonetcheck - Go package to check general network health
 *
 * func: check_url
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
	"net/http"
)

// Status of a URL check
type UrlStat struct {
	Url string
	ResponseCode int
	StatusLine string
	Error error
}

// Make a single URL request and pass a UrlStat into
// the out_queue channel
func check_url (url string, out_queue chan UrlStat) {
	debug_log(DBG_MEDIUM, "check_url: ", url)
	resp, err := http.Head(url)
	if err == nil {
		debug_log(DBG_VERBOSE, "check_url suceess:", resp)
		out_queue <- UrlStat{
			Url: url,
			ResponseCode: resp.StatusCode,
			StatusLine: resp.Status,
		}
	} else {
		debug_log(DBG_VERBOSE, "check_url error: ", err)
		out_queue <- UrlStat{
			Url: url,
			Error: err,
		}
	}
}

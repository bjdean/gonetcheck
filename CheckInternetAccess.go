package gonetcheck

/*********************************************************************
 * gonetcheck - Go package to check general network health
 *
 * func: CheckInternetAccess
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
	"time"
)

// Determine if it looks like this server has access
// to the internet (ie remote servers)
// TODO Implement error checks and return errors
func CheckInternetAccess () (bool, error) {
	// This entire function has a timeout starting
	// when the function is called
	timeout_chan := time.After(2 * time.Second)
	out_queue := make(chan UrlStat)
	for _, url := range test_urls {
		go check_url(url, out_queue)
	}

	var stats []UrlStat
	StatLoop: for {
		select {
			case s := <-out_queue:
				stats = append(stats, s)
				if len(stats) == len(test_urls) {
					break StatLoop
				}
			case <-timeout_chan:
				break StatLoop
		}
	}

	// Check results
	test_count := len(test_urls)
	var success_count int
	for _, stat := range stats {
		debug_log(DBG_MEDIUM, stat)
		if ( stat.Error == nil &&
			stat.ResponseCode < 400 ) {
			success_count += 1
		}
	}
	up_fraction := float32(success_count)/float32(test_count)
	var network_is_up bool
	if up_fraction >= 0.5 {
		network_is_up = true
	}
	debug_log(DBG_QUIET, "Sites up fraction:", up_fraction)
	debug_log(DBG_QUIET, "Network is up:", network_is_up)

	// Return true if network_is_up
	return network_is_up, nil
}

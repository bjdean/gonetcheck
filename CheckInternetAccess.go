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

// The final result to be put on the final_result_chan
type finalResult struct {
	NetworkIsUp bool
	Errors []error
}

// Determine if it looks like this server has access
// to the internet (ie remote servers)
func CheckInternetAccess() (bool, []error) {
	// This entire function has a timeout starting
	// when the function is called
	timeout_chan := time.After(10 * time.Second)

	// All checking goroutines:
	// 1. Register thei existence (ie number of checks) by dropping an int
	//    onto check_count_chan channel
	// 2. Drop either a result or an error onto the result_chan or
	//    error_chan channels
	// result_chan and error_chan channels are buffered to allow the check
	// goroutines to be cleaned up
	check_count_chan := make(chan int, 100)
	result_chan := make(chan bool, 100)
	error_chan := make(chan error, 100)

	// Finally the end-result will be placed on these channels
	final_result_chan := make(chan finalResult)

	// The final_result_check channel acculates and finally
	// calculates the final result
	go final_result_check(
		timeout_chan,
		check_count_chan,
		result_chan,
		error_chan,
		final_result_chan)

	// Run checking goroutines
	go run_url_checks(check_count_chan, result_chan, error_chan)

	// Block until the final_result_chan receives a value
	final_result := <-final_result_chan
	return final_result.NetworkIsUp, final_result.Errors
}

// Run all check_url checks and transpose UrlStat responses into
// the result/error channels
func run_url_checks(
	check_count_chan chan int,
	result_chan chan bool,
	error_chan chan error) {

	check_url_chan := make(chan UrlStat, 100)

	// Launch checks
	for _, url := range test_urls {
		go check_url(url, check_url_chan)
		check_count_chan <- 1
	}

	// Process results into the result_chan
	for {
		stat := <-check_url_chan
		switch stat.Error {
		case nil:
			if stat.ResponseCode < 400 {
				result_chan <- true
			}
		default:
			error_chan <- stat.Error
		}
	}
}

// Collate errors and results from the result_chan and error_chan until the
// timeout_chan fires. Once this occurs calculate the final result and
// place it on the final_result_chan
func final_result_check(
	timeout_chan <-chan time.Time,
	check_count_chan chan int,
	result_chan chan bool,
	error_chan chan error,
	final_result_chan chan finalResult) {

	// Accumulators
	var check_count int
	var success_count int
	var fail_count int
	var errors []error

AccumulatorLoop:
	for {
		select {
		case count := <-check_count_chan:
			check_count += count
		case result := <-result_chan:
			switch result {
				case true: success_count += 1
				default: fail_count += 1
			}
			// If all checks are in, break
			if success_count + fail_count + len(errors) >= check_count {
				break AccumulatorLoop
			}
		case err := <-error_chan:
			errors = append(errors, err)
			// If all checks are in, break
			if success_count + fail_count + len(errors) >= check_count {
				break AccumulatorLoop
			}
		case <-timeout_chan:
			break AccumulatorLoop
		}
		debug_log(
			DBG_VERBOSE,
			"CheckCount = ", check_count, ";",
			"SuccessCount =", success_count, ";",
			"FailCount = ", fail_count, ";",
			"Errors = ", errors, ";",
		)
	}

	// Calculate the final result
	switch errors {
	case nil:
		up_fraction := float32(success_count) / float32(check_count)
		final_result_chan <- finalResult{ up_fraction >= 0.5, nil }
	default:
		final_result_chan <- finalResult{ false, errors }
	}
}

/*************************************************************
 * refactoring:

	// Check results
	test_count := len(test_urls)
	var success_count int
	var err_list []error
	for _, stat := range stats {
		debug_log(DBG_MEDIUM, stat)
		switch stat.Error {
		case nil:
			if stat.ResponseCode < 400 {
				success_count += 1
			}
		default:
			err_list = append(err_list, stat.Error)
		}
	}

	// Calculate network_is_up
	up_fraction := float32(success_count) / float32(test_count)
	var network_is_up bool
	if up_fraction >= 0.5 {
		network_is_up = true
	}
	debug_log(DBG_QUIET, "Sites up fraction:", up_fraction)
	debug_log(DBG_QUIET, "Network is up:", network_is_up)

	// Return true if network_is_up
	return network_is_up, err_list
}


*********************************************************************/

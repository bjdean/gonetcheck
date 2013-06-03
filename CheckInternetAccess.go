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

// The final result to be put on the finalResultChan
type finalResult struct {
	NetworkIsUp bool
	Errors      []error
}

// Determine if it looks like this server has access
// to the internet (ie remote servers)
func CheckInternetAccess(testUrls, testTcpAddrs []string) (bool, []error) {
	// This entire function has a timeout starting
	// when the function is called
	timeoutChan := time.After(10 * time.Second)

	// All checking goroutines:
	// 1. Register thei existence (ie number of checks) by dropping an int
	//    onto checkCountChan channel
	// 2. Drop either a result or an error onto the resultChan or
	//    errorChan channels
	// resultChan and errorChan channels are buffered to allow the check
	// goroutines to be cleaned up
	checkCountChan := make(chan int, 100)
	resultChan := make(chan bool, 100)
	errorChan := make(chan error, 100)

	// Finally the end-result will be placed on these channels
	finalResultChan := make(chan finalResult)

	// The finalResultCheck channel acculates and finally
	// calculates the final result
	go finalResultCheck(
		timeoutChan,
		checkCountChan,
		resultChan,
		errorChan,
		finalResultChan)

	// Run checking goroutines
	go runUrlChecks(testUrls, checkCountChan, resultChan, errorChan)
	go runTcpChecks(testTcpAddrs, checkCountChan, resultChan, errorChan)

	// Block until the finalResultChan receives a value
	finalResult := <-finalResultChan
	return finalResult.NetworkIsUp, finalResult.Errors
}

// Run all checkUrl checks and transpose UrlStat responses into
// the result/error channels
func runUrlChecks(
	testUrls []string,
	checkCountChan chan int,
	resultChan chan bool,
	errorChan chan error) {

	checkUrlChan := make(chan UrlStat, 100)

	// Launch checks
	for _, url := range testUrls {
		go checkUrl(url, checkUrlChan)
		checkCountChan <- 1
	}

	// Process results into the resultChan
	for {
		stat := <-checkUrlChan
		switch stat.Error {
		case nil:
			if stat.ResponseCode < 400 {
				resultChan <- true
			}
		default:
			errorChan <- stat.Error
		}
	}
}

func runTcpChecks(
	testTcpAddrs []string,
	checkCountChan chan int,
	resultChan chan bool,
	errorChan chan error) {

	checkTcpChan := make(chan error, 100)

	// Launch checks
	for _, addr := range testTcpAddrs {
		go checkTcp(addr, checkTcpChan)
		checkCountChan <- 1
	}

	// Process results into the resultChan
	for {
		err := <-checkTcpChan
		switch err {
		case nil:
			resultChan <- true
		default:
			errorChan <- err
		}
	}
}

// Collate errors and results from the resultChan and errorChan until the
// timeoutChan fires. Once this occurs calculate the final result and
// place it on the finalResultChan
func finalResultCheck(
	timeoutChan <-chan time.Time,
	checkCountChan chan int,
	resultChan chan bool,
	errorChan chan error,
	finalResultChan chan finalResult) {

	// Accumulators
	var checkCount int
	var successCount int
	var failCount int
	var errors []error

AccumulatorLoop:
	for {
		select {
		case count := <-checkCountChan:
			checkCount += count
		case result := <-resultChan:
			switch result {
			case true:
				successCount += 1
			default:
				failCount += 1
			}
			// If all checks are in, break
			if successCount+failCount+len(errors) >= checkCount {
				break AccumulatorLoop
			}
		case err := <-errorChan:
			errors = append(errors, err)
			// If all checks are in, break
			if successCount+failCount+len(errors) >= checkCount {
				break AccumulatorLoop
			}
		case <-timeoutChan:
			break AccumulatorLoop
		}
		debugLog(
			DBG_VERBOSE,
			"CheckCount = ", checkCount, ";",
			"SuccessCount =", successCount, ";",
			"FailCount = ", failCount, ";",
			"Errors = ", errors, ";",
		)
	}

	// Calculate the final result
	switch errors {
	case nil:
		upFraction := float32(successCount) / float32(checkCount)
		finalResultChan <- finalResult{upFraction >= 0.5, nil}
	default:
		finalResultChan <- finalResult{false, errors}
	}
}

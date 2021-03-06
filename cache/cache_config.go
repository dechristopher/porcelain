/**
 * MIT License
 *
 * Copyright (c) 2020 Andrew DeChristopher
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package cache

import (
	"log"
	"strconv"
	"strings"
)

var (
	ZoomRangeMin int
	ZoomRangeMax int
)

// ParseZoomRange parses the min and max zoom level configuration values
// from the --range flag on startup
func ParseZoomRange(zoomRange string) {
	minmax := strings.Split(zoomRange, "-")
	if len(minmax) != 2 {
		log.Fatal("Invalid zoom range format. Expected: `X-X`. Example: `0-12`")
	}

	min, err := strconv.Atoi(minmax[0])
	max, err := strconv.Atoi(minmax[1])

	if err != nil {
		log.Fatalf("Invalid zoom range format. Expected integer zoom levels.")
	}

	ZoomRangeMin = min
	ZoomRangeMax = max
}

// IsZoomServed returns whether or not the cache is configured to
// serve tiles at the specified zoom level
func IsZoomServed(zoomLevel int, zoomLevelString ...string) bool {
	if len(zoomLevelString) > 0 {
		zl, err := strconv.Atoi(zoomLevelString[0])
		if err != nil {
			return false
		}
		return evalIsZoomServed(zl)
	}
	return evalIsZoomServed(zoomLevel)
}

// evalIsZoomServed performs the within range check to see if a given
// zoom level is served by the cache
func evalIsZoomServed(zoomLevel int) bool {
	return zoomLevel >= ZoomRangeMin && zoomLevel <= ZoomRangeMax
}

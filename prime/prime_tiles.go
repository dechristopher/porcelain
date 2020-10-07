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

package prime

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/dechristopher/porcelain/cache"
)

// Tiles will prime the tile cache by walking the specified directory
func Tiles(directory string) error {
	_, errCheck := ioutil.ReadDir(directory)
	if errCheck != nil {
		return errCheck
	}

	fmt.Println("Priming cache...")

	numCached := 0

	// walk the path looking for numbered PNGs
	errWalk := filepath.Walk(directory,
		func(path string, info os.FileInfo, err error) error {
			if !info.IsDir() {
				if !strings.Contains(path, ".png") {
					return nil
				}

				key := strings.TrimLeft(strings.TrimRight(path, ".png"),
					directory+"/")

				zl := strings.Split(key, "/")[0]
				zli, err := strconv.Atoi(zl)
				if err != nil {
					return err
				}

				if !cache.IsZoomServed(zli) {
					return nil
				}

				tile, err := ioutil.ReadFile(path)
				if err != nil {
					return err
				}
				err = cache.Cache.Set(key, tile)
				numCached++
				return err
			}
			return nil
		})

	fmt.Printf("Cached %d/%d tiles\n", cache.Cache.Len(), numCached)

	return errWalk
}

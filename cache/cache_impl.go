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

	"github.com/allegro/bigcache/v2"
)

var Cache *bigcache.BigCache

func init() {
	Cache = boot()
}

// Boot will boot a new bigcache for tile storage in memory
func boot() *bigcache.BigCache {
	cache, err := bigcache.NewBigCache(bigcache.Config{
		Shards:           1024 * 64,
		LifeWindow:       -1,
		CleanWindow:      -1,
		MaxEntrySize:     1024 * 1024,
		StatsEnabled:     true,
		Verbose:          true,
		HardMaxCacheSize: 8192,
	})
	if err != nil {
		log.Fatalf("Failed to boot cache, %s", err.Error())
	}
	errSet := cache.Set("test", []byte("test"))
	errDel := cache.Delete("test")
	if errSet != nil {
		log.Fatalf("Failed to boot cache, %s", errSet.Error())
	}
	if errDel != nil {
		log.Fatalf("Failed to boot cache, %s", errDel.Error())
	}
	return cache
}

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

package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/logger"
	recover2 "github.com/gofiber/fiber/v2/middleware/recover"

	"github.com/dechristopher/porcelain/cache"
	"github.com/dechristopher/porcelain/handle"
	"github.com/dechristopher/porcelain/prime"
)

const version = "v0.0.2"

// init starts porcelain's caches and primers
func init() {
	fmt.Println("porcelain " + version)

	tilesDirectory := flag.String("tiles", "tiles",
		"Specify the directory containing rendered tiles")
	zoomRange := flag.String("range", "0-18",
		"Specify the range of zoom levels this cache serves")
	flag.Parse()

	cache.ParseZoomRange(*zoomRange)

	err := prime.Tiles(*tilesDirectory)
	if err != nil {
		log.Fatalf("Failed to prime, %s", err.Error())
	}

	fmt.Printf("Serving zoom levels %s from directory: %s",
		*zoomRange, *tilesDirectory)
}

// main starts the fiber webserver
func main() {
	app := fiber.New(fiber.Config{
		Prefork:       false,
		StrictRouting: false,
		CaseSensitive: false,
		Immutable:     false,
		BodyLimit:     1024 * 1024,
	})

	// recover from panics
	app.Use(recover2.New())

	// compress and optimize for speed
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))

	// allow requests from all origins
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET",
	}))

	// ignore favicon logging
	app.Use(favicon.New())

	// log all requests
	app.Use(logger.New())

	// handle tile requests
	app.Get("/:z/:x/:y.png", handle.Tile)

	// boot porcelain server
	log.Fatal(app.Listen(":2600"))
}

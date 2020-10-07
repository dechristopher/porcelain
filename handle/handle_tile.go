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

package handle

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"

	"github.com/dechristopher/porcelain/cache"
)

// Tile returns a tile hot from cache if it exists
func Tile(ctx *fiber.Ctx) error {
	xyz := fmt.Sprintf("%s/%s/%s",
		ctx.Params("z"), ctx.Params("x"), ctx.Params("y"))
	tile, err := cache.Cache.Get(xyz)
	if err != nil {
		ctx.Status(500)
		return err
	}
	if len(tile) == 0 {
		ctx.Status(404)
		return errors.New("no tile found for request: %s" + xyz)
	}
	ctx.Response().Header.Set("Content-Type", "image/png")
	ctx.Status(200)
	err = ctx.Send(tile)
	return err
}

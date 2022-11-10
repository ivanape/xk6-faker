// MIT License
//
// # Copyright (c) 2022 Radoslav Salov
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package faker

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/dop251/goja"
	"go.k6.io/k6/js/modules"
	"lukechampine.com/frand"
)

type Faker struct {
	vu modules.VU
	*gofakeit.Faker
}

func (f *Faker) Ipv4Address() string {
	return f.IPv4Address()
}

func (f *Faker) Ipv6Address() string {
	return f.IPv6Address()
}

func (f *Faker) HttpStatusCodeSimple() int {
	return f.HTTPStatusCodeSimple()
}

func (f *Faker) HttpStatusCode() int {
	return f.HTTPStatusCode()
}

func (f *Faker) HttpMethod() string {
	return f.HTTPMethod()
}

func (f *Faker) Bs() string {
	return f.BS()
}

func (f *Faker) Uuid() string {
	return f.UUID()
}

func (f *Faker) RgbColor() []int {
	return f.RGBColor()
}

func (f *Faker) ImageJpeg(width int, height int) goja.ArrayBuffer {
	return f.vu.Runtime().NewArrayBuffer(f.Faker.ImageJpeg(width, height))
}

func (f *Faker) ImagePng(width int, height int) goja.ArrayBuffer {
	return f.vu.Runtime().NewArrayBuffer(f.Faker.ImagePng(width, height))
}

func newFaker(vu modules.VU, seed int64) *Faker {
	src := frand.NewSource()

	if seed != 0 {
		src.Seed(seed)
	}

	return &Faker{vu: vu, Faker: gofakeit.NewCustom(src)}
}

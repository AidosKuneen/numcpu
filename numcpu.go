// Copyright (c) 2018 Aidos Developer

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package numcpu

import (
	"os"
	"runtime"
	"strconv"
	"strings"
)

// NumCPU returns the number of logical CPUs usable by the current process.
//Android doesn't return max CPU by runtime.NumCPU(), but just returns number of working CPUs.
//So this checks sysfs directly for linux.
func NumCPU() int {
	dir := "/sys/devices/system/cpu/"
	f, err := os.Open(dir)
	if err != nil {
		return runtime.NumCPU()
	}
	list, err := f.Readdir(-1)
	if err = f.Close(); err != nil {
		panic(err)
	}
	if err != nil {
		return runtime.NumCPU()
	}
	n := 0
	for _, l := range list {
		name := l.Name()
		if l.IsDir() && strings.HasPrefix(name, "cpu") {
			if _, err := strconv.Atoi(name[3:]); err == nil {
				n++
			}
		}
	}
	if n == 0 {
		return runtime.NumCPU()
	}
	return n
}

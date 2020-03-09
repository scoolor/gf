// Copyright 2017 gf Author(https://github.com/gogf/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gcompress_test

import (
	"testing"

	"github.com/gogf/gf/encoding/gcompress"
	"github.com/gogf/gf/test/gtest"
)

func Test_Zlib_UnZlib(t *testing.T) {
	gtest.Case(t, func() {
		src := "hello, world\n"
		dst := []byte{120, 156, 202, 72, 205, 201, 201, 215, 81, 40, 207, 47, 202, 73, 225, 2, 4, 0, 0, 255, 255, 33, 231, 4, 147}
		data, _ := gcompress.Zlib([]byte(src))
		gtest.Assert(data, dst)

		data, _ = gcompress.UnZlib(dst)
		gtest.Assert(data, []byte(src))

		data, _ = gcompress.Zlib(nil)
		gtest.Assert(data, nil)
		data, _ = gcompress.UnZlib(nil)
		gtest.Assert(data, nil)

		data, _ = gcompress.UnZlib(dst[1:])
		gtest.Assert(data, nil)
	})
}
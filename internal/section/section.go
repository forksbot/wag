// Copyright (c) 2018 Timo Savola. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package section

import (
	"io"
	"io/ioutil"

	"github.com/tsavola/wag/internal/loader"
	"github.com/tsavola/wag/internal/module"
	"github.com/tsavola/wag/internal/reader"
)

func Find(
	findID module.SectionID,
	load loader.L,
	sectionMapper func(sectionID byte, r reader.R) (payloadLen uint32, err error),
	customLoader func(reader.R, uint32) error,
) module.SectionID {
	for {
		sectionID, err := load.R.ReadByte()
		if err != nil {
			if err == io.EOF {
				return 0
			}
			panic(err)
		}

		id := module.SectionID(sectionID)

		switch {
		case id == module.SectionCustom:
			var payloadLen uint32

			if sectionMapper != nil {
				payloadLen, err = sectionMapper(sectionID, load.R)
				if err != nil {
					panic(err)
				}
			} else {
				payloadLen = load.Varuint32()
			}

			if customLoader != nil {
				err = customLoader(load.R, payloadLen)
			} else {
				_, err = io.CopyN(ioutil.Discard, load.R, int64(payloadLen))
			}
			if err != nil {
				panic(err)
			}

		case id == findID:
			return id

		default:
			load.R.UnreadByte()
			return id
		}
	}
}

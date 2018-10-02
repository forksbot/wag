// Copyright (c) 2016 Timo Savola. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package x86

import (
	"github.com/tsavola/wag/internal/gen"
	"github.com/tsavola/wag/internal/gen/link"
	"github.com/tsavola/wag/internal/gen/operand"
	"github.com/tsavola/wag/internal/gen/reg"
	"github.com/tsavola/wag/internal/gen/rodata"
	"github.com/tsavola/wag/internal/isa/prop"
	"github.com/tsavola/wag/internal/isa/x86/in"
	"github.com/tsavola/wag/wa"
)

// Convert may allocate registers, use RegResult and update condition flags.
// The source operand may be RegResult or condition flags.
func (MacroAssembler) Convert(f *gen.Func, props uint16, resultType wa.Type, source operand.O) (result operand.O) {
	switch props {
	case prop.ExtendS:
		r, _ := allocResultReg(f, source)
		in.MOVSXD.RegReg(&f.Text, wa.I64, r, r)
		result = operand.Reg(resultType, r, false)

	case prop.ExtendU:
		r, zeroExt := allocResultReg(f, source)
		if !zeroExt {
			in.MOV.RegReg(&f.Text, wa.I32, r, r)
		}
		result = operand.Reg(resultType, r, false)

	case prop.Mote:
		r, _ := allocResultReg(f, source)
		in.CVTS2SSD.RegReg(&f.Text, source.Type, r, r)
		result = operand.Reg(resultType, r, false)

	case prop.TruncS:
		sourceReg, _ := getScratchReg(f, source)
		resultReg := f.Regs.AllocResult(resultType)
		in.CVTTSSD2SI.TypeRegReg(&f.Text, source.Type, resultType, resultReg, sourceReg)
		f.Regs.Free(source.Type, sourceReg)
		result = operand.Reg(resultType, resultReg, true)

	case prop.TruncU:
		sourceReg, _ := getScratchReg(f, source)
		resultReg := f.Regs.AllocResult(resultType)
		if resultType == wa.I32 {
			in.CVTTSSD2SI.TypeRegReg(&f.Text, source.Type, wa.I64, resultReg, sourceReg)
		} else {
			truncFloatToUnsignedI64(f, resultReg, source.Type, sourceReg)
		}
		f.Regs.Free(source.Type, sourceReg)
		result = operand.Reg(resultType, resultReg, false)

	case prop.ConvertS:
		sourceReg, _ := getScratchReg(f, source)
		resultReg := f.Regs.AllocResult(resultType)
		in.CVTSI2SSD.TypeRegReg(&f.Text, resultType, source.Type, resultReg, sourceReg)
		f.Regs.Free(source.Type, sourceReg)
		result = operand.Reg(resultType, resultReg, false)

	case prop.ConvertU:
		sourceReg, zeroExt := getScratchReg(f, source)
		resultReg := f.Regs.AllocResult(resultType)
		if source.Type == wa.I32 {
			if !zeroExt {
				in.MOV.RegReg(&f.Text, wa.I32, sourceReg, sourceReg)
			}
			in.CVTSI2SSD.TypeRegReg(&f.Text, resultType, wa.I64, resultReg, sourceReg)
		} else {
			convertUnsignedI64ToFloat(f, resultType, resultReg, sourceReg)
		}
		f.Regs.Free(source.Type, sourceReg)
		result = operand.Reg(resultType, resultReg, false)

	case prop.Reinterpret:
		sourceReg, _ := getScratchReg(f, source)
		resultReg := f.Regs.AllocResult(resultType)
		if source.Type.Category() == wa.Int {
			in.MOVDQ.RegReg(&f.Text, source.Type, resultReg, sourceReg)
		} else {
			in.MOVDQmr.RegReg(&f.Text, source.Type, sourceReg, resultReg)
		}
		f.Regs.Free(source.Type, sourceReg)
		result = operand.Reg(resultType, resultReg, true)
	}

	return
}

func truncFloatToUnsignedI64(f *gen.Func, target reg.R, sourceType wa.Type, source reg.R) {
	// This algorithm is copied from code generated by gcc and clang:

	truncMaskAddr := rodata.MaskAddr(rodata.MaskTruncBase, sourceType)

	in.MOVAPSD.RegReg(&f.Text, sourceType, RegScratch, source)
	in.SUBSSD.RegMemDisp(&f.Text, sourceType, RegScratch, in.BaseText, truncMaskAddr)
	in.CVTTSSD2SI.TypeRegReg(&f.Text, sourceType, wa.I64, target, RegScratch)
	in.MOV.RegMemDisp(&f.Text, wa.I64, RegScratch, in.BaseText, rodata.Mask80Addr64)
	in.XOR.RegReg(&f.Text, wa.I64, RegScratch, target)
	in.CVTTSSD2SI.TypeRegReg(&f.Text, sourceType, wa.I64, target, source)
	in.UCOMISSD.RegMemDisp(&f.Text, sourceType, source, in.BaseText, truncMaskAddr)
	in.CMOVAE.RegReg(&f.Text, wa.I64, target, RegScratch)
}

func convertUnsignedI64ToFloat(f *gen.Func, targetType wa.Type, target, source reg.R) {
	// This algorithm is copied from code generated by gcc and clang:

	var done link.L
	var huge link.L

	in.TEST.RegReg(&f.Text, wa.I64, source, source)
	in.JScb.Stub8(&f.Text)
	huge.AddSite(f.Text.Addr)

	// max. 63-bit value
	in.CVTSI2SSD.TypeRegReg(&f.Text, targetType, wa.I64, target, source)

	in.JMPcb.Stub8(&f.Text)
	done.AddSite(f.Text.Addr)

	huge.Addr = f.Text.Addr
	isa.UpdateNearBranches(f.Text.Bytes(), &huge)

	// 64-bit value
	in.MOV.RegReg(&f.Text, wa.I64, RegScratch, source)
	in.ANDi.RegImm8(&f.Text, wa.I64, RegScratch, 1)
	in.SHRi.RegImm8(&f.Text, wa.I64, source, 1)
	in.OR.RegReg(&f.Text, wa.I64, source, RegScratch)
	in.CVTSI2SSD.TypeRegReg(&f.Text, targetType, wa.I64, target, source)
	in.ADDSSD.RegReg(&f.Text, targetType, target, target)

	done.Addr = f.Text.Addr
	isa.UpdateNearBranches(f.Text.Bytes(), &done)
}

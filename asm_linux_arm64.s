// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build linux && arm64 && gc
// +build linux
// +build arm64
// +build gc

#include "textflag.h"

// Just jump to package syscall's implementation for all these functions.
// The runtime may know about them.

TEXT ·Syscall6(SB),NOSPLIT,$0-80
	B	syscall·Syscall6(SB)
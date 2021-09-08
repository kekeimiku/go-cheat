// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build gc
// +build gc

#include "textflag.h"

//
// System calls for arm, Linux
//

// Just jump to package syscall's implementation for all these functions.
// The runtime may know about them.

TEXT ·Syscall6(SB),NOSPLIT,$0-40
	B	syscall·Syscall6(SB)
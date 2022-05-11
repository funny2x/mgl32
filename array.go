// Copyright 2016 The G3N Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mgl32

import (
	"unsafe"
)

// ArrayF32 is a slice of float32 with additional convenience methods
type ArrayF32 []float32

// NewArrayF32 creates a returns a new array of floats
// with the specified initial size and capacity
func NewArrayF32(size, capacity int) ArrayF32 {

	return make([]float32, size, capacity)
}

// Bytes returns the size of the array in bytes
func (a *ArrayF32) Bytes() int {

	return len(*a) * int(unsafe.Sizeof(float32(0)))
}

// Size returns the number of float32 elements in the array
func (a *ArrayF32) Size() int {

	return len(*a)
}

// Len returns the number of float32 elements in the array
// It is equivalent to Size()
func (a *ArrayF32) Len() int {

	return len(*a)
}

// Append appends any number of values to the array
func (a *ArrayF32) Append(v ...float32) {

	*a = append(*a, v...)
}

// ArrayU32 is a slice of uint32 with additional convenience methods
type ArrayU32 []uint32

// NewArrayU32 creates a returns a new array of uint32
// with the specified initial size and capacity
func NewArrayU32(size, capacity int) ArrayU32 {

	return make([]uint32, size, capacity)
}

// Bytes returns the size of the array in bytes
func (a *ArrayU32) Bytes() int {

	return len(*a) * int(unsafe.Sizeof(uint32(0)))
}

// Size returns the number of float32 elements in the array
func (a *ArrayU32) Size() int {

	return len(*a)
}

// Len returns the number of float32 elements in the array
func (a *ArrayU32) Len() int {

	return len(*a)
}

// Append appends n elements to the array updating the slice if necessary
func (a *ArrayU32) Append(v ...uint32) {

	*a = append(*a, v...)
}

// ToUint32 converts this array to an array of uint32
func (a ArrayU32) ToUint32() []uint32 {

	return a
}

// BSD 3-Clause License
//
// Copyright © 2009 - The Go Authors. All rights reserved.
//
// Copyright © 2025 - Marius Romeiser. All rights reserved.
//                                _                                                   __
//    ____ ____  ____  ___  _____(_)____     _______  ______  _________  ____  ____  / /
//   / __ `/ _ \/ __ \/ _ \/ ___/ / ___/    / ___/ / / / __ \/ ___/ __ \/ __ \/ __ \/ /
//  / /_/ /  __/ / / /  __/ /  / / /__     (__  ) /_/ / / / / /__/ /_/ / /_/ / /_/ / /
//  \__, /\___/_/ /_/\___/_/  /_/\___/____/____/\__, /_/ /_/\___/ .___/\____/\____/_/
// /____/                           /_____/    /____/          /_/
//
// The "generic_syncpool" package offers a generic wrapper for sync.Pool.
// For more details, you can refer to the sync.Pool documentation at: https://pkg.go.dev/sync.

package generic_syncpool

import (
	"sync"
)

// Pool represents a new generic sync.Pool
type Pool[T any] struct {
	syncpool sync.Pool
}

// New create and returns a new instance of Pool
func New[T any]() *Pool[T] {
	return &Pool[T]{
		syncpool: sync.Pool{
			New: func() any { return new(T) },
		},
	}
}

// A generic wrapper around sync.Pool's Get method.
// https://pkg.go.dev/sync#Pool.Get
func (p *Pool[T]) Get() *T {
	return p.syncpool.Get().(*T)
}

// A generic wrapper around sync.Pool's Put method.
// https://pkg.go.dev/sync#Put.Pool
func (p *Pool[T]) Put(x *T) {
	p.syncpool.Put(x)
}

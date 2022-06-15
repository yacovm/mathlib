/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package stats

import (
	"sync/atomic"

	"github.com/IBM/mathlib/driver"
)

type Curve struct {
	driver.Curve
}

var (
	Pairings          uint32
	AdditionsG1       uint32
	MultiplicationsG1 uint32
	AdditionsG2       uint32
	MultiplicationsG2 uint32
	MultiplicationsGt uint32
	ExponentiationsGt uint32
)

func (c *Curve) Pairing(g2 driver.G2, g1 driver.G1) driver.Gt {
	atomic.AddUint32(&Pairings, 1)
	return &Gt{Gt: c.Pairing(g2, g1)}
}

func (c *Curve) Pairing2(p2a, p2b driver.G2, p1a, p1b driver.G1) driver.Gt {
	atomic.AddUint32(&Pairings, 1)
	return &Gt{Gt: c.Pairing2(p2a, p2b, p1a, p1b)}
}

func (c *Curve) FExp(gt driver.Gt) driver.Gt {
	return &Gt{Gt: c.FExp(gt)}
}

func (c *Curve) GenG1() driver.G1 {
	return &G1{G1: c.GenG1()}
}

func (c *Curve) GenG2() driver.G2 {
	return &G2{G2: c.GenG2()}
}

func (c *Curve) GenGt() driver.Gt {
	return &Gt{Gt: c.GenGt()}
}

func (c *Curve) NewG1() driver.G1 {
	return &G1{G1: c.NewG1()}
}

func (c *Curve) NewG2() driver.G2 {
	return &G2{G2: c.NewG2()}
}

func (c *Curve) NewG1FromCoords(ix, iy driver.Zr) driver.G1 {
	return &G1{G1: c.NewG1FromCoords(ix, iy)}
}

func (c *Curve) NewG1FromBytes(b []byte) driver.G1 {
	return &G1{G1: c.NewG1FromBytes(b)}
}

func (c *Curve) NewG2FromBytes(b []byte) driver.G2 {
	return &G2{G2: c.NewG2FromBytes(b)}
}

func (c *Curve) NewGtFromBytes(b []byte) driver.Gt {
	return &Gt{Gt: c.NewGtFromBytes(b)}
}

func (c *Curve) HashToG1(data []byte) driver.G1 {
	return &G1{G1: c.HashToG1(data)}
}

type G1 struct {
	driver.G1
}

func (g1 *G1) Copy() driver.G1 {
	return &G1{G1: g1.G1.Copy()}
}

func (g1 *G1) Add(Q driver.G1) {
	atomic.AddUint32(&AdditionsG1, 1)
	g1.G1.Add(Q)
}

func (g1 *G1) Mul(z driver.Zr) driver.G1 {
	atomic.AddUint32(&MultiplicationsG1, 1)
	return &G1{G1: g1.G1.Mul(z)}
}

func (g1 *G1) Mul2(e driver.Zr, Q driver.G1, f driver.Zr) driver.G1 {
	atomic.AddUint32(&MultiplicationsG1, 1)
	return &G1{G1: g1.G1.Mul2(e, Q, f)}
}

type G2 struct {
	driver.G2
}

func (g2 *G2) Copy() driver.G2 {
	return &G2{G2: g2.G2.Copy()}
}

func (g2 *G2) Add(Q driver.G2) {
	atomic.AddUint32(&AdditionsG2, 1)
	g2.G2.Add(Q)
}

func (g2 *G2) Mul(z driver.Zr) driver.G2 {
	atomic.AddUint32(&MultiplicationsG2, 1)
	return &G2{G2: g2.G2.Mul(z)}
}

type Gt struct {
	driver.Gt
}

func (gt *Gt) Mul(x driver.Gt) {
	atomic.AddUint32(&MultiplicationsGt, 1)
	gt.Mul(x)
}

func (gt *Gt) Exp(x driver.Zr) driver.Gt {
	atomic.AddUint32(&ExponentiationsGt, 1)
	return &Gt{Gt: gt.Exp(x)}
}

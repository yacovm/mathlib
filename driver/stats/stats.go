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
	return &Gt{Gt: c.Curve.Pairing(g2.(*G2).G2, g1.(*G1).G1)}
}

func (c *Curve) Pairing2(p2a, p2b driver.G2, p1a, p1b driver.G1) driver.Gt {
	atomic.AddUint32(&Pairings, 1)
	return &Gt{Gt: c.Curve.Pairing2(p2a.(*G2).G2, p2b.(*G2).G2, p1a.(*G1).G1, p1b.(*G1).G1)}
}

func (c *Curve) FExp(gt driver.Gt) driver.Gt {
	return &Gt{Gt: c.Curve.FExp(gt.(*Gt).Gt)}
}

func (c *Curve) GenG1() driver.G1 {
	return &G1{G1: c.Curve.GenG1()}
}

func (c *Curve) GenG2() driver.G2 {
	return &G2{G2: c.Curve.GenG2()}
}

func (c *Curve) GenGt() driver.Gt {
	return &Gt{Gt: c.Curve.GenGt()}
}

func (c *Curve) NewG1() driver.G1 {
	return &G1{G1: c.Curve.NewG1()}
}

func (c *Curve) NewG2() driver.G2 {
	return &G2{G2: c.Curve.NewG2()}
}

func (c *Curve) NewG1FromCoords(ix, iy driver.Zr) driver.G1 {
	return &G1{G1: c.Curve.NewG1FromCoords(ix, iy)}
}

func (c *Curve) NewG1FromBytes(b []byte) driver.G1 {
	return &G1{G1: c.Curve.NewG1FromBytes(b)}
}

func (c *Curve) NewG2FromBytes(b []byte) driver.G2 {
	return &G2{G2: c.Curve.NewG2FromBytes(b)}
}

func (c *Curve) NewGtFromBytes(b []byte) driver.Gt {
	return &Gt{Gt: c.Curve.NewGtFromBytes(b)}
}

func (c *Curve) HashToG1(data []byte) driver.G1 {
	return &G1{G1: c.Curve.HashToG1(data)}
}

type G1 struct {
	driver.G1
}

func (g1 *G1) Clone(Q driver.G1) {
	g1.G1.Clone(Q.(*G1).G1)
}

func (g1 *G1) Equals(Q driver.G1) bool {
	return g1.G1.Equals(Q.(*G1).G1)
}

func (g1 *G1) Copy() driver.G1 {
	return &G1{G1: g1.G1.Copy()}
}

func (g1 *G1) Add(Q driver.G1) {
	atomic.AddUint32(&AdditionsG1, 1)
	g1.G1.Add(Q.(*G1).G1)
}

func (g1 *G1) Sub(Q driver.G1) {
	atomic.AddUint32(&AdditionsG1, 1)
	g1.G1.Sub(Q.(*G1).G1)
}

func (g1 *G1) Mul(z driver.Zr) driver.G1 {
	atomic.AddUint32(&MultiplicationsG1, 1)
	return &G1{G1: g1.G1.Mul(z)}
}

func (g1 *G1) Mul2(e driver.Zr, Q driver.G1, f driver.Zr) driver.G1 {
	atomic.AddUint32(&MultiplicationsG1, 1)
	return &G1{G1: g1.G1.Mul2(e, Q.(*G1).G1, f)}
}

type G2 struct {
	driver.G2
}

func (g2 *G2) Clone(Q driver.G2) {
	g2.G2.Clone(Q.(*G2).G2)
}

func (g2 *G2) Equals(Q driver.G2) bool {
	return g2.G2.Equals(Q.(*G2).G2)
}

func (g2 *G2) Copy() driver.G2 {
	return &G2{G2: g2.G2.Copy()}
}

func (g2 *G2) Add(Q driver.G2) {
	atomic.AddUint32(&AdditionsG2, 1)
	g2.G2.Add(Q.(*G2).G2)
}

func (g2 *G2) Sub(Q driver.G2) {
	atomic.AddUint32(&AdditionsG2, 1)
	g2.G2.Sub(Q.(*G2).G2)
}

func (g2 *G2) Mul(z driver.Zr) driver.G2 {
	atomic.AddUint32(&MultiplicationsG2, 1)
	return &G2{G2: g2.G2.Mul(z)}
}

type Gt struct {
	driver.Gt
}

func (gt *Gt) Equals(x driver.Gt) bool {
	return gt.Gt.Equals(x.(*Gt).Gt)
}

func (gt *Gt) Mul(x driver.Gt) {
	atomic.AddUint32(&MultiplicationsGt, 1)
	gt.Gt.Mul(x.(*Gt).Gt)
}

func (gt *Gt) Exp(x driver.Zr) driver.Gt {
	atomic.AddUint32(&ExponentiationsGt, 1)
	return &Gt{Gt: gt.Gt.Exp(x)}
}

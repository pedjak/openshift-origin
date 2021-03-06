// Code generated by "go generate gonum.org/v1/gonum/unit”; DO NOT EDIT.

// Copyright ©2014 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package unit

import (
	"errors"
	"fmt"
	"math"
	"unicode/utf8"
)

// MagneticFluxDensity represents a magnetic flux density in Tesla.
type MagneticFluxDensity float64

const (
	Yottatesla MagneticFluxDensity = 1e24
	Zettatesla MagneticFluxDensity = 1e21
	Exatesla   MagneticFluxDensity = 1e18
	Petatesla  MagneticFluxDensity = 1e15
	Teratesla  MagneticFluxDensity = 1e12
	Gigatesla  MagneticFluxDensity = 1e9
	Megatesla  MagneticFluxDensity = 1e6
	Kilotesla  MagneticFluxDensity = 1e3
	Hectotesla MagneticFluxDensity = 1e2
	Decatesla  MagneticFluxDensity = 1e1
	Tesla      MagneticFluxDensity = 1.0
	Decitesla  MagneticFluxDensity = 1e-1
	Centitesla MagneticFluxDensity = 1e-2
	Millitesla MagneticFluxDensity = 1e-3
	Microtesla MagneticFluxDensity = 1e-6
	Nanotesla  MagneticFluxDensity = 1e-9
	Picotesla  MagneticFluxDensity = 1e-12
	Femtotesla MagneticFluxDensity = 1e-15
	Attotesla  MagneticFluxDensity = 1e-18
	Zeptotesla MagneticFluxDensity = 1e-21
	Yoctotesla MagneticFluxDensity = 1e-24
)

// Unit converts the MagneticFluxDensity to a *Unit
func (m MagneticFluxDensity) Unit() *Unit {
	return New(float64(m), Dimensions{
		CurrentDim: -1,
		MassDim:    1,
		TimeDim:    -2,
	})
}

// MagneticFluxDensity allows MagneticFluxDensity to implement a MagneticFluxDensityer interface
func (m MagneticFluxDensity) MagneticFluxDensity() MagneticFluxDensity {
	return m
}

// From converts the unit into the receiver. From returns an
// error if there is a mismatch in dimension
func (m *MagneticFluxDensity) From(u Uniter) error {
	if !DimensionsMatch(u, Tesla) {
		*m = MagneticFluxDensity(math.NaN())
		return errors.New("Dimension mismatch")
	}
	*m = MagneticFluxDensity(u.Unit().Value())
	return nil
}

func (m MagneticFluxDensity) Format(fs fmt.State, c rune) {
	switch c {
	case 'v':
		if fs.Flag('#') {
			fmt.Fprintf(fs, "%T(%v)", m, float64(m))
			return
		}
		fallthrough
	case 'e', 'E', 'f', 'F', 'g', 'G':
		p, pOk := fs.Precision()
		w, wOk := fs.Width()
		const unit = " T"
		switch {
		case pOk && wOk:
			fmt.Fprintf(fs, "%*.*"+string(c), pos(w-utf8.RuneCount([]byte(unit))), p, float64(m))
		case pOk:
			fmt.Fprintf(fs, "%.*"+string(c), p, float64(m))
		case wOk:
			fmt.Fprintf(fs, "%*"+string(c), pos(w-utf8.RuneCount([]byte(unit))), float64(m))
		default:
			fmt.Fprintf(fs, "%"+string(c), float64(m))
		}
		fmt.Fprint(fs, unit)
	default:
		fmt.Fprintf(fs, "%%!%c(%T=%g T)", c, m, float64(m))
	}
}

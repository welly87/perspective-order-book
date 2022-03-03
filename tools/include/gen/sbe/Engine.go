// Generated SBE (Simple Binary Encoding) message codec

package sbe

import (
	"fmt"
	"io"
	"math"
)

type Engine struct {
	Capacity         uint16
	NumCylinders     uint8
	MaxRpm           uint16
	ManufacturerCode [3]byte
	Fuel             [6]byte
}

func (e *Engine) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint16(_w, e.Capacity); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, e.NumCylinders); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, e.ManufacturerCode[:]); err != nil {
		return err
	}
	return nil
}

func (e *Engine) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if !e.CapacityInActingVersion(actingVersion) {
		e.Capacity = e.CapacityNullValue()
	} else {
		if err := _m.ReadUint16(_r, &e.Capacity); err != nil {
			return err
		}
	}
	if !e.NumCylindersInActingVersion(actingVersion) {
		e.NumCylinders = e.NumCylindersNullValue()
	} else {
		if err := _m.ReadUint8(_r, &e.NumCylinders); err != nil {
			return err
		}
	}
	e.MaxRpm = 9000
	if !e.ManufacturerCodeInActingVersion(actingVersion) {
		for idx := 0; idx < 3; idx++ {
			e.ManufacturerCode[idx] = e.ManufacturerCodeNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, e.ManufacturerCode[:]); err != nil {
			return err
		}
	}
	copy(e.Fuel[:], "Petrol")
	return nil
}

func (e *Engine) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if e.CapacityInActingVersion(actingVersion) {
		if e.Capacity < e.CapacityMinValue() || e.Capacity > e.CapacityMaxValue() {
			return fmt.Errorf("Range check failed on e.Capacity (%v < %v > %v)", e.CapacityMinValue(), e.Capacity, e.CapacityMaxValue())
		}
	}
	if e.NumCylindersInActingVersion(actingVersion) {
		if e.NumCylinders < e.NumCylindersMinValue() || e.NumCylinders > e.NumCylindersMaxValue() {
			return fmt.Errorf("Range check failed on e.NumCylinders (%v < %v > %v)", e.NumCylindersMinValue(), e.NumCylinders, e.NumCylindersMaxValue())
		}
	}
	if e.ManufacturerCodeInActingVersion(actingVersion) {
		for idx := 0; idx < 3; idx++ {
			if e.ManufacturerCode[idx] < e.ManufacturerCodeMinValue() || e.ManufacturerCode[idx] > e.ManufacturerCodeMaxValue() {
				return fmt.Errorf("Range check failed on e.ManufacturerCode[%d] (%v < %v > %v)", idx, e.ManufacturerCodeMinValue(), e.ManufacturerCode[idx], e.ManufacturerCodeMaxValue())
			}
		}
	}
	return nil
}

func EngineInit(e *Engine) {
	e.MaxRpm = 9000
	copy(e.Fuel[:], "Petrol")
	return
}

func (*Engine) EncodedLength() int64 {
	return 6
}

func (*Engine) CapacityMinValue() uint16 {
	return 0
}

func (*Engine) CapacityMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*Engine) CapacityNullValue() uint16 {
	return math.MaxUint16
}

func (*Engine) CapacitySinceVersion() uint16 {
	return 0
}

func (e *Engine) CapacityInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.CapacitySinceVersion()
}

func (*Engine) CapacityDeprecated() uint16 {
	return 0
}

func (*Engine) NumCylindersMinValue() uint8 {
	return 0
}

func (*Engine) NumCylindersMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*Engine) NumCylindersNullValue() uint8 {
	return math.MaxUint8
}

func (*Engine) NumCylindersSinceVersion() uint16 {
	return 0
}

func (e *Engine) NumCylindersInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.NumCylindersSinceVersion()
}

func (*Engine) NumCylindersDeprecated() uint16 {
	return 0
}

func (*Engine) MaxRpmMinValue() uint16 {
	return 0
}

func (*Engine) MaxRpmMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*Engine) MaxRpmNullValue() uint16 {
	return math.MaxUint16
}

func (*Engine) MaxRpmSinceVersion() uint16 {
	return 0
}

func (e *Engine) MaxRpmInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.MaxRpmSinceVersion()
}

func (*Engine) MaxRpmDeprecated() uint16 {
	return 0
}

func (*Engine) ManufacturerCodeMinValue() byte {
	return byte(32)
}

func (*Engine) ManufacturerCodeMaxValue() byte {
	return byte(126)
}

func (*Engine) ManufacturerCodeNullValue() byte {
	return 0
}

func (e *Engine) ManufacturerCodeCharacterEncoding() string {
	return "US-ASCII"
}

func (*Engine) ManufacturerCodeSinceVersion() uint16 {
	return 0
}

func (e *Engine) ManufacturerCodeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ManufacturerCodeSinceVersion()
}

func (*Engine) ManufacturerCodeDeprecated() uint16 {
	return 0
}

func (*Engine) FuelMinValue() byte {
	return byte(32)
}

func (*Engine) FuelMaxValue() byte {
	return byte(126)
}

func (*Engine) FuelNullValue() byte {
	return 0
}

func (*Engine) FuelSinceVersion() uint16 {
	return 0
}

func (e *Engine) FuelInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.FuelSinceVersion()
}

func (*Engine) FuelDeprecated() uint16 {
	return 0
}

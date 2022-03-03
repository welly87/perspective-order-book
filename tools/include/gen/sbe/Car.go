// Generated SBE (Simple Binary Encoding) message codec

package sbe

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type Car struct {
	SerialNumber       uint32
	ModelYear          uint16
	Available          BooleanTypeEnum
	Code               ModelEnum
	SomeNumbers        [5]int32
	VehicleCode        [6]byte
	Extras             OptionalExtras
	Engine             Engine
	FuelFigures        []CarFuelFigures
	PerformanceFigures []CarPerformanceFigures
	Manufacturer       []uint8
	Model              []uint8
}
type CarFuelFigures struct {
	Speed uint16
	Mpg   float32
}
type CarPerformanceFigures struct {
	OctaneRating uint8
	Acceleration []CarPerformanceFiguresAcceleration
}
type CarPerformanceFiguresAcceleration struct {
	Mph     uint16
	Seconds float32
}

func (c *Car) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := c.RangeCheck(c.SbeSchemaVersion(), c.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteUint32(_w, c.SerialNumber); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, c.ModelYear); err != nil {
		return err
	}
	if err := c.Available.Encode(_m, _w); err != nil {
		return err
	}
	if err := c.Code.Encode(_m, _w); err != nil {
		return err
	}
	for idx := 0; idx < 5; idx++ {
		if err := _m.WriteInt32(_w, c.SomeNumbers[idx]); err != nil {
			return err
		}
	}
	if err := _m.WriteBytes(_w, c.VehicleCode[:]); err != nil {
		return err
	}
	if err := c.Extras.Encode(_m, _w); err != nil {
		return err
	}
	if err := c.Engine.Encode(_m, _w); err != nil {
		return err
	}
	var FuelFiguresBlockLength uint16 = 6
	if err := _m.WriteUint16(_w, FuelFiguresBlockLength); err != nil {
		return err
	}
	var FuelFiguresNumInGroup uint16 = uint16(len(c.FuelFigures))
	if err := _m.WriteUint16(_w, FuelFiguresNumInGroup); err != nil {
		return err
	}
	for _, prop := range c.FuelFigures {
		if err := prop.Encode(_m, _w); err != nil {
			return err
		}
	}
	var PerformanceFiguresBlockLength uint16 = 1
	if err := _m.WriteUint16(_w, PerformanceFiguresBlockLength); err != nil {
		return err
	}
	var PerformanceFiguresNumInGroup uint16 = uint16(len(c.PerformanceFigures))
	if err := _m.WriteUint16(_w, PerformanceFiguresNumInGroup); err != nil {
		return err
	}
	for _, prop := range c.PerformanceFigures {
		if err := prop.Encode(_m, _w); err != nil {
			return err
		}
	}
	if err := _m.WriteUint32(_w, uint32(len(c.Manufacturer))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, c.Manufacturer); err != nil {
		return err
	}
	if err := _m.WriteUint32(_w, uint32(len(c.Model))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, c.Model); err != nil {
		return err
	}
	return nil
}

func (c *Car) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !c.SerialNumberInActingVersion(actingVersion) {
		c.SerialNumber = c.SerialNumberNullValue()
	} else {
		if err := _m.ReadUint32(_r, &c.SerialNumber); err != nil {
			return err
		}
	}
	if !c.ModelYearInActingVersion(actingVersion) {
		c.ModelYear = c.ModelYearNullValue()
	} else {
		if err := _m.ReadUint16(_r, &c.ModelYear); err != nil {
			return err
		}
	}
	if c.AvailableInActingVersion(actingVersion) {
		if err := c.Available.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if c.CodeInActingVersion(actingVersion) {
		if err := c.Code.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !c.SomeNumbersInActingVersion(actingVersion) {
		for idx := 0; idx < 5; idx++ {
			c.SomeNumbers[idx] = c.SomeNumbersNullValue()
		}
	} else {
		for idx := 0; idx < 5; idx++ {
			if err := _m.ReadInt32(_r, &c.SomeNumbers[idx]); err != nil {
				return err
			}
		}
	}
	if !c.VehicleCodeInActingVersion(actingVersion) {
		for idx := 0; idx < 6; idx++ {
			c.VehicleCode[idx] = c.VehicleCodeNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, c.VehicleCode[:]); err != nil {
			return err
		}
	}
	if c.ExtrasInActingVersion(actingVersion) {
		if err := c.Extras.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if c.EngineInActingVersion(actingVersion) {
		if err := c.Engine.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if actingVersion > c.SbeSchemaVersion() && blockLength > c.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-c.SbeBlockLength()))
	}

	if c.FuelFiguresInActingVersion(actingVersion) {
		var FuelFiguresBlockLength uint16
		if err := _m.ReadUint16(_r, &FuelFiguresBlockLength); err != nil {
			return err
		}
		var FuelFiguresNumInGroup uint16
		if err := _m.ReadUint16(_r, &FuelFiguresNumInGroup); err != nil {
			return err
		}
		if cap(c.FuelFigures) < int(FuelFiguresNumInGroup) {
			c.FuelFigures = make([]CarFuelFigures, FuelFiguresNumInGroup)
		}
		c.FuelFigures = c.FuelFigures[:FuelFiguresNumInGroup]
		for i := range c.FuelFigures {
			if err := c.FuelFigures[i].Decode(_m, _r, actingVersion, uint(FuelFiguresBlockLength)); err != nil {
				return err
			}
		}
	}

	if c.PerformanceFiguresInActingVersion(actingVersion) {
		var PerformanceFiguresBlockLength uint16
		if err := _m.ReadUint16(_r, &PerformanceFiguresBlockLength); err != nil {
			return err
		}
		var PerformanceFiguresNumInGroup uint16
		if err := _m.ReadUint16(_r, &PerformanceFiguresNumInGroup); err != nil {
			return err
		}
		if cap(c.PerformanceFigures) < int(PerformanceFiguresNumInGroup) {
			c.PerformanceFigures = make([]CarPerformanceFigures, PerformanceFiguresNumInGroup)
		}
		c.PerformanceFigures = c.PerformanceFigures[:PerformanceFiguresNumInGroup]
		for i := range c.PerformanceFigures {
			if err := c.PerformanceFigures[i].Decode(_m, _r, actingVersion, uint(PerformanceFiguresBlockLength)); err != nil {
				return err
			}
		}
	}

	if c.ManufacturerInActingVersion(actingVersion) {
		var ManufacturerLength uint32
		if err := _m.ReadUint32(_r, &ManufacturerLength); err != nil {
			return err
		}
		if cap(c.Manufacturer) < int(ManufacturerLength) {
			c.Manufacturer = make([]uint8, ManufacturerLength)
		}
		c.Manufacturer = c.Manufacturer[:ManufacturerLength]
		if err := _m.ReadBytes(_r, c.Manufacturer); err != nil {
			return err
		}
	}

	if c.ModelInActingVersion(actingVersion) {
		var ModelLength uint32
		if err := _m.ReadUint32(_r, &ModelLength); err != nil {
			return err
		}
		if cap(c.Model) < int(ModelLength) {
			c.Model = make([]uint8, ModelLength)
		}
		c.Model = c.Model[:ModelLength]
		if err := _m.ReadBytes(_r, c.Model); err != nil {
			return err
		}
	}
	if doRangeCheck {
		if err := c.RangeCheck(actingVersion, c.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (c *Car) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if c.SerialNumberInActingVersion(actingVersion) {
		if c.SerialNumber < c.SerialNumberMinValue() || c.SerialNumber > c.SerialNumberMaxValue() {
			return fmt.Errorf("Range check failed on c.SerialNumber (%v < %v > %v)", c.SerialNumberMinValue(), c.SerialNumber, c.SerialNumberMaxValue())
		}
	}
	if c.ModelYearInActingVersion(actingVersion) {
		if c.ModelYear < c.ModelYearMinValue() || c.ModelYear > c.ModelYearMaxValue() {
			return fmt.Errorf("Range check failed on c.ModelYear (%v < %v > %v)", c.ModelYearMinValue(), c.ModelYear, c.ModelYearMaxValue())
		}
	}
	if err := c.Available.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := c.Code.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if c.SomeNumbersInActingVersion(actingVersion) {
		for idx := 0; idx < 5; idx++ {
			if c.SomeNumbers[idx] < c.SomeNumbersMinValue() || c.SomeNumbers[idx] > c.SomeNumbersMaxValue() {
				return fmt.Errorf("Range check failed on c.SomeNumbers[%d] (%v < %v > %v)", idx, c.SomeNumbersMinValue(), c.SomeNumbers[idx], c.SomeNumbersMaxValue())
			}
		}
	}
	if c.VehicleCodeInActingVersion(actingVersion) {
		for idx := 0; idx < 6; idx++ {
			if c.VehicleCode[idx] < c.VehicleCodeMinValue() || c.VehicleCode[idx] > c.VehicleCodeMaxValue() {
				return fmt.Errorf("Range check failed on c.VehicleCode[%d] (%v < %v > %v)", idx, c.VehicleCodeMinValue(), c.VehicleCode[idx], c.VehicleCodeMaxValue())
			}
		}
	}
	for _, prop := range c.FuelFigures {
		if err := prop.RangeCheck(actingVersion, schemaVersion); err != nil {
			return err
		}
	}
	for _, prop := range c.PerformanceFigures {
		if err := prop.RangeCheck(actingVersion, schemaVersion); err != nil {
			return err
		}
	}
	return nil
}

func CarInit(c *Car) {
	return
}

func (c *CarFuelFigures) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint16(_w, c.Speed); err != nil {
		return err
	}
	if err := _m.WriteFloat32(_w, c.Mpg); err != nil {
		return err
	}
	return nil
}

func (c *CarFuelFigures) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint) error {
	if !c.SpeedInActingVersion(actingVersion) {
		c.Speed = c.SpeedNullValue()
	} else {
		if err := _m.ReadUint16(_r, &c.Speed); err != nil {
			return err
		}
	}
	if !c.MpgInActingVersion(actingVersion) {
		c.Mpg = c.MpgNullValue()
	} else {
		if err := _m.ReadFloat32(_r, &c.Mpg); err != nil {
			return err
		}
	}
	if actingVersion > c.SbeSchemaVersion() && blockLength > c.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-c.SbeBlockLength()))
	}
	return nil
}

func (c *CarFuelFigures) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if c.SpeedInActingVersion(actingVersion) {
		if c.Speed < c.SpeedMinValue() || c.Speed > c.SpeedMaxValue() {
			return fmt.Errorf("Range check failed on c.Speed (%v < %v > %v)", c.SpeedMinValue(), c.Speed, c.SpeedMaxValue())
		}
	}
	if c.MpgInActingVersion(actingVersion) {
		if c.Mpg < c.MpgMinValue() || c.Mpg > c.MpgMaxValue() {
			return fmt.Errorf("Range check failed on c.Mpg (%v < %v > %v)", c.MpgMinValue(), c.Mpg, c.MpgMaxValue())
		}
	}
	return nil
}

func CarFuelFiguresInit(c *CarFuelFigures) {
	return
}

func (c *CarPerformanceFigures) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, c.OctaneRating); err != nil {
		return err
	}
	var AccelerationBlockLength uint16 = 6
	if err := _m.WriteUint16(_w, AccelerationBlockLength); err != nil {
		return err
	}
	var AccelerationNumInGroup uint16 = uint16(len(c.Acceleration))
	if err := _m.WriteUint16(_w, AccelerationNumInGroup); err != nil {
		return err
	}
	for _, prop := range c.Acceleration {
		if err := prop.Encode(_m, _w); err != nil {
			return err
		}
	}
	return nil
}

func (c *CarPerformanceFigures) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint) error {
	if !c.OctaneRatingInActingVersion(actingVersion) {
		c.OctaneRating = c.OctaneRatingNullValue()
	} else {
		if err := _m.ReadUint8(_r, &c.OctaneRating); err != nil {
			return err
		}
	}
	if actingVersion > c.SbeSchemaVersion() && blockLength > c.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-c.SbeBlockLength()))
	}

	if c.AccelerationInActingVersion(actingVersion) {
		var AccelerationBlockLength uint16
		if err := _m.ReadUint16(_r, &AccelerationBlockLength); err != nil {
			return err
		}
		var AccelerationNumInGroup uint16
		if err := _m.ReadUint16(_r, &AccelerationNumInGroup); err != nil {
			return err
		}
		if cap(c.Acceleration) < int(AccelerationNumInGroup) {
			c.Acceleration = make([]CarPerformanceFiguresAcceleration, AccelerationNumInGroup)
		}
		c.Acceleration = c.Acceleration[:AccelerationNumInGroup]
		for i := range c.Acceleration {
			if err := c.Acceleration[i].Decode(_m, _r, actingVersion, uint(AccelerationBlockLength)); err != nil {
				return err
			}
		}
	}
	return nil
}

func (c *CarPerformanceFigures) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if c.OctaneRatingInActingVersion(actingVersion) {
		if c.OctaneRating < c.OctaneRatingMinValue() || c.OctaneRating > c.OctaneRatingMaxValue() {
			return fmt.Errorf("Range check failed on c.OctaneRating (%v < %v > %v)", c.OctaneRatingMinValue(), c.OctaneRating, c.OctaneRatingMaxValue())
		}
	}
	for _, prop := range c.Acceleration {
		if err := prop.RangeCheck(actingVersion, schemaVersion); err != nil {
			return err
		}
	}
	return nil
}

func CarPerformanceFiguresInit(c *CarPerformanceFigures) {
	return
}

func (c *CarPerformanceFiguresAcceleration) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint16(_w, c.Mph); err != nil {
		return err
	}
	if err := _m.WriteFloat32(_w, c.Seconds); err != nil {
		return err
	}
	return nil
}

func (c *CarPerformanceFiguresAcceleration) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint) error {
	if !c.MphInActingVersion(actingVersion) {
		c.Mph = c.MphNullValue()
	} else {
		if err := _m.ReadUint16(_r, &c.Mph); err != nil {
			return err
		}
	}
	if !c.SecondsInActingVersion(actingVersion) {
		c.Seconds = c.SecondsNullValue()
	} else {
		if err := _m.ReadFloat32(_r, &c.Seconds); err != nil {
			return err
		}
	}
	if actingVersion > c.SbeSchemaVersion() && blockLength > c.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-c.SbeBlockLength()))
	}
	return nil
}

func (c *CarPerformanceFiguresAcceleration) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if c.MphInActingVersion(actingVersion) {
		if c.Mph < c.MphMinValue() || c.Mph > c.MphMaxValue() {
			return fmt.Errorf("Range check failed on c.Mph (%v < %v > %v)", c.MphMinValue(), c.Mph, c.MphMaxValue())
		}
	}
	if c.SecondsInActingVersion(actingVersion) {
		if c.Seconds < c.SecondsMinValue() || c.Seconds > c.SecondsMaxValue() {
			return fmt.Errorf("Range check failed on c.Seconds (%v < %v > %v)", c.SecondsMinValue(), c.Seconds, c.SecondsMaxValue())
		}
	}
	return nil
}

func CarPerformanceFiguresAccelerationInit(c *CarPerformanceFiguresAcceleration) {
	return
}

func (*Car) SbeBlockLength() (blockLength uint16) {
	return 41
}

func (*Car) SbeTemplateId() (templateId uint16) {
	return 1
}

func (*Car) SbeSchemaId() (schemaId uint16) {
	return 1
}

func (*Car) SbeSchemaVersion() (schemaVersion uint16) {
	return 1
}

func (*Car) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*Car) SerialNumberId() uint16 {
	return 1
}

func (*Car) SerialNumberSinceVersion() uint16 {
	return 0
}

func (c *Car) SerialNumberInActingVersion(actingVersion uint16) bool {
	return actingVersion >= c.SerialNumberSinceVersion()
}

func (*Car) SerialNumberDeprecated() uint16 {
	return 0
}

func (*Car) SerialNumberMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*Car) SerialNumberMinValue() uint32 {
	return 0
}

func (*Car) SerialNumberMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*Car) SerialNumberNullValue() uint32 {
	return math.MaxUint32
}

func (*Car) ModelYearId() uint16 {
	return 2
}

func (*Car) ModelYearSinceVersion() uint16 {
	return 0
}

func (c *Car) ModelYearInActingVersion(actingVersion uint16) bool {
	return actingVersion >= c.ModelYearSinceVersion()
}

func (*Car) ModelYearDeprecated() uint16 {
	return 0
}

func (*Car) ModelYearMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*Car) ModelYearMinValue() uint16 {
	return 0
}

func (*Car) ModelYearMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*Car) ModelYearNullValue() uint16 {
	return math.MaxUint16
}

func (*Car) AvailableId() uint16 {
	return 3
}

func (*Car) AvailableSinceVersion() uint16 {
	return 0
}

func (c *Car) AvailableInActingVersion(actingVersion uint16) bool {
	return actingVersion >= c.AvailableSinceVersion()
}

func (*Car) AvailableDeprecated() uint16 {
	return 0
}

func (*Car) AvailableMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*Car) CodeId() uint16 {
	return 4
}

func (*Car) CodeSinceVersion() uint16 {
	return 0
}

func (c *Car) CodeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= c.CodeSinceVersion()
}

func (*Car) CodeDeprecated() uint16 {
	return 0
}

func (*Car) CodeMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*Car) SomeNumbersId() uint16 {
	return 5
}

func (*Car) SomeNumbersSinceVersion() uint16 {
	return 0
}

func (c *Car) SomeNumbersInActingVersion(actingVersion uint16) bool {
	return actingVersion >= c.SomeNumbersSinceVersion()
}

func (*Car) SomeNumbersDeprecated() uint16 {
	return 0
}

func (*Car) SomeNumbersMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*Car) SomeNumbersMinValue() int32 {
	return math.MinInt32 + 1
}

func (*Car) SomeNumbersMaxValue() int32 {
	return math.MaxInt32
}

func (*Car) SomeNumbersNullValue() int32 {
	return math.MinInt32
}

func (*Car) VehicleCodeId() uint16 {
	return 6
}

func (*Car) VehicleCodeSinceVersion() uint16 {
	return 0
}

func (c *Car) VehicleCodeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= c.VehicleCodeSinceVersion()
}

func (*Car) VehicleCodeDeprecated() uint16 {
	return 0
}

func (*Car) VehicleCodeMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*Car) VehicleCodeMinValue() byte {
	return byte(32)
}

func (*Car) VehicleCodeMaxValue() byte {
	return byte(126)
}

func (*Car) VehicleCodeNullValue() byte {
	return 0
}

func (c *Car) VehicleCodeCharacterEncoding() string {
	return "US-ASCII"
}

func (*Car) ExtrasId() uint16 {
	return 7
}

func (*Car) ExtrasSinceVersion() uint16 {
	return 0
}

func (c *Car) ExtrasInActingVersion(actingVersion uint16) bool {
	return actingVersion >= c.ExtrasSinceVersion()
}

func (*Car) ExtrasDeprecated() uint16 {
	return 0
}

func (*Car) ExtrasMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*Car) EngineId() uint16 {
	return 8
}

func (*Car) EngineSinceVersion() uint16 {
	return 0
}

func (c *Car) EngineInActingVersion(actingVersion uint16) bool {
	return actingVersion >= c.EngineSinceVersion()
}

func (*Car) EngineDeprecated() uint16 {
	return 0
}

func (*Car) EngineMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*CarFuelFigures) SpeedId() uint16 {
	return 10
}

func (*CarFuelFigures) SpeedSinceVersion() uint16 {
	return 0
}

func (c *CarFuelFigures) SpeedInActingVersion(actingVersion uint16) bool {
	return actingVersion >= c.SpeedSinceVersion()
}

func (*CarFuelFigures) SpeedDeprecated() uint16 {
	return 0
}

func (*CarFuelFigures) SpeedMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*CarFuelFigures) SpeedMinValue() uint16 {
	return 0
}

func (*CarFuelFigures) SpeedMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*CarFuelFigures) SpeedNullValue() uint16 {
	return math.MaxUint16
}

func (*CarFuelFigures) MpgId() uint16 {
	return 11
}

func (*CarFuelFigures) MpgSinceVersion() uint16 {
	return 0
}

func (c *CarFuelFigures) MpgInActingVersion(actingVersion uint16) bool {
	return actingVersion >= c.MpgSinceVersion()
}

func (*CarFuelFigures) MpgDeprecated() uint16 {
	return 0
}

func (*CarFuelFigures) MpgMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*CarFuelFigures) MpgMinValue() float32 {
	return -math.MaxFloat32
}

func (*CarFuelFigures) MpgMaxValue() float32 {
	return math.MaxFloat32
}

func (*CarFuelFigures) MpgNullValue() float32 {
	return float32(math.NaN())
}

func (*CarPerformanceFigures) OctaneRatingId() uint16 {
	return 13
}

func (*CarPerformanceFigures) OctaneRatingSinceVersion() uint16 {
	return 0
}

func (c *CarPerformanceFigures) OctaneRatingInActingVersion(actingVersion uint16) bool {
	return actingVersion >= c.OctaneRatingSinceVersion()
}

func (*CarPerformanceFigures) OctaneRatingDeprecated() uint16 {
	return 0
}

func (*CarPerformanceFigures) OctaneRatingMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*CarPerformanceFigures) OctaneRatingMinValue() uint8 {
	return 0
}

func (*CarPerformanceFigures) OctaneRatingMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*CarPerformanceFigures) OctaneRatingNullValue() uint8 {
	return math.MaxUint8
}

func (*CarPerformanceFiguresAcceleration) MphId() uint16 {
	return 15
}

func (*CarPerformanceFiguresAcceleration) MphSinceVersion() uint16 {
	return 0
}

func (c *CarPerformanceFiguresAcceleration) MphInActingVersion(actingVersion uint16) bool {
	return actingVersion >= c.MphSinceVersion()
}

func (*CarPerformanceFiguresAcceleration) MphDeprecated() uint16 {
	return 0
}

func (*CarPerformanceFiguresAcceleration) MphMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*CarPerformanceFiguresAcceleration) MphMinValue() uint16 {
	return 0
}

func (*CarPerformanceFiguresAcceleration) MphMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*CarPerformanceFiguresAcceleration) MphNullValue() uint16 {
	return math.MaxUint16
}

func (*CarPerformanceFiguresAcceleration) SecondsId() uint16 {
	return 16
}

func (*CarPerformanceFiguresAcceleration) SecondsSinceVersion() uint16 {
	return 0
}

func (c *CarPerformanceFiguresAcceleration) SecondsInActingVersion(actingVersion uint16) bool {
	return actingVersion >= c.SecondsSinceVersion()
}

func (*CarPerformanceFiguresAcceleration) SecondsDeprecated() uint16 {
	return 0
}

func (*CarPerformanceFiguresAcceleration) SecondsMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*CarPerformanceFiguresAcceleration) SecondsMinValue() float32 {
	return -math.MaxFloat32
}

func (*CarPerformanceFiguresAcceleration) SecondsMaxValue() float32 {
	return math.MaxFloat32
}

func (*CarPerformanceFiguresAcceleration) SecondsNullValue() float32 {
	return float32(math.NaN())
}

func (*Car) FuelFiguresId() uint16 {
	return 9
}

func (*Car) FuelFiguresSinceVersion() uint16 {
	return 0
}

func (c *Car) FuelFiguresInActingVersion(actingVersion uint16) bool {
	return actingVersion >= c.FuelFiguresSinceVersion()
}

func (*Car) FuelFiguresDeprecated() uint16 {
	return 0
}

func (*CarFuelFigures) SbeBlockLength() (blockLength uint) {
	return 6
}

func (*CarFuelFigures) SbeSchemaVersion() (schemaVersion uint16) {
	return 1
}

func (*Car) PerformanceFiguresId() uint16 {
	return 12
}

func (*Car) PerformanceFiguresSinceVersion() uint16 {
	return 0
}

func (c *Car) PerformanceFiguresInActingVersion(actingVersion uint16) bool {
	return actingVersion >= c.PerformanceFiguresSinceVersion()
}

func (*Car) PerformanceFiguresDeprecated() uint16 {
	return 0
}

func (*CarPerformanceFigures) SbeBlockLength() (blockLength uint) {
	return 1
}

func (*CarPerformanceFigures) SbeSchemaVersion() (schemaVersion uint16) {
	return 1
}

func (*CarPerformanceFigures) AccelerationId() uint16 {
	return 14
}

func (*CarPerformanceFigures) AccelerationSinceVersion() uint16 {
	return 0
}

func (c *CarPerformanceFigures) AccelerationInActingVersion(actingVersion uint16) bool {
	return actingVersion >= c.AccelerationSinceVersion()
}

func (*CarPerformanceFigures) AccelerationDeprecated() uint16 {
	return 0
}

func (*CarPerformanceFiguresAcceleration) SbeBlockLength() (blockLength uint) {
	return 6
}

func (*CarPerformanceFiguresAcceleration) SbeSchemaVersion() (schemaVersion uint16) {
	return 1
}

func (*Car) ManufacturerMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*Car) ManufacturerSinceVersion() uint16 {
	return 0
}

func (c *Car) ManufacturerInActingVersion(actingVersion uint16) bool {
	return actingVersion >= c.ManufacturerSinceVersion()
}

func (*Car) ManufacturerDeprecated() uint16 {
	return 0
}

func (Car) ManufacturerCharacterEncoding() string {
	return "ISO-8859-1"
}

func (Car) ManufacturerHeaderLength() uint64 {
	return 4
}

func (*Car) ModelMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*Car) ModelSinceVersion() uint16 {
	return 0
}

func (c *Car) ModelInActingVersion(actingVersion uint16) bool {
	return actingVersion >= c.ModelSinceVersion()
}

func (*Car) ModelDeprecated() uint16 {
	return 0
}

func (Car) ModelCharacterEncoding() string {
	return "ISO-8859-1"
}

func (Car) ModelHeaderLength() uint64 {
	return 4
}

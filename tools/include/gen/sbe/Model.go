// Generated SBE (Simple Binary Encoding) message codec

package sbe

import (
	"fmt"
	"io"
	"reflect"
)

type ModelEnum byte
type ModelValues struct {
	A         ModelEnum
	B         ModelEnum
	C         ModelEnum
	NullValue ModelEnum
}

var Model = ModelValues{65, 66, 67, 0}

func (m ModelEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, byte(m)); err != nil {
		return err
	}
	return nil
}

func (m *ModelEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*byte)(m)); err != nil {
		return err
	}
	return nil
}

func (m ModelEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(Model)
	for idx := 0; idx < value.NumField(); idx++ {
		if m == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on Model, unknown enumeration value %d", m)
}

func (*ModelEnum) EncodedLength() int64 {
	return 1
}

func (*ModelEnum) ASinceVersion() uint16 {
	return 0
}

func (m *ModelEnum) AInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.ASinceVersion()
}

func (*ModelEnum) ADeprecated() uint16 {
	return 0
}

func (*ModelEnum) BSinceVersion() uint16 {
	return 0
}

func (m *ModelEnum) BInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.BSinceVersion()
}

func (*ModelEnum) BDeprecated() uint16 {
	return 0
}

func (*ModelEnum) CSinceVersion() uint16 {
	return 0
}

func (m *ModelEnum) CInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.CSinceVersion()
}

func (*ModelEnum) CDeprecated() uint16 {
	return 0
}

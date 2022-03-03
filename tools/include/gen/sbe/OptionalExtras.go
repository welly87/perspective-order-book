// Generated SBE (Simple Binary Encoding) message codec

package sbe

import (
	"io"
)

type OptionalExtras [8]bool
type OptionalExtrasChoiceValue uint8
type OptionalExtrasChoiceValues struct {
	SunRoof       OptionalExtrasChoiceValue
	SportsPack    OptionalExtrasChoiceValue
	CruiseControl OptionalExtrasChoiceValue
}

var OptionalExtrasChoice = OptionalExtrasChoiceValues{0, 1, 2}

func (o *OptionalExtras) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	var wireval uint8 = 0
	for k, v := range o {
		if v {
			wireval |= (1 << uint(k))
		}
	}
	return _m.WriteUint8(_w, wireval)
}

func (o *OptionalExtras) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	var wireval uint8

	if err := _m.ReadUint8(_r, &wireval); err != nil {
		return err
	}

	var idx uint
	for idx = 0; idx < 8; idx++ {
		o[idx] = (wireval & (1 << idx)) > 0
	}
	return nil
}

func (OptionalExtras) EncodedLength() int64 {
	return 1
}

func (*OptionalExtras) sunRoofSinceVersion() uint16 {
	return 0
}

func (o *OptionalExtras) sunRoofInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.sunRoofSinceVersion()
}

func (*OptionalExtras) sunRoofDeprecated() uint16 {
	return 0
}

func (*OptionalExtras) sportsPackSinceVersion() uint16 {
	return 0
}

func (o *OptionalExtras) sportsPackInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.sportsPackSinceVersion()
}

func (*OptionalExtras) sportsPackDeprecated() uint16 {
	return 0
}

func (*OptionalExtras) cruiseControlSinceVersion() uint16 {
	return 0
}

func (o *OptionalExtras) cruiseControlInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.cruiseControlSinceVersion()
}

func (*OptionalExtras) cruiseControlDeprecated() uint16 {
	return 0
}

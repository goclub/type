package xtype

import "time"

type OptionString struct {
	string string
	valid bool
}
func (o OptionString) Valid() bool { return o.valid }
func (o OptionString) Unwrap() string { if o.valid {return o.string}; return "" }
func String(s string) OptionString { return OptionString{ s, true} }

type OptionBool struct {
	bool bool
	valid bool
}
func (o OptionBool) Valid() bool { return o.valid }
func (o OptionBool) Unwrap() bool { if o.valid {return o.bool}; return false }
func Bool(b bool) OptionBool { return OptionBool{ b, true} }

type OptionInt8 struct {
	int8 int8
	valid bool
}
func (o OptionInt8) Valid() bool { return o.valid }
func (o OptionInt8) Unwrap() int8 { if o.valid {return o.int8}; return 0 }
func Int8(i int8) OptionInt8 { return OptionInt8{ i, true} }

type OptionUint8 struct {
	uint8 uint8
	valid bool
}
func (o OptionUint8) Valid() bool { return o.valid }
func (o OptionUint8) Unwrap() uint8 { if o.valid {return o.uint8}; return 0 }
func Uint8(i uint8) OptionUint8 { return OptionUint8{ i, true} }


type OptionInt16 struct {
	int16 int16
	valid bool
}
func (o OptionInt16) Valid() bool { return o.valid }
func (o OptionInt16) Unwrap() int16 { if o.valid {return o.int16}; return 0 }
func Int16(i int16) OptionInt16 { return OptionInt16{ i, true} }

type OptionUint16 struct {
	uint16 uint16
	valid bool
}
func (o OptionUint16) Valid() bool { return o.valid }
func (o OptionUint16) Unwrap() uint16 { if o.valid {return o.uint16}; return 0 }
func Uint16(i uint16) OptionUint16 { return OptionUint16{ i, true} }


type OptionInt32 struct {
	int32 int32
	valid bool
}
func (o OptionInt32) Valid() bool { return o.valid }
func (o OptionInt32) Unwrap() int32 { if o.valid {return o.int32}; return 0 }
func Int32(i int32) OptionInt32 { return OptionInt32{ i, true} }

type OptionUint32 struct {
	uint32 uint32
	valid bool
}
func (o OptionUint32) Valid() bool { return o.valid }
func (o OptionUint32) Unwrap() uint32 { if o.valid {return o.uint32}; return 0 }
func Uint32(i uint32) OptionUint32 { return OptionUint32{ i, true} }


type OptionInt64 struct {
	int64 int64
	valid bool
}
func (o OptionInt64) Valid() bool { return o.valid }
func (o OptionInt64) Unwrap() int64 { if o.valid {return o.int64}; return 0 }
func Int64(i int64) OptionInt64 { return OptionInt64{ i, true} }

type OptionUint64 struct {
	uint64 uint64
	valid bool
}
func (o OptionUint64) Valid() bool { return o.valid }
func (o OptionUint64) Unwrap() uint64 { if o.valid {return o.uint64}; return 0 }
func Uint64(i uint64) OptionUint64 { return OptionUint64{ i, true} }

type OptionInt struct {
	int int
	valid bool
}
func (o OptionInt) Valid() bool { return o.valid }
func (o OptionInt) Unwrap() int { if o.valid {return o.int}; return 0 }
func Int(i int) OptionInt { return OptionInt{ i, true} }


type OptionUint struct {
	uint uint
	valid bool
}
func (o OptionUint) Valid() bool { return o.valid }
func (o OptionUint) Unwrap() uint { if o.valid {return o.uint}; return 0 }
func Uint(i uint) OptionUint { return OptionUint{ i, true} }



type OptionFloat32 struct {
	float32 float32
	valid bool
}
func (o OptionFloat32) Valid() bool { return o.valid }
func (o OptionFloat32) Unwrap() float32 { if o.valid {return o.float32}; return 0 }
func Float32(f float32) OptionFloat32 { return OptionFloat32{ f, true} }


type OptionFloat64 struct {
	float64 float64
	valid bool
}
func (o OptionFloat64) Valid() bool { return o.valid }
func (o OptionFloat64) Unwrap() float64 { if o.valid {return o.float64}; return 0 }
func Float64(f float64) OptionFloat64 { return OptionFloat64{ f, true} }

type OptionComplex64 struct {
	complex64 complex64
	valid bool
}
func (o OptionComplex64) Valid() bool { return o.valid }
func (o OptionComplex64) Unwrap() complex64 { if o.valid {return o.complex64}; return 0 }
func Complex64(c complex64) OptionComplex64 { return OptionComplex64{ c, true} }


type OptionComplex128 struct {
	complex128 complex128
	valid bool
}
func (o OptionComplex128) Valid() bool { return o.valid }
func (o OptionComplex128) Unwrap() complex128 { if o.valid {return o.complex128}; return 0 }
func Complex128(c complex128) OptionComplex128 { return OptionComplex128{ c, true} }


type OptionBytes struct {
	bytes []byte
	valid bool
}
func (o OptionBytes) Valid() bool { return o.valid }
func (o OptionBytes) Unwrap() []byte { if o.valid {return o.bytes}; return nil }
func Bytes(bytes []byte) OptionBytes { return OptionBytes{ bytes, true} }


type OptionRunes struct {
	runes []rune
	valid bool
}
func (o OptionRunes) Valid() bool { return o.valid }
func (o OptionRunes) Unwrap() []rune { if o.valid {return o.runes}; return nil }
func Runes(runes []rune) OptionRunes { return OptionRunes{ runes, true} }

type OptionStrings struct {
	strings []string
	valid bool
}
func (o OptionStrings) Valid() bool { return o.valid }
func (o OptionStrings) Unwrap() []string { if o.valid {return o.strings}; return nil }
func Strings(strings []string) OptionStrings { return OptionStrings{ strings, true} }

type OptionTime struct {
	time time.Time
	valid bool
}
func (o OptionTime) Valid() bool { return o.valid }
func (o OptionTime) Unwrap() time.Time { if o.valid {return o.time}; return time.Time{} }
func Time(t time.Time) OptionTime { return OptionTime{ t, true} }

type OptionDuration struct {
	duration time.Duration
	valid bool
}
func (o OptionDuration) Valid() bool { return o.valid }
func (o OptionDuration) Unwrap() time.Duration { if o.valid { return o.duration}; return time.Duration(0) }
func Duration(d time.Duration) OptionDuration { return OptionDuration{d, true} }
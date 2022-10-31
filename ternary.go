package go_tricks

type If bool

func IIF(cond bool, vtrue, vfalse interface{}) interface{} {
	if cond {
		return vtrue
	}
	return vfalse
}

func (t If) String(vtrue, vfalse string) string {
	if t {
		return vtrue
	}
	return vfalse
}

func (t If) Int(vtrue, vfalse int) int {
	if t {
		return vtrue
	}
	return vfalse
}

func (t If) Int8(vtrue, vfalse int8) int8 {
	if t {
		return vtrue
	}
	return vfalse
}

func (t If) Int16(vtrue, vfalse int16) int16 {
	if t {
		return vtrue
	}
	return vfalse
}

func (t If) Int32(vtrue, vfalse int32) int32 {
	if t {
		return vtrue
	}
	return vfalse
}

func (t If) Int64(vtrue, vfalse int64) int64 {
	if t {
		return vtrue
	}
	return vfalse
}

func (t If) Uint(vtrue, vfalse uint) uint {
	if t {
		return vtrue
	}
	return vfalse
}

func (t If) Uint8(vtrue, vfalse uint8) uint8 {
	if t {
		return vtrue
	}
	return vfalse
}

func (t If) Uint16(vtrue, vfalse uint16) uint16 {
	if t {
		return vtrue
	}
	return vfalse
}

func (t If) Uint32(vtrue, vfalse uint32) uint32 {
	if t {
		return vtrue
	}
	return vfalse
}

func (t If) Uint64(vtrue, vfalse uint64) uint64 {
	if t {
		return vtrue
	}
	return vfalse
}

func (t If) Float32(vtrue, vfalse float32) float32 {
	if t {
		return vtrue
	}
	return vfalse
}

func (t If) Float64(vtrue, vfalse float64) float64 {
	if t {
		return vtrue
	}
	return vfalse
}

func (t If) Error(vtrue, vfalse error) error {
	if t {
		return vtrue
	}
	return vfalse
}

func (t If) Bool(vtrue, vfalse bool) bool {
	if t {
		return vtrue
	}
	return vfalse
}

func (t If) Bytes(vtrue, vfalse []byte) []byte {
	if t {
		return vtrue
	}
	return vfalse
}

func (t If) Rune(vtrue, vfalse rune) rune {
	if t {
		return vtrue
	}
	return vfalse
}

func (t If) Byte(vtrue, vfalse byte) byte {
	if t {
		return vtrue
	}
	return vfalse
}

func (t If) Any(vtrue, vfalse interface{}) interface{} {
	if t {
		return vtrue
	}
	return vfalse
}

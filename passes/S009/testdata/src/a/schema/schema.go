package schema

import (
	"strconv"
)

type Schema struct {
	Type         ValueType
	ValidateFunc func()
}

type ValueType int

const (
	TypeInvalid ValueType = iota
	TypeBool
	TypeInt
	TypeFloat
	TypeString
	TypeList
	TypeMap
	TypeSet
)

const _ValueType_name = "TypeInvalidTypeBoolTypeIntTypeFloatTypeStringTypeListTypeMapTypeSettypeObject"

var _ValueType_index = [...]uint8{0, 11, 19, 26, 35, 45, 53, 60, 67, 77}

func (i ValueType) String() string {
	if i < 0 || i >= ValueType(len(_ValueType_index)-1) {
		return "ValueType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ValueType_name[_ValueType_index[i]:_ValueType_index[i+1]]
}

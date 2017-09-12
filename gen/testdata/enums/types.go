// Code generated by thriftrw v1.7.0. DO NOT EDIT.
// @generated

package enums

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go.uber.org/thriftrw/wire"
	"math"
	"strconv"
	"strings"
)

type EmptyEnum int32

func (v *EmptyEnum) UnmarshalText(value []byte) error {
	switch string(value) {
	default:
		return fmt.Errorf("unknown enum value %q for %q", value, "EmptyEnum")
	}
}

func (v EmptyEnum) ToWire() (wire.Value, error) {
	return wire.NewValueI32(int32(v)), nil
}

func (v *EmptyEnum) FromWire(w wire.Value) error {
	*v = (EmptyEnum)(w.GetI32())
	return nil
}

func (v EmptyEnum) String() string {
	w := int32(v)
	return fmt.Sprintf("EmptyEnum(%d)", w)
}

func (v EmptyEnum) Equals(rhs EmptyEnum) bool {
	return v == rhs
}

func (v EmptyEnum) MarshalJSON() ([]byte, error) {
	return ([]byte)(strconv.FormatInt(int64(v), 10)), nil
}

func (v *EmptyEnum) UnmarshalJSON(text []byte) error {
	d := json.NewDecoder(bytes.NewReader(text))
	d.UseNumber()
	t, err := d.Token()
	if err != nil {
		return err
	}

	switch w := t.(type) {
	case json.Number:
		x, err := w.Int64()
		if err != nil {
			return err
		}
		if x > math.MaxInt32 {
			return fmt.Errorf("enum overflow from JSON %q for %q", text, "EmptyEnum")
		}
		if x < math.MinInt32 {
			return fmt.Errorf("enum underflow from JSON %q for %q", text, "EmptyEnum")
		}
		*v = (EmptyEnum)(x)
		return nil
	case string:
		return v.UnmarshalText([]byte(w))
	default:
		return fmt.Errorf("invalid JSON value %q (%T) to unmarshal into %q", t, t, "EmptyEnum")
	}
}

type EnumDefault int32

const (
	EnumDefaultFoo EnumDefault = 0
	EnumDefaultBar EnumDefault = 1
	EnumDefaultBaz EnumDefault = 2
)

func EnumDefault_Values() []EnumDefault {
	return []EnumDefault{
		EnumDefaultFoo,
		EnumDefaultBar,
		EnumDefaultBaz,
	}
}

func (v *EnumDefault) UnmarshalText(value []byte) error {
	switch string(value) {
	case "Foo":
		*v = EnumDefaultFoo
		return nil
	case "Bar":
		*v = EnumDefaultBar
		return nil
	case "Baz":
		*v = EnumDefaultBaz
		return nil
	default:
		return fmt.Errorf("unknown enum value %q for %q", value, "EnumDefault")
	}
}

func (v EnumDefault) ToWire() (wire.Value, error) {
	return wire.NewValueI32(int32(v)), nil
}

func (v *EnumDefault) FromWire(w wire.Value) error {
	*v = (EnumDefault)(w.GetI32())
	return nil
}

func (v EnumDefault) String() string {
	w := int32(v)
	switch w {
	case 0:
		return "Foo"
	case 1:
		return "Bar"
	case 2:
		return "Baz"
	}
	return fmt.Sprintf("EnumDefault(%d)", w)
}

func (v EnumDefault) Equals(rhs EnumDefault) bool {
	return v == rhs
}

func (v EnumDefault) MarshalJSON() ([]byte, error) {
	switch int32(v) {
	case 0:
		return ([]byte)("\"Foo\""), nil
	case 1:
		return ([]byte)("\"Bar\""), nil
	case 2:
		return ([]byte)("\"Baz\""), nil
	}
	return ([]byte)(strconv.FormatInt(int64(v), 10)), nil
}

func (v *EnumDefault) UnmarshalJSON(text []byte) error {
	d := json.NewDecoder(bytes.NewReader(text))
	d.UseNumber()
	t, err := d.Token()
	if err != nil {
		return err
	}

	switch w := t.(type) {
	case json.Number:
		x, err := w.Int64()
		if err != nil {
			return err
		}
		if x > math.MaxInt32 {
			return fmt.Errorf("enum overflow from JSON %q for %q", text, "EnumDefault")
		}
		if x < math.MinInt32 {
			return fmt.Errorf("enum underflow from JSON %q for %q", text, "EnumDefault")
		}
		*v = (EnumDefault)(x)
		return nil
	case string:
		return v.UnmarshalText([]byte(w))
	default:
		return fmt.Errorf("invalid JSON value %q (%T) to unmarshal into %q", t, t, "EnumDefault")
	}
}

type EnumWithDuplicateName int32

const (
	EnumWithDuplicateNameA EnumWithDuplicateName = 0
	EnumWithDuplicateNameB EnumWithDuplicateName = 1
	EnumWithDuplicateNameC EnumWithDuplicateName = 2
	EnumWithDuplicateNameP EnumWithDuplicateName = 3
	EnumWithDuplicateNameQ EnumWithDuplicateName = 4
	EnumWithDuplicateNameR EnumWithDuplicateName = 5
	EnumWithDuplicateNameX EnumWithDuplicateName = 6
	EnumWithDuplicateNameY EnumWithDuplicateName = 7
	EnumWithDuplicateNameZ EnumWithDuplicateName = 8
)

func EnumWithDuplicateName_Values() []EnumWithDuplicateName {
	return []EnumWithDuplicateName{
		EnumWithDuplicateNameA,
		EnumWithDuplicateNameB,
		EnumWithDuplicateNameC,
		EnumWithDuplicateNameP,
		EnumWithDuplicateNameQ,
		EnumWithDuplicateNameR,
		EnumWithDuplicateNameX,
		EnumWithDuplicateNameY,
		EnumWithDuplicateNameZ,
	}
}

func (v *EnumWithDuplicateName) UnmarshalText(value []byte) error {
	switch string(value) {
	case "A":
		*v = EnumWithDuplicateNameA
		return nil
	case "B":
		*v = EnumWithDuplicateNameB
		return nil
	case "C":
		*v = EnumWithDuplicateNameC
		return nil
	case "P":
		*v = EnumWithDuplicateNameP
		return nil
	case "Q":
		*v = EnumWithDuplicateNameQ
		return nil
	case "R":
		*v = EnumWithDuplicateNameR
		return nil
	case "X":
		*v = EnumWithDuplicateNameX
		return nil
	case "Y":
		*v = EnumWithDuplicateNameY
		return nil
	case "Z":
		*v = EnumWithDuplicateNameZ
		return nil
	default:
		return fmt.Errorf("unknown enum value %q for %q", value, "EnumWithDuplicateName")
	}
}

func (v EnumWithDuplicateName) ToWire() (wire.Value, error) {
	return wire.NewValueI32(int32(v)), nil
}

func (v *EnumWithDuplicateName) FromWire(w wire.Value) error {
	*v = (EnumWithDuplicateName)(w.GetI32())
	return nil
}

func (v EnumWithDuplicateName) String() string {
	w := int32(v)
	switch w {
	case 0:
		return "A"
	case 1:
		return "B"
	case 2:
		return "C"
	case 3:
		return "P"
	case 4:
		return "Q"
	case 5:
		return "R"
	case 6:
		return "X"
	case 7:
		return "Y"
	case 8:
		return "Z"
	}
	return fmt.Sprintf("EnumWithDuplicateName(%d)", w)
}

func (v EnumWithDuplicateName) Equals(rhs EnumWithDuplicateName) bool {
	return v == rhs
}

func (v EnumWithDuplicateName) MarshalJSON() ([]byte, error) {
	switch int32(v) {
	case 0:
		return ([]byte)("\"A\""), nil
	case 1:
		return ([]byte)("\"B\""), nil
	case 2:
		return ([]byte)("\"C\""), nil
	case 3:
		return ([]byte)("\"P\""), nil
	case 4:
		return ([]byte)("\"Q\""), nil
	case 5:
		return ([]byte)("\"R\""), nil
	case 6:
		return ([]byte)("\"X\""), nil
	case 7:
		return ([]byte)("\"Y\""), nil
	case 8:
		return ([]byte)("\"Z\""), nil
	}
	return ([]byte)(strconv.FormatInt(int64(v), 10)), nil
}

func (v *EnumWithDuplicateName) UnmarshalJSON(text []byte) error {
	d := json.NewDecoder(bytes.NewReader(text))
	d.UseNumber()
	t, err := d.Token()
	if err != nil {
		return err
	}

	switch w := t.(type) {
	case json.Number:
		x, err := w.Int64()
		if err != nil {
			return err
		}
		if x > math.MaxInt32 {
			return fmt.Errorf("enum overflow from JSON %q for %q", text, "EnumWithDuplicateName")
		}
		if x < math.MinInt32 {
			return fmt.Errorf("enum underflow from JSON %q for %q", text, "EnumWithDuplicateName")
		}
		*v = (EnumWithDuplicateName)(x)
		return nil
	case string:
		return v.UnmarshalText([]byte(w))
	default:
		return fmt.Errorf("invalid JSON value %q (%T) to unmarshal into %q", t, t, "EnumWithDuplicateName")
	}
}

type EnumWithDuplicateValues int32

const (
	EnumWithDuplicateValuesP EnumWithDuplicateValues = 0
	EnumWithDuplicateValuesQ EnumWithDuplicateValues = -1
	EnumWithDuplicateValuesR EnumWithDuplicateValues = 0
)

func EnumWithDuplicateValues_Values() []EnumWithDuplicateValues {
	return []EnumWithDuplicateValues{
		EnumWithDuplicateValuesP,
		EnumWithDuplicateValuesQ,
		EnumWithDuplicateValuesR,
	}
}

func (v *EnumWithDuplicateValues) UnmarshalText(value []byte) error {
	switch string(value) {
	case "P":
		*v = EnumWithDuplicateValuesP
		return nil
	case "Q":
		*v = EnumWithDuplicateValuesQ
		return nil
	case "R":
		*v = EnumWithDuplicateValuesR
		return nil
	default:
		return fmt.Errorf("unknown enum value %q for %q", value, "EnumWithDuplicateValues")
	}
}

func (v EnumWithDuplicateValues) ToWire() (wire.Value, error) {
	return wire.NewValueI32(int32(v)), nil
}

func (v *EnumWithDuplicateValues) FromWire(w wire.Value) error {
	*v = (EnumWithDuplicateValues)(w.GetI32())
	return nil
}

func (v EnumWithDuplicateValues) String() string {
	w := int32(v)
	switch w {
	case 0:
		return "P"
	case -1:
		return "Q"
	}
	return fmt.Sprintf("EnumWithDuplicateValues(%d)", w)
}

func (v EnumWithDuplicateValues) Equals(rhs EnumWithDuplicateValues) bool {
	return v == rhs
}

func (v EnumWithDuplicateValues) MarshalJSON() ([]byte, error) {
	switch int32(v) {
	case 0:
		return ([]byte)("\"P\""), nil
	case -1:
		return ([]byte)("\"Q\""), nil
	}
	return ([]byte)(strconv.FormatInt(int64(v), 10)), nil
}

func (v *EnumWithDuplicateValues) UnmarshalJSON(text []byte) error {
	d := json.NewDecoder(bytes.NewReader(text))
	d.UseNumber()
	t, err := d.Token()
	if err != nil {
		return err
	}

	switch w := t.(type) {
	case json.Number:
		x, err := w.Int64()
		if err != nil {
			return err
		}
		if x > math.MaxInt32 {
			return fmt.Errorf("enum overflow from JSON %q for %q", text, "EnumWithDuplicateValues")
		}
		if x < math.MinInt32 {
			return fmt.Errorf("enum underflow from JSON %q for %q", text, "EnumWithDuplicateValues")
		}
		*v = (EnumWithDuplicateValues)(x)
		return nil
	case string:
		return v.UnmarshalText([]byte(w))
	default:
		return fmt.Errorf("invalid JSON value %q (%T) to unmarshal into %q", t, t, "EnumWithDuplicateValues")
	}
}

type EnumWithValues int32

const (
	EnumWithValuesX EnumWithValues = 123
	EnumWithValuesY EnumWithValues = 456
	EnumWithValuesZ EnumWithValues = 789
)

func EnumWithValues_Values() []EnumWithValues {
	return []EnumWithValues{
		EnumWithValuesX,
		EnumWithValuesY,
		EnumWithValuesZ,
	}
}

func (v *EnumWithValues) UnmarshalText(value []byte) error {
	switch string(value) {
	case "X":
		*v = EnumWithValuesX
		return nil
	case "Y":
		*v = EnumWithValuesY
		return nil
	case "Z":
		*v = EnumWithValuesZ
		return nil
	default:
		return fmt.Errorf("unknown enum value %q for %q", value, "EnumWithValues")
	}
}

func (v EnumWithValues) ToWire() (wire.Value, error) {
	return wire.NewValueI32(int32(v)), nil
}

func (v *EnumWithValues) FromWire(w wire.Value) error {
	*v = (EnumWithValues)(w.GetI32())
	return nil
}

func (v EnumWithValues) String() string {
	w := int32(v)
	switch w {
	case 123:
		return "X"
	case 456:
		return "Y"
	case 789:
		return "Z"
	}
	return fmt.Sprintf("EnumWithValues(%d)", w)
}

func (v EnumWithValues) Equals(rhs EnumWithValues) bool {
	return v == rhs
}

func (v EnumWithValues) MarshalJSON() ([]byte, error) {
	switch int32(v) {
	case 123:
		return ([]byte)("\"X\""), nil
	case 456:
		return ([]byte)("\"Y\""), nil
	case 789:
		return ([]byte)("\"Z\""), nil
	}
	return ([]byte)(strconv.FormatInt(int64(v), 10)), nil
}

func (v *EnumWithValues) UnmarshalJSON(text []byte) error {
	d := json.NewDecoder(bytes.NewReader(text))
	d.UseNumber()
	t, err := d.Token()
	if err != nil {
		return err
	}

	switch w := t.(type) {
	case json.Number:
		x, err := w.Int64()
		if err != nil {
			return err
		}
		if x > math.MaxInt32 {
			return fmt.Errorf("enum overflow from JSON %q for %q", text, "EnumWithValues")
		}
		if x < math.MinInt32 {
			return fmt.Errorf("enum underflow from JSON %q for %q", text, "EnumWithValues")
		}
		*v = (EnumWithValues)(x)
		return nil
	case string:
		return v.UnmarshalText([]byte(w))
	default:
		return fmt.Errorf("invalid JSON value %q (%T) to unmarshal into %q", t, t, "EnumWithValues")
	}
}

// Kinds of records stored in the database.
type RecordType int32

const (
	// Name of the user.
	RecordTypeName RecordType = 0
	// Home address of the user.
	//
	// This record is always present.
	RecordTypeHomeAddress RecordType = 1
	// Home address of the user.
	//
	// This record may not be present.
	RecordTypeWorkAddress RecordType = 2
)

func RecordType_Values() []RecordType {
	return []RecordType{
		RecordTypeName,
		RecordTypeHomeAddress,
		RecordTypeWorkAddress,
	}
}

func (v *RecordType) UnmarshalText(value []byte) error {
	switch string(value) {
	case "NAME":
		*v = RecordTypeName
		return nil
	case "HOME_ADDRESS":
		*v = RecordTypeHomeAddress
		return nil
	case "WORK_ADDRESS":
		*v = RecordTypeWorkAddress
		return nil
	default:
		return fmt.Errorf("unknown enum value %q for %q", value, "RecordType")
	}
}

func (v RecordType) ToWire() (wire.Value, error) {
	return wire.NewValueI32(int32(v)), nil
}

func (v *RecordType) FromWire(w wire.Value) error {
	*v = (RecordType)(w.GetI32())
	return nil
}

func (v RecordType) String() string {
	w := int32(v)
	switch w {
	case 0:
		return "NAME"
	case 1:
		return "HOME_ADDRESS"
	case 2:
		return "WORK_ADDRESS"
	}
	return fmt.Sprintf("RecordType(%d)", w)
}

func (v RecordType) Equals(rhs RecordType) bool {
	return v == rhs
}

func (v RecordType) MarshalJSON() ([]byte, error) {
	switch int32(v) {
	case 0:
		return ([]byte)("\"NAME\""), nil
	case 1:
		return ([]byte)("\"HOME_ADDRESS\""), nil
	case 2:
		return ([]byte)("\"WORK_ADDRESS\""), nil
	}
	return ([]byte)(strconv.FormatInt(int64(v), 10)), nil
}

func (v *RecordType) UnmarshalJSON(text []byte) error {
	d := json.NewDecoder(bytes.NewReader(text))
	d.UseNumber()
	t, err := d.Token()
	if err != nil {
		return err
	}

	switch w := t.(type) {
	case json.Number:
		x, err := w.Int64()
		if err != nil {
			return err
		}
		if x > math.MaxInt32 {
			return fmt.Errorf("enum overflow from JSON %q for %q", text, "RecordType")
		}
		if x < math.MinInt32 {
			return fmt.Errorf("enum underflow from JSON %q for %q", text, "RecordType")
		}
		*v = (RecordType)(x)
		return nil
	case string:
		return v.UnmarshalText([]byte(w))
	default:
		return fmt.Errorf("invalid JSON value %q (%T) to unmarshal into %q", t, t, "RecordType")
	}
}

type RecordTypeValues int32

const (
	RecordTypeValuesFoo RecordTypeValues = 0
	RecordTypeValuesBar RecordTypeValues = 1
)

func RecordTypeValues_Values() []RecordTypeValues {
	return []RecordTypeValues{
		RecordTypeValuesFoo,
		RecordTypeValuesBar,
	}
}

func (v *RecordTypeValues) UnmarshalText(value []byte) error {
	switch string(value) {
	case "FOO":
		*v = RecordTypeValuesFoo
		return nil
	case "BAR":
		*v = RecordTypeValuesBar
		return nil
	default:
		return fmt.Errorf("unknown enum value %q for %q", value, "RecordTypeValues")
	}
}

func (v RecordTypeValues) ToWire() (wire.Value, error) {
	return wire.NewValueI32(int32(v)), nil
}

func (v *RecordTypeValues) FromWire(w wire.Value) error {
	*v = (RecordTypeValues)(w.GetI32())
	return nil
}

func (v RecordTypeValues) String() string {
	w := int32(v)
	switch w {
	case 0:
		return "FOO"
	case 1:
		return "BAR"
	}
	return fmt.Sprintf("RecordTypeValues(%d)", w)
}

func (v RecordTypeValues) Equals(rhs RecordTypeValues) bool {
	return v == rhs
}

func (v RecordTypeValues) MarshalJSON() ([]byte, error) {
	switch int32(v) {
	case 0:
		return ([]byte)("\"FOO\""), nil
	case 1:
		return ([]byte)("\"BAR\""), nil
	}
	return ([]byte)(strconv.FormatInt(int64(v), 10)), nil
}

func (v *RecordTypeValues) UnmarshalJSON(text []byte) error {
	d := json.NewDecoder(bytes.NewReader(text))
	d.UseNumber()
	t, err := d.Token()
	if err != nil {
		return err
	}

	switch w := t.(type) {
	case json.Number:
		x, err := w.Int64()
		if err != nil {
			return err
		}
		if x > math.MaxInt32 {
			return fmt.Errorf("enum overflow from JSON %q for %q", text, "RecordTypeValues")
		}
		if x < math.MinInt32 {
			return fmt.Errorf("enum underflow from JSON %q for %q", text, "RecordTypeValues")
		}
		*v = (RecordTypeValues)(x)
		return nil
	case string:
		return v.UnmarshalText([]byte(w))
	default:
		return fmt.Errorf("invalid JSON value %q (%T) to unmarshal into %q", t, t, "RecordTypeValues")
	}
}

type StructWithOptionalEnum struct {
	E *EnumDefault `json:"e,omitempty"`
}

func (v *StructWithOptionalEnum) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.E != nil {
		w, err = v.E.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _EnumDefault_Read(w wire.Value) (EnumDefault, error) {
	var v EnumDefault
	err := v.FromWire(w)
	return v, err
}

func (v *StructWithOptionalEnum) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TI32 {
				var x EnumDefault
				x, err = _EnumDefault_Read(field.Value)
				v.E = &x
				if err != nil {
					return err
				}

			}
		}
	}

	return nil
}

func (v *StructWithOptionalEnum) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	if v.E != nil {
		fields[i] = fmt.Sprintf("E: %v", *(v.E))
		i++
	}

	return fmt.Sprintf("StructWithOptionalEnum{%v}", strings.Join(fields[:i], ", "))
}

func _EnumDefault_EqualsPtr(lhs, rhs *EnumDefault) bool {
	if lhs != nil && rhs != nil {

		x := *lhs
		y := *rhs
		return x.Equals(y)
	}
	return lhs == nil && rhs == nil
}

func (v *StructWithOptionalEnum) Equals(rhs *StructWithOptionalEnum) bool {
	if !_EnumDefault_EqualsPtr(v.E, rhs.E) {
		return false
	}

	return true
}

func (v *StructWithOptionalEnum) GetE() (o EnumDefault) {
	if v.E != nil {
		return *v.E
	}

	return
}

type LowerCaseEnum int32

const (
	LowerCaseEnumContaining LowerCaseEnum = 0
	LowerCaseEnumLowerCase  LowerCaseEnum = 1
	LowerCaseEnumItems      LowerCaseEnum = 2
)

func LowerCaseEnum_Values() []LowerCaseEnum {
	return []LowerCaseEnum{
		LowerCaseEnumContaining,
		LowerCaseEnumLowerCase,
		LowerCaseEnumItems,
	}
}

func (v *LowerCaseEnum) UnmarshalText(value []byte) error {
	switch string(value) {
	case "containing":
		*v = LowerCaseEnumContaining
		return nil
	case "lower_case":
		*v = LowerCaseEnumLowerCase
		return nil
	case "items":
		*v = LowerCaseEnumItems
		return nil
	default:
		return fmt.Errorf("unknown enum value %q for %q", value, "LowerCaseEnum")
	}
}

func (v LowerCaseEnum) ToWire() (wire.Value, error) {
	return wire.NewValueI32(int32(v)), nil
}

func (v *LowerCaseEnum) FromWire(w wire.Value) error {
	*v = (LowerCaseEnum)(w.GetI32())
	return nil
}

func (v LowerCaseEnum) String() string {
	w := int32(v)
	switch w {
	case 0:
		return "containing"
	case 1:
		return "lower_case"
	case 2:
		return "items"
	}
	return fmt.Sprintf("LowerCaseEnum(%d)", w)
}

func (v LowerCaseEnum) Equals(rhs LowerCaseEnum) bool {
	return v == rhs
}

func (v LowerCaseEnum) MarshalJSON() ([]byte, error) {
	switch int32(v) {
	case 0:
		return ([]byte)("\"containing\""), nil
	case 1:
		return ([]byte)("\"lower_case\""), nil
	case 2:
		return ([]byte)("\"items\""), nil
	}
	return ([]byte)(strconv.FormatInt(int64(v), 10)), nil
}

func (v *LowerCaseEnum) UnmarshalJSON(text []byte) error {
	d := json.NewDecoder(bytes.NewReader(text))
	d.UseNumber()
	t, err := d.Token()
	if err != nil {
		return err
	}

	switch w := t.(type) {
	case json.Number:
		x, err := w.Int64()
		if err != nil {
			return err
		}
		if x > math.MaxInt32 {
			return fmt.Errorf("enum overflow from JSON %q for %q", text, "LowerCaseEnum")
		}
		if x < math.MinInt32 {
			return fmt.Errorf("enum underflow from JSON %q for %q", text, "LowerCaseEnum")
		}
		*v = (LowerCaseEnum)(x)
		return nil
	case string:
		return v.UnmarshalText([]byte(w))
	default:
		return fmt.Errorf("invalid JSON value %q (%T) to unmarshal into %q", t, t, "LowerCaseEnum")
	}
}

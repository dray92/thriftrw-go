// Code generated by thriftrw v1.5.0. DO NOT EDIT.
// @generated

package services

import (
	"errors"
	"fmt"
	"go.uber.org/thriftrw/gen/testdata/exceptions"
	"go.uber.org/thriftrw/gen/testdata/unions"
	"go.uber.org/thriftrw/wire"
	"strings"
)

type KeyValue_GetValue_Args struct {
	Key *Key `json:"key,omitempty"`
}

func (v *KeyValue_GetValue_Args) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	if v.Key != nil {
		w, err = v.Key.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func (v *KeyValue_GetValue_Args) FromWire(w wire.Value) error {
	var err error
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				var x Key
				x, err = _Key_Read(field.Value)
				v.Key = &x
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (v *KeyValue_GetValue_Args) String() string {
	if v == nil {
		return "<nil>"
	}
	var fields [1]string
	i := 0
	if v.Key != nil {
		fields[i] = fmt.Sprintf("Key: %v", *(v.Key))
		i++
	}
	return fmt.Sprintf("KeyValue_GetValue_Args{%v}", strings.Join(fields[:i], ", "))
}

func (v *KeyValue_GetValue_Args) Equals(rhs *KeyValue_GetValue_Args) bool {
	if !_Key_EqualsPtr(v.Key, rhs.Key) {
		return false
	}
	return true
}

func (v *KeyValue_GetValue_Args) MethodName() string {
	return "getValue"
}

func (v *KeyValue_GetValue_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

var KeyValue_GetValue_Helper = struct {
	Args           func(key *Key) *KeyValue_GetValue_Args
	IsException    func(error) bool
	WrapResponse   func(*unions.ArbitraryValue, error) (*KeyValue_GetValue_Result, error)
	UnwrapResponse func(*KeyValue_GetValue_Result) (*unions.ArbitraryValue, error)
}{}

func init() {
	KeyValue_GetValue_Helper.Args = func(key *Key) *KeyValue_GetValue_Args {
		return &KeyValue_GetValue_Args{Key: key}
	}
	KeyValue_GetValue_Helper.IsException = func(err error) bool {
		switch err.(type) {
		case *exceptions.DoesNotExistException:
			return true
		default:
			return false
		}
	}
	KeyValue_GetValue_Helper.WrapResponse = func(success *unions.ArbitraryValue, err error) (*KeyValue_GetValue_Result, error) {
		if err == nil {
			return &KeyValue_GetValue_Result{Success: success}, nil
		}
		switch e := err.(type) {
		case *exceptions.DoesNotExistException:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for KeyValue_GetValue_Result.DoesNotExist")
			}
			return &KeyValue_GetValue_Result{DoesNotExist: e}, nil
		}
		return nil, err
	}
	KeyValue_GetValue_Helper.UnwrapResponse = func(result *KeyValue_GetValue_Result) (success *unions.ArbitraryValue, err error) {
		if result.DoesNotExist != nil {
			err = result.DoesNotExist
			return
		}
		if result.Success != nil {
			success = result.Success
			return
		}
		err = errors.New("expected a non-void result")
		return
	}
}

type KeyValue_GetValue_Result struct {
	Success      *unions.ArbitraryValue            `json:"success,omitempty"`
	DoesNotExist *exceptions.DoesNotExistException `json:"doesNotExist,omitempty"`
}

func (v *KeyValue_GetValue_Result) ToWire() (wire.Value, error) {
	var (
		fields [2]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	if v.Success != nil {
		w, err = v.Success.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 0, Value: w}
		i++
	}
	if v.DoesNotExist != nil {
		w, err = v.DoesNotExist.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}
	if i != 1 {
		return wire.Value{}, fmt.Errorf("KeyValue_GetValue_Result should have exactly one field: got %v fields", i)
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func (v *KeyValue_GetValue_Result) FromWire(w wire.Value) error {
	var err error
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 0:
			if field.Value.Type() == wire.TStruct {
				v.Success, err = _ArbitraryValue_Read(field.Value)
				if err != nil {
					return err
				}
			}
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.DoesNotExist, err = _DoesNotExistException_Read(field.Value)
				if err != nil {
					return err
				}
			}
		}
	}
	count := 0
	if v.Success != nil {
		count++
	}
	if v.DoesNotExist != nil {
		count++
	}
	if count != 1 {
		return fmt.Errorf("KeyValue_GetValue_Result should have exactly one field: got %v fields", count)
	}
	return nil
}

func (v *KeyValue_GetValue_Result) String() string {
	if v == nil {
		return "<nil>"
	}
	var fields [2]string
	i := 0
	if v.Success != nil {
		fields[i] = fmt.Sprintf("Success: %v", v.Success)
		i++
	}
	if v.DoesNotExist != nil {
		fields[i] = fmt.Sprintf("DoesNotExist: %v", v.DoesNotExist)
		i++
	}
	return fmt.Sprintf("KeyValue_GetValue_Result{%v}", strings.Join(fields[:i], ", "))
}

func (v *KeyValue_GetValue_Result) Equals(rhs *KeyValue_GetValue_Result) bool {
	if !((v.Success == nil && rhs.Success == nil) || (v.Success != nil && rhs.Success != nil && v.Success.Equals(rhs.Success))) {
		return false
	}
	if !((v.DoesNotExist == nil && rhs.DoesNotExist == nil) || (v.DoesNotExist != nil && rhs.DoesNotExist != nil && v.DoesNotExist.Equals(rhs.DoesNotExist))) {
		return false
	}
	return true
}

func (v *KeyValue_GetValue_Result) MethodName() string {
	return "getValue"
}

func (v *KeyValue_GetValue_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}

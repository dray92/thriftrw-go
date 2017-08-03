// Code generated by thriftrw v1.5.0. DO NOT EDIT.
// @generated

package typedefs

import (
	"go.uber.org/thriftrw/gen/testdata/enums"
	"go.uber.org/thriftrw/gen/testdata/structs"
	"go.uber.org/thriftrw/thriftreflect"
)

var ThriftModule = &thriftreflect.ThriftModule{Name: "typedefs", Package: "go.uber.org/thriftrw/gen/testdata/typedefs", FilePath: "typedefs.thrift", SHA1: "6c080659c8c233951ce263aca6081a6fea54a386", Includes: []*thriftreflect.ThriftModule{enums.ThriftModule, structs.ThriftModule}, Raw: rawIDL}

const rawIDL = "include \"./structs.thrift\"\ninclude \"./enums.thrift\"\n\ntypedef i64 Timestamp  // alias of primitive\ntypedef string State\n\ntypedef i128 UUID  // alias of struct\n\ntypedef list<Event> EventGroup  // alias fo collection\n\nstruct i128 {\n    1: required i64 high\n    2: required i64 low\n}\n\nstruct Event {\n    1: required UUID uuid  // required typedef\n    2: optional Timestamp time  // optional typedef\n}\n\nstruct Transition {\n    1: required State fromState\n    2: required State toState\n    3: optional EventGroup events\n}\n\ntypedef binary PDF  // alias of []byte\n\ntypedef set<structs.Frame> FrameGroup\n\ntypedef map<structs.Point, structs.Point> PointMap\n\ntypedef set<binary> BinarySet\n\ntypedef map<structs.Edge, structs.Edge> EdgeMap\n\ntypedef enums.EnumWithValues MyEnum\n"

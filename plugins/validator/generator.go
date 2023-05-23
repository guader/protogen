package validator

import (
	"fmt"
	"strings"

	"google.golang.org/protobuf/compiler/protogen"

	"github.com/guader/protogen/pb/validator"
)

const (
	eq operator = "=="
	ne operator = "!="
	lt operator = "<"
	gt operator = ">"
	le operator = "<="
	ge operator = ">="
)

type operator string

func (op operator) opposite() operator {
	return map[operator]operator{
		eq: ne,
		ne: eq,
		lt: ge,
		gt: le,
		le: gt,
		ge: lt,
	}[op]
}

func generateRepeatedRules(f *protogen.Field, r *validator.FieldRulesRepeated) []string {
	if r == nil {
		return nil
	}
	var (
		vs    = [6]*uint64{r.Eq, r.Ne, r.Lt, r.Gt, r.Le, r.Ge}
		ops   = [6]operator{eq, ne, lt, gt, le, ge}
		codes []string
	)
	for i, v := range vs {
		if v == nil {
			continue
		}
		op := ops[i]
		msg := fmt.Sprintf("%s.%s field length %s %v.", f.Parent.Desc.Name(), f.GoName, op, *v)
		code := fmt.Sprintf(`// %s
if len(x.Get%s()) %s %v {
	return errors.New(%#v)
}`,
			msg,
			f.GoName, op.opposite(), *v,
			msg,
		)
		codes = append(codes, code)
	}
	return codes
}

func generateEnumRules(f *protogen.Field, r *validator.FieldRulesEnum) []string {
	if r == nil {
		return nil
	}
	var (
		vss   = [2][]int32{r.In, r.NotIn}
		ops   = [2]string{"in", "not in"}
		conds = [2]string{"!ok", "ok"}
		codes []string
	)
	if r.GetDefined() && r.In == nil {
		vs := make([]int32, len(f.Enum.Values))
		for i, e := range f.Enum.Values {
			vs[i] = int32(e.Desc.Number())
		}
		vss[0] = vs
	}
	for i, vs := range vss {
		if len(vs) == 0 {
			continue
		}
		op := ops[i]
		msg := fmt.Sprintf("%s.%s value %s %#v.", f.Parent.Desc.Name(), f.GoName, op, vs)
		kvs := make([]string, len(vs))
		for j, v := range vs {
			kvs[j] = fmt.Sprintf("%#v: {}", v)
		}
		code := fmt.Sprintf(`// %s
if _, ok := map[%s]struct{}{
	%s,
}[int32(v)]; %s {
	return errors.New(%#v)
}`,
			msg,
			fmt.Sprintf("%T", vs[0]),
			strings.Join(kvs, ","),
			conds[i],
			msg,
		)
		codes = append(codes, code)
	}
	return codes
}

func generateOperatorRules[T comparable](f *protogen.Field, eqVal, neVal, ltVal, gtVal, leVal, geVal *T) []string {
	var (
		vs    = [6]*T{eqVal, neVal, ltVal, gtVal, leVal, geVal}
		ops   = [6]operator{eq, ne, lt, gt, le, ge}
		codes []string
	)
	for i, v := range vs {
		if v == nil {
			continue
		}
		op := ops[i]
		msg := fmt.Sprintf("%s.%s value %s %v.", f.Parent.Desc.Name(), f.GoName, op, *v)
		code := fmt.Sprintf(`// %s
if v %s %v {
	return errors.New(%#v)
}`,
			msg,
			op.opposite(), *v,
			msg,
		)
		codes = append(codes, code)
	}
	return codes
}

func generateStringRules(f *protogen.Field, eqVal, neVal, ltVal, gtVal, leVal, geVal *uint64) []string {
	var (
		vs    = [6]*uint64{eqVal, neVal, ltVal, gtVal, leVal, geVal}
		ops   = [6]operator{eq, ne, lt, gt, le, ge}
		codes []string
	)
	for i, v := range vs {
		if v == nil {
			continue
		}
		op := ops[i]
		msg := fmt.Sprintf("%s.%s value rune count %s %v.", f.Parent.Desc.Name(), f.GoName, op, *v)
		code := fmt.Sprintf(`// %s
if utf8.RuneCountInString(v) %s %v {
	return errors.New(%#v)
}`,
			msg,
			op.opposite(), *v,
			msg,
		)
		codes = append(codes, code)
	}
	return codes
}

func generateExistenceRules[T comparable](f *protogen.Field, inVal, notInVal []T) []string {
	var (
		vss   = [2][]T{inVal, notInVal}
		ops   = [2]string{"in", "not in"}
		conds = [2]string{"!ok", "ok"}
		codes []string
	)
	for i, vs := range vss {
		if len(vs) == 0 {
			continue
		}
		op := ops[i]
		msg := fmt.Sprintf("%s.%s value %s %#v.", f.Parent.Desc.Name(), f.GoName, op, vs)
		kvs := make([]string, len(vs))
		for j, v := range vs {
			kvs[j] = fmt.Sprintf("%#v: {}", v)
		}
		code := fmt.Sprintf(`// %s
if _, ok := map[%s]struct{}{
	%s,
}[v]; %s {
	return errors.New(%#v)
}`,
			msg,
			fmt.Sprintf("%T", vs[0]),
			strings.Join(kvs, ","),
			conds[i],
			msg,
		)
		codes = append(codes, code)
	}
	return codes
}

func generateBytesRules(f *protogen.Field, eqVal, neVal, ltVal, gtVal, leVal, geVal *uint64) []string {
	var (
		vs    = [6]*uint64{eqVal, neVal, ltVal, gtVal, leVal, geVal}
		ops   = [6]operator{eq, ne, lt, gt, le, ge}
		codes []string
	)
	for i, v := range vs {
		if v == nil {
			continue
		}
		op := ops[i]
		msg := fmt.Sprintf("%s.%s value length %s %v.", f.Parent.Desc.Name(), f.GoName, op, *v)
		code := fmt.Sprintf(`// %s
if len(v) %s %v {
	return errors.New(%#v)
}`,
			msg,
			op.opposite(), *v,
			msg,
		)
		codes = append(codes, code)
	}
	return codes
}

func generateMessageRules(f *protogen.Field) string {
	msg := fmt.Sprintf("%s.%s validation.", f.Parent.Desc.Name(), f.GoName)
	return fmt.Sprintf(`// %s
if i, ok := any(v).(interface{ Validate() error }); ok {
	if err := i.Validate(); err != nil {
		return err
	}
}`,
		msg,
	)
}

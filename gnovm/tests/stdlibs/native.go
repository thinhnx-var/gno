// This file is autogenerated from the genstd tool (@/misc/genstd); do not edit.
// To regenerate it, run `go generate` from this directory.

package stdlibs

import (
	"reflect"

	gno "github.com/gnolang/gno/gnovm/pkg/gnolang"
	testlibs_std "github.com/gnolang/gno/gnovm/tests/stdlibs/std"
	testlibs_testing "github.com/gnolang/gno/gnovm/tests/stdlibs/testing"
)

type nativeFunc struct {
	gnoPkg  string
	gnoFunc gno.Name
	params  []gno.FieldTypeExpr
	results []gno.FieldTypeExpr
	f       func(m *gno.Machine)
}

var nativeFuncs = [...]nativeFunc{
	{
		"std",
		"AssertOriginCall",
		[]gno.FieldTypeExpr{},
		[]gno.FieldTypeExpr{},
		func(m *gno.Machine) {
			testlibs_std.AssertOriginCall(
				m,
			)
		},
	},
	{
		"std",
		"IsOriginCall",
		[]gno.FieldTypeExpr{},
		[]gno.FieldTypeExpr{
			{Name: gno.N("r0"), Type: gno.X("bool")},
		},
		func(m *gno.Machine) {
			r0 := testlibs_std.IsOriginCall(
				m,
			)

			m.PushValue(gno.Go2GnoValue(
				m.Alloc,
				m.Store,
				reflect.ValueOf(&r0).Elem(),
			))
		},
	},
	{
		"std",
		"TestCurrentRealm",
		[]gno.FieldTypeExpr{},
		[]gno.FieldTypeExpr{
			{Name: gno.N("r0"), Type: gno.X("string")},
		},
		func(m *gno.Machine) {
			r0 := testlibs_std.TestCurrentRealm(
				m,
			)

			m.PushValue(gno.Go2GnoValue(
				m.Alloc,
				m.Store,
				reflect.ValueOf(&r0).Elem(),
			))
		},
	},
	{
		"std",
		"TestSkipHeights",
		[]gno.FieldTypeExpr{
			{Name: gno.N("p0"), Type: gno.X("int64")},
		},
		[]gno.FieldTypeExpr{},
		func(m *gno.Machine) {
			b := m.LastBlock()
			var (
				p0  int64
				rp0 = reflect.ValueOf(&p0).Elem()
			)

			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 0, "")).TV, rp0)

			testlibs_std.TestSkipHeights(
				m,
				p0)
		},
	},
	{
		"std",
		"ClearStoreCache",
		[]gno.FieldTypeExpr{},
		[]gno.FieldTypeExpr{},
		func(m *gno.Machine) {
			testlibs_std.ClearStoreCache(
				m,
			)
		},
	},
	{
		"std",
		"callerAt",
		[]gno.FieldTypeExpr{
			{Name: gno.N("p0"), Type: gno.X("int")},
		},
		[]gno.FieldTypeExpr{
			{Name: gno.N("r0"), Type: gno.X("string")},
		},
		func(m *gno.Machine) {
			b := m.LastBlock()
			var (
				p0  int
				rp0 = reflect.ValueOf(&p0).Elem()
			)

			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 0, "")).TV, rp0)

			r0 := testlibs_std.X_callerAt(
				m,
				p0)

			m.PushValue(gno.Go2GnoValue(
				m.Alloc,
				m.Store,
				reflect.ValueOf(&r0).Elem(),
			))
		},
	},
	{
		"std",
		"testSetOrigCaller",
		[]gno.FieldTypeExpr{
			{Name: gno.N("p0"), Type: gno.X("string")},
		},
		[]gno.FieldTypeExpr{},
		func(m *gno.Machine) {
			b := m.LastBlock()
			var (
				p0  string
				rp0 = reflect.ValueOf(&p0).Elem()
			)

			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 0, "")).TV, rp0)

			testlibs_std.X_testSetOrigCaller(
				m,
				p0)
		},
	},
	{
		"std",
		"testSetOrigPkgAddr",
		[]gno.FieldTypeExpr{
			{Name: gno.N("p0"), Type: gno.X("string")},
		},
		[]gno.FieldTypeExpr{},
		func(m *gno.Machine) {
			b := m.LastBlock()
			var (
				p0  string
				rp0 = reflect.ValueOf(&p0).Elem()
			)

			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 0, "")).TV, rp0)

			testlibs_std.X_testSetOrigPkgAddr(
				m,
				p0)
		},
	},
	{
		"std",
		"testSetPrevRealm",
		[]gno.FieldTypeExpr{
			{Name: gno.N("p0"), Type: gno.X("string")},
		},
		[]gno.FieldTypeExpr{},
		func(m *gno.Machine) {
			b := m.LastBlock()
			var (
				p0  string
				rp0 = reflect.ValueOf(&p0).Elem()
			)

			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 0, "")).TV, rp0)

			testlibs_std.X_testSetPrevRealm(
				m,
				p0)
		},
	},
	{
		"std",
		"testSetPrevAddr",
		[]gno.FieldTypeExpr{
			{Name: gno.N("p0"), Type: gno.X("string")},
		},
		[]gno.FieldTypeExpr{},
		func(m *gno.Machine) {
			b := m.LastBlock()
			var (
				p0  string
				rp0 = reflect.ValueOf(&p0).Elem()
			)

			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 0, "")).TV, rp0)

			testlibs_std.X_testSetPrevAddr(
				m,
				p0)
		},
	},
	{
		"std",
		"testSetOrigSend",
		[]gno.FieldTypeExpr{
			{Name: gno.N("p0"), Type: gno.X("[]string")},
			{Name: gno.N("p1"), Type: gno.X("[]int64")},
			{Name: gno.N("p2"), Type: gno.X("[]string")},
			{Name: gno.N("p3"), Type: gno.X("[]int64")},
		},
		[]gno.FieldTypeExpr{},
		func(m *gno.Machine) {
			b := m.LastBlock()
			var (
				p0  []string
				rp0 = reflect.ValueOf(&p0).Elem()
				p1  []int64
				rp1 = reflect.ValueOf(&p1).Elem()
				p2  []string
				rp2 = reflect.ValueOf(&p2).Elem()
				p3  []int64
				rp3 = reflect.ValueOf(&p3).Elem()
			)

			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 0, "")).TV, rp0)
			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 1, "")).TV, rp1)
			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 2, "")).TV, rp2)
			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 3, "")).TV, rp3)

			testlibs_std.X_testSetOrigSend(
				m,
				p0, p1, p2, p3)
		},
	},
	{
		"std",
		"testIssueCoins",
		[]gno.FieldTypeExpr{
			{Name: gno.N("p0"), Type: gno.X("string")},
			{Name: gno.N("p1"), Type: gno.X("[]string")},
			{Name: gno.N("p2"), Type: gno.X("[]int64")},
		},
		[]gno.FieldTypeExpr{},
		func(m *gno.Machine) {
			b := m.LastBlock()
			var (
				p0  string
				rp0 = reflect.ValueOf(&p0).Elem()
				p1  []string
				rp1 = reflect.ValueOf(&p1).Elem()
				p2  []int64
				rp2 = reflect.ValueOf(&p2).Elem()
			)

			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 0, "")).TV, rp0)
			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 1, "")).TV, rp1)
			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 2, "")).TV, rp2)

			testlibs_std.X_testIssueCoins(
				m,
				p0, p1, p2)
		},
	},
	{
		"testing",
		"unixNano",
		[]gno.FieldTypeExpr{},
		[]gno.FieldTypeExpr{
			{Name: gno.N("r0"), Type: gno.X("int64")},
		},
		func(m *gno.Machine) {
			r0 := testlibs_testing.X_unixNano()

			m.PushValue(gno.Go2GnoValue(
				m.Alloc,
				m.Store,
				reflect.ValueOf(&r0).Elem(),
			))
		},
	},
}

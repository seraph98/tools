// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typeparams

type (
	Int     int
	Uintptr = uintptr
	String  string
)

func _[AllString ~string, MaybeString ~string | ~int, NotString ~int | byte, NamedString String | Int]() {
	var (
		i int
		r rune
		b byte
		I Int
		U uintptr
		M MaybeString
		N NotString
	)
	const p = 0

	_ = MaybeString(i) // want `conversion from int to string .in MaybeString. yields a string of one rune, not a string of digits`
	_ = MaybeString(r)
	_ = MaybeString(b)
	_ = MaybeString(I) // want `conversion from Int .int. to string .in MaybeString. yields a string of one rune, not a string of digits`
	_ = MaybeString(U) // want `conversion from uintptr to string .in MaybeString. yields a string of one rune, not a string of digits`
	// Type parameters are never constant types, so arguments are always
	// converted to their default type (int versus untyped int, in this case)
	_ = MaybeString(p) // want `conversion from int to string .in MaybeString. yields a string of one rune, not a string of digits`
	// ...even if the type parameter is only strings.
	_ = AllString(p) // want `conversion from int to string .in AllString. yields a string of one rune, not a string of digits`

	_ = NotString(i)
	_ = NotString(r)
	_ = NotString(b)
	_ = NotString(I)
	_ = NotString(U)
	_ = NotString(p)

	_ = NamedString(i) // want `conversion from int to String .string, in NamedString. yields a string of one rune, not a string of digits`
	_ = string(M)      // want `conversion from int .in MaybeString. to string yields a string of one rune, not a string of digits`

	// Note that M is not convertible to rune.
	_ = MaybeString(M) // want `conversion from int .in MaybeString. to string .in MaybeString. yields a string of one rune, not a string of digits`
	_ = NotString(N)   // ok
}

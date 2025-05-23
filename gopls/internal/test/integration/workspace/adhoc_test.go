// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package workspace

import (
	"testing"

	. "golang.org/x/tools/gopls/internal/test/integration"
)

// Test for golang/go#57209: editing a file in an ad-hoc package should not
// trigger conflicting diagnostics.
func TestAdhoc_Edits(t *testing.T) {
	const files = `
-- a.go --
package foo

const X = 1

-- b.go --
package foo

// import "errors"

const Y = X
`

	Run(t, files, func(t *testing.T, env *Env) {
		env.OpenFile("b.go")

		for range 10 {
			env.RegexpReplace("b.go", `// import "errors"`, `import "errors"`)
			env.RegexpReplace("b.go", `import "errors"`, `// import "errors"`)
			env.AfterChange(NoDiagnostics())
		}
	})
}

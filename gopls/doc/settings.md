---
title: "Gopls: Settings"
---

This document describes gopls' configuration settings.

Gopls settings are defined by a JSON object whose valid fields are
described below. These fields are gopls-specific, and generic LSP
clients have no knowledge of them.

Different clients present configuration settings in their user
interfaces in a wide variety of ways.
For example, some expect the user to edit the raw JSON object while
others use a data structure in the editor's configuration language;
still others (such as VS Code) have a graphical configuration system.
Be sure to consult the documentation for how to express configuration
settings in your client.
Some clients also permit settings to be configured differently for
each workspace folder.

Any settings that are experimental or for debugging purposes are
marked as such.

<!--
All settings are uniquely identified by name such as `semanticTokens`
or `templateExtensions`.
However, for convenience of VS Code, each setting also has an
undocumented alias whose form is a dotted path such as
`ui.semanticTokens` or `build.templateExtensions`.
However, only the final segment is actually significant, so
`build.templateExtensions` is equivalent to `templateExtensions`.
All clients but VS Code should use the short form.
-->

<!-- This portion is generated by doc/generate from the ../internal/settings package. -->
<!-- BEGIN User: DO NOT MANUALLY EDIT THIS SECTION -->

* [Build](#build)
* [Formatting](#formatting)
* [UI](#ui)
  * [Completion](#completion)
  * [Diagnostic](#diagnostic)
  * [Documentation](#documentation)
  * [Inlayhint](#inlayhint)
  * [Navigation](#navigation)

<a id='build'></a>
## Build

<a id='buildFlags'></a>
### `buildFlags []string`

buildFlags is the set of flags passed on to the build system when invoked.
It is applied to queries like `go list`, which is used when discovering files.
The most common use is to set `-tags`.

Default: `[]`.

<a id='env'></a>
### `env map[string]string`

env adds environment variables to external commands run by `gopls`, most notably `go list`.

Default: `{}`.

<a id='directoryFilters'></a>
### `directoryFilters []string`

directoryFilters can be used to exclude unwanted directories from the
workspace. By default, all directories are included. Filters are an
operator, `+` to include and `-` to exclude, followed by a path prefix
relative to the workspace folder. They are evaluated in order, and
the last filter that applies to a path controls whether it is included.
The path prefix can be empty, so an initial `-` excludes everything.

DirectoryFilters also supports the `**` operator to match 0 or more directories.

Examples:

Exclude node_modules at current depth: `-node_modules`

Exclude node_modules at any depth: `-**/node_modules`

Include only project_a: `-` (exclude everything), `+project_a`

Include only project_a, but not node_modules inside it: `-`, `+project_a`, `-project_a/node_modules`

Default: `["-**/node_modules"]`.

<a id='templateExtensions'></a>
### `templateExtensions []string`

templateExtensions gives the extensions of file names that are treated
as template files. (The extension
is the part of the file name after the final dot.)

Default: `[]`.

<a id='memoryMode'></a>
### `memoryMode string`

**This setting is experimental and may be deleted.**

obsolete, no effect

Default: `""`.

<a id='expandWorkspaceToModule'></a>
### `expandWorkspaceToModule bool`

**This setting is experimental and may be deleted.**

expandWorkspaceToModule determines which packages are considered
"workspace packages" when the workspace is using modules.

Workspace packages affect the scope of workspace-wide operations. Notably,
gopls diagnoses all packages considered to be part of the workspace after
every keystroke, so by setting "ExpandWorkspaceToModule" to false, and
opening a nested workspace directory, you can reduce the amount of work
gopls has to do to keep your workspace up to date.

Default: `true`.

<a id='standaloneTags'></a>
### `standaloneTags []string`

standaloneTags specifies a set of build constraints that identify
individual Go source files that make up the entire main package of an
executable.

A common example of standalone main files is the convention of using the
directive `//go:build ignore` to denote files that are not intended to be
included in any package, for example because they are invoked directly by
the developer using `go run`.

Gopls considers a file to be a standalone main file if and only if it has
package name "main" and has a build directive of the exact form
"//go:build tag" or "// +build tag", where tag is among the list of tags
configured by this setting. Notably, if the build constraint is more
complicated than a simple tag (such as the composite constraint
`//go:build tag && go1.18`), the file is not considered to be a standalone
main file.

This setting is only supported when gopls is built with Go 1.16 or later.

Default: `["ignore"]`.

<a id='workspaceFiles'></a>
### `workspaceFiles []string`

workspaceFiles configures the set of globs that match files defining the
logical build of the current workspace. Any on-disk changes to any files
matching a glob specified here will trigger a reload of the workspace.

This setting need only be customized in environments with a custom
GOPACKAGESDRIVER.

Default: `[]`.

<a id='formatting'></a>
## Formatting

<a id='local'></a>
### `local string`

local is the equivalent of the `goimports -local` flag, which puts
imports beginning with this string after third-party packages. It should
be the prefix of the import path whose imports should be grouped
separately.

It is used when tidying imports (during an LSP Organize
Imports request) or when inserting new ones (for example,
during completion); an LSP Formatting request merely sorts the
existing imports.

Default: `""`.

<a id='gofumpt'></a>
### `gofumpt bool`

gofumpt indicates if we should run gofumpt formatting.

Default: `false`.

<a id='ui'></a>
## UI

<a id='codelenses'></a>
### `codelenses map[enum]bool`

codelenses overrides the enabled/disabled state of each of gopls'
sources of [Code Lenses](codelenses.md).

Example Usage:

```json5
"gopls": {
...
  "codelenses": {
    "generate": false,  // Don't show the `go generate` lens.
  }
...
}
```

Default: `{"generate":true,"regenerate_cgo":true,"run_govulncheck":false,"tidy":true,"upgrade_dependency":true,"vendor":true}`.

<a id='semanticTokens'></a>
### `semanticTokens bool`

**This setting is experimental and may be deleted.**

semanticTokens controls whether the LSP server will send
semantic tokens to the client.

Default: `false`.

<a id='noSemanticString'></a>
### `noSemanticString bool`

**This setting is experimental and may be deleted.**

noSemanticString turns off the sending of the semantic token 'string'

Deprecated: Use SemanticTokenTypes["string"] = false instead. See
golang/vscode-go#3632

Default: `false`.

<a id='noSemanticNumber'></a>
### `noSemanticNumber bool`

**This setting is experimental and may be deleted.**

noSemanticNumber turns off the sending of the semantic token 'number'

Deprecated: Use SemanticTokenTypes["number"] = false instead. See
golang/vscode-go#3632.

Default: `false`.

<a id='semanticTokenTypes'></a>
### `semanticTokenTypes map[string]bool`

**This setting is experimental and may be deleted.**

semanticTokenTypes configures the semantic token types. It allows
disabling types by setting each value to false.
By default, all types are enabled.

Default: `{}`.

<a id='semanticTokenModifiers'></a>
### `semanticTokenModifiers map[string]bool`

**This setting is experimental and may be deleted.**

semanticTokenModifiers configures the semantic token modifiers. It allows
disabling modifiers by setting each value to false.
By default, all modifiers are enabled.

Default: `{}`.

<a id='completion'></a>
## Completion

<a id='usePlaceholders'></a>
### `usePlaceholders bool`

placeholders enables placeholders for function parameters or struct
fields in completion responses.

Default: `false`.

<a id='completionBudget'></a>
### `completionBudget time.Duration`

**This setting is for debugging purposes only.**

completionBudget is the soft latency goal for completion requests. Most
requests finish in a couple milliseconds, but in some cases deep
completions can take much longer. As we use up our budget we
dynamically reduce the search scope to ensure we return timely
results. Zero means unlimited.

Default: `"100ms"`.

<a id='matcher'></a>
### `matcher enum`

**This is an advanced setting and should not be configured by most `gopls` users.**

matcher sets the algorithm that is used when calculating completion
candidates.

Must be one of:

* `"CaseInsensitive"`
* `"CaseSensitive"`
* `"Fuzzy"`

Default: `"Fuzzy"`.

<a id='experimentalPostfixCompletions'></a>
### `experimentalPostfixCompletions bool`

**This setting is experimental and may be deleted.**

experimentalPostfixCompletions enables artificial method snippets
such as "someSlice.sort!".

Default: `true`.

<a id='completeFunctionCalls'></a>
### `completeFunctionCalls bool`

completeFunctionCalls enables function call completion.

When completing a statement, or when a function return type matches the
expected of the expression being completed, completion may suggest call
expressions (i.e. may include parentheses).

Default: `true`.

<a id='diagnostic'></a>
## Diagnostic

<a id='analyses'></a>
### `analyses map[string]bool`

analyses specify analyses that the user would like to enable or disable.
A map of the names of analysis passes that should be enabled/disabled.
A full list of analyzers that gopls uses can be found in
[analyzers.md](https://github.com/golang/tools/blob/master/gopls/doc/analyzers.md).

Example Usage:

```json5
...
"analyses": {
  "unreachable": false, // Disable the unreachable analyzer.
  "unusedvariable": true  // Enable the unusedvariable analyzer.
}
...
```

Default: `{}`.

<a id='staticcheck'></a>
### `staticcheck bool`

**This setting is experimental and may be deleted.**

staticcheck configures the default set of analyses staticcheck.io.
These analyses are documented on
[Staticcheck's website](https://staticcheck.io/docs/checks/).

The "staticcheck" option has three values:
- false: disable all staticcheck analyzers
- true: enable all staticcheck analyzers
- unset: enable a subset of staticcheck analyzers
  selected by gopls maintainers for runtime efficiency
  and analytic precision.

Regardless of this setting, individual analyzers can be
selectively enabled or disabled using the `analyses` setting.

Default: `false`.

<a id='staticcheckProvided'></a>
### `staticcheckProvided bool`

**This setting is experimental and may be deleted.**


Default: `false`.

<a id='annotations'></a>
### `annotations map[enum]bool`

annotations specifies the various kinds of compiler
optimization details that should be reported as diagnostics
when enabled for a package by the "Toggle compiler
optimization details" (`gopls.gc_details`) command.

(Some users care only about one kind of annotation in their
profiling efforts. More importantly, in large packages, the
number of annotations can sometimes overwhelm the user
interface and exceed the per-file diagnostic limit.)

TODO(adonovan): rename this field to CompilerOptDetail.

Each enum must be one of:

* `"bounds"` controls bounds checking diagnostics.
* `"escape"` controls diagnostics about escape choices.
* `"inline"` controls diagnostics about inlining choices.
* `"nil"` controls nil checks.

Default: `{"bounds":true,"escape":true,"inline":true,"nil":true}`.

<a id='vulncheck'></a>
### `vulncheck enum`

**This setting is experimental and may be deleted.**

vulncheck enables vulnerability scanning.

Must be one of:

* `"Imports"`: In Imports mode, `gopls` will report vulnerabilities that affect packages
directly and indirectly used by the analyzed main module.
* `"Off"`: Disable vulnerability analysis.

Default: `"Off"`.

<a id='diagnosticsDelay'></a>
### `diagnosticsDelay time.Duration`

**This is an advanced setting and should not be configured by most `gopls` users.**

diagnosticsDelay controls the amount of time that gopls waits
after the most recent file modification before computing deep diagnostics.
Simple diagnostics (parsing and type-checking) are always run immediately
on recently modified packages.

This option must be set to a valid duration string, for example `"250ms"`.

Default: `"1s"`.

<a id='diagnosticsTrigger'></a>
### `diagnosticsTrigger enum`

**This setting is experimental and may be deleted.**

diagnosticsTrigger controls when to run diagnostics.

Must be one of:

* `"Edit"`: Trigger diagnostics on file edit and save. (default)
* `"Save"`: Trigger diagnostics only on file save. Events like initial workspace load
or configuration change will still trigger diagnostics.

Default: `"Edit"`.

<a id='analysisProgressReporting'></a>
### `analysisProgressReporting bool`

analysisProgressReporting controls whether gopls sends progress
notifications when construction of its index of analysis facts is taking a
long time. Cancelling these notifications will cancel the indexing task,
though it will restart after the next change in the workspace.

When a package is opened for the first time and heavyweight analyses such as
staticcheck are enabled, it can take a while to construct the index of
analysis facts for all its dependencies. The index is cached in the
filesystem, so subsequent analysis should be faster.

Default: `true`.

<a id='documentation'></a>
## Documentation

<a id='hoverKind'></a>
### `hoverKind enum`

hoverKind controls the information that appears in the hover text.
SingleLine is intended for use only by authors of editor plugins.

Must be one of:

* `"FullDocumentation"`
* `"NoDocumentation"`
* `"SingleLine"`
* `"Structured"` is a misguided experimental setting that returns a JSON
hover format. This setting should not be used, as it will be removed in a
future release of gopls.
* `"SynopsisDocumentation"`

Default: `"FullDocumentation"`.

<a id='linkTarget'></a>
### `linkTarget string`

linkTarget is the base URL for links to Go package
documentation returned by LSP operations such as Hover and
DocumentLinks and in the CodeDescription field of each
Diagnostic.

It might be one of:

* `"godoc.org"`
* `"pkg.go.dev"`

If company chooses to use its own `godoc.org`, its address can be used as well.

Modules matching the GOPRIVATE environment variable will not have
documentation links in hover.

Default: `"pkg.go.dev"`.

<a id='linksInHover'></a>
### `linksInHover enum`

linksInHover controls the presence of documentation links in hover markdown.

Must be one of:

* false: do not show links
* true: show links to the `linkTarget` domain
* `"gopls"`: show links to gopls' internal documentation viewer

Default: `true`.

<a id='inlayhint'></a>
## Inlayhint

<a id='hints'></a>
### `hints map[enum]bool`

**This setting is experimental and may be deleted.**

hints specify inlay hints that users want to see. A full list of hints
that gopls uses can be found in
[inlayHints.md](https://github.com/golang/tools/blob/master/gopls/doc/inlayHints.md).

Default: `{}`.

<a id='navigation'></a>
## Navigation

<a id='importShortcut'></a>
### `importShortcut enum`

importShortcut specifies whether import statements should link to
documentation or go to definitions.

Must be one of:

* `"Both"`
* `"Definition"`
* `"Link"`

Default: `"Both"`.

<a id='symbolMatcher'></a>
### `symbolMatcher enum`

**This is an advanced setting and should not be configured by most `gopls` users.**

symbolMatcher sets the algorithm that is used when finding workspace symbols.

Must be one of:

* `"CaseInsensitive"`
* `"CaseSensitive"`
* `"FastFuzzy"`
* `"Fuzzy"`

Default: `"FastFuzzy"`.

<a id='symbolStyle'></a>
### `symbolStyle enum`

**This is an advanced setting and should not be configured by most `gopls` users.**

symbolStyle controls how symbols are qualified in symbol responses.

Example Usage:

```json5
"gopls": {
...
  "symbolStyle": "Dynamic",
...
}
```

Must be one of:

* `"Dynamic"` uses whichever qualifier results in the highest scoring
match for the given symbol query. Here a "qualifier" is any "/" or "."
delimited suffix of the fully qualified symbol. i.e. "to/pkg.Foo.Field" or
just "Foo.Field".
* `"Full"` is fully qualified symbols, i.e.
"path/to/pkg.Foo.Field".
* `"Package"` is package qualified symbols i.e.
"pkg.Foo.Field".

Default: `"Dynamic"`.

<a id='symbolScope'></a>
### `symbolScope enum`

symbolScope controls which packages are searched for workspace/symbol
requests. When the scope is "workspace", gopls searches only workspace
packages. When the scope is "all", gopls searches all loaded packages,
including dependencies and the standard library.

Must be one of:

* `"all"` matches symbols in any loaded package, including
dependencies.
* `"workspace"` matches symbols in workspace packages only.

Default: `"all"`.

<a id='verboseOutput'></a>
### `verboseOutput bool`

**This setting is for debugging purposes only.**

verboseOutput enables additional debug logging.

Default: `false`.

<!-- END User: DO NOT MANUALLY EDIT THIS SECTION -->

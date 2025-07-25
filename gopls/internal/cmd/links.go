// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmd

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"golang.org/x/tools/gopls/internal/protocol"
	"golang.org/x/tools/internal/tool"
)

// links implements the links verb for gopls.
type links struct {
	JSON bool `flag:"json" help:"emit document links in JSON format"`

	app *Application
}

func (l *links) Name() string      { return "links" }
func (l *links) Parent() string    { return l.app.Name() }
func (l *links) Usage() string     { return "[links-flags] <filename>" }
func (l *links) ShortHelp() string { return "list links in a file" }
func (l *links) DetailedHelp(f *flag.FlagSet) {
	fmt.Fprintf(f.Output(), `
Example: list links contained within a file:

	$ gopls links internal/cmd/check.go

links-flags:
`)
	printFlagDefaults(f)
}

// Run finds all the links within a document
// - if -json is specified, outputs location range and uri
// - otherwise, prints the a list of unique links
func (l *links) Run(ctx context.Context, args ...string) error {
	if len(args) != 1 {
		return tool.CommandLineErrorf("links expects 1 argument")
	}
	cli, _, err := l.app.connect(ctx)
	if err != nil {
		return err
	}
	defer cli.terminate(ctx)

	from := parseSpan(args[0])
	uri := from.URI()

	if _, err := cli.openFile(ctx, uri); err != nil {
		return err
	}
	results, err := cli.server.DocumentLink(ctx, &protocol.DocumentLinkParams{
		TextDocument: protocol.TextDocumentIdentifier{
			URI: uri,
		},
	})
	if err != nil {
		return fmt.Errorf("%v: %v", from, err)
	}
	if l.JSON {
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "\t")
		return enc.Encode(results)
	}
	for _, v := range results {
		fmt.Println(*v.Target)
	}
	return nil
}

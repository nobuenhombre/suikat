#!/bin/bash

## https://pkg.go.dev/golang.org/x/tools/cmd/goimports?tab=doc
goimports -v -w $(go list -f {{.Dir}} ./...)

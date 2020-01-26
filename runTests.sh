#!/bin/bash
go test ./tests/...

# go test -v ./tests/... # -v means verbose, which prints more details about the tests. It prints what the call of t.Logf() and other testing methods
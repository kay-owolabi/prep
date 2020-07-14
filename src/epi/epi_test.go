package main

import (
	chapter5 "epi/arrays"
	chapter4 "epi/primitivetypes"
	"testing"
)

func TestAll(t *testing.T) {
	chapter4.TestParity(t)
	chapter5.Test()
}

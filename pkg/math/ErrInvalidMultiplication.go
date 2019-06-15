// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package math

import (
	"fmt"
)

type ErrInvalidMultiplication struct {
	A interface{}
	B interface{}
}

// Error returns the error formatted as a string.
func (e ErrInvalidMultiplication) Error() string {
	return fmt.Sprintf("cannot multiply %q (%T) and %q (%T)", e.A, e.A, e.B, e.B)
}

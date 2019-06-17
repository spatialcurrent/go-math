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

type ErrInvalidSubtraction struct {
	A interface{}
	B interface{}
}

// Error returns the error formatted as a string.
func (e ErrInvalidSubtraction) Error() string {
	return fmt.Sprintf("cannot subtract %q (%T) from %q (%T)", e.B, e.B, e.A, e.A)
}

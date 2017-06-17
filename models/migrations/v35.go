// Copyright 2017 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package migrations

import (
	"fmt"

	"github.com/go-xorm/xorm"
)

func addCommentIdToAction(x *xorm.Engine) error {
	// Action see models/action.go
	type Action struct {
		CommentID   int64 `xorm:"INDEX"`
	}


	if err := x.Sync2(new(Action)); err != nil {
		return fmt.Errorf("Sync2: %v", err)
	}

	return nil
}

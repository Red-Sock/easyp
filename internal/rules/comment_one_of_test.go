package rules_test

import (
	"errors"
	"testing"

	"github.com/easyp-tech/easyp/internal/core"
	"github.com/easyp-tech/easyp/internal/rules"
)

func TestCommentOneOf_Validate(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		fileName string
		wantErr  error
	}{
		"auth_rpc_comment_is_empty": {
			fileName: invalidAuthProto,
			wantErr:  core.ErrOneOfCommentIsEmpty,
		},
	}

	for name, tc := range tests {
		name, tc := name, tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			r, protos := start(t)

			commentServiceRule := rules.CommentOneOf{}
			err := commentServiceRule.Validate(protos[tc.fileName])
			r.ErrorIs(errors.Join(err...), tc.wantErr)
		})
	}
}

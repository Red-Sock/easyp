package rules

import (
	"github.com/easyp-tech/easyp/internal/lint"
)

var _ lint.Rule = (*CommentEnumValue)(nil)

// CommentEnumValue is a rule for checking enum values comments.
type CommentEnumValue struct{}

// Validate implements core.Rule.
func (c *CommentEnumValue) Validate(protoInfo lint.ProtoInfo) []error {
	var res []error

	for _, enum := range protoInfo.Info.ProtoBody.Enums {
		for _, field := range enum.EnumBody.EnumFields {
			if len(field.Comments) == 0 {
				res = append(res, buildError(field.Meta.Pos, field.Ident, lint.ErrEnumValueCommentIsEmpty))
			}
		}
	}

	if len(res) == 0 {
		return nil
	}

	return res
}
package rules_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yoheimuta/go-protoparser/v4"
	"github.com/yoheimuta/go-protoparser/v4/interpret/unordered"

	"github.com/easyp-tech/easyp/internal/core"
)

const (
	invalidAuthProto = `./../../testdata/auth/service.proto`
	validAuthProto   = `./../../testdata/api/session/v1/session.proto`
)

func start(t testing.TB) (*require.Assertions, map[string]core.ProtoInfo) {
	t.Helper()

	assert := require.New(t)

	protos := map[string]core.ProtoInfo{
		invalidAuthProto: parseFile(t, assert, invalidAuthProto),
		validAuthProto:   parseFile(t, assert, validAuthProto),
	}

	return assert, protos
}

func parseFile(t testing.TB, assert *require.Assertions, path string) core.ProtoInfo {
	t.Helper()

	f, err := os.Open(path)
	assert.NoError(err)
	t.Cleanup(func() { assert.NoError(f.Close()) })

	got, err := protoparser.Parse(f)
	assert.NoError(err)

	res, err := unordered.InterpretProto(got)
	assert.NoError(err)

	return core.ProtoInfo{
		Path: path,
		Info: res,
	}
}
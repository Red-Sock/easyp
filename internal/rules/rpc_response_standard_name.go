package rules

import (
	"go.redsock.ru/protopack/internal/core"
)

var _ core.Rule = (*RPCResponseStandardName)(nil)

// RPCResponseStandardName checks that RPC response type names are RPCNameResponse or ServiceNameRPCNameResponse.
type RPCResponseStandardName struct {
}

// Message implements lint.Rule.
func (r *RPCResponseStandardName) Message() string {
	return "rpc response should have suffix 'Response'"
}

// Validate implements lint.Rule.
func (r *RPCResponseStandardName) Validate(protoInfo core.ProtoInfo) ([]core.Issue, error) {
	var res []core.Issue

	for _, service := range protoInfo.Info.ProtoBody.Services {
		for _, rpc := range service.ServiceBody.RPCs {
			if rpc.RPCResponse.MessageType != rpc.RPCName+"Response" && rpc.RPCResponse.MessageType != service.ServiceName+rpc.RPCName+"Response" {
				res = core.AppendIssue(res, r, rpc.Meta.Pos, rpc.RPCResponse.MessageType, rpc.Comments)
			}
		}
	}

	return res, nil
}

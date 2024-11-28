package config

import (
	"fmt"
	"testing"
)

func TestEval(t *testing.T) {
	var ctx EvalContextImpl = EvalContextImpl{
		url: "http://www.google.com",
		queryParams: map[string]string{
			"data": "hello",
		},
		headers: map[string]string{
			"host": "wwww.google.com",
		},
		operators: map[string]NodeEvalFunc{
			"and":      AndOp,
			"protocol": ProtocolOp,
			"hostname": HostNameOp,
		},
	}

	var rootExpression ExpressionNode = ExpressionNode{
		Operator: "and",
		Children: []*ExpressionNode{
			&ExpressionNode{
				Operator: "protocol",
				Children: []*ExpressionNode{&ExpressionNode{Operator: "", Value: "http", Children: nil}},
			},
			&ExpressionNode{
				Operator: "hostname",
				Children: []*ExpressionNode{&ExpressionNode{Operator: "", Value: "www.google.com", Children: nil}},
			},
		},
	}

	var result = Eval(&ctx, &rootExpression)
	fmt.Println(result)
}

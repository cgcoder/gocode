package config

import (
	"fmt"
	"strings"
)

func SayHello() {
	fmt.Println("Hello, World!")
}

type NodeEvalFunc func(ctx EvalContext, parent *ExpressionNode, node *ExpressionNode, childValues []any) any

type ExpressionNode struct {
	Operator string
	Value    string
	Children []*ExpressionNode
}

type EvalContext interface {
	GetUrl() string
	GetQueryParameter(paramName string) string
	GetHeader(headerKey string) string

	GetEvaluator(operator string) NodeEvalFunc
}

func Eval(ctx EvalContext, root *ExpressionNode) any {
	return EvalNode(ctx, nil, root)
}

func EvalNode(ctx EvalContext, parent *ExpressionNode, node *ExpressionNode) any {
	if len(strings.TrimSpace(node.Operator)) == 0 {
		return node.Value
	}

	var result []any = make([]any, len(node.Children))
	for i := 0; i < len(result); i++ {
		var childResult = EvalNode(ctx, node, node.Children[i])
		result[i] = childResult
	}

	evalFunc := ctx.GetEvaluator((node.Operator))

	return evalFunc(ctx, parent, node, result)
}

type EvalContextImpl struct {
	url         string
	queryParams map[string]string
	headers     map[string]string
	operators   map[string]NodeEvalFunc
}

func (c *EvalContextImpl) GetUrl() string {
	return c.url
}

func (c *EvalContextImpl) GetQueryParameter(paramName string) string {
	val, ok := c.queryParams[paramName]
	if ok {
		return val
	}

	return ""
}

func (c *EvalContextImpl) GetHeader(headerKey string) string {
	val, ok := c.headers[headerKey]
	if ok {
		return val
	}

	return ""

}
func (c *EvalContextImpl) GetEvaluator(operator string) NodeEvalFunc {
	val, ok := c.operators[operator]
	if ok {
		return val
	}

	panic("Missing operator " + operator)
}

func AndOp(ctx EvalContext, parent *ExpressionNode, node *ExpressionNode, childValues []any) any {
	for _, element := range childValues {
		value, ok := element.(bool)
		if !ok || !value {
			return false
		}
	}
	return true
}

func ProtocolOp(ctx EvalContext, parent *ExpressionNode, node *ExpressionNode, childValues []any) any {
	protocol, ok := childValues[0].(string)

	if !ok {
		return false
	}

	return strings.HasPrefix(ctx.GetUrl(), protocol+":")
}

func HostNameOp(ctx EvalContext, parent *ExpressionNode, node *ExpressionNode, childValues []any) any {
	hostToMatch, ok := childValues[0].(string)

	if !ok {
		return false
	}

	url := ctx.GetUrl()
	start := strings.Index(url, "//")
	if start == -1 {
		return false
	}
	start += 2
	end := strings.Index(url[start:], "/")
	if end == -1 {
		end = len(url)
	} else {
		end += start
	}

	host := url[start:end]

	return strings.EqualFold(hostToMatch, host)
}

package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

func main() {
	in := os.Stdin
	if len(os.Args) > 1 {
		f, err := os.Open(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		in = f
	}
	pn, err := ParseReader("", in)
	if err != nil {
		log.Fatal(err)
	}
	ret, err := pn.(ProgramNode).exec()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ret)
}

func toIfaceSlice(v interface{}) []interface{} {
	if v == nil {
		return nil
	}
	return v.([]interface{})
}

var lvalues = make(map[string]int)

// element of statements
type Executor interface {
	exec() error
}

// ProgramNode
type ProgramNode struct {
	statements StatementsNode
	ret        ReturnNode
}

func newProgramNode(stmts StatementsNode, ret ReturnNode) (ProgramNode, error) {
	return ProgramNode{stmts, ret}, nil
}
func (n ProgramNode) exec() (int, error) {
	err := n.statements.exec()
	if err != nil {
		return 0, err
	}
	return n.ret.exec()
}

// StatementsNode
type StatementsNode struct {
	statements []Executor
}

func newStatementsNode(stmts interface{}) (StatementsNode, error) {

	st := toIfaceSlice(stmts)
	ex := make([]Executor, len(st))
	for i, v := range st {
		ex[i] = v.(Executor)
	}
	return StatementsNode{ex}, nil
}
func (n StatementsNode) exec() error {
	for _, v := range n.statements {
		err := v.exec()
		if err != nil {
			return err
		}
	}
	return nil
}

// ReturnNode
type ReturnNode struct {
	arg IdentifierNode
}

func newReturnNode(arg IdentifierNode) (ReturnNode, error) {
	return ReturnNode{arg}, nil
}
func (n ReturnNode) exec() (int, error) {
	v, err := n.arg.exec()
	return v, err
}

// IfNode
type IfNode struct {
	arg        LogicalExpressionNode
	statements StatementsNode
}

func newIfNode(arg LogicalExpressionNode, stmts StatementsNode) (IfNode, error) {
	return IfNode{arg, stmts}, nil
}
func (n IfNode) exec() error {
	cond, err := n.arg.exec()
	if err != nil {
		return err
	}
	if cond {
		err := n.statements.exec()
		return err
	}
	return nil
}

// AssignmentNode
type AssignmentNode struct {
	lvalue string
	rvalue AdditiveExpressionNode
}

func newAssignmentNode(lvalue IdentifierNode, rvalue AdditiveExpressionNode) (AssignmentNode, error) {
	return AssignmentNode{lvalue.val, rvalue}, nil
}
func (n AssignmentNode) exec() error {
	v, err := n.rvalue.exec()
	if err != nil {
		return err
	}
	lvalues[n.lvalue] = v
	return nil
}

// LogicalExpressionNode
type LogicalExpressionNode struct {
	expr PrimaryExpressionNode
}

func newLogicalExpressionNode(expr PrimaryExpressionNode) (LogicalExpressionNode, error) {
	return LogicalExpressionNode{expr}, nil
}
func (n LogicalExpressionNode) exec() (bool, error) {
	ret, err := n.expr.exec()
	b := ret != 0
	return b, err
}

// AdditiveExpressionNode
type AdditiveExpressionNode struct {
	arg1 interface{}
	arg2 PrimaryExpressionNode
	op   string
}

func newAdditiveExpressionNode(arg PrimaryExpressionNode, rest interface{}) (AdditiveExpressionNode, error) {
	var a AdditiveExpressionNode
	var arg1 interface{} = arg

	restSl := toIfaceSlice(rest)
	if len(restSl) == 0 {
		zero, _ := newIntegerNode("0")
		arg2, _ := newPrimaryExpressionNode(zero)
		a = AdditiveExpressionNode{arg1, arg2, "+"}
	}
	for _, v := range restSl {
		restExpr := toIfaceSlice(v)
		arg2 := restExpr[3].(PrimaryExpressionNode)
		op := restExpr[1].(string)
		a = AdditiveExpressionNode{arg1, arg2, op}
		arg1 = a
	}
	return a, nil
}
func (n AdditiveExpressionNode) exec() (int, error) {
	var v, varg1, varg2 int
	var err error
	switch n.arg1.(type) {
	case PrimaryExpressionNode:
		varg1, err = n.arg1.(PrimaryExpressionNode).exec()
	case AdditiveExpressionNode:
		varg1, err = n.arg1.(AdditiveExpressionNode).exec()
	default:
		return 0, errors.New("arg1 has invalid node type while exec AdditiveExpression")
	}
	if err != nil {
		return varg1, err
	}
	varg2, err = n.arg2.exec()
	switch n.op {
	case "+":
		v = varg1 + varg2
	case "-":
		v = varg1 - varg2
	default:
		return 0, errors.New("invalid operation while exec AdditiveExpression")
	}
	return v, err
}

// PrimaryExpressionNode
type PrimaryExpressionNode struct {
	arg interface{}
}

func newPrimaryExpressionNode(arg interface{}) (PrimaryExpressionNode, error) {
	return PrimaryExpressionNode{arg}, nil
}
func (n PrimaryExpressionNode) exec() (int, error) {
	var v int
	var err error
	switch n.arg.(type) {
	case IntegerNode:
		v, err = n.arg.(IntegerNode).exec()
	case IdentifierNode:
		v, err = n.arg.(IdentifierNode).exec()
	default:
		return 0, errors.New("invalid operation while exec AdditiveExpression")
	}
	return v, err
}

// IntegerNode
type IntegerNode struct {
	val int
}

func newIntegerNode(val string) (IntegerNode, error) {
	v, err := strconv.ParseInt(val, 0, 64)
	return IntegerNode{int(v)}, err
}
func (n IntegerNode) exec() (int, error) {
	return n.val, nil
}

// IdentifierNode
type IdentifierNode struct {
	val string
}

func newIdentifierNode(val string) (IdentifierNode, error) {
	return IdentifierNode{val}, nil
}
func (n IdentifierNode) exec() (int, error) {
	v, ok := lvalues[n.val]
	if !ok {
		return 0, errors.New("Identifier " + n.val + " not defined")
	}
	return v, nil
}

var g = &grammar{
	rules: []*rule{
		{
			name: "Input",
			pos:  position{line: 236, col: 1, offset: 5794},
			expr: &actionExpr{
				pos: position{line: 236, col: 15, offset: 5810},
				run: (*parser).callonInput1,
				expr: &seqExpr{
					pos: position{line: 236, col: 15, offset: 5810},
					exprs: []interface{}{
						&stateCodeExpr{
							pos: position{line: 236, col: 15, offset: 5810},
							run: (*parser).callonInput3,
						},
						&labeledExpr{
							pos:   position{line: 236, col: 57, offset: 5852},
							label: "s",
							expr: &ruleRefExpr{
								pos:  position{line: 236, col: 59, offset: 5854},
								name: "Statements",
							},
						},
						&labeledExpr{
							pos:   position{line: 236, col: 71, offset: 5866},
							label: "r",
							expr: &ruleRefExpr{
								pos:  position{line: 236, col: 73, offset: 5868},
								name: "ReturnOp",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 236, col: 82, offset: 5877},
							name: "EOF",
						},
					},
				},
			},
		},
		{
			name: "Statements",
			pos:  position{line: 238, col: 1, offset: 5989},
			expr: &actionExpr{
				pos: position{line: 238, col: 15, offset: 6005},
				run: (*parser).callonStatements1,
				expr: &labeledExpr{
					pos:   position{line: 238, col: 15, offset: 6005},
					label: "s",
					expr: &oneOrMoreExpr{
						pos: position{line: 238, col: 17, offset: 6007},
						expr: &ruleRefExpr{
							pos:  position{line: 238, col: 17, offset: 6007},
							name: "Line",
						},
					},
				},
			},
		},
		{
			name: "Line",
			pos:  position{line: 239, col: 1, offset: 6067},
			expr: &actionExpr{
				pos: position{line: 239, col: 15, offset: 6083},
				run: (*parser).callonLine1,
				expr: &seqExpr{
					pos: position{line: 239, col: 15, offset: 6083},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 239, col: 15, offset: 6083},
							name: "INDENTATION",
						},
						&labeledExpr{
							pos:   position{line: 239, col: 27, offset: 6095},
							label: "s",
							expr: &ruleRefExpr{
								pos:  position{line: 239, col: 29, offset: 6097},
								name: "Statement",
							},
						},
					},
				},
			},
		},
		{
			name: "ReturnOp",
			pos:  position{line: 240, col: 1, offset: 6131},
			expr: &actionExpr{
				pos: position{line: 240, col: 15, offset: 6147},
				run: (*parser).callonReturnOp1,
				expr: &seqExpr{
					pos: position{line: 240, col: 15, offset: 6147},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 240, col: 15, offset: 6147},
							val:        "return",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 240, col: 24, offset: 6156},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 240, col: 26, offset: 6158},
							label: "arg",
							expr: &ruleRefExpr{
								pos:  position{line: 240, col: 30, offset: 6162},
								name: "Identifier",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 240, col: 41, offset: 6173},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "Statement",
			pos:  position{line: 242, col: 1, offset: 6226},
			expr: &choiceExpr{
				pos: position{line: 242, col: 15, offset: 6242},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 242, col: 15, offset: 6242},
						run: (*parser).callonStatement2,
						expr: &seqExpr{
							pos: position{line: 242, col: 15, offset: 6242},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 242, col: 15, offset: 6242},
									label: "s",
									expr: &ruleRefExpr{
										pos:  position{line: 242, col: 17, offset: 6244},
										name: "Assignment",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 242, col: 28, offset: 6255},
									name: "EOL",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 243, col: 7, offset: 6313},
						run: (*parser).callonStatement7,
						expr: &seqExpr{
							pos: position{line: 243, col: 7, offset: 6313},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 243, col: 7, offset: 6313},
									val:        "if",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 243, col: 12, offset: 6318},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 243, col: 14, offset: 6320},
									label: "arg",
									expr: &ruleRefExpr{
										pos:  position{line: 243, col: 18, offset: 6324},
										name: "LogicalExpression",
									},
								},
								&zeroOrOneExpr{
									pos: position{line: 243, col: 36, offset: 6342},
									expr: &ruleRefExpr{
										pos:  position{line: 243, col: 36, offset: 6342},
										name: "_",
									},
								},
								&litMatcher{
									pos:        position{line: 243, col: 39, offset: 6345},
									val:        ":",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 243, col: 43, offset: 6349},
									name: "EOL",
								},
								&ruleRefExpr{
									pos:  position{line: 243, col: 47, offset: 6353},
									name: "INDENT",
								},
								&labeledExpr{
									pos:   position{line: 243, col: 54, offset: 6360},
									label: "s",
									expr: &ruleRefExpr{
										pos:  position{line: 243, col: 56, offset: 6362},
										name: "Statements",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 243, col: 67, offset: 6373},
									name: "DEDENT",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Assignment",
			pos:  position{line: 247, col: 1, offset: 6504},
			expr: &actionExpr{
				pos: position{line: 247, col: 14, offset: 6519},
				run: (*parser).callonAssignment1,
				expr: &seqExpr{
					pos: position{line: 247, col: 14, offset: 6519},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 247, col: 14, offset: 6519},
							label: "lvalue",
							expr: &ruleRefExpr{
								pos:  position{line: 247, col: 21, offset: 6526},
								name: "Identifier",
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 247, col: 32, offset: 6537},
							expr: &ruleRefExpr{
								pos:  position{line: 247, col: 32, offset: 6537},
								name: "_",
							},
						},
						&litMatcher{
							pos:        position{line: 247, col: 35, offset: 6540},
							val:        "=",
							ignoreCase: false,
						},
						&zeroOrOneExpr{
							pos: position{line: 247, col: 39, offset: 6544},
							expr: &ruleRefExpr{
								pos:  position{line: 247, col: 39, offset: 6544},
								name: "_",
							},
						},
						&labeledExpr{
							pos:   position{line: 247, col: 42, offset: 6547},
							label: "rvalue",
							expr: &ruleRefExpr{
								pos:  position{line: 247, col: 49, offset: 6554},
								name: "AdditiveExpression",
							},
						},
					},
				},
			},
		},
		{
			name: "LogicalExpression",
			pos:  position{line: 250, col: 1, offset: 6707},
			expr: &actionExpr{
				pos: position{line: 250, col: 23, offset: 6731},
				run: (*parser).callonLogicalExpression1,
				expr: &labeledExpr{
					pos:   position{line: 250, col: 23, offset: 6731},
					label: "arg",
					expr: &ruleRefExpr{
						pos:  position{line: 250, col: 27, offset: 6735},
						name: "PrimaryExpression",
					},
				},
			},
		},
		{
			name: "AdditiveExpression",
			pos:  position{line: 251, col: 1, offset: 6819},
			expr: &actionExpr{
				pos: position{line: 251, col: 23, offset: 6843},
				run: (*parser).callonAdditiveExpression1,
				expr: &seqExpr{
					pos: position{line: 251, col: 23, offset: 6843},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 251, col: 23, offset: 6843},
							label: "arg",
							expr: &ruleRefExpr{
								pos:  position{line: 251, col: 27, offset: 6847},
								name: "PrimaryExpression",
							},
						},
						&labeledExpr{
							pos:   position{line: 251, col: 45, offset: 6865},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 251, col: 50, offset: 6870},
								expr: &seqExpr{
									pos: position{line: 251, col: 52, offset: 6872},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 251, col: 52, offset: 6872},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 251, col: 54, offset: 6874},
											name: "AddOp",
										},
										&ruleRefExpr{
											pos:  position{line: 251, col: 60, offset: 6880},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 251, col: 62, offset: 6882},
											name: "PrimaryExpression",
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "PrimaryExpression",
			pos:  position{line: 253, col: 1, offset: 7020},
			expr: &actionExpr{
				pos: position{line: 253, col: 23, offset: 7044},
				run: (*parser).callonPrimaryExpression1,
				expr: &labeledExpr{
					pos:   position{line: 253, col: 23, offset: 7044},
					label: "arg",
					expr: &choiceExpr{
						pos: position{line: 253, col: 28, offset: 7049},
						alternatives: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 253, col: 28, offset: 7049},
								name: "Integer",
							},
							&ruleRefExpr{
								pos:  position{line: 253, col: 38, offset: 7059},
								name: "Identifier",
							},
						},
					},
				},
			},
		},
		{
			name: "Integer",
			pos:  position{line: 256, col: 1, offset: 7161},
			expr: &actionExpr{
				pos: position{line: 256, col: 11, offset: 7173},
				run: (*parser).callonInteger1,
				expr: &oneOrMoreExpr{
					pos: position{line: 256, col: 11, offset: 7173},
					expr: &charClassMatcher{
						pos:        position{line: 256, col: 11, offset: 7173},
						val:        "[0-9]",
						ranges:     []rune{'0', '9'},
						ignoreCase: false,
						inverted:   false,
					},
				},
			},
		},
		{
			name: "Identifier",
			pos:  position{line: 257, col: 1, offset: 7250},
			expr: &actionExpr{
				pos: position{line: 257, col: 14, offset: 7265},
				run: (*parser).callonIdentifier1,
				expr: &seqExpr{
					pos: position{line: 257, col: 14, offset: 7265},
					exprs: []interface{}{
						&charClassMatcher{
							pos:        position{line: 257, col: 14, offset: 7265},
							val:        "[a-zA-Z]",
							ranges:     []rune{'a', 'z', 'A', 'Z'},
							ignoreCase: false,
							inverted:   false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 257, col: 23, offset: 7274},
							expr: &charClassMatcher{
								pos:        position{line: 257, col: 23, offset: 7274},
								val:        "[a-zA-Z0-9]",
								ranges:     []rune{'a', 'z', 'A', 'Z', '0', '9'},
								ignoreCase: false,
								inverted:   false,
							},
						},
					},
				},
			},
		},
		{
			name: "AddOp",
			pos:  position{line: 259, col: 1, offset: 7344},
			expr: &actionExpr{
				pos: position{line: 259, col: 9, offset: 7354},
				run: (*parser).callonAddOp1,
				expr: &choiceExpr{
					pos: position{line: 259, col: 11, offset: 7356},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 259, col: 11, offset: 7356},
							val:        "+",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 259, col: 17, offset: 7362},
							val:        "-",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "_",
			pos:  position{line: 261, col: 1, offset: 7423},
			expr: &oneOrMoreExpr{
				pos: position{line: 261, col: 5, offset: 7429},
				expr: &charClassMatcher{
					pos:        position{line: 261, col: 5, offset: 7429},
					val:        "[ \\t]",
					chars:      []rune{' ', '\t'},
					ignoreCase: false,
					inverted:   false,
				},
			},
		},
		{
			name: "EOL",
			pos:  position{line: 263, col: 1, offset: 7439},
			expr: &seqExpr{
				pos: position{line: 263, col: 7, offset: 7447},
				exprs: []interface{}{
					&zeroOrOneExpr{
						pos: position{line: 263, col: 7, offset: 7447},
						expr: &ruleRefExpr{
							pos:  position{line: 263, col: 7, offset: 7447},
							name: "_",
						},
					},
					&zeroOrOneExpr{
						pos: position{line: 263, col: 10, offset: 7450},
						expr: &ruleRefExpr{
							pos:  position{line: 263, col: 10, offset: 7450},
							name: "Comment",
						},
					},
					&choiceExpr{
						pos: position{line: 263, col: 20, offset: 7460},
						alternatives: []interface{}{
							&litMatcher{
								pos:        position{line: 263, col: 20, offset: 7460},
								val:        "\r\n",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 263, col: 29, offset: 7469},
								val:        "\n\r",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 263, col: 38, offset: 7478},
								val:        "\r",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 263, col: 45, offset: 7485},
								val:        "\n",
								ignoreCase: false,
							},
							&ruleRefExpr{
								pos:  position{line: 263, col: 52, offset: 7492},
								name: "EOF",
							},
						},
					},
				},
			},
		},
		{
			name: "Comment",
			pos:  position{line: 265, col: 1, offset: 7500},
			expr: &seqExpr{
				pos: position{line: 265, col: 11, offset: 7512},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 265, col: 11, offset: 7512},
						val:        "//",
						ignoreCase: false,
					},
					&zeroOrMoreExpr{
						pos: position{line: 265, col: 16, offset: 7517},
						expr: &charClassMatcher{
							pos:        position{line: 265, col: 16, offset: 7517},
							val:        "[^\\r\\n]",
							chars:      []rune{'\r', '\n'},
							ignoreCase: false,
							inverted:   true,
						},
					},
				},
			},
		},
		{
			name: "EOF",
			pos:  position{line: 267, col: 1, offset: 7529},
			expr: &notExpr{
				pos: position{line: 267, col: 7, offset: 7537},
				expr: &anyMatcher{
					line: 267, col: 8, offset: 7538,
				},
			},
		},
		{
			name: "INDENTATION",
			pos:  position{line: 269, col: 1, offset: 7543},
			expr: &seqExpr{
				pos: position{line: 269, col: 15, offset: 7559},
				exprs: []interface{}{
					&labeledExpr{
						pos:   position{line: 269, col: 15, offset: 7559},
						label: "spaces",
						expr: &zeroOrMoreExpr{
							pos: position{line: 269, col: 22, offset: 7566},
							expr: &litMatcher{
								pos:        position{line: 269, col: 22, offset: 7566},
								val:        " ",
								ignoreCase: false,
							},
						},
					},
					&andCodeExpr{
						pos: position{line: 269, col: 27, offset: 7571},
						run: (*parser).callonINDENTATION5,
					},
				},
			},
		},
		{
			name: "INDENT",
			pos:  position{line: 271, col: 1, offset: 7647},
			expr: &stateCodeExpr{
				pos: position{line: 271, col: 10, offset: 7658},
				run: (*parser).callonINDENT1,
			},
		},
		{
			name: "DEDENT",
			pos:  position{line: 273, col: 1, offset: 7732},
			expr: &stateCodeExpr{
				pos: position{line: 273, col: 10, offset: 7743},
				run: (*parser).callonDEDENT1,
			},
		},
	},
}

/*c*/
func (state statedict) onInput3() error {
	state["Indentation"] = 0
	return nil
}

func (p *parser) callonInput3() (bool, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	err := p.pt.state.onInput3()
	copyState(state, p.pt.state)
	return true, err
}

func (c *current) onInput1(s, r interface{}) (interface{}, error) {
	return newProgramNode(s.(StatementsNode), r.(ReturnNode))
}

func (p *parser) callonInput1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onInput1(stack["s"], stack["r"])
}

func (c *current) onStatements1(s interface{}) (interface{}, error) {
	return newStatementsNode(s)
}

func (p *parser) callonStatements1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onStatements1(stack["s"])
}

func (c *current) onLine1(s interface{}) (interface{}, error) {
	return s, nil
}

func (p *parser) callonLine1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onLine1(stack["s"])
}

func (c *current) onReturnOp1(arg interface{}) (interface{}, error) {
	return newReturnNode(arg.(IdentifierNode))
}

func (p *parser) callonReturnOp1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onReturnOp1(stack["arg"])
}

func (c *current) onStatement2(s interface{}) (interface{}, error) {
	return s.(AssignmentNode), nil
}

func (p *parser) callonStatement2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onStatement2(stack["s"])
}

func (c *current) onStatement7(arg, s interface{}) (interface{}, error) {
	return newIfNode(arg.(LogicalExpressionNode), s.(StatementsNode))
}

func (p *parser) callonStatement7() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onStatement7(stack["arg"], stack["s"])
}

func (c *current) onAssignment1(lvalue, rvalue interface{}) (interface{}, error) {
	return newAssignmentNode(lvalue.(IdentifierNode), rvalue.(AdditiveExpressionNode))
}

func (p *parser) callonAssignment1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onAssignment1(stack["lvalue"], stack["rvalue"])
}

func (c *current) onLogicalExpression1(arg interface{}) (interface{}, error) {
	return newLogicalExpressionNode(arg.(PrimaryExpressionNode))
}

func (p *parser) callonLogicalExpression1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onLogicalExpression1(stack["arg"])
}

func (c *current) onAdditiveExpression1(arg, rest interface{}) (interface{}, error) {
	return newAdditiveExpressionNode(arg.(PrimaryExpressionNode), rest)
}

func (p *parser) callonAdditiveExpression1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onAdditiveExpression1(stack["arg"], stack["rest"])
}

func (c *current) onPrimaryExpression1(arg interface{}) (interface{}, error) {
	return newPrimaryExpressionNode(arg)
}

func (p *parser) callonPrimaryExpression1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onPrimaryExpression1(stack["arg"])
}

func (c *current) onInteger1() (interface{}, error) {
	return newIntegerNode(string(c.text))
}

func (p *parser) callonInteger1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onInteger1()
}

func (c *current) onIdentifier1() (interface{}, error) {
	return newIdentifierNode(string(c.text))
}

func (p *parser) callonIdentifier1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onIdentifier1()
}

func (c *current) onAddOp1() (interface{}, error) {
	return string(c.text), nil
}

func (p *parser) callonAddOp1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onAddOp1()
}

func (c *current) onINDENTATION5(spaces interface{}) (bool, error) {
	return len(toIfaceSlice(spaces)) == state["Indentation"].(int), nil
}

func (p *parser) callonINDENTATION5() (bool, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onINDENTATION5(stack["spaces"])
}

/*c*/
func (state statedict) onINDENT1() error {
	state["Indentation"] = state["Indentation"].(int) + 4
	return nil
}

func (p *parser) callonINDENT1() (bool, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	err := p.pt.state.onINDENT1()
	copyState(state, p.pt.state)
	return true, err
}

/*c*/
func (state statedict) onDEDENT1() error {
	state["Indentation"] = state["Indentation"].(int) - 4
	return nil
}

func (p *parser) callonDEDENT1() (bool, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	err := p.pt.state.onDEDENT1()
	copyState(state, p.pt.state)
	return true, err
}

var (
	// errNoRule is returned when the grammar to parse has no rule.
	errNoRule = errors.New("grammar has no rule")

	// errInvalidEncoding is returned when the source is not properly
	// utf8-encoded.
	errInvalidEncoding = errors.New("invalid encoding")

	// errNoMatch is returned if no match could be found.
	errNoMatch = errors.New("no match found")

	// State
	state = make(statedict)
)

// Option is a function that can set an option on the parser. It returns
// the previous setting as an Option.
type Option func(*parser) Option

// Debug creates an Option to set the debug flag to b. When set to true,
// debugging information is printed to stdout while parsing.
//
// The default is false.
func Debug(b bool) Option {
	return func(p *parser) Option {
		old := p.debug
		p.debug = b
		return Debug(old)
	}
}

// JDebug creates an Option to set the jdebug flag to b. When set to true,
// debugging information is printed to stdout  as JSON after parsing.
//
// The default is false.
func JDebug(b bool) Option {
	return func(p *parser) Option {
		old := p.jdebug
		p.jdebug = b
		return JDebug(old)
	}
}

// Memoize creates an Option to set the memoize flag to b. When set to true,
// the parser will cache all results so each expression is evaluated only
// once. This guarantees linear parsing time even for pathological cases,
// at the expense of more memory and slower times for typical cases.
//
// The default is false.
func Memoize(b bool) Option {
	return func(p *parser) Option {
		old := p.memoize
		p.memoize = b
		return Memoize(old)
	}
}

// Recover creates an Option to set the recover flag to b. When set to
// true, this causes the parser to recover from panics and convert it
// to an error. Setting it to false can be useful while debugging to
// access the full stack trace.
//
// The default is true.
func Recover(b bool) Option {
	return func(p *parser) Option {
		old := p.recover
		p.recover = b
		return Recover(old)
	}
}

// ParseFile parses the file identified by filename.
func ParseFile(filename string, opts ...Option) (interface{}, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ParseReader(filename, f, opts...)
}

// ParseReader parses the data from r using filename as information in the
// error messages.
func ParseReader(filename string, r io.Reader, opts ...Option) (interface{}, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	return Parse(filename, b, opts...)
}

// Parse parses the data from b using filename as information in the
// error messages.
func Parse(filename string, b []byte, opts ...Option) (interface{}, error) {
	return newParser(filename, b, opts...).parse(g)
}

// position records a position in the text.
type position struct {
	line, col, offset int
}

func (p position) String() string {
	return fmt.Sprintf("%d:%d [%d]", p.line, p.col, p.offset)
}

// savepoint stores all state required to go back to this point in the
// parser.
type savepoint struct {
	position
	rn    rune
	w     int
	state statedict
}

type current struct {
	pos  position // start position of the match
	text []byte   // raw text of the match
}

type statedict map[string]interface{}

// the AST types...

type grammar struct {
	pos   position
	rules []*rule
}

type rule struct {
	pos         position
	name        string
	displayName string
	expr        interface{}
}

type choiceExpr struct {
	pos          position
	alternatives []interface{}
}

type actionExpr struct {
	pos  position
	expr interface{}
	run  func(*parser) (interface{}, error)
}

type seqExpr struct {
	pos   position
	exprs []interface{}
}

type labeledExpr struct {
	pos   position
	label string
	expr  interface{}
}

type expr struct {
	pos  position
	expr interface{}
}

type andExpr expr
type notExpr expr
type zeroOrOneExpr expr
type zeroOrMoreExpr expr
type oneOrMoreExpr expr

type ruleRefExpr struct {
	pos  position
	name string
}

type andCodeExpr struct {
	pos position
	run func(*parser) (bool, error)
}

type notCodeExpr struct {
	pos position
	run func(*parser) (bool, error)
}

type stateCodeExpr struct {
	pos position
	run func(*parser) (bool, error)
}

type litMatcher struct {
	pos        position
	val        string
	ignoreCase bool
}

type charClassMatcher struct {
	pos        position
	val        string
	chars      []rune
	ranges     []rune
	classes    []*unicode.RangeTable
	ignoreCase bool
	inverted   bool
}

type anyMatcher position

// errList cumulates the errors found by the parser.
type errList []error

func (e *errList) add(err error) {
	*e = append(*e, err)
}

func (e errList) err() error {
	if len(e) == 0 {
		return nil
	}
	e.dedupe()
	return e
}

func (e *errList) dedupe() {
	var cleaned []error
	set := make(map[string]bool)
	for _, err := range *e {
		if msg := err.Error(); !set[msg] {
			set[msg] = true
			cleaned = append(cleaned, err)
		}
	}
	*e = cleaned
}

func (e errList) Error() string {
	switch len(e) {
	case 0:
		return ""
	case 1:
		return e[0].Error()
	default:
		var buf bytes.Buffer

		for i, err := range e {
			if i > 0 {
				buf.WriteRune('\n')
			}
			buf.WriteString(err.Error())
		}
		return buf.String()
	}
}

// parserError wraps an error with a prefix indicating the rule in which
// the error occurred. The original error is stored in the Inner field.
type parserError struct {
	Inner  error
	pos    position
	prefix string
}

// Error returns the error message.
func (p *parserError) Error() string {
	return p.prefix + ": " + p.Inner.Error()
}

// newParser creates a parser with the specified input source and options.
func newParser(filename string, b []byte, opts ...Option) *parser {
	p := &parser{
		filename: filename,
		errs:     new(errList),
		data:     b,
		pt:       savepoint{position: position{line: 1}, state: make(statedict)},
		recover:  true,
		trace:    TTrace{},
	}
	p.setOptions(opts)
	return p
}

// setOptions applies the options to the parser.
func (p *parser) setOptions(opts []Option) {
	for _, opt := range opts {
		opt(p)
	}
}

// TDetail group information
type TDetail struct {
	Idx1 int    `json:"idx1"`
	Idx2 int    `json:"idx2"`
	Name string `json:"name"`
}

// TEntry element of trace
type TEntry struct {
	Detail  TDetail  `json:"detail"`
	Calls   []TEntry `json:"calls"`
	IsMatch bool     `json:"ismatch"`
}

// TTrace root of entries
type TTrace struct {
	Entries []TEntry `json:"entries"`
	Errors  string   `json:"errors"`
}

type resultTuple struct {
	v   interface{}
	b   bool
	end savepoint
}

type parser struct {
	filename string
	pt       savepoint
	cur      current

	data []byte
	errs *errList

	recover bool
	debug   bool
	jdebug  bool
	depth   int

	memoize bool
	// memoization table for the packrat algorithm:
	// map[offset in source] map[expression or rule] {value, match}
	memo map[int]map[interface{}]resultTuple

	// rules table, maps the rule identifier to the rule node
	rules map[string]*rule
	// variables stack, map of label to value
	vstack []map[string]interface{}
	// rule stack, allows identification of the current rule in errors
	rstack []*rule

	// stats
	exprCnt int
	// trace calls of rules, if jdebug is true, else nil
	trace  TTrace
	tstack []*TEntry
}

// push a variable set on the vstack.
func (p *parser) pushV() {
	if cap(p.vstack) == len(p.vstack) {
		// create new empty slot in the stack
		p.vstack = append(p.vstack, nil)
	} else {
		// slice to 1 more
		p.vstack = p.vstack[:len(p.vstack)+1]
	}

	// get the last args set
	m := p.vstack[len(p.vstack)-1]
	if m != nil && len(m) == 0 {
		// empty map, all good
		return
	}

	m = make(map[string]interface{})
	p.vstack[len(p.vstack)-1] = m
}

// pop a variable set from the vstack.
func (p *parser) popV() {
	// if the map is not empty, clear it
	m := p.vstack[len(p.vstack)-1]
	if len(m) > 0 {
		// GC that map
		p.vstack[len(p.vstack)-1] = nil
	}
	p.vstack = p.vstack[:len(p.vstack)-1]
}

// push a trace entry on the tstack.
func (p *parser) pushT(e *TEntry) {
	if cap(p.tstack) == len(p.tstack) {
		// create new empty slot in the stack
		p.tstack = append(p.tstack, nil)
	} else {
		// slice to 1 more
		p.tstack = p.tstack[:len(p.tstack)+1]
	}

	// get the last args set
	v := p.tstack[len(p.tstack)-1]
	if v != nil {
		return
	}
	p.tstack[len(p.tstack)-1] = e
}

// pop a trace entry from the tstack.
func (p *parser) popT() *TEntry {
	if len(p.tstack) == 0 {
		return nil
	}
	e := p.tstack[len(p.tstack)-1]
	p.tstack = p.tstack[:len(p.tstack)-1]
	return e
}

func (p *parser) peekT() *TEntry {
	if len(p.tstack) == 0 {
		return nil
	}
	return p.tstack[len(p.tstack)-1]
}

func (p *parser) print(prefix, s string) string {
	if !p.debug {
		return s
	}

	fmt.Printf("%s %d:%d:%d: %s [%#U]\n",
		prefix, p.pt.line, p.pt.col, p.pt.offset, s, p.pt.rn)
	return s
}

func (p *parser) in(s string) string {
	p.depth++
	return p.print(strings.Repeat(" ", p.depth)+">", s)
}

func (p *parser) out(s string) string {
	p.depth--
	return p.print(strings.Repeat(" ", p.depth)+"<", s)
}

func (p *parser) jin(s string) *TEntry {
	if p.jdebug {
		pr := p.peekT()
		e := TEntry{Calls: []TEntry{}}
		e.Detail = TDetail{Idx1: p.pt.offset, Name: s}
		pr.Calls = append(pr.Calls, e)
		w := &(pr.Calls[len(pr.Calls)-1])
		p.pushT(w)
		return w
	}
	return nil
}

func (p *parser) jout(e *TEntry) *TEntry {
	if p.jdebug {
		e := p.popT()
		e.Detail.Idx2 = p.pt.offset
		return e
	}
	return nil
}

func (p *parser) addErr(err error) {
	p.addErrAt(err, p.pt.position)
}

func (p *parser) addErrAt(err error, pos position) {
	var buf bytes.Buffer
	if p.filename != "" {
		buf.WriteString(p.filename)
	}
	if buf.Len() > 0 {
		buf.WriteString(":")
	}
	buf.WriteString(fmt.Sprintf("%d:%d (%d)", pos.line, pos.col, pos.offset))
	if len(p.rstack) > 0 {
		if buf.Len() > 0 {
			buf.WriteString(": ")
		}
		rule := p.rstack[len(p.rstack)-1]
		if rule.displayName != "" {
			buf.WriteString("rule " + rule.displayName)
		} else {
			buf.WriteString("rule " + rule.name)
		}
	}
	pe := &parserError{Inner: err, pos: pos, prefix: buf.String()}
	p.errs.add(pe)
}

// read advances the parser to the next rune.
func (p *parser) read() {
	p.pt.offset += p.pt.w
	rn, n := utf8.DecodeRune(p.data[p.pt.offset:])
	p.pt.rn = rn
	p.pt.w = n
	p.pt.col++
	if rn == '\n' {
		p.pt.line++
		p.pt.col = 0
	}

	if rn == utf8.RuneError {
		if n == 1 {
			p.addErr(errInvalidEncoding)
		}
	}
}

// copy state
func copyState(dst, src statedict) {
	for k, v := range src {
		dst[k] = v
	}
}

// restore parser position to the savepoint pt.
func (p *parser) restore(pt savepoint) {
	if p.debug {
		defer p.out(p.in("restore"))
	}
	if pt.offset == p.pt.offset {
		return
	}
	p.pt = pt
}

// get the slice of bytes from the savepoint start to the current position.
func (p *parser) sliceFrom(start savepoint) []byte {
	return p.data[start.position.offset:p.pt.position.offset]
}

func (p *parser) getMemoized(node interface{}) (resultTuple, bool) {
	if len(p.memo) == 0 {
		return resultTuple{}, false
	}
	m := p.memo[p.pt.offset]
	if len(m) == 0 {
		return resultTuple{}, false
	}
	res, ok := m[node]
	return res, ok
}

func (p *parser) setMemoized(pt savepoint, node interface{}, tuple resultTuple) {
	if p.memo == nil {
		p.memo = make(map[int]map[interface{}]resultTuple)
	}
	m := p.memo[pt.offset]
	if m == nil {
		m = make(map[interface{}]resultTuple)
		p.memo[pt.offset] = m
	}
	m[node] = tuple
}

func (p *parser) buildRulesTable(g *grammar) {
	p.rules = make(map[string]*rule, len(g.rules))
	for _, r := range g.rules {
		p.rules[r.name] = r
	}
}

func (p *parser) parse(g *grammar) (val interface{}, err error) {
	if len(g.rules) == 0 {
		p.addErr(errNoRule)
		return nil, p.errs.err()
	}
	if p.jdebug {
		e := TEntry{}
		p.pushT(&e)
	}
	// TODO : not super critical but this could be generated
	p.buildRulesTable(g)

	if p.recover {
		// panic can be used in action code to stop parsing immediately
		// and return the panic as an error.
		defer func() {
			if e := recover(); e != nil {
				if p.debug {
					defer p.out(p.in("panic handler"))
				}
				val = nil
				switch e := e.(type) {
				case error:
					p.addErr(e)
				default:
					p.addErr(fmt.Errorf("%v", e))
				}
				err = p.errs.err()
			}
		}()
	}

	// start rule is rule [0]
	p.read() // advance to first rune
	val, ok := p.parseRule(g.rules[0])
	if !ok {
		if len(*p.errs) == 0 {
			// make sure this doesn't go out silently
			p.addErr(errNoMatch)
		}
		return nil, p.errs.err()
	}
	if p.jdebug {
		e := p.popT()
		p.trace.Entries = e.Calls
		jtrace, err := json.Marshal(p.trace)
		if err != nil {
			fmt.Println("jdebug: Cant marshal json\n")
		} else {
			fmt.Println(string(jtrace))
		}

	}
	return val, p.errs.err()
}

func (p *parser) parseRule(rule *rule) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseRule " + rule.name))
	}
	var j *TEntry
	if p.jdebug {
		j = p.jin("parseRule " + rule.name)
		defer p.jout(j)
	}

	if p.memoize {
		res, ok := p.getMemoized(rule)
		if ok {
			p.restore(res.end)
			return res.v, res.b
		}
	}
	start := p.pt
	p.rstack = append(p.rstack, rule)
	p.pushV()
	val, ok := p.parseExpr(rule.expr)
	p.popV()
	p.rstack = p.rstack[:len(p.rstack)-1]
	if ok && p.debug {
		p.print(strings.Repeat(" ", p.depth)+"MATCH", string(p.sliceFrom(start)))
	}
	if ok && p.jdebug {
		j.IsMatch = true
	}

	if p.memoize {
		p.setMemoized(start, rule, resultTuple{val, ok, p.pt})
	}
	return val, ok
}

func (p *parser) parseExpr(expr interface{}) (interface{}, bool) {
	var pt savepoint
	var ok bool

	if p.memoize {
		res, ok := p.getMemoized(expr)
		if ok {
			p.restore(res.end)
			return res.v, res.b
		}
		pt = p.pt
	}

	p.exprCnt++
	var val interface{}
	switch expr := expr.(type) {
	case *actionExpr:
		val, ok = p.parseActionExpr(expr)
	case *andCodeExpr:
		val, ok = p.parseAndCodeExpr(expr)
	case *andExpr:
		val, ok = p.parseAndExpr(expr)
	case *anyMatcher:
		val, ok = p.parseAnyMatcher(expr)
	case *charClassMatcher:
		val, ok = p.parseCharClassMatcher(expr)
	case *choiceExpr:
		val, ok = p.parseChoiceExpr(expr)
	case *labeledExpr:
		val, ok = p.parseLabeledExpr(expr)
	case *litMatcher:
		val, ok = p.parseLitMatcher(expr)
	case *notCodeExpr:
		val, ok = p.parseNotCodeExpr(expr)
	case *notExpr:
		val, ok = p.parseNotExpr(expr)
	case *oneOrMoreExpr:
		val, ok = p.parseOneOrMoreExpr(expr)
	case *ruleRefExpr:
		val, ok = p.parseRuleRefExpr(expr)
	case *seqExpr:
		val, ok = p.parseSeqExpr(expr)
	case *stateCodeExpr:
		val, ok = p.parseStateCodeExpr(expr)
	case *zeroOrMoreExpr:
		val, ok = p.parseZeroOrMoreExpr(expr)
	case *zeroOrOneExpr:
		val, ok = p.parseZeroOrOneExpr(expr)
	default:
		panic(fmt.Sprintf("unknown expression type %T", expr))
	}
	if p.memoize {
		p.setMemoized(pt, expr, resultTuple{val, ok, p.pt})
	}
	return val, ok
}

func (p *parser) parseActionExpr(act *actionExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseActionExpr"))
	}

	start := p.pt
	val, ok := p.parseExpr(act.expr)
	if ok {
		p.cur.pos = start.position
		p.cur.text = p.sliceFrom(start)
		actVal, err := act.run(p)
		if err != nil {
			p.addErrAt(err, start.position)
		}
		val = actVal
	}
	if ok && p.debug {
		p.print(strings.Repeat(" ", p.depth)+"MATCH", string(p.sliceFrom(start)))
	}
	return val, ok
}

func (p *parser) parseAndCodeExpr(and *andCodeExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseAndCodeExpr"))
	}

	ok, err := and.run(p)
	if err != nil {
		p.addErr(err)
	}
	return nil, ok
}

func (p *parser) parseAndExpr(and *andExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseAndExpr"))
	}

	pt := p.pt
	pt.state = make(statedict)
	copyState(pt.state, p.pt.state)
	p.pushV()
	_, ok := p.parseExpr(and.expr)
	p.popV()
	p.restore(pt)
	copyState(p.pt.state, pt.state)
	copyState(state, pt.state)
	return nil, ok
}

func (p *parser) parseAnyMatcher(any *anyMatcher) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseAnyMatcher"))
	}

	if p.pt.rn != utf8.RuneError {
		start := p.pt
		p.read()
		return p.sliceFrom(start), true
	}
	return nil, false
}

func (p *parser) parseCharClassMatcher(chr *charClassMatcher) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseCharClassMatcher"))
	}

	cur := p.pt.rn
	// can't match EOF
	if cur == utf8.RuneError {
		return nil, false
	}
	start := p.pt
	if chr.ignoreCase {
		cur = unicode.ToLower(cur)
	}

	// try to match in the list of available chars
	for _, rn := range chr.chars {
		if rn == cur {
			if chr.inverted {
				return nil, false
			}
			p.read()
			return p.sliceFrom(start), true
		}
	}

	// try to match in the list of ranges
	for i := 0; i < len(chr.ranges); i += 2 {
		if cur >= chr.ranges[i] && cur <= chr.ranges[i+1] {
			if chr.inverted {
				return nil, false
			}
			p.read()
			return p.sliceFrom(start), true
		}
	}

	// try to match in the list of Unicode classes
	for _, cl := range chr.classes {
		if unicode.Is(cl, cur) {
			if chr.inverted {
				return nil, false
			}
			p.read()
			return p.sliceFrom(start), true
		}
	}

	if chr.inverted {
		p.read()
		return p.sliceFrom(start), true
	}
	return nil, false
}

func (p *parser) parseChoiceExpr(ch *choiceExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseChoiceExpr"))
	}

	for _, alt := range ch.alternatives {
		p.pushV()
		val, ok := p.parseExpr(alt)
		p.popV()
		if ok {
			return val, ok
		}
	}
	return nil, false
}

func (p *parser) parseLabeledExpr(lab *labeledExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseLabeledExpr"))
	}

	p.pushV()
	val, ok := p.parseExpr(lab.expr)
	p.popV()
	if ok && lab.label != "" {
		m := p.vstack[len(p.vstack)-1]
		m[lab.label] = val
	}
	return val, ok
}

func (p *parser) parseLitMatcher(lit *litMatcher) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseLitMatcher"))
	}

	start := p.pt
	for _, want := range lit.val {
		cur := p.pt.rn
		if lit.ignoreCase {
			cur = unicode.ToLower(cur)
		}
		if cur != want {
			p.restore(start)
			return nil, false
		}
		p.read()
	}
	return p.sliceFrom(start), true
}

func (p *parser) parseNotCodeExpr(not *notCodeExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseNotCodeExpr"))
	}

	ok, err := not.run(p)
	if err != nil {
		p.addErr(err)
	}
	return nil, !ok
}

func (p *parser) parseStateCodeExpr(state *stateCodeExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseStateCodeExpr"))
	}

	_, err := state.run(p)
	if err != nil {
		p.addErr(err)
	}
	return nil, true
}

func (p *parser) parseNotExpr(not *notExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseNotExpr"))
	}

	pt := p.pt
	pt.state = make(statedict)
	copyState(pt.state, p.pt.state)
	p.pushV()
	_, ok := p.parseExpr(not.expr)
	p.popV()
	p.restore(pt)
	copyState(p.pt.state, pt.state)
	copyState(state, pt.state)
	return nil, !ok
}

func (p *parser) parseOneOrMoreExpr(expr *oneOrMoreExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseOneOrMoreExpr"))
	}
	if p.jdebug {
		defer p.jout(p.jin("parseOneOrMoreExpr"))
	}

	var vals []interface{}

	for {
		p.pushV()
		val, ok := p.parseExpr(expr.expr)
		p.popV()
		if !ok {
			if len(vals) == 0 {
				// did not match once, no match
				return nil, false
			}
			return vals, true
		}
		vals = append(vals, val)
	}
}

func (p *parser) parseRuleRefExpr(ref *ruleRefExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseRuleRefExpr " + ref.name))
	}

	if ref.name == "" {
		panic(fmt.Sprintf("%s: invalid rule: missing name", ref.pos))
	}

	rule := p.rules[ref.name]
	if rule == nil {
		p.addErr(fmt.Errorf("undefined rule: %s", ref.name))
		return nil, false
	}
	return p.parseRule(rule)
}

func (p *parser) parseSeqExpr(seq *seqExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseSeqExpr"))
	}

	var vals []interface{}

	pt := p.pt
	pt.state = make(statedict)
	copyState(pt.state, p.pt.state)
	for _, expr := range seq.exprs {
		val, ok := p.parseExpr(expr)
		if !ok {
			p.restore(pt)
			copyState(p.pt.state, pt.state)
			copyState(state, pt.state)
			return nil, false
		}
		vals = append(vals, val)
	}
	return vals, true
}

func (p *parser) parseZeroOrMoreExpr(expr *zeroOrMoreExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseZeroOrMoreExpr"))
	}
	if p.jdebug {
		defer p.jout(p.jin("parseZeroOrMoreExpr"))
	}

	var vals []interface{}

	for {
		p.pushV()
		val, ok := p.parseExpr(expr.expr)
		p.popV()
		if !ok {
			return vals, true
		}
		vals = append(vals, val)
	}
}

func (p *parser) parseZeroOrOneExpr(expr *zeroOrOneExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseZeroOrOneExpr"))
	}
	if p.jdebug {
		defer p.jout(p.jin("parseZeroOrOneExpr"))
	}

	p.pushV()
	val, _ := p.parseExpr(expr.expr)
	p.popV()
	// whether it matched or not, consider it a match
	return val, true
}

func rangeTable(class string) *unicode.RangeTable {
	if rt, ok := unicode.Categories[class]; ok {
		return rt
	}
	if rt, ok := unicode.Properties[class]; ok {
		return rt
	}
	if rt, ok := unicode.Scripts[class]; ok {
		return rt
	}

	// cannot happen
	panic(fmt.Sprintf("invalid Unicode class: %s", class))
}

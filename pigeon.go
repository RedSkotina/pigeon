package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/PuerkitoBio/pigeon/ast"
)

var g = &grammar{
	rules: []*rule{
		{
			name: "Grammar",
			pos:  position{line: 5, col: 1, offset: 22},
			expr: &actionExpr{
				pos: position{line: 5, col: 11, offset: 34},
				run: (*parser).callonGrammar1,
				expr: &seqExpr{
					pos: position{line: 5, col: 11, offset: 34},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 5, col: 11, offset: 34},
							name: "__",
						},
						&labeledExpr{
							pos:   position{line: 5, col: 14, offset: 37},
							label: "initializer",
							expr: &zeroOrOneExpr{
								pos: position{line: 5, col: 26, offset: 49},
								expr: &seqExpr{
									pos: position{line: 5, col: 28, offset: 51},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 5, col: 28, offset: 51},
											name: "Initializer",
										},
										&ruleRefExpr{
											pos:  position{line: 5, col: 40, offset: 63},
											name: "__",
										},
									},
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 5, col: 46, offset: 69},
							label: "rules",
							expr: &oneOrMoreExpr{
								pos: position{line: 5, col: 52, offset: 75},
								expr: &seqExpr{
									pos: position{line: 5, col: 54, offset: 77},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 5, col: 54, offset: 77},
											name: "Rule",
										},
										&ruleRefExpr{
											pos:  position{line: 5, col: 59, offset: 82},
											name: "__",
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 5, col: 65, offset: 88},
							name: "EOF",
						},
					},
				},
			},
		},
		{
			name: "Initializer",
			pos:  position{line: 24, col: 1, offset: 548},
			expr: &actionExpr{
				pos: position{line: 24, col: 15, offset: 564},
				run: (*parser).callonInitializer1,
				expr: &seqExpr{
					pos: position{line: 24, col: 15, offset: 564},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 24, col: 15, offset: 564},
							label: "code",
							expr: &ruleRefExpr{
								pos:  position{line: 24, col: 20, offset: 569},
								name: "CodeBlock",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 24, col: 30, offset: 579},
							name: "EOS",
						},
					},
				},
			},
		},
		{
			name: "Rule",
			pos:  position{line: 28, col: 1, offset: 613},
			expr: &actionExpr{
				pos: position{line: 28, col: 8, offset: 622},
				run: (*parser).callonRule1,
				expr: &seqExpr{
					pos: position{line: 28, col: 8, offset: 622},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 28, col: 8, offset: 622},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 28, col: 13, offset: 627},
								name: "IdentifierName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 28, col: 28, offset: 642},
							name: "__",
						},
						&labeledExpr{
							pos:   position{line: 28, col: 31, offset: 645},
							label: "display",
							expr: &zeroOrOneExpr{
								pos: position{line: 28, col: 39, offset: 653},
								expr: &seqExpr{
									pos: position{line: 28, col: 41, offset: 655},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 28, col: 41, offset: 655},
											name: "StringLiteral",
										},
										&ruleRefExpr{
											pos:  position{line: 28, col: 55, offset: 669},
											name: "__",
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 28, col: 61, offset: 675},
							name: "RuleDefOp",
						},
						&ruleRefExpr{
							pos:  position{line: 28, col: 71, offset: 685},
							name: "__",
						},
						&labeledExpr{
							pos:   position{line: 28, col: 74, offset: 688},
							label: "expr",
							expr: &ruleRefExpr{
								pos:  position{line: 28, col: 79, offset: 693},
								name: "Expression",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 28, col: 90, offset: 704},
							name: "EOS",
						},
					},
				},
			},
		},
		{
			name: "Expression",
			pos:  position{line: 41, col: 1, offset: 1001},
			expr: &ruleRefExpr{
				pos:  position{line: 41, col: 14, offset: 1016},
				name: "ChoiceExpr",
			},
		},
		{
			name: "ChoiceExpr",
			pos:  position{line: 43, col: 1, offset: 1030},
			expr: &actionExpr{
				pos: position{line: 43, col: 14, offset: 1045},
				run: (*parser).callonChoiceExpr1,
				expr: &seqExpr{
					pos: position{line: 43, col: 14, offset: 1045},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 43, col: 14, offset: 1045},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 43, col: 20, offset: 1051},
								name: "ActionExpr",
							},
						},
						&labeledExpr{
							pos:   position{line: 43, col: 31, offset: 1062},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 43, col: 36, offset: 1067},
								expr: &seqExpr{
									pos: position{line: 43, col: 38, offset: 1069},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 43, col: 38, offset: 1069},
											name: "__",
										},
										&litMatcher{
											pos:        position{line: 43, col: 41, offset: 1072},
											val:        "/",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 43, col: 45, offset: 1076},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 43, col: 48, offset: 1079},
											name: "ActionExpr",
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
			name: "ActionExpr",
			pos:  position{line: 58, col: 1, offset: 1499},
			expr: &actionExpr{
				pos: position{line: 58, col: 14, offset: 1514},
				run: (*parser).callonActionExpr1,
				expr: &seqExpr{
					pos: position{line: 58, col: 14, offset: 1514},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 58, col: 14, offset: 1514},
							label: "expr",
							expr: &ruleRefExpr{
								pos:  position{line: 58, col: 19, offset: 1519},
								name: "SeqExpr",
							},
						},
						&labeledExpr{
							pos:   position{line: 58, col: 27, offset: 1527},
							label: "code",
							expr: &zeroOrOneExpr{
								pos: position{line: 58, col: 32, offset: 1532},
								expr: &seqExpr{
									pos: position{line: 58, col: 34, offset: 1534},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 58, col: 34, offset: 1534},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 58, col: 37, offset: 1537},
											name: "CodeBlock",
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
			name: "SeqExpr",
			pos:  position{line: 72, col: 1, offset: 1817},
			expr: &actionExpr{
				pos: position{line: 72, col: 11, offset: 1829},
				run: (*parser).callonSeqExpr1,
				expr: &seqExpr{
					pos: position{line: 72, col: 11, offset: 1829},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 72, col: 11, offset: 1829},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 72, col: 17, offset: 1835},
								name: "LabeledExpr",
							},
						},
						&labeledExpr{
							pos:   position{line: 72, col: 29, offset: 1847},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 72, col: 34, offset: 1852},
								expr: &seqExpr{
									pos: position{line: 72, col: 36, offset: 1854},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 72, col: 36, offset: 1854},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 72, col: 39, offset: 1857},
											name: "LabeledExpr",
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
			name: "LabeledExpr",
			pos:  position{line: 85, col: 1, offset: 2221},
			expr: &choiceExpr{
				pos: position{line: 85, col: 15, offset: 2237},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 85, col: 15, offset: 2237},
						run: (*parser).callonLabeledExpr2,
						expr: &seqExpr{
							pos: position{line: 85, col: 15, offset: 2237},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 85, col: 15, offset: 2237},
									label: "label",
									expr: &ruleRefExpr{
										pos:  position{line: 85, col: 21, offset: 2243},
										name: "Identifier",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 85, col: 32, offset: 2254},
									name: "__",
								},
								&litMatcher{
									pos:        position{line: 85, col: 35, offset: 2257},
									val:        ":",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 85, col: 39, offset: 2261},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 85, col: 42, offset: 2264},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 85, col: 47, offset: 2269},
										name: "PrefixedExpr",
									},
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 91, col: 5, offset: 2448},
						name: "PrefixedExpr",
					},
				},
			},
		},
		{
			name: "PrefixedExpr",
			pos:  position{line: 93, col: 1, offset: 2464},
			expr: &choiceExpr{
				pos: position{line: 93, col: 16, offset: 2481},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 93, col: 16, offset: 2481},
						run: (*parser).callonPrefixedExpr2,
						expr: &seqExpr{
							pos: position{line: 93, col: 16, offset: 2481},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 93, col: 16, offset: 2481},
									label: "op",
									expr: &ruleRefExpr{
										pos:  position{line: 93, col: 19, offset: 2484},
										name: "PrefixedOp",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 93, col: 30, offset: 2495},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 93, col: 33, offset: 2498},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 93, col: 38, offset: 2503},
										name: "SuffixedExpr",
									},
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 105, col: 5, offset: 2798},
						name: "SuffixedExpr",
					},
				},
			},
		},
		{
			name: "PrefixedOp",
			pos:  position{line: 107, col: 1, offset: 2814},
			expr: &actionExpr{
				pos: position{line: 107, col: 14, offset: 2829},
				run: (*parser).callonPrefixedOp1,
				expr: &choiceExpr{
					pos: position{line: 107, col: 16, offset: 2831},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 107, col: 16, offset: 2831},
							val:        "&",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 107, col: 22, offset: 2837},
							val:        "!",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "SuffixedExpr",
			pos:  position{line: 111, col: 1, offset: 2884},
			expr: &choiceExpr{
				pos: position{line: 111, col: 16, offset: 2901},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 111, col: 16, offset: 2901},
						run: (*parser).callonSuffixedExpr2,
						expr: &seqExpr{
							pos: position{line: 111, col: 16, offset: 2901},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 111, col: 16, offset: 2901},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 111, col: 21, offset: 2906},
										name: "PrimaryExpr",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 111, col: 33, offset: 2918},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 111, col: 36, offset: 2921},
									label: "op",
									expr: &ruleRefExpr{
										pos:  position{line: 111, col: 39, offset: 2924},
										name: "SuffixedOp",
									},
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 130, col: 5, offset: 3473},
						name: "PrimaryExpr",
					},
				},
			},
		},
		{
			name: "SuffixedOp",
			pos:  position{line: 132, col: 1, offset: 3489},
			expr: &actionExpr{
				pos: position{line: 132, col: 14, offset: 3504},
				run: (*parser).callonSuffixedOp1,
				expr: &choiceExpr{
					pos: position{line: 132, col: 16, offset: 3506},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 132, col: 16, offset: 3506},
							val:        "?",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 132, col: 22, offset: 3512},
							val:        "*",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 132, col: 28, offset: 3518},
							val:        "+",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "PrimaryExpr",
			pos:  position{line: 136, col: 1, offset: 3564},
			expr: &choiceExpr{
				pos: position{line: 136, col: 15, offset: 3580},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 136, col: 15, offset: 3580},
						name: "LitMatcher",
					},
					&ruleRefExpr{
						pos:  position{line: 136, col: 28, offset: 3593},
						name: "CharClassMatcher",
					},
					&ruleRefExpr{
						pos:  position{line: 136, col: 47, offset: 3612},
						name: "AnyMatcher",
					},
					&ruleRefExpr{
						pos:  position{line: 136, col: 60, offset: 3625},
						name: "RuleRefExpr",
					},
					&ruleRefExpr{
						pos:  position{line: 136, col: 74, offset: 3639},
						name: "SemanticPredExpr",
					},
					&actionExpr{
						pos: position{line: 136, col: 93, offset: 3658},
						run: (*parser).callonPrimaryExpr7,
						expr: &seqExpr{
							pos: position{line: 136, col: 93, offset: 3658},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 136, col: 93, offset: 3658},
									val:        "(",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 136, col: 97, offset: 3662},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 136, col: 100, offset: 3665},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 136, col: 105, offset: 3670},
										name: "Expression",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 136, col: 116, offset: 3681},
									name: "__",
								},
								&litMatcher{
									pos:        position{line: 136, col: 119, offset: 3684},
									val:        ")",
									ignoreCase: false,
								},
							},
						},
					},
				},
			},
		},
		{
			name: "RuleRefExpr",
			pos:  position{line: 139, col: 1, offset: 3716},
			expr: &actionExpr{
				pos: position{line: 139, col: 15, offset: 3732},
				run: (*parser).callonRuleRefExpr1,
				expr: &seqExpr{
					pos: position{line: 139, col: 15, offset: 3732},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 139, col: 15, offset: 3732},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 139, col: 20, offset: 3737},
								name: "IdentifierName",
							},
						},
						&notExpr{
							pos: position{line: 139, col: 35, offset: 3752},
							expr: &seqExpr{
								pos: position{line: 139, col: 38, offset: 3755},
								exprs: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 139, col: 38, offset: 3755},
										name: "__",
									},
									&zeroOrOneExpr{
										pos: position{line: 139, col: 41, offset: 3758},
										expr: &seqExpr{
											pos: position{line: 139, col: 43, offset: 3760},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 139, col: 43, offset: 3760},
													name: "StringLiteral",
												},
												&ruleRefExpr{
													pos:  position{line: 139, col: 57, offset: 3774},
													name: "__",
												},
											},
										},
									},
									&ruleRefExpr{
										pos:  position{line: 139, col: 63, offset: 3780},
										name: "RuleDefOp",
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "SemanticPredExpr",
			pos:  position{line: 144, col: 1, offset: 3901},
			expr: &actionExpr{
				pos: position{line: 144, col: 20, offset: 3922},
				run: (*parser).callonSemanticPredExpr1,
				expr: &seqExpr{
					pos: position{line: 144, col: 20, offset: 3922},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 144, col: 20, offset: 3922},
							label: "op",
							expr: &ruleRefExpr{
								pos:  position{line: 144, col: 23, offset: 3925},
								name: "SemanticPredOp",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 144, col: 38, offset: 3940},
							name: "__",
						},
						&labeledExpr{
							pos:   position{line: 144, col: 41, offset: 3943},
							label: "code",
							expr: &ruleRefExpr{
								pos:  position{line: 144, col: 46, offset: 3948},
								name: "CodeBlock",
							},
						},
					},
				},
			},
		},
		{
			name: "SemanticPredOp",
			pos:  position{line: 160, col: 1, offset: 4394},
			expr: &actionExpr{
				pos: position{line: 160, col: 18, offset: 4413},
				run: (*parser).callonSemanticPredOp1,
				expr: &choiceExpr{
					pos: position{line: 160, col: 20, offset: 4415},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 160, col: 20, offset: 4415},
							val:        "#",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 160, col: 26, offset: 4421},
							val:        "&",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 160, col: 32, offset: 4427},
							val:        "!",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "RuleDefOp",
			pos:  position{line: 164, col: 1, offset: 4473},
			expr: &choiceExpr{
				pos: position{line: 164, col: 13, offset: 4487},
				alternatives: []interface{}{
					&litMatcher{
						pos:        position{line: 164, col: 13, offset: 4487},
						val:        "=",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 164, col: 19, offset: 4493},
						val:        "<-",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 164, col: 26, offset: 4500},
						val:        "←",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 164, col: 37, offset: 4511},
						val:        "⟵",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "SourceChar",
			pos:  position{line: 166, col: 1, offset: 4523},
			expr: &anyMatcher{
				line: 166, col: 14, offset: 4538,
			},
		},
		{
			name: "Comment",
			pos:  position{line: 167, col: 1, offset: 4541},
			expr: &choiceExpr{
				pos: position{line: 167, col: 11, offset: 4553},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 167, col: 11, offset: 4553},
						name: "MultiLineComment",
					},
					&ruleRefExpr{
						pos:  position{line: 167, col: 30, offset: 4572},
						name: "SingleLineComment",
					},
				},
			},
		},
		{
			name: "MultiLineComment",
			pos:  position{line: 168, col: 1, offset: 4591},
			expr: &seqExpr{
				pos: position{line: 168, col: 20, offset: 4612},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 168, col: 20, offset: 4612},
						val:        "/*",
						ignoreCase: false,
					},
					&zeroOrMoreExpr{
						pos: position{line: 168, col: 25, offset: 4617},
						expr: &seqExpr{
							pos: position{line: 168, col: 27, offset: 4619},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 168, col: 27, offset: 4619},
									expr: &litMatcher{
										pos:        position{line: 168, col: 28, offset: 4620},
										val:        "*/",
										ignoreCase: false,
									},
								},
								&ruleRefExpr{
									pos:  position{line: 168, col: 33, offset: 4625},
									name: "SourceChar",
								},
							},
						},
					},
					&litMatcher{
						pos:        position{line: 168, col: 47, offset: 4639},
						val:        "*/",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "MultiLineCommentNoLineTerminator",
			pos:  position{line: 169, col: 1, offset: 4645},
			expr: &seqExpr{
				pos: position{line: 169, col: 36, offset: 4682},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 169, col: 36, offset: 4682},
						val:        "/*",
						ignoreCase: false,
					},
					&zeroOrMoreExpr{
						pos: position{line: 169, col: 41, offset: 4687},
						expr: &seqExpr{
							pos: position{line: 169, col: 43, offset: 4689},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 169, col: 43, offset: 4689},
									expr: &choiceExpr{
										pos: position{line: 169, col: 46, offset: 4692},
										alternatives: []interface{}{
											&litMatcher{
												pos:        position{line: 169, col: 46, offset: 4692},
												val:        "*/",
												ignoreCase: false,
											},
											&ruleRefExpr{
												pos:  position{line: 169, col: 53, offset: 4699},
												name: "EOL",
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 169, col: 59, offset: 4705},
									name: "SourceChar",
								},
							},
						},
					},
					&litMatcher{
						pos:        position{line: 169, col: 73, offset: 4719},
						val:        "*/",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "SingleLineComment",
			pos:  position{line: 170, col: 1, offset: 4725},
			expr: &seqExpr{
				pos: position{line: 170, col: 21, offset: 4747},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 170, col: 21, offset: 4747},
						val:        "//",
						ignoreCase: false,
					},
					&zeroOrMoreExpr{
						pos: position{line: 170, col: 26, offset: 4752},
						expr: &seqExpr{
							pos: position{line: 170, col: 28, offset: 4754},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 170, col: 28, offset: 4754},
									expr: &ruleRefExpr{
										pos:  position{line: 170, col: 29, offset: 4755},
										name: "EOL",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 170, col: 33, offset: 4759},
									name: "SourceChar",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Identifier",
			pos:  position{line: 172, col: 1, offset: 4776},
			expr: &actionExpr{
				pos: position{line: 172, col: 14, offset: 4791},
				run: (*parser).callonIdentifier1,
				expr: &labeledExpr{
					pos:   position{line: 172, col: 14, offset: 4791},
					label: "ident",
					expr: &ruleRefExpr{
						pos:  position{line: 172, col: 20, offset: 4797},
						name: "IdentifierName",
					},
				},
			},
		},
		{
			name: "IdentifierName",
			pos:  position{line: 180, col: 1, offset: 5024},
			expr: &actionExpr{
				pos: position{line: 180, col: 18, offset: 5043},
				run: (*parser).callonIdentifierName1,
				expr: &seqExpr{
					pos: position{line: 180, col: 18, offset: 5043},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 180, col: 18, offset: 5043},
							name: "IdentifierStart",
						},
						&zeroOrMoreExpr{
							pos: position{line: 180, col: 34, offset: 5059},
							expr: &ruleRefExpr{
								pos:  position{line: 180, col: 34, offset: 5059},
								name: "IdentifierPart",
							},
						},
					},
				},
			},
		},
		{
			name: "IdentifierStart",
			pos:  position{line: 183, col: 1, offset: 5144},
			expr: &charClassMatcher{
				pos:        position{line: 183, col: 19, offset: 5164},
				val:        "[\\pL_]",
				chars:      []rune{'_'},
				classes:    []*unicode.RangeTable{rangeTable("L")},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "IdentifierPart",
			pos:  position{line: 184, col: 1, offset: 5172},
			expr: &choiceExpr{
				pos: position{line: 184, col: 18, offset: 5191},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 184, col: 18, offset: 5191},
						name: "IdentifierStart",
					},
					&charClassMatcher{
						pos:        position{line: 184, col: 36, offset: 5209},
						val:        "[\\p{Nd}]",
						classes:    []*unicode.RangeTable{rangeTable("Nd")},
						ignoreCase: false,
						inverted:   false,
					},
				},
			},
		},
		{
			name: "LitMatcher",
			pos:  position{line: 186, col: 1, offset: 5221},
			expr: &actionExpr{
				pos: position{line: 186, col: 14, offset: 5236},
				run: (*parser).callonLitMatcher1,
				expr: &seqExpr{
					pos: position{line: 186, col: 14, offset: 5236},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 186, col: 14, offset: 5236},
							label: "lit",
							expr: &ruleRefExpr{
								pos:  position{line: 186, col: 18, offset: 5240},
								name: "StringLiteral",
							},
						},
						&labeledExpr{
							pos:   position{line: 186, col: 32, offset: 5254},
							label: "ignore",
							expr: &zeroOrOneExpr{
								pos: position{line: 186, col: 39, offset: 5261},
								expr: &litMatcher{
									pos:        position{line: 186, col: 39, offset: 5261},
									val:        "i",
									ignoreCase: false,
								},
							},
						},
					},
				},
			},
		},
		{
			name: "StringLiteral",
			pos:  position{line: 199, col: 1, offset: 5673},
			expr: &choiceExpr{
				pos: position{line: 199, col: 17, offset: 5691},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 199, col: 17, offset: 5691},
						run: (*parser).callonStringLiteral2,
						expr: &choiceExpr{
							pos: position{line: 199, col: 19, offset: 5693},
							alternatives: []interface{}{
								&seqExpr{
									pos: position{line: 199, col: 19, offset: 5693},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 199, col: 19, offset: 5693},
											val:        "\"",
											ignoreCase: false,
										},
										&zeroOrMoreExpr{
											pos: position{line: 199, col: 23, offset: 5697},
											expr: &ruleRefExpr{
												pos:  position{line: 199, col: 23, offset: 5697},
												name: "DoubleStringChar",
											},
										},
										&litMatcher{
											pos:        position{line: 199, col: 41, offset: 5715},
											val:        "\"",
											ignoreCase: false,
										},
									},
								},
								&seqExpr{
									pos: position{line: 199, col: 47, offset: 5721},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 199, col: 47, offset: 5721},
											val:        "'",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 199, col: 51, offset: 5725},
											name: "SingleStringChar",
										},
										&litMatcher{
											pos:        position{line: 199, col: 68, offset: 5742},
											val:        "'",
											ignoreCase: false,
										},
									},
								},
								&seqExpr{
									pos: position{line: 199, col: 74, offset: 5748},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 199, col: 74, offset: 5748},
											val:        "`",
											ignoreCase: false,
										},
										&zeroOrMoreExpr{
											pos: position{line: 199, col: 78, offset: 5752},
											expr: &ruleRefExpr{
												pos:  position{line: 199, col: 78, offset: 5752},
												name: "RawStringChar",
											},
										},
										&litMatcher{
											pos:        position{line: 199, col: 93, offset: 5767},
											val:        "`",
											ignoreCase: false,
										},
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 201, col: 5, offset: 5842},
						run: (*parser).callonStringLiteral18,
						expr: &choiceExpr{
							pos: position{line: 201, col: 7, offset: 5844},
							alternatives: []interface{}{
								&seqExpr{
									pos: position{line: 201, col: 9, offset: 5846},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 201, col: 9, offset: 5846},
											val:        "\"",
											ignoreCase: false,
										},
										&zeroOrMoreExpr{
											pos: position{line: 201, col: 13, offset: 5850},
											expr: &ruleRefExpr{
												pos:  position{line: 201, col: 13, offset: 5850},
												name: "DoubleStringChar",
											},
										},
										&choiceExpr{
											pos: position{line: 201, col: 33, offset: 5870},
											alternatives: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 201, col: 33, offset: 5870},
													name: "EOL",
												},
												&ruleRefExpr{
													pos:  position{line: 201, col: 39, offset: 5876},
													name: "EOF",
												},
											},
										},
									},
								},
								&seqExpr{
									pos: position{line: 201, col: 51, offset: 5888},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 201, col: 51, offset: 5888},
											val:        "'",
											ignoreCase: false,
										},
										&zeroOrOneExpr{
											pos: position{line: 201, col: 55, offset: 5892},
											expr: &ruleRefExpr{
												pos:  position{line: 201, col: 55, offset: 5892},
												name: "SingleStringChar",
											},
										},
										&choiceExpr{
											pos: position{line: 201, col: 75, offset: 5912},
											alternatives: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 201, col: 75, offset: 5912},
													name: "EOL",
												},
												&ruleRefExpr{
													pos:  position{line: 201, col: 81, offset: 5918},
													name: "EOF",
												},
											},
										},
									},
								},
								&seqExpr{
									pos: position{line: 201, col: 91, offset: 5928},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 201, col: 91, offset: 5928},
											val:        "`",
											ignoreCase: false,
										},
										&zeroOrMoreExpr{
											pos: position{line: 201, col: 95, offset: 5932},
											expr: &ruleRefExpr{
												pos:  position{line: 201, col: 95, offset: 5932},
												name: "RawStringChar",
											},
										},
										&ruleRefExpr{
											pos:  position{line: 201, col: 110, offset: 5947},
											name: "EOF",
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
			name: "DoubleStringChar",
			pos:  position{line: 205, col: 1, offset: 6053},
			expr: &choiceExpr{
				pos: position{line: 205, col: 20, offset: 6074},
				alternatives: []interface{}{
					&seqExpr{
						pos: position{line: 205, col: 20, offset: 6074},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 205, col: 20, offset: 6074},
								expr: &choiceExpr{
									pos: position{line: 205, col: 23, offset: 6077},
									alternatives: []interface{}{
										&litMatcher{
											pos:        position{line: 205, col: 23, offset: 6077},
											val:        "\"",
											ignoreCase: false,
										},
										&litMatcher{
											pos:        position{line: 205, col: 29, offset: 6083},
											val:        "\\",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 205, col: 36, offset: 6090},
											name: "EOL",
										},
									},
								},
							},
							&ruleRefExpr{
								pos:  position{line: 205, col: 42, offset: 6096},
								name: "SourceChar",
							},
						},
					},
					&seqExpr{
						pos: position{line: 205, col: 55, offset: 6109},
						exprs: []interface{}{
							&litMatcher{
								pos:        position{line: 205, col: 55, offset: 6109},
								val:        "\\",
								ignoreCase: false,
							},
							&ruleRefExpr{
								pos:  position{line: 205, col: 60, offset: 6114},
								name: "DoubleStringEscape",
							},
						},
					},
				},
			},
		},
		{
			name: "SingleStringChar",
			pos:  position{line: 206, col: 1, offset: 6134},
			expr: &choiceExpr{
				pos: position{line: 206, col: 20, offset: 6155},
				alternatives: []interface{}{
					&seqExpr{
						pos: position{line: 206, col: 20, offset: 6155},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 206, col: 20, offset: 6155},
								expr: &choiceExpr{
									pos: position{line: 206, col: 23, offset: 6158},
									alternatives: []interface{}{
										&litMatcher{
											pos:        position{line: 206, col: 23, offset: 6158},
											val:        "'",
											ignoreCase: false,
										},
										&litMatcher{
											pos:        position{line: 206, col: 29, offset: 6164},
											val:        "\\",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 206, col: 36, offset: 6171},
											name: "EOL",
										},
									},
								},
							},
							&ruleRefExpr{
								pos:  position{line: 206, col: 42, offset: 6177},
								name: "SourceChar",
							},
						},
					},
					&seqExpr{
						pos: position{line: 206, col: 55, offset: 6190},
						exprs: []interface{}{
							&litMatcher{
								pos:        position{line: 206, col: 55, offset: 6190},
								val:        "\\",
								ignoreCase: false,
							},
							&ruleRefExpr{
								pos:  position{line: 206, col: 60, offset: 6195},
								name: "SingleStringEscape",
							},
						},
					},
				},
			},
		},
		{
			name: "RawStringChar",
			pos:  position{line: 207, col: 1, offset: 6215},
			expr: &seqExpr{
				pos: position{line: 207, col: 17, offset: 6233},
				exprs: []interface{}{
					&notExpr{
						pos: position{line: 207, col: 17, offset: 6233},
						expr: &litMatcher{
							pos:        position{line: 207, col: 18, offset: 6234},
							val:        "`",
							ignoreCase: false,
						},
					},
					&ruleRefExpr{
						pos:  position{line: 207, col: 22, offset: 6238},
						name: "SourceChar",
					},
				},
			},
		},
		{
			name: "DoubleStringEscape",
			pos:  position{line: 209, col: 1, offset: 6252},
			expr: &choiceExpr{
				pos: position{line: 209, col: 22, offset: 6275},
				alternatives: []interface{}{
					&choiceExpr{
						pos: position{line: 209, col: 24, offset: 6277},
						alternatives: []interface{}{
							&litMatcher{
								pos:        position{line: 209, col: 24, offset: 6277},
								val:        "\"",
								ignoreCase: false,
							},
							&ruleRefExpr{
								pos:  position{line: 209, col: 30, offset: 6283},
								name: "CommonEscapeSequence",
							},
						},
					},
					&actionExpr{
						pos: position{line: 210, col: 7, offset: 6313},
						run: (*parser).callonDoubleStringEscape5,
						expr: &choiceExpr{
							pos: position{line: 210, col: 9, offset: 6315},
							alternatives: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 210, col: 9, offset: 6315},
									name: "SourceChar",
								},
								&ruleRefExpr{
									pos:  position{line: 210, col: 22, offset: 6328},
									name: "EOL",
								},
								&ruleRefExpr{
									pos:  position{line: 210, col: 28, offset: 6334},
									name: "EOF",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "SingleStringEscape",
			pos:  position{line: 213, col: 1, offset: 6402},
			expr: &choiceExpr{
				pos: position{line: 213, col: 22, offset: 6425},
				alternatives: []interface{}{
					&choiceExpr{
						pos: position{line: 213, col: 24, offset: 6427},
						alternatives: []interface{}{
							&litMatcher{
								pos:        position{line: 213, col: 24, offset: 6427},
								val:        "'",
								ignoreCase: false,
							},
							&ruleRefExpr{
								pos:  position{line: 213, col: 30, offset: 6433},
								name: "CommonEscapeSequence",
							},
						},
					},
					&actionExpr{
						pos: position{line: 214, col: 7, offset: 6463},
						run: (*parser).callonSingleStringEscape5,
						expr: &choiceExpr{
							pos: position{line: 214, col: 9, offset: 6465},
							alternatives: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 214, col: 9, offset: 6465},
									name: "SourceChar",
								},
								&ruleRefExpr{
									pos:  position{line: 214, col: 22, offset: 6478},
									name: "EOL",
								},
								&ruleRefExpr{
									pos:  position{line: 214, col: 28, offset: 6484},
									name: "EOF",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "CommonEscapeSequence",
			pos:  position{line: 218, col: 1, offset: 6554},
			expr: &choiceExpr{
				pos: position{line: 218, col: 24, offset: 6579},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 218, col: 24, offset: 6579},
						name: "SingleCharEscape",
					},
					&ruleRefExpr{
						pos:  position{line: 218, col: 43, offset: 6598},
						name: "OctalEscape",
					},
					&ruleRefExpr{
						pos:  position{line: 218, col: 57, offset: 6612},
						name: "HexEscape",
					},
					&ruleRefExpr{
						pos:  position{line: 218, col: 69, offset: 6624},
						name: "LongUnicodeEscape",
					},
					&ruleRefExpr{
						pos:  position{line: 218, col: 89, offset: 6644},
						name: "ShortUnicodeEscape",
					},
				},
			},
		},
		{
			name: "SingleCharEscape",
			pos:  position{line: 219, col: 1, offset: 6664},
			expr: &choiceExpr{
				pos: position{line: 219, col: 20, offset: 6685},
				alternatives: []interface{}{
					&litMatcher{
						pos:        position{line: 219, col: 20, offset: 6685},
						val:        "a",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 219, col: 26, offset: 6691},
						val:        "b",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 219, col: 32, offset: 6697},
						val:        "n",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 219, col: 38, offset: 6703},
						val:        "f",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 219, col: 44, offset: 6709},
						val:        "r",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 219, col: 50, offset: 6715},
						val:        "t",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 219, col: 56, offset: 6721},
						val:        "v",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 219, col: 62, offset: 6727},
						val:        "\\",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "OctalEscape",
			pos:  position{line: 220, col: 1, offset: 6733},
			expr: &choiceExpr{
				pos: position{line: 220, col: 15, offset: 6749},
				alternatives: []interface{}{
					&seqExpr{
						pos: position{line: 220, col: 15, offset: 6749},
						exprs: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 220, col: 15, offset: 6749},
								name: "OctalDigit",
							},
							&ruleRefExpr{
								pos:  position{line: 220, col: 26, offset: 6760},
								name: "OctalDigit",
							},
							&ruleRefExpr{
								pos:  position{line: 220, col: 37, offset: 6771},
								name: "OctalDigit",
							},
						},
					},
					&actionExpr{
						pos: position{line: 221, col: 7, offset: 6789},
						run: (*parser).callonOctalEscape6,
						expr: &seqExpr{
							pos: position{line: 221, col: 7, offset: 6789},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 221, col: 7, offset: 6789},
									name: "OctalDigit",
								},
								&choiceExpr{
									pos: position{line: 221, col: 20, offset: 6802},
									alternatives: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 221, col: 20, offset: 6802},
											name: "SourceChar",
										},
										&ruleRefExpr{
											pos:  position{line: 221, col: 33, offset: 6815},
											name: "EOL",
										},
										&ruleRefExpr{
											pos:  position{line: 221, col: 39, offset: 6821},
											name: "EOF",
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
			name: "HexEscape",
			pos:  position{line: 224, col: 1, offset: 6885},
			expr: &choiceExpr{
				pos: position{line: 224, col: 13, offset: 6899},
				alternatives: []interface{}{
					&seqExpr{
						pos: position{line: 224, col: 13, offset: 6899},
						exprs: []interface{}{
							&litMatcher{
								pos:        position{line: 224, col: 13, offset: 6899},
								val:        "x",
								ignoreCase: false,
							},
							&ruleRefExpr{
								pos:  position{line: 224, col: 17, offset: 6903},
								name: "HexDigit",
							},
							&ruleRefExpr{
								pos:  position{line: 224, col: 26, offset: 6912},
								name: "HexDigit",
							},
						},
					},
					&actionExpr{
						pos: position{line: 225, col: 7, offset: 6928},
						run: (*parser).callonHexEscape6,
						expr: &seqExpr{
							pos: position{line: 225, col: 7, offset: 6928},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 225, col: 7, offset: 6928},
									val:        "x",
									ignoreCase: false,
								},
								&choiceExpr{
									pos: position{line: 225, col: 13, offset: 6934},
									alternatives: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 225, col: 13, offset: 6934},
											name: "SourceChar",
										},
										&ruleRefExpr{
											pos:  position{line: 225, col: 26, offset: 6947},
											name: "EOL",
										},
										&ruleRefExpr{
											pos:  position{line: 225, col: 32, offset: 6953},
											name: "EOF",
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
			name: "LongUnicodeEscape",
			pos:  position{line: 228, col: 1, offset: 7023},
			expr: &choiceExpr{
				pos: position{line: 229, col: 5, offset: 7051},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 229, col: 5, offset: 7051},
						run: (*parser).callonLongUnicodeEscape2,
						expr: &seqExpr{
							pos: position{line: 229, col: 5, offset: 7051},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 229, col: 5, offset: 7051},
									val:        "U",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 229, col: 9, offset: 7055},
									name: "HexDigit",
								},
								&ruleRefExpr{
									pos:  position{line: 229, col: 18, offset: 7064},
									name: "HexDigit",
								},
								&ruleRefExpr{
									pos:  position{line: 229, col: 27, offset: 7073},
									name: "HexDigit",
								},
								&ruleRefExpr{
									pos:  position{line: 229, col: 36, offset: 7082},
									name: "HexDigit",
								},
								&ruleRefExpr{
									pos:  position{line: 229, col: 45, offset: 7091},
									name: "HexDigit",
								},
								&ruleRefExpr{
									pos:  position{line: 229, col: 54, offset: 7100},
									name: "HexDigit",
								},
								&ruleRefExpr{
									pos:  position{line: 229, col: 63, offset: 7109},
									name: "HexDigit",
								},
								&ruleRefExpr{
									pos:  position{line: 229, col: 72, offset: 7118},
									name: "HexDigit",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 232, col: 7, offset: 7223},
						run: (*parser).callonLongUnicodeEscape13,
						expr: &seqExpr{
							pos: position{line: 232, col: 7, offset: 7223},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 232, col: 7, offset: 7223},
									val:        "U",
									ignoreCase: false,
								},
								&choiceExpr{
									pos: position{line: 232, col: 13, offset: 7229},
									alternatives: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 232, col: 13, offset: 7229},
											name: "SourceChar",
										},
										&ruleRefExpr{
											pos:  position{line: 232, col: 26, offset: 7242},
											name: "EOL",
										},
										&ruleRefExpr{
											pos:  position{line: 232, col: 32, offset: 7248},
											name: "EOF",
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
			name: "ShortUnicodeEscape",
			pos:  position{line: 235, col: 1, offset: 7314},
			expr: &choiceExpr{
				pos: position{line: 236, col: 5, offset: 7343},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 236, col: 5, offset: 7343},
						run: (*parser).callonShortUnicodeEscape2,
						expr: &seqExpr{
							pos: position{line: 236, col: 5, offset: 7343},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 236, col: 5, offset: 7343},
									val:        "u",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 236, col: 9, offset: 7347},
									name: "HexDigit",
								},
								&ruleRefExpr{
									pos:  position{line: 236, col: 18, offset: 7356},
									name: "HexDigit",
								},
								&ruleRefExpr{
									pos:  position{line: 236, col: 27, offset: 7365},
									name: "HexDigit",
								},
								&ruleRefExpr{
									pos:  position{line: 236, col: 36, offset: 7374},
									name: "HexDigit",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 239, col: 7, offset: 7479},
						run: (*parser).callonShortUnicodeEscape9,
						expr: &seqExpr{
							pos: position{line: 239, col: 7, offset: 7479},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 239, col: 7, offset: 7479},
									val:        "u",
									ignoreCase: false,
								},
								&choiceExpr{
									pos: position{line: 239, col: 13, offset: 7485},
									alternatives: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 239, col: 13, offset: 7485},
											name: "SourceChar",
										},
										&ruleRefExpr{
											pos:  position{line: 239, col: 26, offset: 7498},
											name: "EOL",
										},
										&ruleRefExpr{
											pos:  position{line: 239, col: 32, offset: 7504},
											name: "EOF",
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
			name: "OctalDigit",
			pos:  position{line: 243, col: 1, offset: 7572},
			expr: &charClassMatcher{
				pos:        position{line: 243, col: 14, offset: 7587},
				val:        "[0-7]",
				ranges:     []rune{'0', '7'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "DecimalDigit",
			pos:  position{line: 244, col: 1, offset: 7594},
			expr: &charClassMatcher{
				pos:        position{line: 244, col: 16, offset: 7611},
				val:        "[0-9]",
				ranges:     []rune{'0', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "HexDigit",
			pos:  position{line: 245, col: 1, offset: 7618},
			expr: &charClassMatcher{
				pos:        position{line: 245, col: 12, offset: 7631},
				val:        "[0-9a-f]i",
				ranges:     []rune{'0', '9', 'a', 'f'},
				ignoreCase: true,
				inverted:   false,
			},
		},
		{
			name: "CharClassMatcher",
			pos:  position{line: 247, col: 1, offset: 7644},
			expr: &choiceExpr{
				pos: position{line: 247, col: 20, offset: 7665},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 247, col: 20, offset: 7665},
						run: (*parser).callonCharClassMatcher2,
						expr: &seqExpr{
							pos: position{line: 247, col: 20, offset: 7665},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 247, col: 20, offset: 7665},
									val:        "[",
									ignoreCase: false,
								},
								&zeroOrMoreExpr{
									pos: position{line: 247, col: 24, offset: 7669},
									expr: &choiceExpr{
										pos: position{line: 247, col: 26, offset: 7671},
										alternatives: []interface{}{
											&ruleRefExpr{
												pos:  position{line: 247, col: 26, offset: 7671},
												name: "ClassCharRange",
											},
											&ruleRefExpr{
												pos:  position{line: 247, col: 43, offset: 7688},
												name: "ClassChar",
											},
											&seqExpr{
												pos: position{line: 247, col: 55, offset: 7700},
												exprs: []interface{}{
													&litMatcher{
														pos:        position{line: 247, col: 55, offset: 7700},
														val:        "\\",
														ignoreCase: false,
													},
													&ruleRefExpr{
														pos:  position{line: 247, col: 60, offset: 7705},
														name: "UnicodeClassEscape",
													},
												},
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 247, col: 82, offset: 7727},
									val:        "]",
									ignoreCase: false,
								},
								&zeroOrOneExpr{
									pos: position{line: 247, col: 86, offset: 7731},
									expr: &litMatcher{
										pos:        position{line: 247, col: 86, offset: 7731},
										val:        "i",
										ignoreCase: false,
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 251, col: 5, offset: 7842},
						run: (*parser).callonCharClassMatcher15,
						expr: &seqExpr{
							pos: position{line: 251, col: 5, offset: 7842},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 251, col: 5, offset: 7842},
									val:        "[",
									ignoreCase: false,
								},
								&zeroOrMoreExpr{
									pos: position{line: 251, col: 9, offset: 7846},
									expr: &seqExpr{
										pos: position{line: 251, col: 11, offset: 7848},
										exprs: []interface{}{
											&notExpr{
												pos: position{line: 251, col: 11, offset: 7848},
												expr: &ruleRefExpr{
													pos:  position{line: 251, col: 14, offset: 7851},
													name: "EOL",
												},
											},
											&ruleRefExpr{
												pos:  position{line: 251, col: 20, offset: 7857},
												name: "SourceChar",
											},
										},
									},
								},
								&choiceExpr{
									pos: position{line: 251, col: 36, offset: 7873},
									alternatives: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 251, col: 36, offset: 7873},
											name: "EOL",
										},
										&ruleRefExpr{
											pos:  position{line: 251, col: 42, offset: 7879},
											name: "EOF",
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
			name: "ClassCharRange",
			pos:  position{line: 255, col: 1, offset: 7993},
			expr: &seqExpr{
				pos: position{line: 255, col: 18, offset: 8012},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 255, col: 18, offset: 8012},
						name: "ClassChar",
					},
					&litMatcher{
						pos:        position{line: 255, col: 28, offset: 8022},
						val:        "-",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 255, col: 32, offset: 8026},
						name: "ClassChar",
					},
				},
			},
		},
		{
			name: "ClassChar",
			pos:  position{line: 256, col: 1, offset: 8037},
			expr: &choiceExpr{
				pos: position{line: 256, col: 13, offset: 8051},
				alternatives: []interface{}{
					&seqExpr{
						pos: position{line: 256, col: 13, offset: 8051},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 256, col: 13, offset: 8051},
								expr: &choiceExpr{
									pos: position{line: 256, col: 16, offset: 8054},
									alternatives: []interface{}{
										&litMatcher{
											pos:        position{line: 256, col: 16, offset: 8054},
											val:        "]",
											ignoreCase: false,
										},
										&litMatcher{
											pos:        position{line: 256, col: 22, offset: 8060},
											val:        "\\",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 256, col: 29, offset: 8067},
											name: "EOL",
										},
									},
								},
							},
							&ruleRefExpr{
								pos:  position{line: 256, col: 35, offset: 8073},
								name: "SourceChar",
							},
						},
					},
					&seqExpr{
						pos: position{line: 256, col: 48, offset: 8086},
						exprs: []interface{}{
							&litMatcher{
								pos:        position{line: 256, col: 48, offset: 8086},
								val:        "\\",
								ignoreCase: false,
							},
							&ruleRefExpr{
								pos:  position{line: 256, col: 53, offset: 8091},
								name: "CharClassEscape",
							},
						},
					},
				},
			},
		},
		{
			name: "CharClassEscape",
			pos:  position{line: 257, col: 1, offset: 8108},
			expr: &choiceExpr{
				pos: position{line: 257, col: 19, offset: 8128},
				alternatives: []interface{}{
					&choiceExpr{
						pos: position{line: 257, col: 21, offset: 8130},
						alternatives: []interface{}{
							&litMatcher{
								pos:        position{line: 257, col: 21, offset: 8130},
								val:        "]",
								ignoreCase: false,
							},
							&ruleRefExpr{
								pos:  position{line: 257, col: 27, offset: 8136},
								name: "CommonEscapeSequence",
							},
						},
					},
					&actionExpr{
						pos: position{line: 258, col: 7, offset: 8166},
						run: (*parser).callonCharClassEscape5,
						expr: &seqExpr{
							pos: position{line: 258, col: 7, offset: 8166},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 258, col: 7, offset: 8166},
									expr: &litMatcher{
										pos:        position{line: 258, col: 8, offset: 8167},
										val:        "p",
										ignoreCase: false,
									},
								},
								&choiceExpr{
									pos: position{line: 258, col: 14, offset: 8173},
									alternatives: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 258, col: 14, offset: 8173},
											name: "SourceChar",
										},
										&ruleRefExpr{
											pos:  position{line: 258, col: 27, offset: 8186},
											name: "EOL",
										},
										&ruleRefExpr{
											pos:  position{line: 258, col: 33, offset: 8192},
											name: "EOF",
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
			name: "UnicodeClassEscape",
			pos:  position{line: 262, col: 1, offset: 8262},
			expr: &seqExpr{
				pos: position{line: 262, col: 22, offset: 8285},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 262, col: 22, offset: 8285},
						val:        "p",
						ignoreCase: false,
					},
					&choiceExpr{
						pos: position{line: 263, col: 7, offset: 8299},
						alternatives: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 263, col: 7, offset: 8299},
								name: "SingleCharUnicodeClass",
							},
							&actionExpr{
								pos: position{line: 264, col: 7, offset: 8329},
								run: (*parser).callonUnicodeClassEscape5,
								expr: &seqExpr{
									pos: position{line: 264, col: 7, offset: 8329},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 264, col: 7, offset: 8329},
											expr: &litMatcher{
												pos:        position{line: 264, col: 8, offset: 8330},
												val:        "{",
												ignoreCase: false,
											},
										},
										&choiceExpr{
											pos: position{line: 264, col: 14, offset: 8336},
											alternatives: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 264, col: 14, offset: 8336},
													name: "SourceChar",
												},
												&ruleRefExpr{
													pos:  position{line: 264, col: 27, offset: 8349},
													name: "EOL",
												},
												&ruleRefExpr{
													pos:  position{line: 264, col: 33, offset: 8355},
													name: "EOF",
												},
											},
										},
									},
								},
							},
							&actionExpr{
								pos: position{line: 265, col: 7, offset: 8427},
								run: (*parser).callonUnicodeClassEscape13,
								expr: &seqExpr{
									pos: position{line: 265, col: 7, offset: 8427},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 265, col: 7, offset: 8427},
											val:        "{",
											ignoreCase: false,
										},
										&labeledExpr{
											pos:   position{line: 265, col: 11, offset: 8431},
											label: "ident",
											expr: &ruleRefExpr{
												pos:  position{line: 265, col: 17, offset: 8437},
												name: "IdentifierName",
											},
										},
										&litMatcher{
											pos:        position{line: 265, col: 32, offset: 8452},
											val:        "}",
											ignoreCase: false,
										},
									},
								},
							},
							&actionExpr{
								pos: position{line: 271, col: 7, offset: 8635},
								run: (*parser).callonUnicodeClassEscape19,
								expr: &seqExpr{
									pos: position{line: 271, col: 7, offset: 8635},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 271, col: 7, offset: 8635},
											val:        "{",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 271, col: 11, offset: 8639},
											name: "IdentifierName",
										},
										&choiceExpr{
											pos: position{line: 271, col: 28, offset: 8656},
											alternatives: []interface{}{
												&litMatcher{
													pos:        position{line: 271, col: 28, offset: 8656},
													val:        "]",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 271, col: 34, offset: 8662},
													name: "EOL",
												},
												&ruleRefExpr{
													pos:  position{line: 271, col: 40, offset: 8668},
													name: "EOF",
												},
											},
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
			name: "SingleCharUnicodeClass",
			pos:  position{line: 275, col: 1, offset: 8755},
			expr: &charClassMatcher{
				pos:        position{line: 275, col: 26, offset: 8782},
				val:        "[LMNCPZS]",
				chars:      []rune{'L', 'M', 'N', 'C', 'P', 'Z', 'S'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "AnyMatcher",
			pos:  position{line: 277, col: 1, offset: 8795},
			expr: &actionExpr{
				pos: position{line: 277, col: 14, offset: 8810},
				run: (*parser).callonAnyMatcher1,
				expr: &litMatcher{
					pos:        position{line: 277, col: 14, offset: 8810},
					val:        ".",
					ignoreCase: false,
				},
			},
		},
		{
			name: "CodeBlock",
			pos:  position{line: 282, col: 1, offset: 8890},
			expr: &choiceExpr{
				pos: position{line: 282, col: 13, offset: 8904},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 282, col: 13, offset: 8904},
						run: (*parser).callonCodeBlock2,
						expr: &seqExpr{
							pos: position{line: 282, col: 13, offset: 8904},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 282, col: 13, offset: 8904},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 282, col: 17, offset: 8908},
									name: "Code",
								},
								&litMatcher{
									pos:        position{line: 282, col: 22, offset: 8913},
									val:        "}",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 286, col: 5, offset: 9016},
						run: (*parser).callonCodeBlock7,
						expr: &seqExpr{
							pos: position{line: 286, col: 5, offset: 9016},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 286, col: 5, offset: 9016},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 286, col: 9, offset: 9020},
									name: "Code",
								},
								&ruleRefExpr{
									pos:  position{line: 286, col: 14, offset: 9025},
									name: "EOF",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Code",
			pos:  position{line: 290, col: 1, offset: 9094},
			expr: &zeroOrMoreExpr{
				pos: position{line: 290, col: 8, offset: 9103},
				expr: &choiceExpr{
					pos: position{line: 290, col: 10, offset: 9105},
					alternatives: []interface{}{
						&oneOrMoreExpr{
							pos: position{line: 290, col: 10, offset: 9105},
							expr: &seqExpr{
								pos: position{line: 290, col: 12, offset: 9107},
								exprs: []interface{}{
									&notExpr{
										pos: position{line: 290, col: 12, offset: 9107},
										expr: &charClassMatcher{
											pos:        position{line: 290, col: 13, offset: 9108},
											val:        "[{}]",
											chars:      []rune{'{', '}'},
											ignoreCase: false,
											inverted:   false,
										},
									},
									&ruleRefExpr{
										pos:  position{line: 290, col: 18, offset: 9113},
										name: "SourceChar",
									},
								},
							},
						},
						&seqExpr{
							pos: position{line: 290, col: 34, offset: 9129},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 290, col: 34, offset: 9129},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 290, col: 38, offset: 9133},
									name: "Code",
								},
								&litMatcher{
									pos:        position{line: 290, col: 43, offset: 9138},
									val:        "}",
									ignoreCase: false,
								},
							},
						},
					},
				},
			},
		},
		{
			name: "__",
			pos:  position{line: 292, col: 1, offset: 9148},
			expr: &zeroOrMoreExpr{
				pos: position{line: 292, col: 6, offset: 9155},
				expr: &choiceExpr{
					pos: position{line: 292, col: 8, offset: 9157},
					alternatives: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 292, col: 8, offset: 9157},
							name: "Whitespace",
						},
						&ruleRefExpr{
							pos:  position{line: 292, col: 21, offset: 9170},
							name: "EOL",
						},
						&ruleRefExpr{
							pos:  position{line: 292, col: 27, offset: 9176},
							name: "Comment",
						},
					},
				},
			},
		},
		{
			name: "_",
			pos:  position{line: 293, col: 1, offset: 9188},
			expr: &zeroOrMoreExpr{
				pos: position{line: 293, col: 5, offset: 9194},
				expr: &choiceExpr{
					pos: position{line: 293, col: 7, offset: 9196},
					alternatives: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 293, col: 7, offset: 9196},
							name: "Whitespace",
						},
						&ruleRefExpr{
							pos:  position{line: 293, col: 20, offset: 9209},
							name: "MultiLineCommentNoLineTerminator",
						},
					},
				},
			},
		},
		{
			name: "Whitespace",
			pos:  position{line: 295, col: 1, offset: 9248},
			expr: &charClassMatcher{
				pos:        position{line: 295, col: 14, offset: 9263},
				val:        "[ \\t\\r]",
				chars:      []rune{' ', '\t', '\r'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "EOL",
			pos:  position{line: 296, col: 1, offset: 9272},
			expr: &litMatcher{
				pos:        position{line: 296, col: 7, offset: 9280},
				val:        "\n",
				ignoreCase: false,
			},
		},
		{
			name: "EOS",
			pos:  position{line: 297, col: 1, offset: 9286},
			expr: &choiceExpr{
				pos: position{line: 297, col: 7, offset: 9294},
				alternatives: []interface{}{
					&seqExpr{
						pos: position{line: 297, col: 7, offset: 9294},
						exprs: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 297, col: 7, offset: 9294},
								name: "__",
							},
							&litMatcher{
								pos:        position{line: 297, col: 10, offset: 9297},
								val:        ";",
								ignoreCase: false,
							},
						},
					},
					&seqExpr{
						pos: position{line: 297, col: 16, offset: 9303},
						exprs: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 297, col: 16, offset: 9303},
								name: "_",
							},
							&zeroOrOneExpr{
								pos: position{line: 297, col: 18, offset: 9305},
								expr: &ruleRefExpr{
									pos:  position{line: 297, col: 18, offset: 9305},
									name: "SingleLineComment",
								},
							},
							&ruleRefExpr{
								pos:  position{line: 297, col: 37, offset: 9324},
								name: "EOL",
							},
						},
					},
					&seqExpr{
						pos: position{line: 297, col: 43, offset: 9330},
						exprs: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 297, col: 43, offset: 9330},
								name: "__",
							},
							&ruleRefExpr{
								pos:  position{line: 297, col: 46, offset: 9333},
								name: "EOF",
							},
						},
					},
				},
			},
		},
		{
			name: "EOF",
			pos:  position{line: 299, col: 1, offset: 9340},
			expr: &notExpr{
				pos: position{line: 299, col: 7, offset: 9348},
				expr: &anyMatcher{
					line: 299, col: 8, offset: 9349,
				},
			},
		},
	},
}

func (c *current) onGrammar1(initializer, rules interface{}) (interface{}, error) {

	pos := c.astPos()

	// create the grammar, assign its initializer
	g := ast.NewGrammar(pos)
	initSlice := toIfaceSlice(initializer)
	if len(initSlice) > 0 {
		g.Init = initSlice[0].(*ast.CodeBlock)
	}

	rulesSlice := toIfaceSlice(rules)
	g.Rules = make([]*ast.Rule, len(rulesSlice))
	for i, duo := range rulesSlice {
		g.Rules[i] = duo.([]interface{})[0].(*ast.Rule)
	}

	return g, nil
}

func (p *parser) callonGrammar1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onGrammar1(stack["initializer"], stack["rules"])
}

func (c *current) onInitializer1(code interface{}) (interface{}, error) {

	return code, nil
}

func (p *parser) callonInitializer1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onInitializer1(stack["code"])
}

func (c *current) onRule1(name, display, expr interface{}) (interface{}, error) {

	pos := c.astPos()

	rule := ast.NewRule(pos, name.(*ast.Identifier))
	displaySlice := toIfaceSlice(display)
	if len(displaySlice) > 0 {
		rule.DisplayName = displaySlice[0].(*ast.StringLit)
	}
	rule.Expr = expr.(ast.Expression)

	return rule, nil
}

func (p *parser) callonRule1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onRule1(stack["name"], stack["display"], stack["expr"])
}

func (c *current) onChoiceExpr1(first, rest interface{}) (interface{}, error) {

	restSlice := toIfaceSlice(rest)
	if len(restSlice) == 0 {
		return first, nil
	}

	pos := c.astPos()
	choice := ast.NewChoiceExpr(pos)
	choice.Alternatives = []ast.Expression{first.(ast.Expression)}
	for _, sl := range restSlice {
		choice.Alternatives = append(choice.Alternatives, sl.([]interface{})[3].(ast.Expression))
	}
	return choice, nil
}

func (p *parser) callonChoiceExpr1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onChoiceExpr1(stack["first"], stack["rest"])
}

func (c *current) onActionExpr1(expr, code interface{}) (interface{}, error) {

	if code == nil {
		return expr, nil
	}

	pos := c.astPos()
	act := ast.NewActionExpr(pos)
	act.Expr = expr.(ast.Expression)
	codeSlice := toIfaceSlice(code)
	act.Code = codeSlice[1].(*ast.CodeBlock)

	return act, nil
}

func (p *parser) callonActionExpr1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onActionExpr1(stack["expr"], stack["code"])
}

func (c *current) onSeqExpr1(first, rest interface{}) (interface{}, error) {

	restSlice := toIfaceSlice(rest)
	if len(restSlice) == 0 {
		return first, nil
	}
	seq := ast.NewSeqExpr(c.astPos())
	seq.Exprs = []ast.Expression{first.(ast.Expression)}
	for _, sl := range restSlice {
		seq.Exprs = append(seq.Exprs, sl.([]interface{})[1].(ast.Expression))
	}
	return seq, nil
}

func (p *parser) callonSeqExpr1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSeqExpr1(stack["first"], stack["rest"])
}

func (c *current) onLabeledExpr2(label, expr interface{}) (interface{}, error) {

	pos := c.astPos()
	lab := ast.NewLabeledExpr(pos)
	lab.Label = label.(*ast.Identifier)
	lab.Expr = expr.(ast.Expression)
	return lab, nil
}

func (p *parser) callonLabeledExpr2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onLabeledExpr2(stack["label"], stack["expr"])
}

func (c *current) onPrefixedExpr2(op, expr interface{}) (interface{}, error) {

	pos := c.astPos()
	opStr := op.(string)
	if opStr == "&" {
		and := ast.NewAndExpr(pos)
		and.Expr = expr.(ast.Expression)
		return and, nil
	}

	not := ast.NewNotExpr(pos)
	not.Expr = expr.(ast.Expression)
	return not, nil
}

func (p *parser) callonPrefixedExpr2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onPrefixedExpr2(stack["op"], stack["expr"])
}

func (c *current) onPrefixedOp1() (interface{}, error) {

	return string(c.text), nil
}

func (p *parser) callonPrefixedOp1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onPrefixedOp1()
}

func (c *current) onSuffixedExpr2(expr, op interface{}) (interface{}, error) {

	pos := c.astPos()
	opStr := op.(string)
	switch opStr {
	case "?":
		zero := ast.NewZeroOrOneExpr(pos)
		zero.Expr = expr.(ast.Expression)
		return zero, nil
	case "*":
		zero := ast.NewZeroOrMoreExpr(pos)
		zero.Expr = expr.(ast.Expression)
		return zero, nil
	case "+":
		one := ast.NewOneOrMoreExpr(pos)
		one.Expr = expr.(ast.Expression)
		return one, nil
	default:
		return nil, errors.New("unknown operator: " + opStr)
	}
}

func (p *parser) callonSuffixedExpr2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSuffixedExpr2(stack["expr"], stack["op"])
}

func (c *current) onSuffixedOp1() (interface{}, error) {

	return string(c.text), nil
}

func (p *parser) callonSuffixedOp1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSuffixedOp1()
}

func (c *current) onPrimaryExpr7(expr interface{}) (interface{}, error) {

	return expr, nil
}

func (p *parser) callonPrimaryExpr7() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onPrimaryExpr7(stack["expr"])
}

func (c *current) onRuleRefExpr1(name interface{}) (interface{}, error) {

	ref := ast.NewRuleRefExpr(c.astPos())
	ref.Name = name.(*ast.Identifier)
	return ref, nil
}

func (p *parser) callonRuleRefExpr1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onRuleRefExpr1(stack["name"])
}

func (c *current) onSemanticPredExpr1(op, code interface{}) (interface{}, error) {

	opStr := op.(string)
	if opStr == "&" {
		and := ast.NewAndCodeExpr(c.astPos())
		and.Code = code.(*ast.CodeBlock)
		return and, nil
	} else if opStr == "#" {
		state := ast.NewStateCodeExpr(c.astPos())
		state.Code = code.(*ast.CodeBlock)
		return state, nil
	}

	not := ast.NewNotCodeExpr(c.astPos())
	not.Code = code.(*ast.CodeBlock)
	return not, nil
}

func (p *parser) callonSemanticPredExpr1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSemanticPredExpr1(stack["op"], stack["code"])
}

func (c *current) onSemanticPredOp1() (interface{}, error) {

	return string(c.text), nil
}

func (p *parser) callonSemanticPredOp1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSemanticPredOp1()
}

func (c *current) onIdentifier1(ident interface{}) (interface{}, error) {

	astIdent := ast.NewIdentifier(c.astPos(), string(c.text))
	if reservedWords[astIdent.Val] {
		return astIdent, errors.New("identifier is a reserved word")
	}
	return astIdent, nil
}

func (p *parser) callonIdentifier1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onIdentifier1(stack["ident"])
}

func (c *current) onIdentifierName1() (interface{}, error) {

	return ast.NewIdentifier(c.astPos(), string(c.text)), nil
}

func (p *parser) callonIdentifierName1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onIdentifierName1()
}

func (c *current) onLitMatcher1(lit, ignore interface{}) (interface{}, error) {

	rawStr := lit.(*ast.StringLit).Val
	s, err := strconv.Unquote(rawStr)
	if err != nil {
		// an invalid string literal raises an error in the escape rules,
		// so simply replace the literal with an empty string here to
		// avoid a cascade of errors.
		s = ""
	}
	m := ast.NewLitMatcher(c.astPos(), s)
	m.IgnoreCase = ignore != nil
	return m, nil
}

func (p *parser) callonLitMatcher1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onLitMatcher1(stack["lit"], stack["ignore"])
}

func (c *current) onStringLiteral2() (interface{}, error) {

	return ast.NewStringLit(c.astPos(), string(c.text)), nil
}

func (p *parser) callonStringLiteral2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onStringLiteral2()
}

func (c *current) onStringLiteral18() (interface{}, error) {

	return ast.NewStringLit(c.astPos(), "``"), errors.New("string literal not terminated")
}

func (p *parser) callonStringLiteral18() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onStringLiteral18()
}

func (c *current) onDoubleStringEscape5() (interface{}, error) {

	return nil, errors.New("invalid escape character")
}

func (p *parser) callonDoubleStringEscape5() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onDoubleStringEscape5()
}

func (c *current) onSingleStringEscape5() (interface{}, error) {

	return nil, errors.New("invalid escape character")
}

func (p *parser) callonSingleStringEscape5() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSingleStringEscape5()
}

func (c *current) onOctalEscape6() (interface{}, error) {

	return nil, errors.New("invalid octal escape")
}

func (p *parser) callonOctalEscape6() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onOctalEscape6()
}

func (c *current) onHexEscape6() (interface{}, error) {

	return nil, errors.New("invalid hexadecimal escape")
}

func (p *parser) callonHexEscape6() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onHexEscape6()
}

func (c *current) onLongUnicodeEscape2() (interface{}, error) {

	return validateUnicodeEscape(string(c.text), "invalid Unicode escape")

}

func (p *parser) callonLongUnicodeEscape2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onLongUnicodeEscape2()
}

func (c *current) onLongUnicodeEscape13() (interface{}, error) {

	return nil, errors.New("invalid Unicode escape")
}

func (p *parser) callonLongUnicodeEscape13() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onLongUnicodeEscape13()
}

func (c *current) onShortUnicodeEscape2() (interface{}, error) {

	return validateUnicodeEscape(string(c.text), "invalid Unicode escape")

}

func (p *parser) callonShortUnicodeEscape2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onShortUnicodeEscape2()
}

func (c *current) onShortUnicodeEscape9() (interface{}, error) {

	return nil, errors.New("invalid Unicode escape")
}

func (p *parser) callonShortUnicodeEscape9() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onShortUnicodeEscape9()
}

func (c *current) onCharClassMatcher2() (interface{}, error) {

	pos := c.astPos()
	cc := ast.NewCharClassMatcher(pos, string(c.text))
	return cc, nil
}

func (p *parser) callonCharClassMatcher2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onCharClassMatcher2()
}

func (c *current) onCharClassMatcher15() (interface{}, error) {

	return ast.NewCharClassMatcher(c.astPos(), "[]"), errors.New("character class not terminated")
}

func (p *parser) callonCharClassMatcher15() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onCharClassMatcher15()
}

func (c *current) onCharClassEscape5() (interface{}, error) {

	return nil, errors.New("invalid escape character")
}

func (p *parser) callonCharClassEscape5() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onCharClassEscape5()
}

func (c *current) onUnicodeClassEscape5() (interface{}, error) {
	return nil, errors.New("invalid Unicode class escape")
}

func (p *parser) callonUnicodeClassEscape5() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onUnicodeClassEscape5()
}

func (c *current) onUnicodeClassEscape13(ident interface{}) (interface{}, error) {

	if !unicodeClasses[ident.(*ast.Identifier).Val] {
		return nil, errors.New("invalid Unicode class escape")
	}
	return nil, nil

}

func (p *parser) callonUnicodeClassEscape13() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onUnicodeClassEscape13(stack["ident"])
}

func (c *current) onUnicodeClassEscape19() (interface{}, error) {

	return nil, errors.New("Unicode class not terminated")

}

func (p *parser) callonUnicodeClassEscape19() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onUnicodeClassEscape19()
}

func (c *current) onAnyMatcher1() (interface{}, error) {

	any := ast.NewAnyMatcher(c.astPos(), ".")
	return any, nil
}

func (p *parser) callonAnyMatcher1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onAnyMatcher1()
}

func (c *current) onCodeBlock2() (interface{}, error) {

	pos := c.astPos()
	cb := ast.NewCodeBlock(pos, string(c.text))
	return cb, nil
}

func (p *parser) callonCodeBlock2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onCodeBlock2()
}

func (c *current) onCodeBlock7() (interface{}, error) {

	return nil, errors.New("code block not terminated")
}

func (p *parser) callonCodeBlock7() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onCodeBlock7()
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

// ShowMaxPos creates an Option to set the showMaxPos flag to b. When set to true,
// max position is printed to stdout after parsing.
//
// The default is false.
func ShowMaxPos(b bool) Option {
	return func(p *parser) Option {
		old := p.showMaxPos
		p.showMaxPos = b
		return ShowMaxPos(old)
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
	// max position of parser cursor
	maxpos position
	// show max position
	showMaxPos bool
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
	p.tstack = append(p.tstack, e)
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
		p.popT()
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
		p.maxpos.line = p.pt.line
		p.maxpos.col = 0
	}

	if rn == utf8.RuneError {
		if n == 1 {
			p.addErr(errInvalidEncoding)
		}
	}
	p.maxpos.col = max(p.maxpos.col, p.pt.col)
	p.maxpos.offset = p.pt.offset
}

// maximum of 2 numbers
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
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
	if p.showMaxPos {
		fmt.Printf("maxpos = [%d:%d:%d]\n", p.maxpos.line, p.maxpos.col, p.maxpos.offset)
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
	if !ok {
		if len(*p.errs) == 0 {
			// make sure this doesn't go out silently
			p.addErr(errNoMatch)
		}
		return nil, p.errs.err()
	}

	return val, p.errs.err()
}

func (p *parser) parseRule(rule *rule) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseRule " + rule.name))
	}
	var j *TEntry
	if p.jdebug {
		j = p.jin("Rule " + rule.name)
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
		defer p.jout(p.jin("OneOrMoreExpr"))
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
		defer p.jout(p.jin("ZeroOrMoreExpr"))
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
		defer p.jout(p.jin("ZeroOrOneExpr"))
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

package main

import (
	"errors"
	"strconv"

	"github.com/mna/pigeon/ast"
)

var g = &grammar{
	rules: []*rule{
		{
			name: "Grammar",
			pos:  position{line: 5, col: 1, offset: 18},
			expr: &actionExpr{
				pos: position{line: 5, col: 11, offset: 30},
				run: (*parser).callonGrammar1,
				expr: &seqExpr{
					pos: position{line: 5, col: 11, offset: 30},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 5, col: 11, offset: 30},
							name: "__",
						},
						&labeledExpr{
							pos:   position{line: 5, col: 14, offset: 33},
							label: "initializer",
							expr: &zeroOrOneExpr{
								pos: position{line: 5, col: 28, offset: 47},
								expr: &seqExpr{
									pos: position{line: 5, col: 28, offset: 47},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 5, col: 28, offset: 47},
											name: "Initializer",
										},
										&ruleRefExpr{
											pos:  position{line: 5, col: 40, offset: 59},
											name: "__",
										},
									},
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 5, col: 46, offset: 65},
							label: "rules",
							expr: &oneOrMoreExpr{
								pos: position{line: 5, col: 54, offset: 73},
								expr: &seqExpr{
									pos: position{line: 5, col: 54, offset: 73},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 5, col: 54, offset: 73},
											name: "Rule",
										},
										&ruleRefExpr{
											pos:  position{line: 5, col: 59, offset: 78},
											name: "__",
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
			name: "Initializer",
			pos:  position{line: 24, col: 1, offset: 521},
			expr: &actionExpr{
				pos: position{line: 24, col: 15, offset: 537},
				run: (*parser).callonInitializer1,
				expr: &seqExpr{
					pos: position{line: 24, col: 15, offset: 537},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 24, col: 15, offset: 537},
							label: "code",
							expr: &ruleRefExpr{
								pos:  position{line: 24, col: 20, offset: 542},
								name: "CodeBlock",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 24, col: 30, offset: 552},
							name: "EOS",
						},
					},
				},
			},
		},
		{
			name: "Rule",
			pos:  position{line: 28, col: 1, offset: 582},
			expr: &actionExpr{
				pos: position{line: 28, col: 8, offset: 591},
				run: (*parser).callonRule1,
				expr: &seqExpr{
					pos: position{line: 28, col: 8, offset: 591},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 28, col: 8, offset: 591},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 28, col: 13, offset: 596},
								name: "IdentifierName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 28, col: 28, offset: 611},
							name: "__",
						},
						&labeledExpr{
							pos:   position{line: 28, col: 31, offset: 614},
							label: "display",
							expr: &zeroOrOneExpr{
								pos: position{line: 28, col: 41, offset: 624},
								expr: &seqExpr{
									pos: position{line: 28, col: 41, offset: 624},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 28, col: 41, offset: 624},
											name: "StringLiteral",
										},
										&ruleRefExpr{
											pos:  position{line: 28, col: 55, offset: 638},
											name: "__",
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 28, col: 61, offset: 644},
							name: "RuleDefOp",
						},
						&ruleRefExpr{
							pos:  position{line: 28, col: 71, offset: 654},
							name: "__",
						},
						&labeledExpr{
							pos:   position{line: 28, col: 74, offset: 657},
							label: "expr",
							expr: &ruleRefExpr{
								pos:  position{line: 28, col: 79, offset: 662},
								name: "Expression",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 28, col: 90, offset: 673},
							name: "EOS",
						},
					},
				},
			},
		},
		{
			name: "Expression",
			pos:  position{line: 41, col: 1, offset: 957},
			expr: &ruleRefExpr{
				pos:  position{line: 41, col: 14, offset: 972},
				name: "ChoiceExpr",
			},
		},
		{
			name: "ChoiceExpr",
			pos:  position{line: 43, col: 1, offset: 984},
			expr: &actionExpr{
				pos: position{line: 43, col: 14, offset: 999},
				run: (*parser).callonChoiceExpr1,
				expr: &seqExpr{
					pos: position{line: 43, col: 14, offset: 999},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 43, col: 14, offset: 999},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 43, col: 20, offset: 1005},
								name: "ActionExpr",
							},
						},
						&labeledExpr{
							pos:   position{line: 43, col: 31, offset: 1016},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 43, col: 38, offset: 1023},
								expr: &seqExpr{
									pos: position{line: 43, col: 38, offset: 1023},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 43, col: 38, offset: 1023},
											name: "__",
										},
										&litMatcher{
											pos:        position{line: 43, col: 41, offset: 1026},
											val:        "/",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 43, col: 45, offset: 1030},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 43, col: 48, offset: 1033},
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
			pos:  position{line: 58, col: 1, offset: 1438},
			expr: &actionExpr{
				pos: position{line: 58, col: 14, offset: 1453},
				run: (*parser).callonActionExpr1,
				expr: &seqExpr{
					pos: position{line: 58, col: 14, offset: 1453},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 58, col: 14, offset: 1453},
							label: "expr",
							expr: &ruleRefExpr{
								pos:  position{line: 58, col: 19, offset: 1458},
								name: "SeqExpr",
							},
						},
						&labeledExpr{
							pos:   position{line: 58, col: 27, offset: 1466},
							label: "code",
							expr: &zeroOrOneExpr{
								pos: position{line: 58, col: 34, offset: 1473},
								expr: &seqExpr{
									pos: position{line: 58, col: 34, offset: 1473},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 58, col: 34, offset: 1473},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 58, col: 37, offset: 1476},
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
			pos:  position{line: 72, col: 1, offset: 1742},
			expr: &actionExpr{
				pos: position{line: 72, col: 11, offset: 1754},
				run: (*parser).callonSeqExpr1,
				expr: &seqExpr{
					pos: position{line: 72, col: 11, offset: 1754},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 72, col: 11, offset: 1754},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 72, col: 17, offset: 1760},
								name: "LabeledExpr",
							},
						},
						&labeledExpr{
							pos:   position{line: 72, col: 29, offset: 1772},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 72, col: 36, offset: 1779},
								expr: &seqExpr{
									pos: position{line: 72, col: 36, offset: 1779},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 72, col: 36, offset: 1779},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 72, col: 39, offset: 1782},
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
			pos:  position{line: 85, col: 1, offset: 2133},
			expr: &choiceExpr{
				pos: position{line: 85, col: 15, offset: 2149},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 85, col: 15, offset: 2149},
						run: (*parser).callonLabeledExpr2,
						expr: &seqExpr{
							pos: position{line: 85, col: 15, offset: 2149},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 85, col: 15, offset: 2149},
									label: "label",
									expr: &ruleRefExpr{
										pos:  position{line: 85, col: 21, offset: 2155},
										name: "Identifier",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 85, col: 32, offset: 2166},
									name: "__",
								},
								&litMatcher{
									pos:        position{line: 85, col: 35, offset: 2169},
									val:        ":",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 85, col: 39, offset: 2173},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 85, col: 42, offset: 2176},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 85, col: 47, offset: 2181},
										name: "PrefixedExpr",
									},
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 91, col: 5, offset: 2354},
						name: "PrefixedExpr",
					},
				},
			},
		},
		{
			name: "PrefixedExpr",
			pos:  position{line: 93, col: 1, offset: 2368},
			expr: &choiceExpr{
				pos: position{line: 93, col: 16, offset: 2385},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 93, col: 16, offset: 2385},
						run: (*parser).callonPrefixedExpr2,
						expr: &seqExpr{
							pos: position{line: 93, col: 16, offset: 2385},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 93, col: 16, offset: 2385},
									label: "op",
									expr: &ruleRefExpr{
										pos:  position{line: 93, col: 19, offset: 2388},
										name: "PrefixedOp",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 93, col: 30, offset: 2399},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 93, col: 33, offset: 2402},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 93, col: 38, offset: 2407},
										name: "SuffixedExpr",
									},
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 104, col: 5, offset: 2689},
						name: "SuffixedExpr",
					},
				},
			},
		},
		{
			name: "PrefixedOp",
			pos:  position{line: 106, col: 1, offset: 2703},
			expr: &actionExpr{
				pos: position{line: 106, col: 14, offset: 2718},
				run: (*parser).callonPrefixedOp1,
				expr: &choiceExpr{
					pos: position{line: 106, col: 16, offset: 2720},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 106, col: 16, offset: 2720},
							val:        "&",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 106, col: 22, offset: 2726},
							val:        "!",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "SuffixedExpr",
			pos:  position{line: 110, col: 1, offset: 2768},
			expr: &choiceExpr{
				pos: position{line: 110, col: 16, offset: 2785},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 110, col: 16, offset: 2785},
						run: (*parser).callonSuffixedExpr2,
						expr: &seqExpr{
							pos: position{line: 110, col: 16, offset: 2785},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 110, col: 16, offset: 2785},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 110, col: 21, offset: 2790},
										name: "PrimaryExpr",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 110, col: 33, offset: 2802},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 110, col: 36, offset: 2805},
									label: "op",
									expr: &ruleRefExpr{
										pos:  position{line: 110, col: 39, offset: 2808},
										name: "SuffixedOp",
									},
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 129, col: 5, offset: 3338},
						name: "PrimaryExpr",
					},
				},
			},
		},
		{
			name: "SuffixedOp",
			pos:  position{line: 131, col: 1, offset: 3352},
			expr: &actionExpr{
				pos: position{line: 131, col: 14, offset: 3367},
				run: (*parser).callonSuffixedOp1,
				expr: &choiceExpr{
					pos: position{line: 131, col: 16, offset: 3369},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 131, col: 16, offset: 3369},
							val:        "?",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 131, col: 22, offset: 3375},
							val:        "*",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 131, col: 28, offset: 3381},
							val:        "+",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "PrimaryExpr",
			pos:  position{line: 135, col: 1, offset: 3423},
			expr: &choiceExpr{
				pos: position{line: 135, col: 15, offset: 3439},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 135, col: 15, offset: 3439},
						name: "LitMatcher",
					},
					&ruleRefExpr{
						pos:  position{line: 135, col: 28, offset: 3452},
						name: "CharClassMatcher",
					},
					&ruleRefExpr{
						pos:  position{line: 135, col: 47, offset: 3471},
						name: "AnyMatcher",
					},
					&ruleRefExpr{
						pos:  position{line: 135, col: 60, offset: 3484},
						name: "RuleRefExpr",
					},
					&ruleRefExpr{
						pos:  position{line: 135, col: 74, offset: 3498},
						name: "SemanticPredExpr",
					},
					&actionExpr{
						pos: position{line: 135, col: 93, offset: 3517},
						run: (*parser).callonPrimaryExpr7,
						expr: &seqExpr{
							pos: position{line: 135, col: 93, offset: 3517},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 135, col: 93, offset: 3517},
									val:        "(",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 135, col: 97, offset: 3521},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 135, col: 100, offset: 3524},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 135, col: 105, offset: 3529},
										name: "Expression",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 135, col: 116, offset: 3540},
									name: "__",
								},
								&litMatcher{
									pos:        position{line: 135, col: 119, offset: 3543},
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
			pos:  position{line: 138, col: 1, offset: 3572},
			expr: &actionExpr{
				pos: position{line: 138, col: 15, offset: 3588},
				run: (*parser).callonRuleRefExpr1,
				expr: &seqExpr{
					pos: position{line: 138, col: 15, offset: 3588},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 138, col: 15, offset: 3588},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 138, col: 20, offset: 3593},
								name: "IdentifierName",
							},
						},
						&notExpr{
							pos: position{line: 138, col: 35, offset: 3608},
							expr: &seqExpr{
								pos: position{line: 138, col: 38, offset: 3611},
								exprs: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 138, col: 38, offset: 3611},
										name: "__",
									},
									&zeroOrOneExpr{
										pos: position{line: 138, col: 43, offset: 3616},
										expr: &seqExpr{
											pos: position{line: 138, col: 43, offset: 3616},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 138, col: 43, offset: 3616},
													name: "StringLiteral",
												},
												&ruleRefExpr{
													pos:  position{line: 138, col: 57, offset: 3630},
													name: "__",
												},
											},
										},
									},
									&ruleRefExpr{
										pos:  position{line: 138, col: 63, offset: 3636},
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
			pos:  position{line: 143, col: 1, offset: 3752},
			expr: &actionExpr{
				pos: position{line: 143, col: 20, offset: 3773},
				run: (*parser).callonSemanticPredExpr1,
				expr: &seqExpr{
					pos: position{line: 143, col: 20, offset: 3773},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 143, col: 20, offset: 3773},
							label: "op",
							expr: &ruleRefExpr{
								pos:  position{line: 143, col: 23, offset: 3776},
								name: "SemanticPredOp",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 143, col: 38, offset: 3791},
							name: "__",
						},
						&labeledExpr{
							pos:   position{line: 143, col: 41, offset: 3794},
							label: "code",
							expr: &ruleRefExpr{
								pos:  position{line: 143, col: 46, offset: 3799},
								name: "CodeBlock",
							},
						},
					},
				},
			},
		},
		{
			name: "SemanticPredOp",
			pos:  position{line: 154, col: 1, offset: 4076},
			expr: &actionExpr{
				pos: position{line: 154, col: 18, offset: 4095},
				run: (*parser).callonSemanticPredOp1,
				expr: &choiceExpr{
					pos: position{line: 154, col: 20, offset: 4097},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 154, col: 20, offset: 4097},
							val:        "&",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 154, col: 26, offset: 4103},
							val:        "!",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "RuleDefOp",
			pos:  position{line: 158, col: 1, offset: 4145},
			expr: &choiceExpr{
				pos: position{line: 158, col: 13, offset: 4159},
				alternatives: []interface{}{
					&litMatcher{
						pos:        position{line: 158, col: 13, offset: 4159},
						val:        "=",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 158, col: 19, offset: 4165},
						val:        "<-",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 158, col: 26, offset: 4172},
						val:        "←",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 158, col: 37, offset: 4183},
						val:        "⟵",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "SourceChar",
			pos:  position{line: 160, col: 1, offset: 4193},
			expr: &anyMatcher{
				line: 160, col: 14, offset: 4208,
			},
		},
		{
			name: "Comment",
			pos:  position{line: 161, col: 1, offset: 4210},
			expr: &choiceExpr{
				pos: position{line: 161, col: 11, offset: 4222},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 161, col: 11, offset: 4222},
						name: "MultiLineComment",
					},
					&ruleRefExpr{
						pos:  position{line: 161, col: 30, offset: 4241},
						name: "SingleLineComment",
					},
				},
			},
		},
		{
			name: "MultiLineComment",
			pos:  position{line: 162, col: 1, offset: 4259},
			expr: &seqExpr{
				pos: position{line: 162, col: 20, offset: 4280},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 162, col: 20, offset: 4280},
						val:        "/*",
						ignoreCase: false,
					},
					&zeroOrMoreExpr{
						pos: position{line: 162, col: 27, offset: 4287},
						expr: &seqExpr{
							pos: position{line: 162, col: 27, offset: 4287},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 162, col: 27, offset: 4287},
									expr: &litMatcher{
										pos:        position{line: 162, col: 28, offset: 4288},
										val:        "*/",
										ignoreCase: false,
									},
								},
								&ruleRefExpr{
									pos:  position{line: 162, col: 33, offset: 4293},
									name: "SourceChar",
								},
							},
						},
					},
					&litMatcher{
						pos:        position{line: 162, col: 47, offset: 4307},
						val:        "*/",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "MultiLineCommentNoLineTerminator",
			pos:  position{line: 163, col: 1, offset: 4312},
			expr: &seqExpr{
				pos: position{line: 163, col: 36, offset: 4349},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 163, col: 36, offset: 4349},
						val:        "/*",
						ignoreCase: false,
					},
					&zeroOrMoreExpr{
						pos: position{line: 163, col: 43, offset: 4356},
						expr: &seqExpr{
							pos: position{line: 163, col: 43, offset: 4356},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 163, col: 43, offset: 4356},
									expr: &choiceExpr{
										pos: position{line: 163, col: 46, offset: 4359},
										alternatives: []interface{}{
											&litMatcher{
												pos:        position{line: 163, col: 46, offset: 4359},
												val:        "*/",
												ignoreCase: false,
											},
											&ruleRefExpr{
												pos:  position{line: 163, col: 53, offset: 4366},
												name: "EOL",
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 163, col: 59, offset: 4372},
									name: "SourceChar",
								},
							},
						},
					},
					&litMatcher{
						pos:        position{line: 163, col: 73, offset: 4386},
						val:        "*/",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "SingleLineComment",
			pos:  position{line: 164, col: 1, offset: 4391},
			expr: &seqExpr{
				pos: position{line: 164, col: 21, offset: 4413},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 164, col: 21, offset: 4413},
						val:        "//",
						ignoreCase: false,
					},
					&zeroOrMoreExpr{
						pos: position{line: 164, col: 28, offset: 4420},
						expr: &seqExpr{
							pos: position{line: 164, col: 28, offset: 4420},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 164, col: 28, offset: 4420},
									expr: &ruleRefExpr{
										pos:  position{line: 164, col: 29, offset: 4421},
										name: "EOL",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 164, col: 33, offset: 4425},
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
			pos:  position{line: 166, col: 1, offset: 4440},
			expr: &ruleRefExpr{
				pos:  position{line: 166, col: 14, offset: 4455},
				name: "IdentifierName",
			},
		},
		{
			name: "IdentifierName",
			pos:  position{line: 167, col: 1, offset: 4470},
			expr: &actionExpr{
				pos: position{line: 167, col: 18, offset: 4489},
				run: (*parser).callonIdentifierName1,
				expr: &seqExpr{
					pos: position{line: 167, col: 18, offset: 4489},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 167, col: 18, offset: 4489},
							name: "IdentifierStart",
						},
						&zeroOrMoreExpr{
							pos: position{line: 167, col: 34, offset: 4505},
							expr: &ruleRefExpr{
								pos:  position{line: 167, col: 34, offset: 4505},
								name: "IdentifierPart",
							},
						},
					},
				},
			},
		},
		{
			name: "IdentifierStart",
			pos:  position{line: 170, col: 1, offset: 4587},
			expr: &charClassMatcher{
				pos:        position{line: 170, col: 19, offset: 4607},
				val:        "[a-z_]i",
				chars:      []rune{'_'},
				ranges:     []rune{'a', 'z'},
				ignoreCase: true,
				inverted:   false,
			},
		},
		{
			name: "IdentifierPart",
			pos:  position{line: 171, col: 1, offset: 4615},
			expr: &choiceExpr{
				pos: position{line: 171, col: 18, offset: 4634},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 171, col: 18, offset: 4634},
						name: "IdentifierStart",
					},
					&charClassMatcher{
						pos:        position{line: 171, col: 36, offset: 4652},
						val:        "[0-9]",
						ranges:     []rune{'0', '9'},
						ignoreCase: false,
						inverted:   false,
					},
				},
			},
		},
		{
			name: "LitMatcher",
			pos:  position{line: 173, col: 1, offset: 4659},
			expr: &actionExpr{
				pos: position{line: 173, col: 14, offset: 4674},
				run: (*parser).callonLitMatcher1,
				expr: &seqExpr{
					pos: position{line: 173, col: 14, offset: 4674},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 173, col: 14, offset: 4674},
							label: "lit",
							expr: &ruleRefExpr{
								pos:  position{line: 173, col: 18, offset: 4678},
								name: "StringLiteral",
							},
						},
						&labeledExpr{
							pos:   position{line: 173, col: 32, offset: 4692},
							label: "ignore",
							expr: &zeroOrOneExpr{
								pos: position{line: 173, col: 39, offset: 4699},
								expr: &litMatcher{
									pos:        position{line: 173, col: 39, offset: 4699},
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
			pos:  position{line: 183, col: 1, offset: 4925},
			expr: &actionExpr{
				pos: position{line: 183, col: 17, offset: 4943},
				run: (*parser).callonStringLiteral1,
				expr: &choiceExpr{
					pos: position{line: 183, col: 19, offset: 4945},
					alternatives: []interface{}{
						&seqExpr{
							pos: position{line: 183, col: 19, offset: 4945},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 183, col: 19, offset: 4945},
									val:        "\"",
									ignoreCase: false,
								},
								&zeroOrMoreExpr{
									pos: position{line: 183, col: 23, offset: 4949},
									expr: &ruleRefExpr{
										pos:  position{line: 183, col: 23, offset: 4949},
										name: "DoubleStringChar",
									},
								},
								&litMatcher{
									pos:        position{line: 183, col: 41, offset: 4967},
									val:        "\"",
									ignoreCase: false,
								},
							},
						},
						&seqExpr{
							pos: position{line: 183, col: 47, offset: 4973},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 183, col: 47, offset: 4973},
									val:        "'",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 183, col: 51, offset: 4977},
									name: "SingleStringChar",
								},
								&litMatcher{
									pos:        position{line: 183, col: 68, offset: 4994},
									val:        "'",
									ignoreCase: false,
								},
							},
						},
						&seqExpr{
							pos: position{line: 183, col: 74, offset: 5000},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 183, col: 74, offset: 5000},
									val:        "`",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 183, col: 78, offset: 5004},
									name: "RawStringChar",
								},
								&litMatcher{
									pos:        position{line: 183, col: 92, offset: 5018},
									val:        "`",
									ignoreCase: false,
								},
							},
						},
					},
				},
			},
		},
		{
			name: "DoubleStringChar",
			pos:  position{line: 186, col: 1, offset: 5089},
			expr: &choiceExpr{
				pos: position{line: 186, col: 20, offset: 5110},
				alternatives: []interface{}{
					&seqExpr{
						pos: position{line: 186, col: 20, offset: 5110},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 186, col: 20, offset: 5110},
								expr: &choiceExpr{
									pos: position{line: 186, col: 23, offset: 5113},
									alternatives: []interface{}{
										&litMatcher{
											pos:        position{line: 186, col: 23, offset: 5113},
											val:        "\"",
											ignoreCase: false,
										},
										&litMatcher{
											pos:        position{line: 186, col: 29, offset: 5119},
											val:        "\\",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 186, col: 36, offset: 5126},
											name: "EOL",
										},
									},
								},
							},
							&ruleRefExpr{
								pos:  position{line: 186, col: 42, offset: 5132},
								name: "SourceChar",
							},
						},
					},
					&seqExpr{
						pos: position{line: 186, col: 55, offset: 5145},
						exprs: []interface{}{
							&litMatcher{
								pos:        position{line: 186, col: 55, offset: 5145},
								val:        "\\",
								ignoreCase: false,
							},
							&ruleRefExpr{
								pos:  position{line: 186, col: 60, offset: 5150},
								name: "DoubleStringEscape",
							},
						},
					},
				},
			},
		},
		{
			name: "SingleStringChar",
			pos:  position{line: 187, col: 1, offset: 5169},
			expr: &choiceExpr{
				pos: position{line: 187, col: 20, offset: 5190},
				alternatives: []interface{}{
					&seqExpr{
						pos: position{line: 187, col: 20, offset: 5190},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 187, col: 20, offset: 5190},
								expr: &choiceExpr{
									pos: position{line: 187, col: 23, offset: 5193},
									alternatives: []interface{}{
										&litMatcher{
											pos:        position{line: 187, col: 23, offset: 5193},
											val:        "'",
											ignoreCase: false,
										},
										&litMatcher{
											pos:        position{line: 187, col: 29, offset: 5199},
											val:        "\\",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 187, col: 36, offset: 5206},
											name: "EOL",
										},
									},
								},
							},
							&ruleRefExpr{
								pos:  position{line: 187, col: 42, offset: 5212},
								name: "SourceChar",
							},
						},
					},
					&seqExpr{
						pos: position{line: 187, col: 55, offset: 5225},
						exprs: []interface{}{
							&litMatcher{
								pos:        position{line: 187, col: 55, offset: 5225},
								val:        "\\",
								ignoreCase: false,
							},
							&ruleRefExpr{
								pos:  position{line: 187, col: 60, offset: 5230},
								name: "SingleStringEscape",
							},
						},
					},
				},
			},
		},
		{
			name: "RawStringChar",
			pos:  position{line: 188, col: 1, offset: 5249},
			expr: &seqExpr{
				pos: position{line: 188, col: 17, offset: 5267},
				exprs: []interface{}{
					&notExpr{
						pos: position{line: 188, col: 17, offset: 5267},
						expr: &litMatcher{
							pos:        position{line: 188, col: 18, offset: 5268},
							val:        "`",
							ignoreCase: false,
						},
					},
					&ruleRefExpr{
						pos:  position{line: 188, col: 22, offset: 5272},
						name: "SourceChar",
					},
				},
			},
		},
		{
			name: "DoubleStringEscape",
			pos:  position{line: 190, col: 1, offset: 5284},
			expr: &choiceExpr{
				pos: position{line: 190, col: 22, offset: 5307},
				alternatives: []interface{}{
					&litMatcher{
						pos:        position{line: 190, col: 22, offset: 5307},
						val:        "'",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 190, col: 28, offset: 5313},
						name: "CommonEscapeSequence",
					},
				},
			},
		},
		{
			name: "SingleStringEscape",
			pos:  position{line: 191, col: 1, offset: 5334},
			expr: &choiceExpr{
				pos: position{line: 191, col: 22, offset: 5357},
				alternatives: []interface{}{
					&litMatcher{
						pos:        position{line: 191, col: 22, offset: 5357},
						val:        "\"",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 191, col: 28, offset: 5363},
						name: "CommonEscapeSequence",
					},
				},
			},
		},
		{
			name: "CommonEscapeSequence",
			pos:  position{line: 193, col: 1, offset: 5385},
			expr: &choiceExpr{
				pos: position{line: 193, col: 24, offset: 5410},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 193, col: 24, offset: 5410},
						name: "SingleCharEscape",
					},
					&ruleRefExpr{
						pos:  position{line: 193, col: 43, offset: 5429},
						name: "OctalEscape",
					},
					&ruleRefExpr{
						pos:  position{line: 193, col: 57, offset: 5443},
						name: "HexEscape",
					},
					&ruleRefExpr{
						pos:  position{line: 193, col: 69, offset: 5455},
						name: "LongUnicodeEscape",
					},
					&ruleRefExpr{
						pos:  position{line: 193, col: 89, offset: 5475},
						name: "ShortUnicodeEscape",
					},
				},
			},
		},
		{
			name: "SingleCharEscape",
			pos:  position{line: 194, col: 1, offset: 5494},
			expr: &choiceExpr{
				pos: position{line: 194, col: 20, offset: 5515},
				alternatives: []interface{}{
					&litMatcher{
						pos:        position{line: 194, col: 20, offset: 5515},
						val:        "a",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 194, col: 26, offset: 5521},
						val:        "b",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 194, col: 32, offset: 5527},
						val:        "n",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 194, col: 38, offset: 5533},
						val:        "f",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 194, col: 44, offset: 5539},
						val:        "r",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 194, col: 50, offset: 5545},
						val:        "t",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 194, col: 56, offset: 5551},
						val:        "v",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 194, col: 62, offset: 5557},
						val:        "\\",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "OctalEscape",
			pos:  position{line: 195, col: 1, offset: 5562},
			expr: &seqExpr{
				pos: position{line: 195, col: 15, offset: 5578},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 195, col: 15, offset: 5578},
						name: "OctalDigit",
					},
					&ruleRefExpr{
						pos:  position{line: 195, col: 26, offset: 5589},
						name: "OctalDigit",
					},
					&ruleRefExpr{
						pos:  position{line: 195, col: 37, offset: 5600},
						name: "OctalDigit",
					},
				},
			},
		},
		{
			name: "HexEscape",
			pos:  position{line: 196, col: 1, offset: 5611},
			expr: &seqExpr{
				pos: position{line: 196, col: 13, offset: 5625},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 196, col: 13, offset: 5625},
						val:        "x",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 196, col: 17, offset: 5629},
						name: "HexDigit",
					},
					&ruleRefExpr{
						pos:  position{line: 196, col: 26, offset: 5638},
						name: "HexDigit",
					},
				},
			},
		},
		{
			name: "LongUnicodeEscape",
			pos:  position{line: 197, col: 1, offset: 5647},
			expr: &seqExpr{
				pos: position{line: 197, col: 21, offset: 5669},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 197, col: 21, offset: 5669},
						val:        "U",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 197, col: 25, offset: 5673},
						name: "HexDigit",
					},
					&ruleRefExpr{
						pos:  position{line: 197, col: 34, offset: 5682},
						name: "HexDigit",
					},
					&ruleRefExpr{
						pos:  position{line: 197, col: 43, offset: 5691},
						name: "HexDigit",
					},
					&ruleRefExpr{
						pos:  position{line: 197, col: 52, offset: 5700},
						name: "HexDigit",
					},
					&ruleRefExpr{
						pos:  position{line: 197, col: 61, offset: 5709},
						name: "HexDigit",
					},
					&ruleRefExpr{
						pos:  position{line: 197, col: 70, offset: 5718},
						name: "HexDigit",
					},
					&ruleRefExpr{
						pos:  position{line: 197, col: 79, offset: 5727},
						name: "HexDigit",
					},
					&ruleRefExpr{
						pos:  position{line: 197, col: 88, offset: 5736},
						name: "HexDigit",
					},
				},
			},
		},
		{
			name: "ShortUnicodeEscape",
			pos:  position{line: 198, col: 1, offset: 5745},
			expr: &seqExpr{
				pos: position{line: 198, col: 22, offset: 5768},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 198, col: 22, offset: 5768},
						val:        "u",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 198, col: 26, offset: 5772},
						name: "HexDigit",
					},
					&ruleRefExpr{
						pos:  position{line: 198, col: 35, offset: 5781},
						name: "HexDigit",
					},
					&ruleRefExpr{
						pos:  position{line: 198, col: 44, offset: 5790},
						name: "HexDigit",
					},
					&ruleRefExpr{
						pos:  position{line: 198, col: 53, offset: 5799},
						name: "HexDigit",
					},
				},
			},
		},
		{
			name: "OctalDigit",
			pos:  position{line: 200, col: 1, offset: 5809},
			expr: &charClassMatcher{
				pos:        position{line: 200, col: 14, offset: 5824},
				val:        "[0-7]",
				ranges:     []rune{'0', '7'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "DecimalDigit",
			pos:  position{line: 201, col: 1, offset: 5830},
			expr: &charClassMatcher{
				pos:        position{line: 201, col: 16, offset: 5847},
				val:        "[0-9]",
				ranges:     []rune{'0', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "HexDigit",
			pos:  position{line: 202, col: 1, offset: 5853},
			expr: &charClassMatcher{
				pos:        position{line: 202, col: 12, offset: 5866},
				val:        "[0-9a-f]i",
				ranges:     []rune{'0', '9', 'a', 'f'},
				ignoreCase: true,
				inverted:   false,
			},
		},
		{
			name: "CharClassMatcher",
			pos:  position{line: 204, col: 1, offset: 5877},
			expr: &actionExpr{
				pos: position{line: 204, col: 20, offset: 5898},
				run: (*parser).callonCharClassMatcher1,
				expr: &seqExpr{
					pos: position{line: 204, col: 20, offset: 5898},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 204, col: 20, offset: 5898},
							val:        "[",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 204, col: 26, offset: 5904},
							expr: &choiceExpr{
								pos: position{line: 204, col: 26, offset: 5904},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 204, col: 26, offset: 5904},
										name: "ClassCharRange",
									},
									&ruleRefExpr{
										pos:  position{line: 204, col: 43, offset: 5921},
										name: "ClassChar",
									},
									&seqExpr{
										pos: position{line: 204, col: 55, offset: 5933},
										exprs: []interface{}{
											&litMatcher{
												pos:        position{line: 204, col: 55, offset: 5933},
												val:        "\\",
												ignoreCase: false,
											},
											&ruleRefExpr{
												pos:  position{line: 204, col: 60, offset: 5938},
												name: "UnicodeClassEscape",
											},
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 204, col: 82, offset: 5960},
							val:        "]",
							ignoreCase: false,
						},
						&zeroOrOneExpr{
							pos: position{line: 204, col: 86, offset: 5964},
							expr: &litMatcher{
								pos:        position{line: 204, col: 86, offset: 5964},
								val:        "i",
								ignoreCase: false,
							},
						},
					},
				},
			},
		},
		{
			name: "ClassCharRange",
			pos:  position{line: 209, col: 1, offset: 6069},
			expr: &seqExpr{
				pos: position{line: 209, col: 18, offset: 6088},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 209, col: 18, offset: 6088},
						name: "ClassChar",
					},
					&litMatcher{
						pos:        position{line: 209, col: 28, offset: 6098},
						val:        "-",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 209, col: 32, offset: 6102},
						name: "ClassChar",
					},
				},
			},
		},
		{
			name: "ClassChar",
			pos:  position{line: 210, col: 1, offset: 6112},
			expr: &choiceExpr{
				pos: position{line: 210, col: 13, offset: 6126},
				alternatives: []interface{}{
					&seqExpr{
						pos: position{line: 210, col: 13, offset: 6126},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 210, col: 13, offset: 6126},
								expr: &choiceExpr{
									pos: position{line: 210, col: 16, offset: 6129},
									alternatives: []interface{}{
										&litMatcher{
											pos:        position{line: 210, col: 16, offset: 6129},
											val:        "]",
											ignoreCase: false,
										},
										&litMatcher{
											pos:        position{line: 210, col: 22, offset: 6135},
											val:        "\\",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 210, col: 29, offset: 6142},
											name: "EOL",
										},
									},
								},
							},
							&ruleRefExpr{
								pos:  position{line: 210, col: 35, offset: 6148},
								name: "SourceChar",
							},
						},
					},
					&seqExpr{
						pos: position{line: 210, col: 48, offset: 6161},
						exprs: []interface{}{
							&litMatcher{
								pos:        position{line: 210, col: 48, offset: 6161},
								val:        "\\",
								ignoreCase: false,
							},
							&ruleRefExpr{
								pos:  position{line: 210, col: 53, offset: 6166},
								name: "CharClassEscape",
							},
						},
					},
				},
			},
		},
		{
			name: "CharClassEscape",
			pos:  position{line: 211, col: 1, offset: 6182},
			expr: &choiceExpr{
				pos: position{line: 211, col: 19, offset: 6202},
				alternatives: []interface{}{
					&litMatcher{
						pos:        position{line: 211, col: 19, offset: 6202},
						val:        "]",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 211, col: 25, offset: 6208},
						name: "CommonEscapeSequence",
					},
				},
			},
		},
		{
			name: "UnicodeClassEscape",
			pos:  position{line: 213, col: 1, offset: 6230},
			expr: &seqExpr{
				pos: position{line: 213, col: 22, offset: 6253},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 213, col: 22, offset: 6253},
						val:        "p",
						ignoreCase: false,
					},
					&choiceExpr{
						pos: position{line: 213, col: 28, offset: 6259},
						alternatives: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 213, col: 28, offset: 6259},
								name: "SingleCharUnicodeClass",
							},
							&seqExpr{
								pos: position{line: 213, col: 53, offset: 6284},
								exprs: []interface{}{
									&litMatcher{
										pos:        position{line: 213, col: 53, offset: 6284},
										val:        "{",
										ignoreCase: false,
									},
									&ruleRefExpr{
										pos:  position{line: 213, col: 57, offset: 6288},
										name: "UnicodeClass",
									},
									&litMatcher{
										pos:        position{line: 213, col: 70, offset: 6301},
										val:        "}",
										ignoreCase: false,
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
			pos:  position{line: 214, col: 1, offset: 6307},
			expr: &charClassMatcher{
				pos:        position{line: 214, col: 26, offset: 6334},
				val:        "[LMNCPZS]",
				chars:      []rune{'L', 'M', 'N', 'C', 'P', 'Z', 'S'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "UnicodeClass",
			pos:  position{line: 215, col: 1, offset: 6344},
			expr: &oneOrMoreExpr{
				pos: position{line: 215, col: 16, offset: 6361},
				expr: &charClassMatcher{
					pos:        position{line: 215, col: 16, offset: 6361},
					val:        "[a-z_]i",
					chars:      []rune{'_'},
					ranges:     []rune{'a', 'z'},
					ignoreCase: true,
					inverted:   false,
				},
			},
		},
		{
			name: "AnyMatcher",
			pos:  position{line: 217, col: 1, offset: 6371},
			expr: &actionExpr{
				pos: position{line: 217, col: 14, offset: 6386},
				run: (*parser).callonAnyMatcher1,
				expr: &litMatcher{
					pos:        position{line: 217, col: 14, offset: 6386},
					val:        ".",
					ignoreCase: false,
				},
			},
		},
		{
			name: "CodeBlock",
			pos:  position{line: 222, col: 1, offset: 6461},
			expr: &actionExpr{
				pos: position{line: 222, col: 13, offset: 6475},
				run: (*parser).callonCodeBlock1,
				expr: &seqExpr{
					pos: position{line: 222, col: 13, offset: 6475},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 222, col: 13, offset: 6475},
							val:        "{",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 222, col: 17, offset: 6479},
							name: "Code",
						},
						&litMatcher{
							pos:        position{line: 222, col: 22, offset: 6484},
							val:        "}",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "Code",
			pos:  position{line: 228, col: 1, offset: 6582},
			expr: &zeroOrMoreExpr{
				pos: position{line: 228, col: 10, offset: 6593},
				expr: &choiceExpr{
					pos: position{line: 228, col: 10, offset: 6593},
					alternatives: []interface{}{
						&oneOrMoreExpr{
							pos: position{line: 228, col: 12, offset: 6595},
							expr: &seqExpr{
								pos: position{line: 228, col: 12, offset: 6595},
								exprs: []interface{}{
									&notExpr{
										pos: position{line: 228, col: 12, offset: 6595},
										expr: &charClassMatcher{
											pos:        position{line: 228, col: 13, offset: 6596},
											val:        "[{}]",
											chars:      []rune{'{', '}'},
											ignoreCase: false,
											inverted:   false,
										},
									},
									&ruleRefExpr{
										pos:  position{line: 228, col: 18, offset: 6601},
										name: "SourceChar",
									},
								},
							},
						},
						&seqExpr{
							pos: position{line: 228, col: 34, offset: 6617},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 228, col: 34, offset: 6617},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 228, col: 38, offset: 6621},
									name: "Code",
								},
								&litMatcher{
									pos:        position{line: 228, col: 43, offset: 6626},
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
			pos:  position{line: 230, col: 1, offset: 6634},
			expr: &zeroOrMoreExpr{
				pos: position{line: 230, col: 8, offset: 6643},
				expr: &choiceExpr{
					pos: position{line: 230, col: 8, offset: 6643},
					alternatives: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 230, col: 8, offset: 6643},
							name: "Whitespace",
						},
						&ruleRefExpr{
							pos:  position{line: 230, col: 21, offset: 6656},
							name: "EOL",
						},
						&ruleRefExpr{
							pos:  position{line: 230, col: 27, offset: 6662},
							name: "Comment",
						},
					},
				},
			},
		},
		{
			name: "_",
			pos:  position{line: 231, col: 1, offset: 6673},
			expr: &zeroOrMoreExpr{
				pos: position{line: 231, col: 7, offset: 6681},
				expr: &choiceExpr{
					pos: position{line: 231, col: 7, offset: 6681},
					alternatives: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 231, col: 7, offset: 6681},
							name: "Whitespace",
						},
						&ruleRefExpr{
							pos:  position{line: 231, col: 20, offset: 6694},
							name: "MultiLineCommentNoLineTerminator",
						},
					},
				},
			},
		},
		{
			name: "Whitespace",
			pos:  position{line: 233, col: 1, offset: 6731},
			expr: &charClassMatcher{
				pos:        position{line: 233, col: 14, offset: 6746},
				val:        "[ \\t\\r]",
				chars:      []rune{' ', '\t', '\r'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "EOL",
			pos:  position{line: 234, col: 1, offset: 6754},
			expr: &litMatcher{
				pos:        position{line: 234, col: 7, offset: 6762},
				val:        "\n",
				ignoreCase: false,
			},
		},
		{
			name: "EOS",
			pos:  position{line: 235, col: 1, offset: 6767},
			expr: &choiceExpr{
				pos: position{line: 235, col: 7, offset: 6775},
				alternatives: []interface{}{
					&seqExpr{
						pos: position{line: 235, col: 7, offset: 6775},
						exprs: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 235, col: 7, offset: 6775},
								name: "__",
							},
							&litMatcher{
								pos:        position{line: 235, col: 10, offset: 6778},
								val:        ";",
								ignoreCase: false,
							},
						},
					},
					&seqExpr{
						pos: position{line: 235, col: 16, offset: 6784},
						exprs: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 235, col: 16, offset: 6784},
								name: "_",
							},
							&zeroOrOneExpr{
								pos: position{line: 235, col: 18, offset: 6786},
								expr: &ruleRefExpr{
									pos:  position{line: 235, col: 18, offset: 6786},
									name: "SingleLineComment",
								},
							},
							&ruleRefExpr{
								pos:  position{line: 235, col: 37, offset: 6805},
								name: "EOL",
							},
						},
					},
					&seqExpr{
						pos: position{line: 235, col: 43, offset: 6811},
						exprs: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 235, col: 43, offset: 6811},
								name: "__",
							},
							&ruleRefExpr{
								pos:  position{line: 235, col: 46, offset: 6814},
								name: "EOF",
							},
						},
					},
				},
			},
		},
		{
			name: "EOF",
			pos:  position{line: 237, col: 1, offset: 6819},
			expr: &notExpr{
				pos: position{line: 237, col: 7, offset: 6827},
				expr: &anyMatcher{
					line: 237, col: 8, offset: 6828,
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
		return nil, err
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

func (c *current) onStringLiteral1() (interface{}, error) {
	return ast.NewStringLit(c.astPos(), string(c.text)), nil
}

func (p *parser) callonStringLiteral1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onStringLiteral1()
}

func (c *current) onCharClassMatcher1() (interface{}, error) {
	pos := c.astPos()
	cc := ast.NewCharClassMatcher(pos, string(c.text))
	return cc, nil
}

func (p *parser) callonCharClassMatcher1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onCharClassMatcher1()
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

func (c *current) onCodeBlock1() (interface{}, error) {
	pos := c.astPos()
	cb := ast.NewCodeBlock(pos, string(c.text))
	return cb, nil
}

func (p *parser) callonCodeBlock1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onCodeBlock1()
}

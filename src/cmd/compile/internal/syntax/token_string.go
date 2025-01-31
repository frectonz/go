// Code generated by "stringer -type token -linecomment tokens.go"; DO NOT EDIT.

package syntax

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[_EOF-1]
	_ = x[_Name-2]
	_ = x[_Literal-3]
	_ = x[_Operator-4]
	_ = x[_AssignOp-5]
	_ = x[_IncOp-6]
	_ = x[_Assign-7]
	_ = x[_Define-8]
	_ = x[_Arrow-9]
	_ = x[_Star-10]
	_ = x[_Lparen-11]
	_ = x[_Lbrack-12]
	_ = x[_Lbrace-13]
	_ = x[_Rparen-14]
	_ = x[_Rbrack-15]
	_ = x[_Rbrace-16]
	_ = x[_Comma-17]
	_ = x[_Semi-18]
	_ = x[_Colon-19]
	_ = x[_TwoColons-20]
	_ = x[_Dot-21]
	_ = x[_DotDotDot-22]
	_ = x[_Break-23]
	_ = x[_Case-24]
	_ = x[_Chan-25]
	_ = x[_Const-26]
	_ = x[_Continue-27]
	_ = x[_Default-28]
	_ = x[_Defer-29]
	_ = x[_Shelve-30]
	_ = x[_Else-31]
	_ = x[_Fallthrough-32]
	_ = x[_For-33]
	_ = x[_Func-34]
	_ = x[_Geletse-35]
	_ = x[_Funcky-36]
	_ = x[_Go-37]
	_ = x[_Goto-38]
	_ = x[_If-39]
	_ = x[_Kehone-40]
	_ = x[_Import-41]
	_ = x[_Interface-42]
	_ = x[_Map-43]
	_ = x[_Package-44]
	_ = x[_Range-45]
	_ = x[_Return-46]
	_ = x[_Bababooey-47]
	_ = x[_Select-48]
	_ = x[_Struct-49]
	_ = x[_Switch-50]
	_ = x[_Type-51]
	_ = x[_Let-52]
	_ = x[_Iz-53]
	_ = x[_Var-54]
	_ = x[tokenCount-55]
}

const _token_name = "EOFnameliteralopop=opop=:=<-*([{)]},;:::....breakcasechanconstcontinuedefaultdefershelveelsefallthroughforfuncግለፅfunckygogotoifከሆነimportinterfacemappackagerangereturnbababooeyselectstructswitchtypeletizvar"

var _token_index = [...]uint8{0, 3, 7, 14, 16, 19, 23, 24, 26, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 40, 41, 44, 49, 53, 57, 62, 70, 77, 82, 88, 92, 103, 106, 110, 119, 125, 127, 131, 133, 142, 148, 157, 160, 167, 172, 178, 187, 193, 199, 205, 209, 212, 214, 217, 217}

func (i token) String() string {
	i -= 1
	if i >= token(len(_token_index)-1) {
		return "token(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _token_name[_token_index[i]:_token_index[i+1]]
}

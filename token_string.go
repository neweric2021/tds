// Code generated by "stringer -type=token"; DO NOT EDIT.

package tds

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[noneToken-0]
	_ = x[capabilityReqToken-1]
	_ = x[capabilityResToken-2]
	_ = x[paramFmt2Token-32]
	_ = x[languageToken-33]
	_ = x[orderBy2Token-34]
	_ = x[wideColumnFmtToken-97]
	_ = x[dynamic2Token-98]
	_ = x[msgToken-101]
	_ = x[returnStatusToken-121]
	_ = x[curCloseToken-128]
	_ = x[curDeleteToken-129]
	_ = x[curFetchToken-130]
	_ = x[curFmtToken-131]
	_ = x[curOpenToken-132]
	_ = x[curDeclareToken-134]
	_ = x[logoutToken-113]
	_ = x[tableNameToken-164]
	_ = x[columnInfoToken-165]
	_ = x[optionCmdToken-166]
	_ = x[cmpRowNameToken-167]
	_ = x[cmpRowFmtToken-168]
	_ = x[orderByToken-169]
	_ = x[infoToken-171]
	_ = x[loginAckToken-173]
	_ = x[controlToken-174]
	_ = x[rowToken-209]
	_ = x[cmpRowToken-211]
	_ = x[paramToken-215]
	_ = x[capabilitiesToken-226]
	_ = x[envChangeToken-227]
	_ = x[sqlMessageToken-229]
	_ = x[dbRPCToken-230]
	_ = x[dynamicToken-231]
	_ = x[paramFmtToken-236]
	_ = x[authToken-237]
	_ = x[columnFmtToken-238]
	_ = x[doneToken-253]
	_ = x[doneProcToken-254]
	_ = x[doneInProcToken-255]
}

const _token_name = "noneTokencapabilityReqTokencapabilityResTokenparamFmt2TokenlanguageTokenorderBy2TokenwideColumnFmtTokendynamic2TokenmsgTokenlogoutTokenreturnStatusTokencurCloseTokencurDeleteTokencurFetchTokencurFmtTokencurOpenTokencurDeclareTokentableNameTokencolumnInfoTokenoptionCmdTokencmpRowNameTokencmpRowFmtTokenorderByTokeninfoTokenloginAckTokencontrolTokenrowTokencmpRowTokenparamTokencapabilitiesTokenenvChangeTokensqlMessageTokendbRPCTokendynamicTokenparamFmtTokenauthTokencolumnFmtTokendoneTokendoneProcTokendoneInProcToken"

var _token_map = map[token]string{
	0:   _token_name[0:9],
	1:   _token_name[9:27],
	2:   _token_name[27:45],
	32:  _token_name[45:59],
	33:  _token_name[59:72],
	34:  _token_name[72:85],
	97:  _token_name[85:103],
	98:  _token_name[103:116],
	101: _token_name[116:124],
	113: _token_name[124:135],
	121: _token_name[135:152],
	128: _token_name[152:165],
	129: _token_name[165:179],
	130: _token_name[179:192],
	131: _token_name[192:203],
	132: _token_name[203:215],
	134: _token_name[215:230],
	164: _token_name[230:244],
	165: _token_name[244:259],
	166: _token_name[259:273],
	167: _token_name[273:288],
	168: _token_name[288:302],
	169: _token_name[302:314],
	171: _token_name[314:323],
	173: _token_name[323:336],
	174: _token_name[336:348],
	209: _token_name[348:356],
	211: _token_name[356:367],
	215: _token_name[367:377],
	226: _token_name[377:394],
	227: _token_name[394:408],
	229: _token_name[408:423],
	230: _token_name[423:433],
	231: _token_name[433:445],
	236: _token_name[445:458],
	237: _token_name[458:467],
	238: _token_name[467:481],
	253: _token_name[481:490],
	254: _token_name[490:503],
	255: _token_name[503:518],
}

func (i token) String() string {
	if str, ok := _token_map[i]; ok {
		return str
	}
	return "token(" + strconv.FormatInt(int64(i), 10) + ")"
}

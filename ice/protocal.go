package ice

//报文版本
type ProtocolVersion struct {
	Major byte
	Minor byte
}

//
type ProtocalEncoding struct {
	Major byte
	Minor byte
}

//编码版本
type EncodingVersion struct {
	Major byte
	Minor byte
}

var _pversion = &ProtocolVersion{
	Major: 1,
	Minor: 0,
}
var _peversion = &ProtocalEncoding{
	Major: 1,
	Minor: 0,
}
var _eversion = &EncodingVersion{
	Major: 1,
	Minor: 1,
}

func GetDefaultProtocolVersion() *ProtocolVersion {
	return _pversion
}

func GetDefaultProtocalEncodingVersion() *ProtocalEncoding {
	return _peversion
}

func GetDefaultEncodingVersion() *EncodingVersion {
	return _eversion
}

package ice

var magic = []byte{0x49, 0x63, 0x65, 0x50}
var msgType byte = 0
var requestHead = []byte{magic[0],
	magic[1],
	magic[2],
	magic[3],
	GetDefaultProtocolVersion().Major,
	GetDefaultProtocolVersion().Minor,
	GetDefaultProtocalEncodingVersion().Major,
	GetDefaultProtocalEncodingVersion().Minor,
	msgType,
	0} // Compression status. 0：不支持压缩，1：支持压缩，2：已经压缩
//还要设置压缩位，见requestHdr里面关于压缩位设置。如果压缩的话，第10位设置为2，并且11 12 13 14设置为压缩后的长度。

var connHead = []byte{magic[0],
	magic[1],
	magic[2],
	magic[3],
	GetDefaultProtocolVersion().Major,
	GetDefaultProtocolVersion().Minor,
	GetDefaultProtocalEncodingVersion().Major,
	GetDefaultProtocalEncodingVersion().Minor,
	msgType,
	0}

func GetHead() *[]byte {
	return &requestHead
}
func GetConnHead() *[]byte {
	return &connHead
}

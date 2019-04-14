package IceInternal

import "sync"

type ProtocalEncoding struct {
	Major byte
	Minor byte
}
var _peversion *ProtocalEncoding
var _peonce sync.Once = sync.Once{}
func GetDefaultProtocalEncodingVersion() *ProtocalEncoding{
	_peonce.Do(func() {
		_peversion = &ProtocalEncoding{
			Major:1,
			Minor:0,
		}
	})
	return _peversion
}

package ice

import "sync"

//编码版本
type EncodingVersion struct {
	Major byte
	Minor byte
}
var _eversion *EncodingVersion
var _eonce sync.Once = sync.Once{}
func GetDefaultEncodingVersion() *EncodingVersion{
	_eonce.Do(func() {
		_eversion = &EncodingVersion{
			Major:1,
			Minor:1,
		}
	})
	return _eversion
}
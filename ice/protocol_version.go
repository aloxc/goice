package ice

import "sync"

type ProtocolVersion struct {
	Major byte
	Minor byte
}
var _pversion *ProtocolVersion
var _ponce sync.Once = sync.Once{}
func GetDefaultProtocolVersion() *ProtocolVersion{
	_ponce.Do(func() {
		_pversion = &ProtocolVersion{
			Major:1,
			Minor:0,
		}
	})
	return _pversion
}
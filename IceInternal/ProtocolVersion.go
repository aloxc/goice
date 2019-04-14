package IceInternal
type ProtocolVersion struct {
	Major byte
	Minor byte
}
func GetDefaultProtocolVersion() *ProtocolVersion{
	return &ProtocolVersion{
		Major:1,
		Minor:1,
	}
}
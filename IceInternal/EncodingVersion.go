package IceInternal

//编码版本
type EncodingVersion struct {
	Major byte
	Minor byte
}
func GetDefaultEncodingVersion() *EncodingVersion{
	return &EncodingVersion{
		Major:1,
		Minor:1,
	}
}
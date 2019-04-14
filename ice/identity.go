package ice

type Identity struct {
	Name string
	Category string
}

var identityMap = make(map[string]*Identity)

func GetIdentity(name string) *Identity {
	if v,ok := identityMap[name];ok{
		return v
	}
	return nil
}
func NewIdentity(name,category string) *Identity {
	if v,ok := identityMap[name];ok{
		return v
	}
	identity := &Identity{
		Name:name,
		Category:category,
	}
	identityMap[name] = identity
	return identity
}
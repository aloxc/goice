package ice

type Identity struct {
	name     string
	category string
}

var identityMap = make(map[string]*Identity)

func GetIdentity(name, category string) *Identity {
	if v, ok := identityMap[name]; ok {
		return v
	}
	identity := &Identity{
		name:     name,
		category: category,
	}
	identityMap[name] = identity
	return identity
}
func (this *Identity) GetIdentityName() string {
	return this.name
}

func (this *Identity) GetIdentityCategory() string {
	return this.category
}

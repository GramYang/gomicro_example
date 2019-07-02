package config

type Profiles interface {
	GetInclude() string
}

type defaultProfiles struct {
	Include string
}

func (p defaultProfiles) GetInclude() string {
	return p.Include
}

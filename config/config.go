package config

type Configuration interface {
	GetString(name string) (configureValue string, found bool)
	GetInt(name string) (configureValue int, found bool)
	GetBool(name string) (configureValue bool, found bool)
	GetFloat(name string) (configureValue float64, found bool)

	GetStringDefault(name, defVal string) (configValue string)
	GetIntDefault(name string, defVal int) (configValue int)
	GetBoolDefault(name string, defVal bool) (configValue bool)
	GetFloatDefault(name string, defVal float64) (configValue float64)

	GetSection(sectionName string) (section Configuration, floud bool)
}

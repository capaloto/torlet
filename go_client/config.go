package tor

const (
	DefaultSocksAddress = "127.0.0.1:9050"

	DefaultControlAddress = "127.0.0.1:9051"

	DefaultControlPassword = "suqamadiq"

	DefaultVerbose = false
)

type Config struct {
	SocksAddress    string `toml:"socks-address"`
	ControlAddress  string `toml:"control-address"`
	ControlPassword string `toml:"control-password"`
	Verbose         bool   `toml:"verbose"`
}

func NewConfig() *Config {
	return &Config{
		SocksAddress:    DefaultSocksAddress,
		ControlAddress:  DefaultControlAddress,
		ControlPassword: DefaultControlPassword,
	}
}

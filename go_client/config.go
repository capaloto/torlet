package tor

const (
	DefaultSocksAddress = "127.0.0.1:9050"

	DefaultControlAddress = "127.0.0.1:9051"

	DefaultControlPassword = "suqamadiq"
)

type Config struct {
	SocksAddress    string `toml:"socks-address"`
	ControlAddress  string `toml:"control-address"`
	ControlPassword string `toml:"control-password"`
}

func NewConfig() *Config {
	return &Config{
		SocksAddress:    DefaultSocksAddress,
		ControlAddress:  DefaultControlAddress,
		ControlPassword: DefaultControlPassword,
	}
}

package shared

import (
	"log"

	"github.com/spf13/viper"
)

const ActiveStatus = "A"
const InactiveStatus = "IA"

type api struct {
	Port       string `mapstructure:"port"`
	SystemCode string `mapstructure:"system_code"`
}

type pgConn struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Name     string `mapstructure:"name"`
	Schema   string `mapstructure:"schema"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
}

type auth struct {
	OtpExpireMinute uint `mapstructure:"otp_expire_minute"`
}

type jwt struct {
	PrivateKey string `mapstructure:"private_key"`
	PublicKey  string `mapstructure:"public_key"`
}

type busApi struct {
	Url string `mapstructure:"api_url"`
}

type messagePro struct {
	Url   string `mapstructure:"api_url"`
	Key   string `mapstructure:"api_key"`
	Phone string `mapstructure:"phone"`
}

type qPay struct {
	Url          string `mapstructure:"url"`
	ClientId     string `mapstructure:"client_id"`
	ClientSecret string `mapstructure:"client_secret"`
	InvoiceCode  string `mapstructure:"invoice_code"`
	CallbackUrl  string `mapstructure:"callback_url"`
}

type rabbitmq struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
}

type chimegeApi struct {
	Url   string `mapstructure:"url"`
	Token string `mapstructure:"token"`
}

type aptamsApi struct {
	Url   string `mapstructure:"url"`
	Token string `mapstructure:"token"`
}

type ecommerce struct {
	Key         string `mapstructure:"key"`
	Terminal    string `mapstructure:"terminal"`
	PosTerminal string `mapstructure:"posterminal"`
	Merchant    string `mapstructure:"merchant"`
	Token       string `mapstructure:"token"`
	Url         string `mapstructure:"url"`
}

type config struct {
	DB         pgConn     `mapstructure:"database"`
	Api        api        `mapstructure:"api"`
	Jwt        jwt        `mapstructure:"jwt"`
	BusApi     busApi     `mapstructure:"busapi"`
	MessagePro messagePro `mapstructure:"messagepro"`
	Auth       auth       `mapstructure:"auth"`
	Qpay       qPay       `mapstructure:"qpay"`
	RabbitMQ   rabbitmq   `mapstructure:"rabbitmq"`
	ChimegeApi chimegeApi `mapstructure:"chimegeapi"`
	AptamsApi  aptamsApi  `mapstructure:"aptamsapi"`
	Ecommerce  ecommerce  `mapstructure:"ecommerce"`
}

var Config *config = &config{}

func LoadConfig() {
	viper.SetConfigFile("config.toml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Could not read config")
	}

	if err := viper.Unmarshal(Config); err != nil {
		log.Fatal("Could not load config")
	}
}

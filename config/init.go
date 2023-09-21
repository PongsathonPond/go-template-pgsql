package config

import (
	"fmt"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type Config struct {
	// Gin config
	GinMode string `env:"GIN_MODE,required"`
	// App config
	AppEnvironment    string `env:"APP_ENV" envDefault:"local"`
	AppName           string `env:"APP_NAME" envDefault:"service"`
	AppTimeZone       string `env:"APP_TIMEZONE" envDefault:"Asia/Bangkok"`
	AppMaxAllowed     int    `env:"APP_MAX_ALLOWED" envDefault:"500"`
	AppLogLevel       int    `env:"APP_LOG_LEVEL" envDefault:"0"`
	AppLogEnableColor bool   `env:"APP_LOG_ENABLE_COLOR" envDefault:"true"`
	AppPort           string `env:"PORT" envDefault:"8080"`

	// gRPC config
	GRPCServerHost string `env:"GRPC_SERVER_HOST" envDefault:":50051"`
	GRPCAuthorHost string `env:"GRPC_AUTHOR_HOST" envDefault:":50052"`
	GRPCMediaHost  string `env:"GRPC_MEDIA_HOST" envDefault:":50054"`

	// Jaeger config
	JaegerAgentHost string `env:"JAEGER_AGENT_HOST" envDefault:"localhost"`
	JaegerAgentPort string `env:"JAEGER_AGENT_PORT" envDefault:"6831"`
	JaegerLogType   string `env:"JAEGER_LOG_TYPE" envDefault:"stdLogger"`

	// Swagger config
	SwaggerInfoHost     string `env:"SWAGGER_INFO_HOST" envDefault:"localhost:8080"`
	SwaggerInfoBasePath string `env:"SWAGGER_INFO_BASE_PATH" envDefault:"/api/v1"`

	// MongoDB config
	MongoDBEndpoint string `env:"MONGODB_ENDPOINT" envDefault:"mongodb://localhost:27017"`
	MongoDBName     string `env:"MONGODB_NAME" envDefault:"service"`
	MongoDBTimeout  int    `env:"MONGODB_TIMEOUT" envDefault:"30"`

	// Redis config
	RedisEndpoint string `env:"REDIS_ENDPOINT" envDefault:"localhost:6379"`
	RedisPassword string `env:"REDIS_PASSWORD" envDefault:""`
	RedisTTL      int64  `env:"REDIS_TTL" envDefault:"480"` // minutes

	// Token config
	AccessTokenSigningKey  string `env:"ACCESS_TOKEN_SIGNING_KEY,required"`
	RefreshTokenSigningKey string `env:"REFRESH_TOKEN_SIGNING_KEY,required"`
	AccessTokenTTL         int    `env:"ACCESS_TOKEN_TTL" envDefault:"30"`   // minutes
	RefreshTokenTTL        int    `env:"REFRESH_TOKEN_TTL" envDefault:"60"`  // minutes
	ActivateTokenTTL       int    `env:"ACTIVATE_TOKEN_TTL" envDefault:"72"` // hours

	// MessageBroker config
	MessageBrokerBackOffTime  int      `env:"MESSAGE_BROKER_BACKOFF_TIME" envDefault:"2"`
	MessageBrokerMaximumRetry int      `env:"MESSAGE_BROKER_MAXIMUM_RETRY" envDefault:"3"`
	MessageBrokerEndpoint     []string `env:"MESSAGE_BROKER_ENDPOINT" envDefault:"localhost:9094"`
	MessageBrokerGroup        string   `env:"MESSAGE_BROKER_GROUP" envDefault:"my-group"`
	MessageBrokerVersion      string   `env:"MESSAGE_BROKER_VERSION" envDefault:"2.6.1"`
	MessageBrokerDebug        bool     `env:"MESSAGE_BROKER_DEBUG" envDefault:"true"`

	// Invitation config
	ActivationDomain string `env:"ACTIVATION_DOMAIN" envDefault:"https://cms.idev.com"`

	//Pgsql config
	Host     string `env:"HOST_PGSQL" envDefault:"localhost" `
	Port     int    `env:"PORT_PGSQL" envDefault:"5433" `
	User     string `env:"USER_PGSQL" envDefault:"adminidev" `
	Password string `env:"PASSWORD_PGSQL" envDefault:"g9Npp7VTaCb4z0WjMs7TNOVUc3SC7ZD9xmJaLqCDef9jdy2X8y" `
	Timeout  int    `env:"TIMEOUT_PGSQL" envDefault:"10" `
}

func New() *Config {
	conf := &Config{}
	_ = godotenv.Load(fmt.Sprintf(".env"))
	if err := env.Parse(conf); err != nil {
		panic(err)
	}
	return conf
}

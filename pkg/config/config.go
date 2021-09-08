package config

type HttpServerConfig struct {
	Http Http
	Grpc Grpc
}

type GrpcServerConfig struct {
	Grpc Grpc
	Vk   Vk
	//Db   Db
}

type Http struct {
	Port string `required:"true"`
}

type Grpc struct {
	Port string `required:"true"`
}

type Vk struct {
	Token string `required:"true"`
}

type Db struct {
	Dsn string `required:"true"`
}

package main

type Config struct {
	DB
	Redis
	Level     string
	Domain    string
	StartCron bool       `yaml:"start_cron"`
	CloudApis []CloudApi `yaml:"cloud_api"`
	Dingtalk
}
type Redis struct {
	Addr     string
	Password string
	DB       int
}
type DB struct {
	Host     string
	User     string
	Password string
	Database string
}
type CloudApi struct {
	Name      string
	Type      string
	Region    string
	ProjectId string `yaml:"project_id"`
	Ak        string
	Sk        string
}
type Dingtalk struct {
	Token string
	Key   string
}

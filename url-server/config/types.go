package config

// Server contain Server IP/Port configuration parameters
type Server struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

// DB contain Databse related configuration parameters
type DB struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"db_name"`
}

// JWT contain JWT related configuration parameters
type JWT struct {
	Duration         int    `yaml:"duration_minutes"`
	RefreshDuration  int    `yaml:"refresh_duration_minutes"`
	SigningAlgorithm string `yaml:"signing_algorithm"`
	SigningKeyEnv    string `yaml:"signing_key_env"`
}

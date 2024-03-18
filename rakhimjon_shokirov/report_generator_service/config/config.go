package config

type Config struct {
	HTTPPort    string    `env:"HTTP_PORT"`
	ReportMysql *DBConfig `env:",prefix=MYSQL_REPORT_"`
	LedgerMysql *DBConfig `env:",prefix=MYSQL_LEDGER_"`
}

type DBConfig struct {
	Host     string `env:"HOST,required"`
	Port     string `env:"PORT,required"`
	Database string `env:"DATABASE,required"`
	User     string `env:"USER,required"`
	Password string `env:"PASSWORD,required"`
}

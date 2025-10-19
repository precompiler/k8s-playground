package mysql


// Allowed values: `standalone` or `replication`.
type MysqlArchitecture string

const (
	// standalone.
	MysqlArchitecture_STANDALONE MysqlArchitecture = "STANDALONE"
	// replication.
	MysqlArchitecture_REPLICATION MysqlArchitecture = "REPLICATION"
)


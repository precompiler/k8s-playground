package mysql


type MysqlAuth struct {
	Password *string `field:"required" json:"password" yaml:"password"`
	Username *string `field:"required" json:"username" yaml:"username"`
	// Values that are not available in values.schema.json will not be code generated. You can add such values to this property.
	AdditionalValues *map[string]interface{} `field:"optional" json:"additionalValues" yaml:"additionalValues"`
	CreateDatabase *bool `field:"optional" json:"createDatabase" yaml:"createDatabase"`
	Database *string `field:"optional" json:"database" yaml:"database"`
	ReplicationPassword *string `field:"optional" json:"replicationPassword" yaml:"replicationPassword"`
	ReplicationUser *string `field:"optional" json:"replicationUser" yaml:"replicationUser"`
	// Defaults to a random 10-character alphanumeric string if not set.
	// Default: a random 10-character alphanumeric string if not set.
	//
	RootPassword *string `field:"optional" json:"rootPassword" yaml:"rootPassword"`
}


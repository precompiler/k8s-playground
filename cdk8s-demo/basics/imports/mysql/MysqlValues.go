package mysql


type MysqlValues struct {
	// Values that are not available in values.schema.json will not be code generated. You can add such values to this property.
	AdditionalValues *map[string]interface{} `field:"optional" json:"additionalValues" yaml:"additionalValues"`
	// Allowed values: `standalone` or `replication`.
	Architecture MysqlArchitecture `field:"optional" json:"architecture" yaml:"architecture"`
	Auth *MysqlAuth `field:"optional" json:"auth" yaml:"auth"`
	Common *map[string]interface{} `field:"optional" json:"common" yaml:"common"`
	Global *map[string]interface{} `field:"optional" json:"global" yaml:"global"`
	Primary *MysqlPrimary `field:"optional" json:"primary" yaml:"primary"`
	Secondary *MysqlSecondary `field:"optional" json:"secondary" yaml:"secondary"`
}


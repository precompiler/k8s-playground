package mysql


type MysqlPrimaryContainerSecurityContext struct {
	// Values that are not available in values.schema.json will not be code generated. You can add such values to this property.
	AdditionalValues *map[string]interface{} `field:"optional" json:"additionalValues" yaml:"additionalValues"`
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	RunAsUser *float64 `field:"optional" json:"runAsUser" yaml:"runAsUser"`
}


package mysql


type MysqlPrimary struct {
	// Values that are not available in values.schema.json will not be code generated. You can add such values to this property.
	AdditionalValues *map[string]interface{} `field:"optional" json:"additionalValues" yaml:"additionalValues"`
	ContainerSecurityContext *MysqlPrimaryContainerSecurityContext `field:"optional" json:"containerSecurityContext" yaml:"containerSecurityContext"`
	Persistence *MysqlPrimaryPersistence `field:"optional" json:"persistence" yaml:"persistence"`
	PodSecurityContext *MysqlPrimaryPodSecurityContext `field:"optional" json:"podSecurityContext" yaml:"podSecurityContext"`
}


package mysql


type MysqlSecondary struct {
	// Values that are not available in values.schema.json will not be code generated. You can add such values to this property.
	AdditionalValues *map[string]interface{} `field:"optional" json:"additionalValues" yaml:"additionalValues"`
	ContainerSecurityContext *MysqlSecondaryContainerSecurityContext `field:"optional" json:"containerSecurityContext" yaml:"containerSecurityContext"`
	Persistence *MysqlSecondaryPersistence `field:"optional" json:"persistence" yaml:"persistence"`
	PodSecurityContext *MysqlSecondaryPodSecurityContext `field:"optional" json:"podSecurityContext" yaml:"podSecurityContext"`
}


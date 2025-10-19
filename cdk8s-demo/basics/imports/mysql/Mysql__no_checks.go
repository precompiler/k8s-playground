//go:build no_runtime_type_checking

package mysql

// Building without runtime type checking enabled, so all the below just return nil

func validateMysql_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewMysqlParameters(scope constructs.Construct, id *string, props *MysqlProps) error {
	return nil
}


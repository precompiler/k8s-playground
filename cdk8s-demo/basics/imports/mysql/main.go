// mysql
package mysql

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"mysql.Mysql",
		reflect.TypeOf((*Mysql)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Mysql{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterEnum(
		"mysql.MysqlArchitecture",
		reflect.TypeOf((*MysqlArchitecture)(nil)).Elem(),
		map[string]interface{}{
			"STANDALONE": MysqlArchitecture_STANDALONE,
			"REPLICATION": MysqlArchitecture_REPLICATION,
		},
	)
	_jsii_.RegisterStruct(
		"mysql.MysqlAuth",
		reflect.TypeOf((*MysqlAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"mysql.MysqlPrimary",
		reflect.TypeOf((*MysqlPrimary)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"mysql.MysqlPrimaryContainerSecurityContext",
		reflect.TypeOf((*MysqlPrimaryContainerSecurityContext)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"mysql.MysqlPrimaryPersistence",
		reflect.TypeOf((*MysqlPrimaryPersistence)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"mysql.MysqlPrimaryPodSecurityContext",
		reflect.TypeOf((*MysqlPrimaryPodSecurityContext)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"mysql.MysqlProps",
		reflect.TypeOf((*MysqlProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"mysql.MysqlSecondary",
		reflect.TypeOf((*MysqlSecondary)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"mysql.MysqlSecondaryContainerSecurityContext",
		reflect.TypeOf((*MysqlSecondaryContainerSecurityContext)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"mysql.MysqlSecondaryPersistence",
		reflect.TypeOf((*MysqlSecondaryPersistence)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"mysql.MysqlSecondaryPodSecurityContext",
		reflect.TypeOf((*MysqlSecondaryPodSecurityContext)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"mysql.MysqlValues",
		reflect.TypeOf((*MysqlValues)(nil)).Elem(),
	)
}

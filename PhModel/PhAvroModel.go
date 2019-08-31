package PhModel

import (
	"encoding/json"
	"github.com/elodina/go-avro"
	"reflect"
)

type PhAvroModel interface {
	GenSchema(data interface{}) PhAvroModel
	GenRecord(data interface{}) (record *avro.GenericRecord, err error)
	Inject(data map[string]interface{}) error
}

type PhSchemaModel struct {
	schemaString string
}

/** 暂不支持嵌套结构 */
func (model *PhSchemaModel) GenSchema(data interface{}) PhAvroModel {
	schemaMap := make(map[string]interface{})
	schemaMap["type"] = "record"
	schemaMap["namespace"] = "com.pharbers.kafka.schema"

	schemaName := reflect.TypeOf(data).Elem().Name()
	schemaMap["name"] = schemaName

	var fields = make([]map[string]string, 0)
	v := reflect.ValueOf(data).Elem()
	for j := 0; j < v.NumField(); j++ {
		// 跳过指针类型
		if v.Field(j).Kind().String() == "ptr" {
			continue
		}
		fieldName := v.Type().Field(j).Name
		fieldType := v.Field(j).Type().Name()
		fields = append(fields, map[string]string{"name": fieldName, "type": fieldType})
	}
	schemaMap["fields"] = fields
	schemaBytes, _ := json.Marshal(schemaMap)

	model.schemaString = string(schemaBytes)
	return model
}

/** 暂不支持嵌套结构 */
func (model *PhSchemaModel) GenRecord(data interface{}) (record *avro.GenericRecord, err error) {
	if model.schemaString == "" {
		model.GenSchema(data)
	}

	schema, err := avro.ParseSchema(model.schemaString)
	if err != nil {
		return
	}
	record = avro.NewGenericRecord(schema)

	v := reflect.ValueOf(data).Elem()
	for j := 0; j < v.NumField(); j++ {
		// 跳过指针类型
		if v.Field(j).Kind().String() == "ptr" {
			continue
		}
		fieldName := v.Type().Field(j).Name
		record.Set(fieldName, v.Field(j).Interface())
	}

	return
}

func (model *PhSchemaModel) Inject(data map[string]interface{}) error {
	return nil
}

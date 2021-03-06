package decimal

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//MarshalBSONValue mongo-driver
func (d Decimal) MarshalBSONValue() (bsontype.Type, []byte, error) {
	tmp, _ := primitive.ParseDecimal128(d.String())
	return bson.MarshalValue(tmp)
}

//UnmarshalBSONValue mongo-driver
func (d Decimal) UnmarshalBSONValue(btype bsontype.Type, value []byte) error {
	str := ""
	rv := bson.RawValue{Type: btype, Value: value}
	rv.Unmarshal(&str)
	x, err := NewFromString(str)
	if err == nil {
		d = x
	}
	return err
}

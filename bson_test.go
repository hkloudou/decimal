package decimal

import (
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Test_xx(t *testing.T) {
	tmp, _ := primitive.ParseDecimal128("23")
	bson.MarshalValue(tmp)
	t.Log(bson.Marshal(tmp))

}

// func (f Field) MarshalBSONValue() (bsontype.Type, []byte, error) {
// 	return f.Type, f.Value, nil
// }

// func (f *Field) UnmarshalBSONValue(btype bsontype.Type, value []byte) error {
// 	f.Type = btype
// 	f.Value = value
// 	return nil
// }

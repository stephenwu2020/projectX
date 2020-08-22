package collections

type Users struct {
	Uid  uint32 `json:"uid" bson:"uid"`
	Name string `json:"name" bson:"name"`
	Sex  uint8  `json:"sex" bson:"sex"`
}

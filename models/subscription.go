package models

type Subscription struct {
	ID          string `json:"id,omitempty" bson:"_id,omitempty"`
	TypeCode    string `json:"type_code,omitempty" bson:"type_code,omitempty"`
	NumOfSource uint   `json:"num_of_source,omitempty" bson:"num_of_source,omitempty"`
	NumOfKey    uint   `json:"num_of_key,omitempty" bson:"num_of_key,omitempty"`
}

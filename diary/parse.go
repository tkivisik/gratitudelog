package diary

import "encoding/json"

func ParseJSON(input []byte) (*Diary, error) {
	var diary Diary
	err := json.Unmarshal(input, &diary)
	return &diary, err
}

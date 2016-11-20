package diary

import "encoding/json"

func ParseJSON(input string) (*Diary, error) {
	var diary Diary
	err := json.Unmarshal([]byte(input), &diary)
	return &diary, err
}

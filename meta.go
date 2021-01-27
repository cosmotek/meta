package meta

type JSONMeta map[string]string

func (a JSONMeta) Value() (driver.Value, error) {
	return json.Marshal(a)
}

func (a *JSONMeta) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}

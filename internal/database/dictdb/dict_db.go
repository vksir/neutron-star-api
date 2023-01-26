package dictdb

import (
	"neutron-star-api/internal/database"
)

type KeyValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func Get(key string) (*KeyValue, error) {
	kv := KeyValue{}
	err := database.DB.QueryRow(
		"select * from t_dict where key = $1", key).Scan(&kv.Key, &kv.Value)
	if err != nil {
		return nil, err
	}
	return &kv, nil
}

func Set(key, value string) error {
	_, err := database.DB.Exec(
		"insert into t_dict (key, value) values ($1, $2)", key, value)
	return err
}

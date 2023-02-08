package models

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
)

type Event struct {
	ID        string `gorm:"column:id;primaryKey;type:char(64)"`
	Pubkey    string `gorm:"column:pubkey;index;type:char(64)"`
	CreatedAt int64  `gorm:"column:created_at;index"`
	Kind      int    `gorm:"column:kind;index"`
	Tags      string `gorm:"column:tags;type:text"`
	Content   string `gorm:"column:content;type:text"`
	Sig       string `gorm:"column:sig;type:char(128)"`
}

func (Event) TableName() string {
	return "events"
}

func (event Event) CalculateID() (string, error) {
	buf, err := event.Serialize()
	if err != nil {
		return "", err
	}
	digest := sha256.New()
	digest.Write(buf)
	return hex.EncodeToString(digest.Sum(nil)), nil
}
func (event Event) Serialize() ([]byte, error) {
	var tags [][]string
	if event.Tags != "" {
		json.Unmarshal([]byte(event.Tags), &tags)
	}

	rawData := []interface{}{0, event.Pubkey, event.CreatedAt, event.Kind, tags, event.Content}
	buf, err := json.Marshal(rawData)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

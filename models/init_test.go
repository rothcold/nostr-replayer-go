package models

import (
	"encoding/json"
	"log"
	"testing"
)

func TestGenerateTables(t *testing.T) {
	GenerateTables()
}

func TestInsert(t *testing.T) {
	tags := [][]string{{"a", "1"}, {"b", "2"}}
	tagsBuf, err := json.Marshal(tags)
	if err != nil {
		t.Fatal(err)
	}
	event := Event{
		ID:        "8412a80c494287045dce358e6d9110c88f8c989e70f5d266dee2623195a14ebd",
		Pubkey:    "5b0e8da6fdfba663038690b37d216d8345a623cc33e111afd0f738ed7792bc54",
		CreatedAt: 1675872970,
		Kind:      1,
		Tags:      string(tagsBuf),
		Content:   "Matcha is such an incredibly pretty color. \n\nhttps://nostr.build/i/nostr.build_299a036695652982553c601f837de092f96ef82e266362bf8ff71467556666c6.jpg",
		Sig:       "53838abe3a564e91fd7b5bac927032fa20c7f86138b5e27700bf5833f40d4e1c883bb49fb0faedc71d53441b7a352de4afae461ba90f5c4d27081f95a8f5001d",
	}
	err = database.Create(event).Error
	if err != nil {
		t.Fatal(err)
	}
	var newEvent Event
	err = database.Model(newEvent).Where("id", "8412a80c494287045dce358e6d9110c88f8c989e70f5d266dee2623195a14ebd").First(&newEvent).Error
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("%v", newEvent)
	var newTags [][]string
	err = json.Unmarshal([]byte(newEvent.Tags), &newTags)
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("%v", newTags)
	// err = database.Delete(newEvent, "id =?", "8412a80c494287045dce358e6d9110c88f8c989e70f5d266dee2623195a14ebd").Error
	// if err != nil {
	// 	t.Fatal(err)
	// }
}

func TestCalculateID(t *testing.T) {
	event := Event{
		Pubkey:    "3f770d65d3a764a9c5cb503ae123e62ec7598ad035d836e2a810f3877a745b24",
		CreatedAt: 1675875193,
		Kind:      1,
		Tags:      "[[\"e\",\"8d4c990851482898ed060e6789b2bd2ee99869041d813d8305036224f130221f\",\"\",\"reply\"],[\"p\",\"ecfa3c5c82d589c867c044056f75d6cff794f1886d5ebcdd48ad851da47adae4\"]]",
		Content:   "Nope?",
	}
	id, err := event.CalculateID()
	if err != nil {
		t.Fail()
		log.Fatal(err)
	}

	if id != "2a446f4822c79096b993e4653d4004fbfa06bd954b58141adccf6bd6e1d9a219" {
		t.Fail()
		log.Fatal("Invalid ID")
	}

	event2 := Event{
		CreatedAt: 1675835576,
		Pubkey:    "fa984bd7dbb282f07e16e7ae87b26a2a7b9b90b7246a44771f0cf5ae58018f52",
		Kind:      1,
		Content:   "Are you still using Twitter?\n\nWhy?\n\nI’m curious 🧡",
		Tags:      "[]",
	}
	id, err = event2.CalculateID()
	if err != nil {
		t.Fail()
		log.Fatal(err)
	}
	if id != "79eeb6aa19a5a94fba6402fb4dedc3f385a9a8ca956f638043300ecca5e8bd0b" {
		t.Fail()
		log.Fatal("Invalid ID")
	}
}

package db

import (
	"gorm.io/gorm"
)

type SCHEDULE struct {
	gorm.Model
	Key        			string            
	Overseers	  		string		
	Comments    		string
	ChecklistJSON		string
}

type CHECKLIST struct {
	gorm.Model
	ChecklistItem    string   		`json:"ChecklistItem"`
}

type ChecklistItem struct {
	Category       string          `json:"category"`
	SwedishValues  []SwedishValue  `json:"swedish_values"`
	EnglishValues  []EnglishValue  `json:"english_values"`
}

type SwedishValue struct {
	ChecklistItemID string   `json:"ChecklistItemID"`
	Prompt          string   `json:"prompt"`
}

type EnglishValue struct {
	ChecklistItemID string   `json:"ChecklistItemID"`
	Prompt          string   `json:"prompt"`
}

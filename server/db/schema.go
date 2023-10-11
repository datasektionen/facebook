package db

import (
	"gorm.io/gorm"
)

type SCHEDULE struct {
	Key        			string            
	Overseers	  		string		
	Comments    		string
	Checklist			int
}

type CHECKLIST struct {
	ChecklistID 		int
	Prompts []PROMPTS `gorm:"foreignKey:ChecklistID"`
}

type PROMPTS struct {
	gorm.Model
	Category			string
	Prompt				string
	Value				bool
}

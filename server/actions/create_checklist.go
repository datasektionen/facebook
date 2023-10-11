package actions

import (
	"github.com/gin-gonic/gin"

	database "github.com/datasektionen/facebook/server/db"
)



func CreateChecklist(c *gin.Context){
    database.InitDB()
    db := database.GetDB()

	checklist := database.CHECKLIST{
		ChecklistID: 1,
		Prompts: []database.PROMPTS{
			{
				Category: "Allmänt",
				Prompt:   "Ta bort dekoration och tejp som använts under eventet / Remove decorations and tape used during the event",
				Value:    false,
			},
			{
				Category: "Allmänt",
				Prompt:   "Rensa bort allt skräp som uppkommit under eventet. (Kolla bakom soffor och i fönstret) / Clear away all the trash that has accumulated during the event. (Check behind sofas and in the windows)",
				Value:    false,
			},
			{
				Category: "Allmänt",
				Prompt:   "Sopa alla golvytor / Sweep all floor surfaces",
				Value:    false,
			},
			{
				Category: "Allmänt",
				Prompt:   "!Återställ uteplatsen utanför META! / !Restore the outdoor patio outside of META!",
				Value:    false,
			},
			{
				Category: "Allmänt",
				Prompt:   "Svabba & skrapa alla golvytor (ska alltid göras efter pub) / Mop and scrape all floor surfaces (should always be done after the pub)",
				Value:    false,
			},
			{
				Category: "Allmänt",
				Prompt:   "Skölj ur och lägg burkar och petflaskor i pant-kärlen / Rinse and place cans and plastic bottles in the recycling bins",
				Value:    false,
			},
			{
				Category: "Allmänt",
				Prompt:   "Töm papperskorgar och byt säckar / Empty trash cans and replace bags",
				Value:    false,
			},
			{
				Category: "Allmänt",
				Prompt:   "Släng kartonger och glas i soprummet / Dispose of cardboard and glass in the garbage room",
				Value:    false,
			},
			{
				Category: "Allmänt",
				Prompt:   "Släng kapsyler och annan metall (äckliga konserver slängs i soprummet) / Dispose of caps and other metals (disgusting cans go in the garbage room)",
				Value:    false,
			},
			{
				Category: "Allmänt",
				Prompt:   "Städa extraordinärt äckel (spyor och dyl.) / Clean up extraordinary filth (vomit, etc.)",
				Value:    false,
			},
		},
	}


	db.Create(checklist)
}
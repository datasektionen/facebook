package actions

import (
	"github.com/gin-gonic/gin"
	"fmt"

	"encoding/json"

	database "github.com/datasektionen/facebook/server/db"
)


func CreateChecklist(c *gin.Context) {
	database.InitDB()
	db := database.GetDB()

	fmt.Println("CREATE CHECKLIST CODE RUNNING!")

	checklistTemplate := []database.ChecklistItem{
		{
			Category: "Hallen & Toaletter",
			SwedishValues: []database.SwedishValue{
				{
					Prompt:           "Plocka bort stolar och bord som använts i entrén",
					ChecklistItemID: "A1",
				},
				{
					Prompt:           "Ställ tillbaka saker ordentligt i städskrubben",
					ChecklistItemID: "A2",
				},
				{
					Prompt:           "!Se till att brandcentralen inte blockeras!",
					ChecklistItemID: "A3",
				},
				{
					Prompt: "!Återställ dörröppnaren så att den funkar dagen efter. Båda knapparna på apparaten ska vara tryckta uppåt för att återställa!",
					ChecklistItemID: "A4",
				},
				{
					Prompt: "!Se till att Entrédörren är stängd och låst!",
					ChecklistItemID: "A5",
				},
				{
					Prompt:           "Se till att det inte ser förjävligt ut",
					ChecklistItemID: "A6",
				},
			},
			EnglishValues: []database.EnglishValue{
				{
					Prompt:           "Remove chairs and tables that have been used at the entrance",
					ChecklistItemID: "A1",
				},
				{
					Prompt:           "Put things back properly in the cleaning closet",
					ChecklistItemID: "A2",
				},
				{
					Prompt:           "!Ensure that the fire center is not blocked!",
					ChecklistItemID: "A3",
				},
				{
					Prompt: "!Reset the door opener so that it works the day after. Both buttons on the device should be pushed upward to reset!",
					ChecklistItemID: "A4",
				},
				{
					Prompt: "!Make sure that the entrance door is closed and locked!",
					ChecklistItemID: "A5",
				},
				{
					Prompt:           "Make sure it doesn't look terrible",
					ChecklistItemID: "A6",
				},
			},
		},
		{
			Category: "Allmänt",
			SwedishValues: []database.SwedishValue{
				{
					Prompt:           "Ta bort dekoration och tejp som använts under eventet",
					ChecklistItemID: "B1",
				},
				{
					Prompt:           "Rensa bort allt skräp som uppkommit under eventet. (Kolla bakom soffor och i fönstret)",
					ChecklistItemID: "B2",
				},
				{
					Prompt:           "!Återställ uteplatsen utanför META!",
					ChecklistItemID: "B3",
				},
				{
					Prompt:           "Sopa alla golvytor",
					ChecklistItemID: "B4",
				},
				{
					Prompt:           "Svabba & skrapa alla golvytor (ska alltid göras efter pub)",
					ChecklistItemID: "B5",
				},
				{
					Prompt:           "Skölj ur och lägg burkar och petflaskor i pant-kärlen",
					ChecklistItemID: "B6",
				},
				{
					Prompt:           "Töm papperskorgar och byt säckar",
					ChecklistItemID: "B7",
				},
				{
					Prompt:           "Släng kartonger och glas i soprummet",
					ChecklistItemID: "B8",
				},
				{
					Prompt:           "Släng kapsyler och annan metall (äckliga konserver slängs i soprummet)",
					ChecklistItemID: "B9",
				},
				{
					Prompt:           "Städa extraordinärt äckel (spyor och dyl.)",
					ChecklistItemID: "B10",
				},
			},
			EnglishValues: []database.EnglishValue{
				{
					Prompt:           "Remove decorations and tape used during the event",
					ChecklistItemID: "B1",
				},
				{
					Prompt:           "Clear away all the trash that has accumulated during the event. (Check behind sofas and in the windows)",
					ChecklistItemID: "B2",
				},
				{
					Prompt:           "!Restore the outdoor patio outside of META!",
					ChecklistItemID: "B3",
				},
				{
					Prompt:           "Sweep all floor surfaces",
					ChecklistItemID: "B4",
				},
				{
					Prompt:           "Mop and scrape all floor surfaces (should always be done after the pub)",
					ChecklistItemID: "B5",
				},
				{
					Prompt:           "Rinse and place cans and plastic bottles in the recycling bins",
					ChecklistItemID: "B6",
				},
				{
					Prompt:           "Empty trash cans and replace bags",
					ChecklistItemID: "B7",
				},
				{
					Prompt:           "Dispose of cardboard and glass in the garbage room",
					ChecklistItemID: "B8",
				},
				{
					Prompt:           "Dispose of caps and other metals (disgusting cans go in the garbage room)",
					ChecklistItemID: "B9",
				},
				{
					Prompt:           "Clean up extraordinary filth (vomit, etc.)",
					ChecklistItemID: "B10",
				},
			},
		},
		{
			Category: "Bakom baren",
			SwedishValues: []database.SwedishValue{
				{
					Prompt:           "Töm diskställen på allt som är tort",
					ChecklistItemID: "C1",
				},
				{
					Prompt:           "Diska & plocka undan barutrustningen",
					ChecklistItemID: "C2",
				},
				{
					Prompt:           "!Diska och plocka undan all disk, ställ inte in blöta saker i skåpen!",
					ChecklistItemID: "C3",
				},
				{
					Prompt:           "!Se till att isskoporna ligger i lådan ovanför ismaskinen med stängt lock!",
					ChecklistItemID: "C4",
				},
				{
					Prompt:           "Städa ur kylarna från saker ni använt (inkluderar även eventkylen). Ätbar mat kan ställas i gratiskylen men om den inte äts upp är ni ansvariga för att städa bort det sen",
					ChecklistItemID: "C5",
				},
				{
					Prompt:           "!Se till att Brutus är avstängd och tömd!",
					ChecklistItemID: "C6",
				},
				{
					Prompt:           "Se till att diskhoarna är rena",
					ChecklistItemID: "C7",
				},
				{
					Prompt:           "Torka av spritdiket utvändigt och invändigt, inklusive under gallret",
					ChecklistItemID: "C8",
				},
				{
					Prompt:           "Torka av bänkytorna och bardisken med trasa + såpa/diskmedel",
					ChecklistItemID: "C9",
				},
				{
					Prompt:           "Återställ bänkytorna",
					ChecklistItemID: "C10",
				},
				{
					Prompt:           "!Se till att ölkylarna är stängda och låsta!",
					ChecklistItemID: "C11",
				},
				{
					Prompt:           "!Se till att spritdiket är avstängt (dra ur sladden) och dörrarna är öppna (om den var på)!",
					ChecklistItemID: "C12",
				},
				{
					Prompt:           "!Töm, gör rent, stäng av och lämna diskmaskinen öppen enligt instruktioner!",
					ChecklistItemID: "C13",
				},
			},
			EnglishValues: []database.EnglishValue{
				{
					Prompt:           "Empty the dish racks of everything that is dry",
					ChecklistItemID: "C1",
				},
				{
					Prompt:           "Wash and put away the bar equipment",
					ChecklistItemID: "C2",
				},
				{
					Prompt:           "!Wash and put away all dishes; do not place wet items in the cabinets!",
					ChecklistItemID: "C3",
				},
				{
					Prompt:           "!Ensure that the ice scoops are in the box above the ice machine with the lid closed!",
					ChecklistItemID: "C4",
				},
				{
					Prompt:           "Clean out the refrigerators of items you've used (including the event fridge). Edible food can be placed in the free fridge, but if it's not eaten, you are responsible for cleaning it up later",
					ChecklistItemID: "C5",
				},
				{
					Prompt:           "!Ensure that Brutus is turned off and emptied!",
					ChecklistItemID: "C6",
				},
				{
					Prompt:           "Make sure the sinks are clean",
					ChecklistItemID: "C7",
				},
				{
					Prompt:           "Wipe the spirit trough outside and inside, including under the grate",
					ChecklistItemID: "C8",
				},
				{
					Prompt:           "Wipe the countertops and the bar counter with a cloth + soap/dish detergent",
					ChecklistItemID: "C9",
				},
				{
					Prompt:           "Restore the countertops",
					ChecklistItemID: "C10",
				},
				{
					Prompt:           "!Ensure that the beer fridges are closed and locked!",
					ChecklistItemID: "C11",
				},
				{
					Prompt:           "!Ensure that the liquore ditch is turned off (unplug it) and the doors are open (if it was on)!",
					ChecklistItemID: "C12",
				},
				{
					Prompt:           "!Empty, clean, turn off, and leave the dishwasher open according to instructions!",
					ChecklistItemID: "C13",
				},
			},
		},
		{
			Category: "Huvudrummet",
			SwedishValues: []database.SwedishValue{
				{
					Prompt:           "Torka av bord och stolar med trasa + såpa/diskmedel",
					ChecklistItemID: "D1",
				},
				{
					Prompt:           "Ställ tillbaka bord och stolar enligt bordsplaceringen (stolar upp-och-ner på bord)",
					ChecklistItemID: "D2",
				},
				{
					Prompt:           "Plocka undan saker från bord och andra ytor (inkl. soffor)",
					ChecklistItemID: "D3",
				},
				{
					Prompt:           "!Plocka undan saker (stolar/bord) från den tejpade ytan framför mikrovågsugnarna!",
					ChecklistItemID: "D4",
				},
				{
					Prompt:           "!Se till att nödutgång är stängd och låst!",
					ChecklistItemID: "D5",
				},
			},
			EnglishValues: []database.EnglishValue{
				{
					Prompt:           "Wipe tables and chairs with a cloth + soap/dish detergent",
					ChecklistItemID: "D1",
				},
				{
					Prompt:           "Place tables and chairs back according to the table arrangement (chairs upside down on the tables)",
					ChecklistItemID: "D2",
				},
				{
					Prompt:           "Clear away items from tables and other surfaces (including sofas)",
					ChecklistItemID: "D3",
				},
				{
					Prompt:           "!Remove items (chairs/tables) from the taped area in front of the microwaves!",
					ChecklistItemID: "D4",
				},
				{
					Prompt:           "!Ensure that the emergency exit is closed and locked!",
					ChecklistItemID: "D5",
				},
			},
		},
		{
			Category: "Catwalken",
			SwedishValues: []database.SwedishValue{
				{
					Prompt: "Plocka undan saker från catwalkborden",
					ChecklistItemID: "E1",
				},
				{
					Prompt: "Se till att utklädningskläder som använts viks ihop och läggs ner i rätt låda",
					ChecklistItemID: "E2",
				},
				{
					Prompt: "Se till att det inte ligger kartonger och annat skräp kvar nedanför trappan",
					ChecklistItemID: "E3",
				},
				{
					Prompt: "Se till att ljud & ljus-systemen är avstängda",
					ChecklistItemID: "E4",
				},
				{
					Prompt: "Se till att det inte ser förjävligt ut",
					ChecklistItemID: "E5",
				},
			},
			EnglishValues: []database.EnglishValue{
				{
					Prompt: "Clear away items from the catwalk tables",
					ChecklistItemID: "E1",
				},
				{
					Prompt: "Ensure that costumes that have been used are folded and placed in the correct box",
					ChecklistItemID: "E2",
				},
				{
					Prompt: "Make sure there are no cardboard boxes and other trash left below the stairs",
					ChecklistItemID: "E3",
				},
				{
					Prompt: "Ensure that the sound and lighting systems are turned off",
					ChecklistItemID: "E4",
				},
				{
					Prompt: "Make sure it doesn't look terrible",
					ChecklistItemID: "E5",
				},
			},
		},
		{
			Category: "Mötesrummet",
			SwedishValues: []database.SwedishValue{
				{
					Prompt: "Se till att det inte ser förjävligt ut",
					ChecklistItemID: "F1",
				},
				{
					Prompt: "!Se till att fönstret i mötesrummet är stängt och låst!",
					ChecklistItemID: "F2",
				},
			},
			EnglishValues: []database.EnglishValue{
				{
					Prompt: "Make sure it doesn't look terrible",
					ChecklistItemID: "F1",
				},
				{
					Prompt: "!Ensure that the window in the meeting room is closed and locked!",
					ChecklistItemID: "F2",
				},
			},
		},
	}

	// Iterate over each checklist item and insert into the database
	for _, checklist := range checklistTemplate {

		checklist_string, _ := json.Marshal(checklist)
		checklistStringLiteral := string(checklist_string)

		// fmt.Println(checklistStringLiteral)

		data_entry := database.CHECKLIST{
			ChecklistItem: checklistStringLiteral,
		}


		err := db.Create(&data_entry).Error
		if err != nil {
			c.JSON(500, gin.H{
				"error": err,
			})
			return
		}
	}

	c.JSON(200, gin.H{
		"message": "Checklist items inserted successfully",
	})
}

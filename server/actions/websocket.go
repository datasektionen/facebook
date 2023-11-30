package actions

import (
	"net/http"
	"fmt"
	"gorm.io/gorm"
	"github.com/gin-gonic/gin"

	"github.com/gorilla/websocket"

	"gorm.io/gorm/clause"
	"encoding/json"
	database "github.com/datasektionen/facebook/server/db"
)

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		// allowedOrigin := "http://your-nextjs-app-domain"
		// requestOrigin := r.Header.Get("Origin")
		// return allowedOrigin == requestOrigin
		return true 
	},
}

var clients = make(map[*websocket.Conn]bool)

func SendWebSocketMessageToClients(c *gin.Context) {
    // Broadcast the message to all connected clients
    fmt.Println("SendWebSocketMessageToClients called")

    for client := range clients {
        fmt.Println("Sending message to client")
        err := client.WriteMessage(websocket.TextMessage, []byte("this is from an API!"))
        if err != nil {
            fmt.Println("Error sending message:", err)
            // Handle the error, e.g., remove the client from the list
            delete(clients, client)
        }
    }
}

func SendWebSocketDataToClients(c *gin.Context) {
    // Broadcast the message to all connected clients
    fmt.Println("SendWebSocketMessageToClients called")

	data_marshal, err := json.Marshal(GetData("sokior"))
	if err != nil {
		fmt.Println("Error marshaling to JSON:", err)
		return
	}

    for client := range clients {
        fmt.Println("Sending message to client")
        // err := client.WriteMessage(websocket.TextMessage, []byte("this is from an API!"))
        err := client.WriteMessage(websocket.TextMessage, data_marshal)
        if err != nil {
            fmt.Println("Error sending message:", err)
            // Handle the error, e.g., remove the client from the list
            delete(clients, client)
        }
    }
}

type SCHEDULE_new struct {
	gorm.Model
	Key          string
	Overseers    string
	Comments     string
	ChecklistJSON map[string]bool
}

type CombinedData struct {
	ScheduleItems        []SCHEDULE_new      `json:"scheduleItems"`
	ParsedChecklistItems []database.ChecklistItem `json:"parsedChecklistItems"`
}

func GetData(quary string) CombinedData{


	database.InitDB()
    db := database.GetDB()

	var schedules []database.SCHEDULE
    if err := db.Find(&schedules).Error; err != nil {
		fmt.Println("ERR getting schedules from database")
    }

	var scheduleItems []SCHEDULE_new
	for _, schedule := range schedules {
		if schedule.Key != quary {
			continue
		}
	
		var checklistJSON map[string]bool // Change here
		checklist_unmarshal := json.Unmarshal([]byte(schedule.ChecklistJSON), &checklistJSON)
	
		fmt.Println("checklist_marshal: ", checklist_unmarshal)
	
		scheduleItems = append(scheduleItems, SCHEDULE_new{
			Key:           schedule.Key,
			Overseers:     schedule.Overseers,
			Comments:      schedule.Comments,
			ChecklistJSON: checklistJSON,
		})
	}

	var checklistItems []database.CHECKLIST
	if err := db.Preload(clause.Associations).Find(&checklistItems).Error; err != nil {
		fmt.Println("ERR getting CHECKLIST from database")
	}

    // Now, loop through each item and parse the JSON
    var parsedChecklistItems []database.ChecklistItem
    for _, item := range checklistItems {
        var checklistItem database.ChecklistItem

        // Assuming ChecklistItem field in CHECKLIST contains the JSON data
        if err := json.Unmarshal([]byte(item.ChecklistItem), &checklistItem); err != nil {
			fmt.Println("ERR getting CHECKLIST from database")
        }

        parsedChecklistItems = append(parsedChecklistItems, checklistItem)
    }

	combinedData := CombinedData{
		ScheduleItems:        scheduleItems,
		ParsedChecklistItems: parsedChecklistItems,
	}

	return combinedData
}

func SendWebsocket(c *gin.Context) {

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Print("upgrade failed: ", err)
		return
	}


	data_marshal, err := json.Marshal(GetData("sokior"))
	if err != nil {
		fmt.Println("Error marshaling to JSON:", err)
		return
	}
	err = conn.WriteMessage(websocket.TextMessage, data_marshal)
	if err != nil {
		fmt.Println("Error sending message:", err)
		// Handle the error, e.g., remove the client from the list
		delete(clients, conn)
	}

	clients[conn] = true

}
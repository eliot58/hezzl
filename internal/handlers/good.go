package handlers

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/eliot58/hezzl/internal/database"
	"github.com/gin-gonic/gin"
)

func CreateGood(c *gin.Context) {
	projectId := c.Query("projectId")

	var data map[string]interface{}

	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var maxPriority sql.NullInt64
	err := database.Db.QueryRow("SELECT MAX(priority) FROM goods").Scan(&maxPriority)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var id int
	var description string
	var priority int
	var removed bool
	var createdAt time.Time
	

	if maxPriority.Valid {
        database.Db.QueryRow("INSERT INTO goods (project_id, name, priority) VALUES ($1, $2, $3) RETURNING id", projectId, data["name"], maxPriority.Int64+1).Scan(&id)
    } else {
		database.Db.QueryRow("INSERT INTO goods (project_id, name, priority) VALUES ($1, $2, $3) RETURNING id", projectId, data["name"], 1).Scan(&id)
	}

	row := database.Db.QueryRow("SELECT description, priority, removed, created_at FROM goods WHERE id = $1", id)
	row.Scan(&description, &priority, &removed, &createdAt)

	createdData := map[string]interface{}{
        "id":          id,
        "description": description,
        "priority":    priority,
        "removed":     removed,
        "created_at":  createdAt,
    }

    c.JSON(http.StatusOK, createdData)

}

func UpdateGood(c *gin.Context) {
	// id := c.Query("id")
	// projectId := c.Query("projectId")
	var data map[string]interface{}

	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}


}


func DeleteGood(c *gin.Context) {

}


func GetGoods(c *gin.Context) {

}

func Reprioritize(c *gin.Context) {

}

package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/kamran0812/notes-backend/db"
	"github.com/kamran0812/notes-backend/internal/model"
)

// Get ALL
func GetNotes(c *fiber.Ctx) error {
	db := db.Conn
	var notes []model.Note

	db.Find(&notes)
	if len(notes) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No notes present", "data": nil})
	}
	return c.JSON(fiber.Map{"status": "success", "message": "Notes Found", "data": notes})
}

// Create
func CreateNotes(c *fiber.Ctx) error {
	db := db.Conn
	note := new(model.Note)

	err := c.BodyParser(note)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}
	note.ID = uuid.New()
	err = db.Create(&note).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create note", "data": err})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Created Note", "data": note})
}

// Get by ID
func GetNote(c *fiber.Ctx) error {
	db := db.Conn
	var note model.Note

	id := c.Params("noteId")

	db.Find(&note, "id = ?", id)

	if note.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No note present", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Notes Found", "data": note})
}

// Update
func UpdateNote(c *fiber.Ctx) error {
	type updateNote struct {
		Title    string `json:"title"`
		SubTitle string `json:"sub_title"`
		Text     string `json:"Text"`
	}
	db := db.Conn
	var note model.Note

	id := c.Params("noteId")

	db.Find(&note, "id = ?", id)

	if note.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No note present", "data": nil})
	}

	var updateNoteData updateNote
	err := c.BodyParser(&updateNoteData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	note.Title = updateNoteData.Title
	note.Text = updateNoteData.Text

	db.Save(&note)

	return c.JSON(fiber.Map{"status": "success", "message": "Notes Found", "data": note})
}

// Delete
func DeleteNote(c *fiber.Ctx) error {
	db := db.Conn
	var note model.Note

	id := c.Params("noteId")

	db.Find(&note, "id = ?", id)

	if note.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No note present", "data": nil})
	}

	err := db.Delete(&note, "id = ?", id).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Failed to delete note", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Deleted Note"})
}

func SearchNotes(c *fiber.Ctx) error {
	db := db.Conn
	var notes []model.Note

	searchQuery := c.Query("query")

	db.Where("title ~* ?", searchQuery).Or(db.Where("text ~* ?", searchQuery)).Find(&notes)

	if len(notes) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "No notes found",
			"data":    nil,
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Notes Found",
		"data":    notes,
	})
}

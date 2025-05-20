package controllers

import (
	"github.com/Andrianns/andrian-universe-service-v1/app/clients"
	"github.com/gofiber/fiber/v2"
)

type DocumentController struct {
	drive clients.GoogleDrive
}

func NewDocumentController(d clients.GoogleDrive) *DocumentController {
	return &DocumentController{drive: d}
}

// UploadCV handles POST /cv
func (ctrl *DocumentController) UploadCV(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "file is required",
		})
	}
	folderID, err := ctrl.drive.EnsureFolder("CV")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to access or create CV folder",
		})
	}

	fileID, err := ctrl.drive.UploadFile(file, folderID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to upload file to Google Drive",
		})
	}

	publicURL, err := ctrl.drive.ShareFile(fileID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to share uploaded file",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "CV uploaded successfully",
		"url":     publicURL,
	})
}

// GetCV handles GET /cv
func (ctrl *DocumentController) GetCV(c *fiber.Ctx) error {
	var req struct {
		FileName string `json:"fileName"`
	}

	if err := c.BodyParser(&req); err != nil || req.FileName == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid or missing fileName in request body",
		})
	}
	url, err := ctrl.drive.GetFileURLByName(req.FileName, "CV")
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "CV file not found",
		})
	}

	return c.JSON(fiber.Map{
		"url": url,
	})
}

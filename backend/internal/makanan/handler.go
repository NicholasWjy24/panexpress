package makanan

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"panexpress.com/internal/middleware"
)

const maxMakananImageSize = 5 << 20

type Handler struct {
	service *Service
}

func MakananHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) RegisterRoutes(router *gin.Engine) {
	authorized := router.Group("/api")
	authorized.Use(middleware.AuthMiddleware())
	authorized.GET("/makanan", h.GetAllMakanan)
	authorized.GET("/makanan/:id/image", h.GetMakananImage)
	authorized.POST("/makanan", h.RegisterMakanan)
	authorized.PUT("/makanan/:id", h.UpdateMakanan)
	authorized.PUT("/makanan/:id/image", h.UpdateMakananImage)
	authorized.DELETE("/makanan/:id", h.DeleteMakanan)
}

func (h *Handler) GetAllMakanan(c *gin.Context) {

	menus, err := h.service.GetAllMakanan()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, menus)
}

func (h *Handler) RegisterMakanan(c *gin.Context) {
	input, err := bindRegisterInput(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	makanan, err := h.service.Register(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": fmt.Sprintf(
			"Makanan %s created successfully!",
			makanan.MknnNm,
		),
	})
}

func (h *Handler) GetMakananImage(c *gin.Context) {
	image, err := h.service.GetMakananImage(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Makanan image not found",
		})
		return
	}

	if len(image) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Makanan has no image",
		})
		return
	}

	c.Data(http.StatusOK, http.DetectContentType(image), image)
}

func (h *Handler) UpdateMakanan(
	c *gin.Context,
) {

	id := c.Param("id")

	var input Makanan

	if err := c.ShouldBindJSON(&input); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid data layout",
		})

		return
	}

	if err := h.service.UpdateMakanan(
		id,
		&input,
	); err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Makanan updated successfully",
	})
}

func (h *Handler) UpdateMakananImage(
	c *gin.Context,
) {

	image, err := readImageFile(c, "image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := h.service.UpdateMakananImage(
		c.Param("id"),
		image,
	); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Makanan image updated successfully",
	})
}

func (h *Handler) DeleteMakanan(
	c *gin.Context,
) {

	id := c.Param("id")

	if err := h.service.DeleteMakanan(id); err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Makanan deleted successfully",
	})
}

func bindRegisterInput(c *gin.Context) (RegisterInputMakanan, error) {
	if !strings.HasPrefix(c.ContentType(), "multipart/form-data") {
		var input RegisterInputMakanan
		err := c.ShouldBindJSON(&input)
		return input, err
	}

	price, err := strconv.ParseFloat(c.PostForm("mknnpr"), 64)
	if err != nil {
		return RegisterInputMakanan{}, fmt.Errorf("invalid mknnpr")
	}

	quantity, err := strconv.Atoi(c.PostForm("mknnqt"))
	if err != nil {
		return RegisterInputMakanan{}, fmt.Errorf("invalid mknnqt")
	}

	status, err := strconv.ParseBool(c.DefaultPostForm("mknnst", "false"))
	if err != nil {
		return RegisterInputMakanan{}, fmt.Errorf("invalid mknnst")
	}

	input := RegisterInputMakanan{
		MknnNm: c.PostForm("mknnnm"),
		MknnTp: c.PostForm("mknntp"),
		MknnPr: price,
		MknnQt: quantity,
		MknnSt: status,
	}

	if input.MknnNm == "" || input.MknnTp == "" {
		return RegisterInputMakanan{}, fmt.Errorf("mknnnm and mknntp are required")
	}

	file, err := c.FormFile("image")
	if err == nil {
		image, err := readImageFromHeader(file)
		if err != nil {
			return RegisterInputMakanan{}, err
		}

		input.MknnIg = image
	}

	return input, nil
}

func readImageFile(c *gin.Context, fieldName string) ([]byte, error) {
	file, err := c.FormFile(fieldName)
	if err != nil {
		return nil, fmt.Errorf("image file is required")
	}

	return readImageFromHeader(file)
}

func readImageFromHeader(file *multipart.FileHeader) ([]byte, error) {
	if file.Size > maxMakananImageSize {
		return nil, fmt.Errorf("image max size is 5MB")
	}

	openedFile, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer openedFile.Close()

	image, err := io.ReadAll(openedFile)
	if err != nil {
		return nil, err
	}

	contentType := http.DetectContentType(image)
	if !strings.HasPrefix(contentType, "image/") {
		return nil, fmt.Errorf("file must be an image")
	}

	return image, nil
}

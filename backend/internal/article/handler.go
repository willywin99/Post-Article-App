package article

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

type Handler struct {
    Repo *Repository
    DB   *gorm.DB
}

func NewHandler(db *gorm.DB) *Handler {
    return &Handler{Repo: NewRepository(db), DB: db}
}

func (h *Handler) Create(c *gin.Context) {
    var req Article
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := h.Repo.Create(&req); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, req)
}

func (h *Handler) List(c *gin.Context) {
    limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
    offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))
    status := c.Query("status")
    res, err := h.Repo.List(limit, offset, status)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, res)
}

func (h *Handler) Get(c *gin.Context) {
    idStr := c.Param("id")
    id, _ := strconv.Atoi(idStr)
    a, err := h.Repo.GetByID(uint(id))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, a)
}

func (h *Handler) Update(c *gin.Context) {
    idStr := c.Param("id")
    id, _ := strconv.Atoi(idStr)
    var payload Article
    if err := c.ShouldBindJSON(&payload); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    payload.ID = uint(id)
    if err := h.Repo.Update(&payload); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "updated"})
}

func (h *Handler) Delete(c *gin.Context) {
    idStr := c.Param("id")
    id, _ := strconv.Atoi(idStr)
    if err := h.Repo.Delete(uint(id)); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}

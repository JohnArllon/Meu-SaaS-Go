package technicians

import (
	"net/http"

	"github.com/JohnArllon/Meu-SaaS-Go/internal/domain"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	repo *Repository
}

func NewHandler(repo *Repository) *Handler {
	return &Handler{repo}
}

func (h *Handler) CreateTechnician(c *gin.Context) {
	var t domain.Technician

	// Tenta transformar o JSON que vem da internet na nossa Struct
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}
	// Chama o repositório que você já testou e sabe que funciona!
	if err := h.repo.Create(c.Request.Context(), t); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao salvar no banco"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Técnico criado comn sucesso!"})
}

func (h *Handler) ListTechnicians(c *gin.Context) {
	list, err := h.repo.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar técnicos"})
		return
	}
	c.JSON(http.StatusOK, list)
}

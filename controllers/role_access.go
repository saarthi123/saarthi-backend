package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/saarthi123/saarthi-backend/config"
	"github.com/saarthi123/saarthi-backend/models"
	
)

// RolePermissionRequest represents the incoming JSON for updating permissions
type RolePermissionRequest struct {
	RoleID        string   `json:"roleId"`
	PermissionIDs []string `json:"permissionIds"`
}

// ─────────────────────────────────────────────────────────────
// ✅ Get All Roles with Their Permissions
// ─────────────────────────────────────────────────────────────
func GetRoles(c *gin.Context) {
	var roles []models.Role

	if err := config.DB.Preload("Permissions").Find(&roles).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch roles"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"roles": roles})
}

// ─────────────────────────────────────────────────────────────
// ✅ Dynamically Update Role Permissions
// ─────────────────────────────────────────────────────────────
func UpdateRolePermissions(c *gin.Context) {
	var req RolePermissionRequest

	if err := c.ShouldBindJSON(&req); err != nil || req.RoleID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role or permission data"})
		return
	}

	var role models.Role
	if err := config.DB.Preload("Permissions").First(&role, "id = ?", req.RoleID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Role not found"})
		return
	}

	// Fetch the new permissions
	var newPermissions []models.Permission
	if len(req.PermissionIDs) > 0 {
		if err := config.DB.Where("id IN ?", req.PermissionIDs).Find(&newPermissions).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching permissions"})
			return
		}
	}

	// Dynamically replace role-permission relationships
	if err := config.DB.Model(&role).Association("Permissions").Replace(&newPermissions); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update role permissions"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":     "Role permissions updated successfully",
		"role":        role.Name,
		"permissions": newPermissions,
	})
}


// func GetAllRoles(c *gin.Context) {
// 	var roles []models.Role
// 	if err := config.DB.Find(&roles).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch roles"})
// 		return
// 	}
// c.JSON(http.StatusOK, gin.H{"roles": roles}) // ✅ Also valid
// }

// func GetRolePermissions(c *gin.Context) {
// 	roleID := c.Param("id")
// 	var role models.Role
// 	if err := config.DB.Preload("Permissions").First(&role, "id = ?", roleID).Error; err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "Role not found"})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"permissions": role.Permissions})
// }

// controllers/role_controller.go
func GetAllRoles(c *gin.Context) {
	var roles []models.Role
	if err := config.DB.Find(&roles).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch roles"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"roles": roles})
}

func GetRolePermissions(c *gin.Context) {
	roleID := c.Param("role")
	var role models.Role
	if err := config.DB.Preload("Permissions").First(&role, "id = ?", roleID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Role not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"role": role})
}

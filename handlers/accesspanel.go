package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

// RolePermissions represents a role and its associated permissions
type RolePermissions struct {
    Role        string   `json:"role"`
    Permissions []string `json:"permissions"`
}

// Simulated in-memory store (replace with DB in production)
var rolePermissionsMap = map[string][]string{
    "admin":   {"create_user", "delete_user", "view_reports"},
    "student": {"view_courses", "submit_assignments"},
    "teacher": {"create_course", "grade_assignments"},
}

// GET /roles
func getAllRoles(c *gin.Context) {
    roles := make([]RolePermissions, 0, len(rolePermissionsMap))
    for role, perms := range rolePermissionsMap {
        roles = append(roles, RolePermissions{Role: role, Permissions: perms})
    }

    // ✅ Corrected: removed trailing comma after c.JSON
    c.JSON(http.StatusOK, gin.H{
        "data": roles,
    })
}

// GET /roles/:role
func getRolePermissions(c *gin.Context) {
    role := c.Param("role")
    perms, exists := rolePermissionsMap[role]
    if !exists {
        c.JSON(http.StatusNotFound, gin.H{"error": "role not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "data": RolePermissions{
            Role:        role,
            Permissions: perms,
        },
    })
}

type updatePermissionsRequest struct {
    Permissions []string `json:"permissions"`
}

// PUT /roles/:role
func updateRolePermissions(c *gin.Context) {
    role := c.Param("role")
    _, exists := rolePermissionsMap[role]
    if !exists {
        c.JSON(http.StatusNotFound, gin.H{"error": "role not found"})
        return
    }

    var req updatePermissionsRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
        return
    }

    rolePermissionsMap[role] = req.Permissions
    c.JSON(http.StatusOK, gin.H{
        "message":     "permissions updated",
        "role":        role,
        "permissions": req.Permissions,
    })
}

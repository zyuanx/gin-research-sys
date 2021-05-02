package controllers

import (
	"gin-research-sys/controllers/req"
	"gin-research-sys/controllers/res"
	"gin-research-sys/models"
	"gin-research-sys/services"
	"github.com/gin-gonic/gin"
	"strconv"
)

var roleServices = services.NewRoleService()

type IRoleController interface {
	List(ctx *gin.Context)
	Create(ctx *gin.Context)
	Retrieve(ctx *gin.Context)
	Update(ctx *gin.Context)
	PartialUpdate(ctx *gin.Context)
	Destroy(ctx *gin.Context)
}
type RoleController struct{}

func NewRoleController() RoleController {
	return RoleController{}
}

// List
// @Summary list all role
// @Description
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Param pagination query req.PaginationQuery true "PaginationQuery"
// @Success 200 {object} res.Result "成功后返回值"
// @Failure 400 {object} res.Fail
// @Router /api/role [get]
func (r RoleController) List(ctx *gin.Context) {
	var _ []models.Role
	pg := req.PaginationQuery{}
	if err := ctx.ShouldBindQuery(&pg); err != nil {
		res.Success(ctx, nil, err.Error())
		return
	}
	var roleList []res.RoleListResponse
	var total int64
	err := roleServices.List(pg.Page, pg.Size, &roleList, &total)
	if err != nil {
		res.Success(ctx, nil, err.Error())
	}
	res.Success(ctx, gin.H{
		"page":    pg.Page,
		"size":    pg.Size,
		"results": roleList,
		"total":   total,
	}, "")
}

// Create
// @Summary create a new role
// @Description get string by ID
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Param role body req.CreateRoleValidate true "角色"
// @Success 200 {object} res.Result "成功后返回值"
// @Router /api/role [post]
func (r RoleController) Create(ctx *gin.Context) {
	createRoleValidate := req.CreateRoleValidate{}
	if err := ctx.ShouldBindJSON(&createRoleValidate); err != nil {
		res.Fail(ctx, gin.H{}, err.Error())
		return
	}
	role := models.Role{
		Title: createRoleValidate.Title,
		Desc:  createRoleValidate.Desc,
	}
	if err := roleServices.Create(&role); err != nil {
		res.Fail(ctx, gin.H{}, err.Error())
		return
	}
	res.Success(ctx, gin.H{"role": role}, "")
}

func (r RoleController) Retrieve(ctx *gin.Context) {
	idString := ctx.Param("id")
	role := models.Role{}
	id, err := strconv.Atoi(idString)
	if err != nil {
		res.Fail(ctx, gin.H{}, err.Error())
		return
	}
	err = roleServices.Retrieve(&role, id)
	if err != nil {
		res.Fail(ctx, gin.H{}, err.Error())
		return
	}
	res.Success(ctx, gin.H{"role": role}, "")
}

func (r RoleController) Update(ctx *gin.Context) {
	panic("implement me")
}

func (r RoleController) PartialUpdate(ctx *gin.Context) {
	panic("implement me")
}

func (r RoleController) Destroy(ctx *gin.Context) {
	panic("implement me")
}

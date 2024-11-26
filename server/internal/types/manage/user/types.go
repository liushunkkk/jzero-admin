// Code generated by jzero. DO NOT EDIT.
package manageuser

import (
	"time"
)

var (
	_ = time.Now()
)

type AddRequest struct {
	Username   string   `json:"username"`
	UserGender string   `json:"userGender,default=1"`
	NickName   string   `json:"nickName"`
	UserPhone  string   `json:"userPhone"`
	UserEmail  string   `json:"userEmail"`
	Status     string   `json:"status,default=1"`
	Password   string   `json:"password"`
	UserRoles  []string `json:"userRoles"`
}

type AddResponse struct {
}

type DeleteRequest struct {
	Ids []uint64 `json:"ids"`
}

type DeleteResponse struct {
}

type EditRequest struct {
	Id         uint64   `json:"id"`
	Username   string   `json:"username"`
	UserGender string   `json:"userGender"`
	NickName   string   `json:"nickName"`
	UserPhone  string   `json:"userPhone"`
	UserEmail  string   `json:"userEmail"`
	Status     string   `json:"status"`
	UserRoles  []string `json:"userRoles"`
}

type EditResponse struct {
}

type Empty struct {
}

type ListRequest struct {
	PageRequest
	Username   string `form:"username,optional"`
	UserGender string `form:"userGender,optional"`
	NickName   string `form:"nickName,optional"`
	UserPhone  string `form:"userPhone,optional"`
	UserEmail  string `form:"userEmail,optional"`
	Status     string `form:"status,optional"`
}

type ListResponse struct {
	PageResponse
	Records []ManageUser `json:"records"`
}

type ManageUser struct {
	Id         uint64   `json:"id"`
	Username   string   `json:"username"`
	UserGender string   `json:"userGender"`
	NickName   string   `json:"nickName"`
	UserPhone  string   `json:"userPhone"`
	UserEmail  string   `json:"userEmail"`
	UserRoles  []string `json:"userRoles"`
	Status     string   `json:"status"`
	CreateTime string   `json:"createTime"`
	UpdateTime string   `json:"updateTime"`
}

type PageRequest struct {
	Current int `form:"current,default=1,optional"`
	Size    int `form:"size,default=10,optional"`
}

type PageResponse struct {
	Current int   `json:"current"`
	Size    int   `json:"size"`
	Total   int64 `json:"total"`
}
// Code generated by jzero. DO NOT EDIT.

package handler

import (
	"net/http"
	"strings"

	casbinutil "github.com/casbin/casbin/v2/util"
)

var routesCodesMap = map[string]string{
	"POST:/api/auth/code-login":          "auth:codeLogin",
	"POST:/api/auth/error":               "auth:error",
	"GET:/api/auth/getUserInfo":          "auth:getUserInfo",
	"POST:/api/auth/pwd-login":           "auth:pwdLogin",
	"POST:/api/auth/refreshToken":        "auth:refreshToken",
	"POST:/api/auth/register":            "auth:register",
	"POST:/api/auth/resetPassword":       "auth:resetPassword",
	"GET:/api/auth/sendVerificationCode": "auth:sendVerificationCode",
	"POST:/api/manage/addMenu":           "manage:menu:add",
	"POST:/api/manage/addRole":           "manage:role:add",
	"POST:/api/manage/addUser":           "manage:user:add",
	"POST:/api/manage/deleteMenu":        "manage:menu:delete",
	"POST:/api/manage/deleteRole":        "manage:role:delete",
	"POST:/api/manage/deleteUser":        "manage:user:delete",
	"POST:/api/manage/editMenu":          "manage:menu:edit",
	"POST:/api/manage/editRole":          "manage:role:edit",
	"POST:/api/manage/editUser":          "manage:user:edit",
	"GET:/api/manage/getAllPages":        "manage:menu:getAllPages",
	"GET:/api/manage/getAllRoles":        "manage:role:getAll",
	"GET:/api/manage/getMenuList/v2":     "manage:menu:list",
	"GET:/api/manage/getMenuTree":        "manage:menu:tree",
	"GET:/api/manage/getRoleHome":        "manage:role:getHome",
	"GET:/api/manage/getRoleList":        "manage:role:list",
	"GET:/api/manage/getRoleMenus":       "manage:role:getMenus",
	"GET:/api/manage/getUserList":        "manage:user:list",
	"POST:/api/manage/setRoleMenus":      "manage:role:setMenus",
	"POST:/api/manage/updateRoleHome":    "manage:role:updateHome",
	"GET:/api/route/getConstantRoutes":   "route:getConstantRoutes",
	"GET:/api/route/getUserRoutes":       "route:getUserRoutes",
	"GET:/api/route/isRouteExist":        "route:isRouteExist",
	"GET:/api/version":                   "version:get",
}

func Route2Code(r *http.Request) string {
	for k, v := range routesCodesMap {
		if splits := strings.Split(k, ":"); len(splits) >= 2 && splits[0] == strings.ToUpper(r.Method) {
			if casbinutil.KeyMatch2(r.URL.Path, strings.Join(splits[1:], ":")) {
				return v
			}
		}
	}
	return "unknown_code"
}

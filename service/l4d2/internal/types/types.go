// Code generated by goctl. DO NOT EDIT.
package types

type ServerInfo struct {
	Host    string `json:"host"`
	Version string `json:"version"`
	Listen  string `json:"listen"`
	Public  string `json:"public"`
	Os      string `json:"os"`
	Map     string `json:"map"`
	Player  string `json:"player"`
}

type MapFile struct {
	Id        int64  `json:"id"`
	FileName  string `json:"fileName"`
	GroupID   int64  `json:"groupId"`
	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
}

type MapGroup struct {
	Id           int64  `json:"id"`
	Name         string `json:"name"`
	StartName    string `json:"startName"`
	LastPlayTime int64  `json:"lastPlayTime"`
	CreatedAt    int64  `json:"createdAt"`
	UpdatedAt    int64  `json:"updatedAt"`
}

type User struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Right    int    `json:"right"`
}

type Plugin struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Enabled     bool   `json:"enabled"`
}

type CommandTemplate struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Params      []string `json:"params"`
	Command     string   `json:"command"`
}

type GetServerStatusRequest struct {
}

type GetServerStatusResponse struct {
	Code int        `json:"code"`
	Msg  string     `json:"msg"`
	Data ServerInfo `json:"data"`
}

type GetServerPluginsRequest struct {
}

type GetServerPluginsResponse struct {
	Code int      `json:"code"`
	Msg  string   `json:"msg"`
	Data []Plugin `json:"data"`
}

type GetMapGroupRequest struct {
}

type GetMapGroupResponse struct {
	Code int        `json:"code"`
	Msg  string     `json:"msg"`
	Data []MapGroup `json:"data"`
}

type GetMapGroupByIDRequest struct {
	ID int64 `path:"id"`
}

type GetMapGroupByIDResponse struct {
	Code int       `json:"code"`
	Msg  string    `json:"msg"`
	Data *MapGroup `json:"data"`
}

type GetMapFilesByGroupIDRequest struct {
	ID int64 `path:"id"`
}

type GetMapFilesByGroupIDResponse struct {
	Code int       `json:"code"`
	Msg  string    `json:"msg"`
	Data []MapFile `json:"data"`
}

type DeleteMapGroupByIDRequest struct {
	ID int64 `path:"id"`
}

type DeleteMapGroupByIDResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type UploadMapFileRequest struct {
	GroupID int64 `path:"groupID"`
}

type UploadMapFileResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type ExecuteCommandRequest struct {
	Command string `json:"command"`
}

type ExecuteCommandResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data string `json:"data"`
}

type EnablePluginRequest struct {
	PluginNames []string `json:"pluginNames"`
}

type EnablePluginResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type DisablePluginRequest struct {
	PluginNames []string `json:"pluginNames"`
}

type DisablePluginResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type GetCommandTemplateReqeust struct {
}

type GetCommandTemplateResponse struct {
	Code int               `json:"code"`
	Msg  string            `json:"msg"`
	Data []CommandTemplate `json:"data"`
}

type LoginReplyData struct {
	Id           int64  `json:"id"`
	AccessToken  string `json:"accessToken"`
	AccessExpire int64  `json:"accessExpire"`
	RefreshAfter int64  `json:"refreshAfter"`
}

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginReply struct {
	Code int             `json:"code"`
	Msg  string          `json:"msg"`
	Data *LoginReplyData `json:"data,optional"`
}

package types

type ApiUrl struct {
	Url    string
	Method string
}

const (
	GetVersion    = "version"
	UploadBom     = "upload_bom"
	CreateProject = "create_project"
	ListProjects  = "list_projects"
	RemoveProject = "remove_project"
	CreateTeam    = "create_team"
	ListTeams     = "list_teams"
	GenTeamKey    = "gen_team_key"
	SelfTeam      = "self_team"
	CheckProcess  = "check_process"
)

var ApiUrls = map[string]ApiUrl{
	GetVersion: {
		Url:    "/api/version",
		Method: "GET",
	},
	UploadBom: {
		Url:    "/api/v1/bom",
		Method: "PUT",
	},
	CreateProject: {
		Url:    "/api/v1/project",
		Method: "PUT",
	},
	ListProjects: {
		Url:    "/api/v1/project",
		Method: "GET",
	},
	RemoveProject: {
		Url:    "/api/v1/project/{id}",
		Method: "DELETE",
	},
	CreateTeam: {
		Url:    "/api/v1/team",
		Method: "PUT",
	},
	ListTeams: {
		Url:    "/api/v1/team/visible",
		Method: "GET",
	},
	GenTeamKey: {
		Url:    "/api/v1/team/{id}/key",
		Method: "PUT",
	},
	SelfTeam: {
		Url:    "/api/v1/team/self",
		Method: "GET",
	},
	CheckProcess: {
		Url:    "/api/v1/event/token/{id}",
		Method: "GET",
	},
}

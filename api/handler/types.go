package handler

type ServiceInfo struct {
	ServiceName  string `json:"service_name"`
	Description  string `json:"description"`
	Language     string `json:"language"`
	TmplExpandId string `json:"tmpl_expand_id"`
}

type Pager struct {
	PageNumber int64 `json:"page_number"`
	PageSize   int64 `json:"page_size"`
	Total      int64 `json:"total"`
}

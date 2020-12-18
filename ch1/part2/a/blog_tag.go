package main

type BlogTagModel struct {
	Id int32 `json:"id"` // id

	Name string `json:"name"` // 标签名称

	CreatedOn int32 `json:"created_on"` // 创建时间

	CreatedBy string `json:"created_by"` // 创建人

	ModifiedOn int32 `json:"modified_on"` // 修改时间

	ModifiedBy string `json:"modified_by"` // 修改人

	DeleteOn int32 `json:"delete_on"` // 删除时间

	IsDel int8 `json:"is_del"` // 是否删除，0 未删除 1 已删除

	State int8 `json:"state"` // 状态，0 禁用 1 启用

}

func (model *BlogTagModel) TableName() string {
	return "blog_tag"
}

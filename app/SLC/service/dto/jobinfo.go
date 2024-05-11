package dto

import (
	"SmartLinkProject/app/SLC/models"

	"SmartLinkProject/common/dto"
	common "SmartLinkProject/common/models"
)

type JobinfoGetPageReq struct {
	dto.Pagination `search:"-"`
	JobId          int `form:"jobId" search:"type:exact;column:job_id;table:jobinfo" comment:"任务ID"`
	DeviceId       int `form:"deviceId" search:"type:exact;column:device_id;table:jobinfo" comment:"设备ID"`
	TypeId         int `form:"typeId" search:"type:exact;column:type_id;table:jobinfo" comment:"类型ID"`
	LocId          int `form:"locId" search:"type:exact;column:loc_id;table:jobinfo" comment:"位置ID"`

	// Start           int    `form:"start"`
	// Length          int    `form:"length"`
	// Author          string `form:"author"`          // 作者
	// TriggerStatus   string `form:"triggerStatus"`   // 触发状态
	// JobDesc         string `form:"jobDesc"`         // 任务描述
	// ExecutorHandler string `form:"executorHandler"` // 执行器

	JLocJoin    `search:"type:left;on:loc_id:loc_id;table:jobinfo;join:sys_loc"`
	JDeviceJoin `search:"type:left;on:device_id:device_id;table:jobinfo;join:device"`
	JTypeJoin   `search:"type:left;on:type_id:type_id;table:jobinfo;join:device_type"`
	JobinfoOrder
}

type JobinfoOrder struct {
	JobIdOrder    string `search:"type:order;column:job_id;table:jobinfo" form:"jobIdOrder"`
	DeviceIdOrder string `search:"type:order;column:device_id;table:jobinfo" form:"deviceIdOrder"`
	TypeIdOrder   string `search:"type:order;column:type_id;table:jobinfo" form:"typeIdOrder"`
	LocIdOrder    string `search:"type:order;column:loc_id;table:jobinfo" form:"locIdOrder"`
}

type JLocJoin struct {
	LocId string `search:"type:contains;column:loc_path;table:sys_loc" form:"locId"`
}

type JDeviceJoin struct {
	DeviceId string `search:"type:exact;column:device_id;table:device" form:"deviceId"`
}

type JTypeJoin struct {
	TypeId string `search:"type:contains;column:type_path;table:device_type" form:"typeId"`
}

func (m *JobinfoGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type JobinfoInsertReq struct {
	JobId    int `json:"jobId" comment:"任务ID"` // 任务ID
	DeviceId int `json:"deviceId" comment:"设备ID" vd:"$>0" `
	TypeId   int `json:"typeId" comment:"类型ID" vd:"$>0"` // 设备类型ID
	LocId    int `json:"locId" comment:"位置ID" vd:"$>0"`  // 设备位置ID
}

// type JobinfoInsertRemoteReq struct {
// 	Author                  string `form:"author"` // 作者
// 	AlarmEmail              string `form:"alarmEmail"`
// 	TriggerStatus           string `form:"triggerStatus"` // 触发状态
// 	JobDesc                 string `form:"jobDesc"`       // 任务描述
// 	ScheduleConf            string `form:"scheduleConf"`
// 	CronGen_display         string `form:"cronGen_display"`
// 	Schedule_conf_CRON      string `form:"schedule_conf_CRON"`
// 	Schedule_conf_FIX_RATE  string `form:"schedule_conf_FIX_RATE"`
// 	Schedule_conf_FIX_DELAY string `form:"schedule_conf_FIX_DELAY"`
// 	GlueType                string `form:"glueType"`
// 	ExecutorHandler         string `form:"executorHandler"` // 执行器
// 	ExecutorParam           string `form:"executorParam"`
// 	ExecutorRouteStrategy   string `form:"executorRouteStrategy"`
// 	ChildJobId              string `form:"childJobId"`
// 	MisfireStrategy         string `form:"misfireStrategy"`
// 	ExecutorBlockStrategy   string `form:"executorBlockStrategy"`
// 	ExecutorTimeout         string `form:"executorTimeout"`
// 	ExecutorFailRetryCount  string `form:"executorFailRetryCount"`
// 	GlueRemark              string `form:"glueRemark"`
// 	GlueSource              string `form:"glueSource"`
// }

func (d *JobinfoInsertReq) Generate(model *models.Jobinfo) {
	if d.JobId != 0 {
		model.JobId = d.JobId
	}
	model.DeviceId = d.DeviceId
	model.TypeId = d.TypeId
	model.LocId = d.LocId
}

func (s *JobinfoInsertReq) GetId() interface{} {
	return s.JobId
}

type JobinfoUpdateReq struct {
	JobId    int `json:"jobId" comment:"任务ID"` // 任务ID
	DeviceId int `json:"deviceId" comment:"设备ID" vd:"$>0" `
	TypeId   int `json:"typeId" comment:"类型ID" vd:"$>0"` // 设备类型ID
	LocId    int `json:"locId" comment:"位置ID" vd:"$>0"`  // 设备位置ID
}

func (d *JobinfoUpdateReq) Generate(model *models.Jobinfo) {
	if d.JobId != 0 {
		model.JobId = d.JobId
	}
	model.TypeId = d.TypeId
	model.LocId = d.LocId
	model.DeviceId = d.DeviceId
}

func (d *JobinfoUpdateReq) GetId() interface{} {
	return d.JobId
}

type JobinfoById struct {
	dto.ObjectById
}

func (s *JobinfoById) GetId() interface{} {
	if len(s.Ids) > 0 {
		s.Ids = append(s.Ids, s.Id)
		return s.Ids
	}
	return s.Id
}

func (s *JobinfoById) GenerateM() (common.ActiveRecord, error) {
	return &models.Jobinfo{}, nil
}

package repository

import (
	"fmt"
	"task/internal/service"
)

type Task struct {
	TaskID    int64 `gorm:"primarykey"` // id
	UserID    int64 `gorm:"index"`      // 用户id
	Status    int   `gorm:"default:0"`
	Title     string
	Content   string `gorm:"type:longtext"`
	StartTime int64
	EndTime   int64
}

func (*Task) CreateTask(req *service.TaskRequest) (err error) {
	task := &Task{
		UserID:    req.UserID,
		Title:     req.Title,
		Content:   req.Content,
		Status:    int(req.Status),
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
	}

	if err = DB.Create(&task).Error; err != nil {
		fmt.Println("Insert Task Error:" + err.Error())
	}

	return
}

func (*Task) ShowTask(req *service.TaskRequest) (taskList []Task, err error) {
	if req.PageSize == 0 {
		req.PageSize = 15
	}
	err = DB.Model(&Task{}).Where("user_id=?", req.UserID).Limit(int(req.PageSize)).Offset(int((req.PageNum - 1) * req.PageSize)).Find(&taskList).Error

	if err != nil {
		return nil, err
	}

	return taskList, nil
}

func (*Task) SearchTask(req *service.TaskRequest) (taskList []Task, err error) {
	if req.PageSize == 0 {
		req.PageSize = 15
	}

	err = DB.Model(&Task{}).Where("user_id=?", req.UserID).
		Where("title LIKE ? OR content LIKE ?", "%"+req.Info+"%", "%"+req.Info+"%").
		Limit(int(req.PageSize)).Offset(int((req.PageNum - 1) * req.PageSize)).
		Find(&taskList).Error

	if err != nil {
		return nil, err
	}

	return taskList, nil
}
func (*Task) UpdateTask(req *service.TaskRequest) (err error) {
	taskUpdateMap := make(map[string]interface{})
	taskUpdateMap["title"] = req.Title
	taskUpdateMap["content"] = req.Content
	taskUpdateMap["status"] = int(req.Status)
	taskUpdateMap["start_time"] = req.StartTime
	taskUpdateMap["end_time"] = req.EndTime
	err = DB.Model(&Task{}).
		Where("task_id=?", req.TaskID).Updates(&taskUpdateMap).Error

	return
}

func (*Task) UpdateAllTask(req *service.TaskRequest) (err error) {
	var tasks []Task
	err = DB.Model(&Task{}).Where("user_id=?", req.UserID).Find(&tasks).Error
	if err != nil {
		return err
	}
	for _, task := range tasks {
		task.Status = int(req.Status)
		DB.Save(&task)
	}
	return nil
}
func (*Task) DeleteTask(req *service.TaskRequest) (err error) {
	return DB.Model(&Task{}).Where("task_id = ? AND user_id = ?", req.TaskID, req.UserID).Delete(Task{}).Error
}
func (*Task) DeleteAllTask(req *service.TaskRequest) (err error) {
	if req.Status == -1 {
		err = DB.Model(&Task{}).Where("user_id = ?", req.UserID).Delete([]Task{}).Error
	} else {
		err = DB.Model(&Task{}).Where("user_id = ?", req.UserID).Where("status=?", req.Status).Delete([]Task{}).Error
	}
	return
}

// BuildTask 视图返回
func BuildTask(item Task) *service.TaskModel {
	taskModel := service.TaskModel{
		TaskID:    item.TaskID,
		UserID:    item.UserID,
		Status:    int64(item.Status),
		Title:     item.Title,
		Content:   item.Content,
		StartTime: item.StartTime,
		EndTime:   item.EndTime,
	}
	return &taskModel
}

func BuildTasks(items []Task) (taskList []*service.TaskModel) {
	for _, item := range items {
		task := BuildTask(item)
		taskList = append(taskList, task)
	}
	return taskList
}

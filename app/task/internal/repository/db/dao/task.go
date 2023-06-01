package dao

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"grpc-todolist/app/task/internal/repository/db/model"
	taskPb "grpc-todolist/idl/pb/task"
)

type TaskDao struct {
	*gorm.DB
}

func NewTaskDao(ctx context.Context) *TaskDao {
	return &TaskDao{NewDBClient(ctx)}
}

func (dao *TaskDao) CreateTask(req *taskPb.TaskRequest) (err error) {
	task := &model.Task{
		UserID:    req.UserID,
		Title:     req.Title,
		Content:   req.Content,
		Status:    int(req.Status),
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
	}

	if err = dao.Create(&task).Error; err != nil {
		fmt.Println("Insert Task Error:" + err.Error())
	}

	return
}

func (dao *TaskDao) ShowTask(req *taskPb.TaskRequest) (taskList []model.Task, err error) {
	if req.PageSize == 0 {
		req.PageSize = 15
	}
	err = dao.Model(&model.Task{}).Where("user_id=?", req.UserID).Limit(int(req.PageSize)).Offset(int((req.PageNum - 1) * req.PageSize)).Find(&taskList).Error

	if err != nil {
		return nil, err
	}

	return taskList, nil
}

func (dao *TaskDao) SearchTask(req *taskPb.TaskRequest) (taskList []model.Task, err error) {
	if req.PageSize == 0 {
		req.PageSize = 15
	}

	err = dao.Model(&model.Task{}).Where("user_id=?", req.UserID).
		Where("title LIKE ? OR content LIKE ?", "%"+req.Info+"%", "%"+req.Info+"%").
		Limit(int(req.PageSize)).Offset(int((req.PageNum - 1) * req.PageSize)).
		Find(&taskList).Error

	if err != nil {
		return nil, err
	}

	return taskList, nil
}
func (dao *TaskDao) UpdateTask(req *taskPb.TaskRequest) (err error) {
	taskUpdateMap := make(map[string]interface{})
	taskUpdateMap["title"] = req.Title
	taskUpdateMap["content"] = req.Content
	taskUpdateMap["status"] = int(req.Status)
	taskUpdateMap["start_time"] = req.StartTime
	taskUpdateMap["end_time"] = req.EndTime
	err = dao.Model(&model.Task{}).
		Where("task_id=?", req.TaskID).Updates(&taskUpdateMap).Error

	return
}

func (dao *TaskDao) UpdateAllTask(req *taskPb.TaskRequest) (err error) {
	var tasks []model.Task
	err = dao.Model(&model.Task{}).Where("user_id=?", req.UserID).Find(&tasks).Error
	if err != nil {
		return err
	}
	for _, task := range tasks {
		task.Status = int(req.Status)
		dao.Save(&task)
	}
	return nil
}
func (dao *TaskDao) DeleteTask(req *taskPb.TaskRequest) (err error) {
	return dao.Model(&model.Task{}).Where("task_id = ? AND user_id = ?", req.TaskID, req.UserID).Delete(model.Task{}).Error
}
func (dao *TaskDao) DeleteAllTask(req *taskPb.TaskRequest) (err error) {
	if req.Status == -1 {
		err = dao.Model(&model.Task{}).Where("user_id = ?", req.UserID).Delete([]model.Task{}).Error
	} else {
		err = dao.Model(&model.Task{}).Where("user_id = ?", req.UserID).Where("status=?", req.Status).Delete([]model.Task{}).Error
	}
	return
}

// BuildTask 视图返回
func BuildTask(item model.Task) *taskPb.TaskModel {
	taskModel := taskPb.TaskModel{
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

func BuildTasks(items []model.Task) (taskList []*taskPb.TaskModel) {
	for _, item := range items {
		task := BuildTask(item)
		taskList = append(taskList, task)
	}
	return taskList
}

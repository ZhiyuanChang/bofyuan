package game

type TaskInfo struct {
	TaskId int
	state  int
}

type ModUniqueTask struct {
	MyTaskInfo map[int]*TaskInfo
}

func (self *ModUniqueTask) IsTaskFInish(taskId int) bool {
	task, ok := self.MyTaskInfo[taskId]
	if !ok {
		return false
	}
	return task.state == TASK_STATE_FINISH
}

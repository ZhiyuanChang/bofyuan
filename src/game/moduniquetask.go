package game

import "sync"

type TaskInfo struct {
	TaskId int
	state  int
}

type ModUniqueTask struct {
	MyTaskInfo map[int]*TaskInfo //保证数据的独占性 保证数据安全（保证map的原子性） 当map出现同时读写的时候 会报错fatal error，所以要加sync.map
	Locker     sync.RWMutex
}

func NewModUniqueTask() *ModUniqueTask {
	return &ModUniqueTask{
		MyTaskInfo: map[int]*TaskInfo{},
		Locker:     sync.RWMutex{},
	}
}

func (self *ModUniqueTask) IsTaskFInish(taskId int) bool {
	if taskId == 10001 || taskId == 10002 {
		return true
	}
	task, ok := self.MyTaskInfo[taskId]
	if !ok {
		return false
	}
	return task.state == TASK_STATE_FINISH
}

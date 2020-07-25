package models

// Task 被认为是有状态的, 所以某个任务应该以 struct 定义.
type Task interface {
	Func()
}

// MockTask 用于测试其 Func() 能否传递给 Pool.Submit()
type MockTask struct {
}

func (mt *MockTask) Func() {

}

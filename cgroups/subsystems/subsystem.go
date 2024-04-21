package subsystems

// ResourceConfig 内存，CPU限制
type ResourceConfig struct {
	MemoryLimit string
	CpuShare    string
	CpuSet      string
}

type Subsystem interface {
	// Name 返回名字
	Name() string
	// Set 设置cgroup中的限制
	Set(path string, res *ResourceConfig) error
	// Apply 添加进程
	Apply(path string, pid int) error
	// Remove 移除
	Remove(path string) error
}

var (
	SubsystemsIns = []Subsystem{
		&CpusetSubSystem{},
		&MemorySubSystem{},
		&CpuSubSystem{},
	}
)

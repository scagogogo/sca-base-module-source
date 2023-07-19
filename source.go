package sca_base_module_source

import "context"

// Source 输入源的定义
type Source[Config any] interface {

	// Init 可以在此进行一些初始化
	Init(ctx context.Context, config Config) error

	// Run 执行源输入任务
	Run(ctx context.Context) error

	// GetConfig 获取source的配置
	GetConfig() Config

	// Close 拉取完毕，执行一些回收
	Close(ctx context.Context) error
}

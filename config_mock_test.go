// Code generated by github.com/efritz/go-mockgen 0.1.0; DO NOT EDIT.
// This file was generated by robots at
// 2019-06-17T20:19:16-05:00
// using the command
// $ go-mockgen -f github.com/go-nacelle/config -i Config -o config_mock_test.go

package config

import "sync"

// MockConfig is a mock implementation of the Config interface (from the
// package github.com/go-nacelle/config) used for unit testing.
type MockConfig struct {
	// AssetsFunc is an instance of a mock function object controlling the
	// behavior of the method Assets.
	AssetsFunc *ConfigAssetsFunc
	// DumpFunc is an instance of a mock function object controlling the
	// behavior of the method Dump.
	DumpFunc *ConfigDumpFunc
	// LoadFunc is an instance of a mock function object controlling the
	// behavior of the method Load.
	LoadFunc *ConfigLoadFunc
	// MustLoadFunc is an instance of a mock function object controlling the
	// behavior of the method MustLoad.
	MustLoadFunc *ConfigMustLoadFunc
}

// NewMockConfig creates a new mock of the Config interface. All methods
// return zero values for all results, unless overwritten.
func NewMockConfig() *MockConfig {
	return &MockConfig{
		AssetsFunc: &ConfigAssetsFunc{
			defaultHook: func() []string {
				return nil
			},
		},
		DumpFunc: &ConfigDumpFunc{
			defaultHook: func() map[string]string {
				return nil
			},
		},
		LoadFunc: &ConfigLoadFunc{
			defaultHook: func(interface{}, ...TagModifier) error {
				return nil
			},
		},
		MustLoadFunc: &ConfigMustLoadFunc{
			defaultHook: func(interface{}, ...TagModifier) {
				return
			},
		},
	}
}

// NewMockConfigFrom creates a new mock of the MockConfig interface. All
// methods delegate to the given implementation, unless overwritten.
func NewMockConfigFrom(i Config) *MockConfig {
	return &MockConfig{
		AssetsFunc: &ConfigAssetsFunc{
			defaultHook: i.Assets,
		},
		DumpFunc: &ConfigDumpFunc{
			defaultHook: i.Dump,
		},
		LoadFunc: &ConfigLoadFunc{
			defaultHook: i.Load,
		},
		MustLoadFunc: &ConfigMustLoadFunc{
			defaultHook: i.MustLoad,
		},
	}
}

// ConfigAssetsFunc describes the behavior when the Assets method of the
// parent MockConfig instance is invoked.
type ConfigAssetsFunc struct {
	defaultHook func() []string
	hooks       []func() []string
	history     []ConfigAssetsFuncCall
	mutex       sync.Mutex
}

// Assets delegates to the next hook function in the queue and stores the
// parameter and result values of this invocation.
func (m *MockConfig) Assets() []string {
	r0 := m.AssetsFunc.nextHook()()
	m.AssetsFunc.appendCall(ConfigAssetsFuncCall{r0})
	return r0
}

// SetDefaultHook sets function that is called when the Assets method of the
// parent MockConfig instance is invoked and the hook queue is empty.
func (f *ConfigAssetsFunc) SetDefaultHook(hook func() []string) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// Assets method of the parent MockConfig instance invokes the hook at the
// front of the queue and discards it. After the queue is empty, the default
// hook function is invoked for any future action.
func (f *ConfigAssetsFunc) PushHook(hook func() []string) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultDefaultHook with a function that returns
// the given values.
func (f *ConfigAssetsFunc) SetDefaultReturn(r0 []string) {
	f.SetDefaultHook(func() []string {
		return r0
	})
}

// PushReturn calls PushDefaultHook with a function that returns the given
// values.
func (f *ConfigAssetsFunc) PushReturn(r0 []string) {
	f.PushHook(func() []string {
		return r0
	})
}

func (f *ConfigAssetsFunc) nextHook() func() []string {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *ConfigAssetsFunc) appendCall(r0 ConfigAssetsFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of ConfigAssetsFuncCall objects describing the
// invocations of this function.
func (f *ConfigAssetsFunc) History() []ConfigAssetsFuncCall {
	f.mutex.Lock()
	history := make([]ConfigAssetsFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// ConfigAssetsFuncCall is an object that describes an invocation of method
// Assets on an instance of MockConfig.
type ConfigAssetsFuncCall struct {
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 []string
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c ConfigAssetsFuncCall) Args() []interface{} {
	return []interface{}{}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c ConfigAssetsFuncCall) Results() []interface{} {
	return []interface{}{c.Result0}
}

// ConfigDumpFunc describes the behavior when the Dump method of the parent
// MockConfig instance is invoked.
type ConfigDumpFunc struct {
	defaultHook func() map[string]string
	hooks       []func() map[string]string
	history     []ConfigDumpFuncCall
	mutex       sync.Mutex
}

// Dump delegates to the next hook function in the queue and stores the
// parameter and result values of this invocation.
func (m *MockConfig) Dump() map[string]string {
	r0 := m.DumpFunc.nextHook()()
	m.DumpFunc.appendCall(ConfigDumpFuncCall{r0})
	return r0
}

// SetDefaultHook sets function that is called when the Dump method of the
// parent MockConfig instance is invoked and the hook queue is empty.
func (f *ConfigDumpFunc) SetDefaultHook(hook func() map[string]string) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// Dump method of the parent MockConfig instance invokes the hook at the
// front of the queue and discards it. After the queue is empty, the default
// hook function is invoked for any future action.
func (f *ConfigDumpFunc) PushHook(hook func() map[string]string) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultDefaultHook with a function that returns
// the given values.
func (f *ConfigDumpFunc) SetDefaultReturn(r0 map[string]string) {
	f.SetDefaultHook(func() map[string]string {
		return r0
	})
}

// PushReturn calls PushDefaultHook with a function that returns the given
// values.
func (f *ConfigDumpFunc) PushReturn(r0 map[string]string) {
	f.PushHook(func() map[string]string {
		return r0
	})
}

func (f *ConfigDumpFunc) nextHook() func() map[string]string {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *ConfigDumpFunc) appendCall(r0 ConfigDumpFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of ConfigDumpFuncCall objects describing the
// invocations of this function.
func (f *ConfigDumpFunc) History() []ConfigDumpFuncCall {
	f.mutex.Lock()
	history := make([]ConfigDumpFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// ConfigDumpFuncCall is an object that describes an invocation of method
// Dump on an instance of MockConfig.
type ConfigDumpFuncCall struct {
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 map[string]string
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c ConfigDumpFuncCall) Args() []interface{} {
	return []interface{}{}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c ConfigDumpFuncCall) Results() []interface{} {
	return []interface{}{c.Result0}
}

// ConfigLoadFunc describes the behavior when the Load method of the parent
// MockConfig instance is invoked.
type ConfigLoadFunc struct {
	defaultHook func(interface{}, ...TagModifier) error
	hooks       []func(interface{}, ...TagModifier) error
	history     []ConfigLoadFuncCall
	mutex       sync.Mutex
}

// Load delegates to the next hook function in the queue and stores the
// parameter and result values of this invocation.
func (m *MockConfig) Load(v0 interface{}, v1 ...TagModifier) error {
	r0 := m.LoadFunc.nextHook()(v0, v1...)
	m.LoadFunc.appendCall(ConfigLoadFuncCall{v0, v1, r0})
	return r0
}

// SetDefaultHook sets function that is called when the Load method of the
// parent MockConfig instance is invoked and the hook queue is empty.
func (f *ConfigLoadFunc) SetDefaultHook(hook func(interface{}, ...TagModifier) error) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// Load method of the parent MockConfig instance invokes the hook at the
// front of the queue and discards it. After the queue is empty, the default
// hook function is invoked for any future action.
func (f *ConfigLoadFunc) PushHook(hook func(interface{}, ...TagModifier) error) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultDefaultHook with a function that returns
// the given values.
func (f *ConfigLoadFunc) SetDefaultReturn(r0 error) {
	f.SetDefaultHook(func(interface{}, ...TagModifier) error {
		return r0
	})
}

// PushReturn calls PushDefaultHook with a function that returns the given
// values.
func (f *ConfigLoadFunc) PushReturn(r0 error) {
	f.PushHook(func(interface{}, ...TagModifier) error {
		return r0
	})
}

func (f *ConfigLoadFunc) nextHook() func(interface{}, ...TagModifier) error {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *ConfigLoadFunc) appendCall(r0 ConfigLoadFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of ConfigLoadFuncCall objects describing the
// invocations of this function.
func (f *ConfigLoadFunc) History() []ConfigLoadFuncCall {
	f.mutex.Lock()
	history := make([]ConfigLoadFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// ConfigLoadFuncCall is an object that describes an invocation of method
// Load on an instance of MockConfig.
type ConfigLoadFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 interface{}
	// Arg1 is a slice containing the values of the variadic arguments
	// passed to this method invocation.
	Arg1 []TagModifier
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 error
}

// Args returns an interface slice containing the arguments of this
// invocation. The variadic slice argument is flattened in this array such
// that one positional argument and three variadic arguments would result in
// a slice of four, not two.
func (c ConfigLoadFuncCall) Args() []interface{} {
	trailing := []interface{}{}
	for _, val := range c.Arg1 {
		trailing = append(trailing, val)
	}

	return append([]interface{}{c.Arg0}, trailing...)
}

// Results returns an interface slice containing the results of this
// invocation.
func (c ConfigLoadFuncCall) Results() []interface{} {
	return []interface{}{c.Result0}
}

// ConfigMustLoadFunc describes the behavior when the MustLoad method of the
// parent MockConfig instance is invoked.
type ConfigMustLoadFunc struct {
	defaultHook func(interface{}, ...TagModifier)
	hooks       []func(interface{}, ...TagModifier)
	history     []ConfigMustLoadFuncCall
	mutex       sync.Mutex
}

// MustLoad delegates to the next hook function in the queue and stores the
// parameter and result values of this invocation.
func (m *MockConfig) MustLoad(v0 interface{}, v1 ...TagModifier) {
	m.MustLoadFunc.nextHook()(v0, v1...)
	m.MustLoadFunc.appendCall(ConfigMustLoadFuncCall{v0, v1})
	return
}

// SetDefaultHook sets function that is called when the MustLoad method of
// the parent MockConfig instance is invoked and the hook queue is empty.
func (f *ConfigMustLoadFunc) SetDefaultHook(hook func(interface{}, ...TagModifier)) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// MustLoad method of the parent MockConfig instance invokes the hook at the
// front of the queue and discards it. After the queue is empty, the default
// hook function is invoked for any future action.
func (f *ConfigMustLoadFunc) PushHook(hook func(interface{}, ...TagModifier)) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultDefaultHook with a function that returns
// the given values.
func (f *ConfigMustLoadFunc) SetDefaultReturn() {
	f.SetDefaultHook(func(interface{}, ...TagModifier) {
		return
	})
}

// PushReturn calls PushDefaultHook with a function that returns the given
// values.
func (f *ConfigMustLoadFunc) PushReturn() {
	f.PushHook(func(interface{}, ...TagModifier) {
		return
	})
}

func (f *ConfigMustLoadFunc) nextHook() func(interface{}, ...TagModifier) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *ConfigMustLoadFunc) appendCall(r0 ConfigMustLoadFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of ConfigMustLoadFuncCall objects describing
// the invocations of this function.
func (f *ConfigMustLoadFunc) History() []ConfigMustLoadFuncCall {
	f.mutex.Lock()
	history := make([]ConfigMustLoadFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// ConfigMustLoadFuncCall is an object that describes an invocation of
// method MustLoad on an instance of MockConfig.
type ConfigMustLoadFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 interface{}
	// Arg1 is a slice containing the values of the variadic arguments
	// passed to this method invocation.
	Arg1 []TagModifier
}

// Args returns an interface slice containing the arguments of this
// invocation. The variadic slice argument is flattened in this array such
// that one positional argument and three variadic arguments would result in
// a slice of four, not two.
func (c ConfigMustLoadFuncCall) Args() []interface{} {
	trailing := []interface{}{}
	for _, val := range c.Arg1 {
		trailing = append(trailing, val)
	}

	return append([]interface{}{c.Arg0}, trailing...)
}

// Results returns an interface slice containing the results of this
// invocation.
func (c ConfigMustLoadFuncCall) Results() []interface{} {
	return []interface{}{}
}

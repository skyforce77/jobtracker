package util

import (
	"container/list"

	"github.com/skyforce77/jobtracker/providers"
	"github.com/yuin/gopher-lua"
	"layeh.com/gopher-luar"
)

//Filter filters jobs from a lua script
func Filter(script string, fn func(*providers.Job)) func(job *providers.Job) {
	return func(job *providers.Job) {
		L := lua.NewState()
		defer L.Close()
		if err := L.DoFile(script); err != nil {
			panic(err)
		}
		if err := L.CallByParam(lua.P{
			Fn: L.GetGlobal("filter"),
			NRet: 1,
			Protect: true,
		}, luar.New(L, job)); err != nil {
			panic(err)
		}
		ret := L.Get(-1)
		L.Pop(1)
		if ret == lua.LTrue {
			fn(job)
		}
	}
}

// FilterCollect helps you recovering a list of jobs from a Provider
func FilterCollect(provider providers.Provider, script string) *list.List {
	lst := list.New()
	provider.RetrieveJobs(Filter(script, func(job *providers.Job) {
		lst.PushBack(job)
	}))
	return lst
}
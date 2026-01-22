package snapshot

import (
	"bytes"
	"compress/gzip"
	"container/list"
	"encoding/gob"
	"io"
	"os"
	"sync"

	"github.com/skyforce77/jobtracker/providers"
)

// Snapshot helps you to collect and store your scrapped jobs
type Snapshot struct {
	file    string
	content *list.List
	lock    sync.Mutex
}

// ProviderFailure refers to an error occurred in a provider
type ProviderFailure struct {
	Provider *providers.Provider
	Error error
}

// RetrieveJobs replays your snapshot jobs
func (snapshot *Snapshot) RetrieveJobs(fn func(job *providers.Job)) error {
	node := snapshot.content.Front()
	for node != nil {
		fn(node.Value.(*providers.Job))
		node = node.Next()
	}
	return nil
}

// NewSnapshot returns a new provider
func NewSnapshot(fileName string) *Snapshot {
	snapshot := &Snapshot{
		fileName, list.New(),
		sync.Mutex{}}

	file, err := os.Open(snapshot.file)
	if err != nil {
		return snapshot
	}

	gz, err := gzip.NewReader(file)
	if err != nil {
		return snapshot
	}
	defer gz.Close()

	data, err := io.ReadAll(gz)
	if err != nil {
		return snapshot
	}

	e := gob.NewDecoder(bytes.NewReader(data))
	var jobs []*providers.Job
	if err := e.Decode(&jobs); err != nil {
		return snapshot
	}

	for _, j := range jobs {
		snapshot.content.PushBack(j)
	}

	return snapshot
}

// Collector returns a thread-safe function to use with a provider
//
// Example :
//
//	snap := snapshot.NewSnapshot("./test")
//	p.RetrieveJobs(snap.Collector())
func (snapshot *Snapshot) Collector() func(job *providers.Job) {
	return func(job *providers.Job) {
		snapshot.lock.Lock()
		snapshot.content.PushBack(job)
		snapshot.lock.Unlock()
	}
}

// Save compresses your collected data and saves it
func (snapshot *Snapshot) Save() error {
	var arr []*providers.Job

	providers.IterateOver(snapshot.content, func(job *providers.Job) {
		arr = append(arr, job)
	})

	var b bytes.Buffer
	e := gob.NewEncoder(&b)
	if err := e.Encode(&arr); err != nil {
		return err
	}

	var gb bytes.Buffer
	gz := gzip.NewWriter(&gb)
	if _, err := gz.Write(b.Bytes()); err != nil {
		return err
	}
	if err := gz.Flush(); err != nil {
		return err
	}
	if err := gz.Close(); err != nil {
		return err
	}

	err := os.WriteFile(snapshot.file, gb.Bytes(), 0644)
	return err
}

// Erase removes every entry of a snapshot
func (snapshot *Snapshot) Erase() {
	snapshot.content = list.New()
}

// CollectFrom a slice of providers
func (snapshot *Snapshot) CollectFrom(pro []providers.Provider, collector func(job *providers.Job), fn func(int, *ProviderFailure)) {

	cn := make(chan *ProviderFailure, 8)

	for _, p := range pro {
		sp := p
		go func() {
			err := sp.RetrieveJobs(collector)
			cn <- &ProviderFailure{&p, err}
		}()
	}

	i := 0
	for i != len(pro) {
		err := <-cn
		fn(i, err)
		i++
	}
}

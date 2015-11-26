package repository

import (
	"log"
	"sync"
	"time"

	"github.com/ArtemTeleshev/go-scheduler"
)

const (
	SCHEDULER_NAME = "repository"

	EVENT_NAME_CLEAR  = "clear"
	EVENT_NAME_UPDATE = "update"
)

func NewRepository(name string) *Repository { // {{{
	repository := &Repository{
		name:      name,
		data:      make(Storages, 0),
		scheduler: scheduler.NewScheduler(SCHEDULER_NAME),
	}

	return repository
} // }}}

type Repository struct {
	sync.Mutex

	name      string
	data      Storages
	version   int
	scheduler *scheduler.Scheduler
}

func (this *Repository) Commit(storage *Storage) { // {{{
	this.Lock()
	defer this.Unlock()

	this.version++
	this.data[this.version] = storage
} // }}}

func (this *Repository) Clear() error { // {{{
	log.Println("Repository.Clear ... [started]")

	if len(this.data) > 1 {
		this.Lock()
		defer this.Unlock()

		for version, _ := range this.data {
			if version < this.version {
				delete(this.data, version)
				log.Println("Deleted", version)
			} else {
				log.Println("Not deleted", version)
			}
		}
	}

	return nil
} // }}}

func (this *Repository) RegisterEvent(event *scheduler.Event) { // {{{
	this.scheduler.Set(event)
} // }}}

func (this *Repository) RegisterEventClear(duration time.Duration) { // {{{
	this.RegisterEvent(scheduler.NewPeriodicEvent(EVENT_NAME_CLEAR, duration, scheduler.Action(this.Clear)))
} // }}}

func (this *Repository) RegisterEventUpdate(duration time.Duration, action scheduler.Action) { // {{{
	this.RegisterEvent(scheduler.NewPeriodicEvent(EVENT_NAME_UPDATE, duration, action))
	this.ExecuteEvent(EVENT_NAME_UPDATE)
} // }}}

func (this *Repository) ExecuteEvent(eventName string) { // {{{
	if this.scheduler.Has(eventName) {
		this.scheduler.Get(eventName).Execute()
	}
} // }}}

func (this *Repository) SchedulerStart() { // {{{
	this.scheduler.Start()
} // }}}

func (this *Repository) Version() int { // {{{
	return this.version
} // }}}

func (this *Repository) Storage() *Storage { // {{{
	return this.data[this.version]
} // }}}

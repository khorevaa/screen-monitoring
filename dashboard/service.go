package dashboard

import (
	"encoding/json"
	"sync"
)

type Widget struct {
	ID      string `json:"id"`
	Width   int    `json:"width"`
	Height  int    `json:"height"`
	Content string `json:"content"`
	Url     string `json:"url"`
}

func (w *Widget) ToJson() ([]byte, error) {
	return json.Marshal(w)
}

type PageContent struct {
	Widgets []Widget `json:"widgets"`
}

type RegisterResponse struct {
	Success bool
}

type DashboardService interface {
	GetPages() (pc PageContent, err error)
	Register(widget Widget) (pr RegisterResponse, err error)
	Init() ([]error, bool)
}

type dashboardService struct {
	sync.RWMutex
	migrator Migrator
	dbManager DatabaseManager
}

func NewDashboardService(migrator Migrator, dbManager DatabaseManager) DashboardService {
	return &dashboardService{
		migrator: migrator,
		dbManager: dbManager,
	}
}

func (d dashboardService) Init() ([]error, bool) {
	return d.migrator.Up()
}

func (d dashboardService) GetPages() (pc PageContent, err error) {
	d.Lock()
	defer d.Unlock()
	widgets, err := d.dbManager.GetAll(10,0)
	pc = PageContent{
		Widgets: widgets,
	}
	return
}

func (d *dashboardService) Register(widget Widget) (pr RegisterResponse, err error) {
	d.Lock()
	defer d.Unlock()
	_, err = d.dbManager.InsertOrUpdateWidget(&widget)
	if err != nil {
		pr = RegisterResponse{Success: true}
		return
	}
	pr = RegisterResponse{Success: false}
	return
}

package pool

import (
	"sync"

	"github.com/elastic/go-elasticsearch/v7"
)

type Pool struct {
	workers map[*elasticsearch.Client]bool
	m       *sync.Mutex
}

func New() (*Pool, error) {
	p := &Pool{}
	p.workers = make(map[*elasticsearch.Client]bool)

	for i := 0; i < 5; i++ {
		c, err := elasticsearch.NewDefaultClient()
		if err != nil {
			return nil, err
		}
		p.workers[c] = false
	}

	p.m = &sync.Mutex{}

	return p, nil
}

func (p *Pool) GetWorker() *elasticsearch.Client {
	p.m.Lock()
	defer p.m.Unlock()
	for k, v := range p.workers {
		if !v {
			p.workers[k] = true
			return k
		}
	}
	return nil
}

func (p *Pool) ReleaseWorker(client *elasticsearch.Client) {
	p.m.Lock()

	p.workers[client] = false

	p.m.Unlock()
}

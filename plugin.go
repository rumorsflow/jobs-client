package jobsclient

import (
	"github.com/go-redis/redis/v8"
	"github.com/hibiken/asynq"
	contracts "github.com/rumorsflow/contracts/redis"
)

const PluginName = "jobs_client"

type Plugin struct {
	client *asynq.Client
}

func (p *Plugin) Init(client redis.UniversalClient) error {
	p.client = asynq.NewClient(contracts.NewProxy(client))
	return nil
}

// Name returns user-friendly plugin name
func (p *Plugin) Name() string {
	return PluginName
}

// Provides declares factory methods.
func (p *Plugin) Provides() []any {
	return []any{
		p.ServiceClient,
	}
}

func (p *Plugin) ServiceClient() *asynq.Client {
	return p.client
}

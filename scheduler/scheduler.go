package scheduler

import (
	"github.com/adamdecaf/aws-ecs-nginx-proxy/ecs"
	"github.com/adamdecaf/aws-ecs-nginx-proxy/nginx"
)

// `Scheduler` is designed to watch over the ECS api for changes and call out to the nginx
// runner to repaint and reload the config.
type Scheduler struct {

}

// `Run()` starts the scheduler on it's interval of checks
func (s *Scheduler) Run() *error {
	return nil
}

// `New(nginx, ecs)` creates a Scheduler instance linked to an nginx process runner
// and an ECS cluster.
func New(nginx nginx.Nginx, ecs ecs.ECSClient) (*Scheduler, error) {
	return nil, nil
}

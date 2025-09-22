package core

import "github.com/v1adis1av28/level2/tasks/task16/internal/models"

type QueueManager interface {
	AddJob(job models.Job)
}

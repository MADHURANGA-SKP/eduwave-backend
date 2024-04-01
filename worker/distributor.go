package worker

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
)

// DistributeTaskSendVerifyEmail distributes the task for sending a verification email
func (distributor *SimpleTaskDistributor) DistributeTaskSendVerifyEmail(
	ctx context.Context,
	payload *PayloadSendVerifyEmail,
	opts ...asynq.Option,
) error {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal task payload: %w", err)
	}

	// Create a new task with the payload
	task := asynq.NewTask(TaskSendVerifyEmail, jsonPayload, opts...)

	// Process the task immediately instead of enqueuing
	err = distributor.ProcessTaskSendVerifyEmail(ctx, task)
	if err != nil {
		return fmt.Errorf("failed to process task: %w", err)
	}

	log.Info().Str("type", task.Type()).Bytes("payload", task.Payload()).
		Msg("task processed")
	return nil
}

package cluster

import (
	"context"
	"time"

	"github.com/filanov/bm-inventory/internal/common"
	"github.com/filanov/bm-inventory/models"
	logutil "github.com/filanov/bm-inventory/pkg/log"
	"github.com/go-openapi/swag"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

type PrepareConfig struct {
	InstallationTimeout time.Duration `envconfig:"PREPARE_FOR_INSTALLATION_TIMEOUT" default:"10m"`
}

type prepare struct {
	baseState
	PrepareConfig
}

var _ StateAPI = (*prepare)(nil)

func NewPrepareForInstallation(cfg PrepareConfig, log logrus.FieldLogger, db *gorm.DB) *prepare {
	return &prepare{
		baseState: baseState{
			log: log,
			db:  db,
		},
		PrepareConfig: cfg,
	}
}

func (p *prepare) RefreshStatus(ctx context.Context, c *common.Cluster, _ *gorm.DB) (*common.Cluster, error) {
	// can happen if the service was rebooted or somehow the async part crashed.
	if time.Since(time.Time(c.StatusUpdatedAt)) > p.InstallationTimeout {
		return updateClusterStatus(logutil.FromContext(ctx, p.log), p.db,
			*c.ID, swag.StringValue(c.Status), models.ClusterStatusError, statusInfoPreparingForInstallationTimeout)
	}
	return c, nil
}

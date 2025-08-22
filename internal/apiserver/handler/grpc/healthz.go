// Copyright 2025 武晓晨 <wuxc.eng@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/iWuxc/miniblog. The professional
// version of this repository is https://github.com/onexstack/onex.

package grpc

import (
	"context"
	"github.com/iWuxc/miniblog/internal/pkg/log"
	apiv1 "github.com/iWuxc/miniblog/pkg/api/apiserver/v1"
	"google.golang.org/protobuf/types/known/emptypb"
	"time"
)

func (h *Handler) Healthz(ctx context.Context, rq *emptypb.Empty) (*apiv1.HealthzResponse, error) {
	log.W(ctx).Infow("Healthz handler is called", "method", "Healthz", "status", "healthy")
	return &apiv1.HealthzResponse{
		Status:    apiv1.ServiceStatus_Healthy,
		Timestamp: time.Now().Format(time.DateTime),
	}, nil
}

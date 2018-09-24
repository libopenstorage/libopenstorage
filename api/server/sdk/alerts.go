/*
Package sdk is the gRPC implementation of the SDK gRPC server
Copyright 2018 Portworx

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package sdk

import (
	"context"
	"sync"
	"time"

	"github.com/libopenstorage/openstorage/alerts"
	"github.com/libopenstorage/openstorage/api"
	"github.com/libopenstorage/openstorage/pkg/proto/time"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// AlertsServer implements api.OpenStorageAlertsServer.
// In order to use this server implementation just have
// AlertsServer pointer properly instantiated with a valid
// alerts.Reader.
type AlertsServer struct {
	// Reader holds pointer to alerts Reader
	Reader alerts.Reader
}

func getOpts(opts []*api.SdkAlertsOption) []alerts.Option {
	var options []alerts.Option

	for _, opt := range opts {
		switch opt.GetOpt().(type) {
		case *api.SdkAlertsOption_MinSeverityType:
			options = append(options,
				alerts.NewMinSeverityOption(opt.GetMinSeverityType()))
		case *api.SdkAlertsOption_IsCleared:
			options = append(options,
				alerts.NewFlagCheckOption(opt.GetIsCleared()))
		case *api.SdkAlertsOption_TimeSpan:
			options = append(options,
				alerts.NewTimeSpanOption(
					prototime.TimestampToTime(opt.GetTimeSpan().GetStartTime()),
					prototime.TimestampToTime(opt.GetTimeSpan().GetEndTime())))
		case *api.SdkAlertsOption_CountSpan:
			options = append(options,
				alerts.NewCountSpanOption(
					opt.GetCountSpan().GetMinCount(),
					opt.GetCountSpan().GetMaxCount()))
		}
	}

	return options
}

func getFilters(queries []*api.SdkAlertsQuery) []alerts.Filter {
	var filters []alerts.Filter

	// range over all queries
	for _, x := range queries {
		switch x.GetQuery().(type) {
		case *api.SdkAlertsQuery_ResourceTypeQuery:
			q := x.GetResourceTypeQuery()
			filters = append(filters,
				alerts.NewResourceTypeFilter(
					q.ResourceType,
					getOpts(x.GetOpts())...))
		case *api.SdkAlertsQuery_AlertTypeQuery:
			q := x.GetAlertTypeQuery()
			filters = append(filters,
				alerts.NewAlertTypeFilter(
					q.AlertType,
					q.ResourceType,
					getOpts(x.GetOpts())...))
		case *api.SdkAlertsQuery_ResourceIdQuery:
			q := x.GetResourceIdQuery()
			filters = append(filters,
				alerts.NewResourceIDFilter(
					q.ResourceId,
					q.AlertType,
					q.ResourceType,
					getOpts(x.GetOpts())...))
		}
	}

	return filters
}

// Enumerate implements api.OpenStorageAlertsServer for AlertsServer.
// Input context should ideally have a deadline, in which case, a
// graceful exit is ensured within that deadline.
func (g *AlertsServer) Enumerate(ctx context.Context,
	request *api.SdkAlertsEnumerateRequest) (*api.SdkAlertsEnumerateResponse, error) {

	// if input has deadline, ensure graceful exit within that deadline.
	deadline, ok := ctx.Deadline()
	var cancel context.CancelFunc
	if ok {
		// create a new context that will get done on deadline
		ctx, cancel = context.WithTimeout(ctx, deadline.Sub(time.Now()))
		defer cancel()
	}

	group, _ := errgroup.WithContext(ctx)
	errChan := make(chan error)

	resp := new(api.SdkAlertsEnumerateResponse)
	var mu sync.Mutex

	queries := request.GetQueries()
	if queries == nil {
		return nil, status.Error(codes.InvalidArgument, "must provide a query")
	}

	filters := getFilters(queries)

	// spawn err-group process.
	// collect output using mutex.
	group.Go(func() error {
		if out, err := g.Reader.Enumerate(filters...); err != nil {
			return err
		} else {
			mu.Lock()
			resp.Alerts = append(resp.Alerts, out...)
			mu.Unlock()
			return nil
		}
	})

	// wait for err-group processes to be done
	go func() {
		errChan <- group.Wait()
	}()

	// wait only as long as context deadline allows
	select {
	case err := <-errChan:
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		} else {
			return resp, nil
		}
	case <-ctx.Done():
		return nil, status.Error(codes.DeadlineExceeded,
			"deadline is reached, server side func exiting")
	}
}

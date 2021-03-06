package entry

import (
	"context"

	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/spiffe/spire/pkg/common/telemetry"
	"github.com/spiffe/spire/pkg/server/api"
	"github.com/spiffe/spire/pkg/server/api/rpccontext"
	"github.com/spiffe/spire/pkg/server/plugin/datastore"
	"github.com/spiffe/spire/proto/spire-next/api/server/entry/v1"
	"github.com/spiffe/spire/proto/spire-next/types"
	"github.com/spiffe/spire/proto/spire/common"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// RegisterService registers the entry service on the gRPC server.
func RegisterService(s *grpc.Server, service *Service) {
	entry.RegisterEntryServer(s, service)
}

// Config is the service configuration
type Config struct {
	Datastore datastore.DataStore
}

// New creates a new entry service
func New(config Config) *Service {
	return &Service{
		ds: config.Datastore,
	}
}

// Service implements the v1 entry service
type Service struct {
	ds datastore.DataStore
}

func (s *Service) ListEntries(ctx context.Context, req *entry.ListEntriesRequest) (*entry.ListEntriesResponse, error) {
	return nil, status.Error(codes.Unimplemented, "method ListEntries not implemented")
}

func (s *Service) GetEntry(ctx context.Context, req *entry.GetEntryRequest) (*types.Entry, error) {
	log := rpccontext.Logger(ctx)

	if req.Id == "" {
		log.Error("Invalid request: missing ID")
		return nil, status.Error(codes.InvalidArgument, "missing ID")
	}
	log = log.WithField(telemetry.RegistrationID, req.Id)
	dsResp, err := s.ds.FetchRegistrationEntry(ctx, &datastore.FetchRegistrationEntryRequest{
		EntryId: req.Id,
	})
	if err != nil {
		log.WithError(err).Error("Failed to fetch entry")
		return nil, status.Errorf(codes.Internal, "failed to fetch entry: %v", err)
	}

	if dsResp.Entry == nil {
		log.Error("Entry not found")
		return nil, status.Error(codes.NotFound, "entry not found")
	}

	entry, err := api.RegistrationEntryToProto(dsResp.Entry)
	if err != nil {
		log.WithError(err).Error("Failed to convert entry")
		return nil, status.Errorf(codes.Internal, "failed to convert entry: %v", err)
	}
	applyMask(entry, req.OutputMask)

	return entry, nil
}

func (s *Service) BatchCreateEntry(ctx context.Context, req *entry.BatchCreateEntryRequest) (*entry.BatchCreateEntryResponse, error) {
	var results []*entry.BatchCreateEntryResponse_Result
	for _, eachEntry := range req.Entries {
		results = append(results, s.createEntry(ctx, eachEntry, req.OutputMask))
	}

	return &entry.BatchCreateEntryResponse{
		Results: results,
	}, nil
}

func (s *Service) createEntry(ctx context.Context, e *types.Entry, outputMask *types.EntryMask) *entry.BatchCreateEntryResponse_Result {
	log := rpccontext.Logger(ctx)

	cEntry, err := api.ProtoToRegistrationEntry(e)
	if err != nil {
		log.WithError(err).Error("Invalid request: failed to convert entry")
		return &entry.BatchCreateEntryResponse_Result{
			Status: api.CreateStatus(codes.InvalidArgument, "failed to convert entry: %v", err),
		}
	}

	log = log.WithField(telemetry.SPIFFEID, cEntry.SpiffeId)

	// Validates that there is no similar entry
	if isUniqueStatus := s.isEntryUnique(ctx, cEntry); isUniqueStatus != nil {
		return &entry.BatchCreateEntryResponse_Result{
			Status: isUniqueStatus,
		}
	}

	// Create entry
	resp, err := s.ds.CreateRegistrationEntry(ctx, &datastore.CreateRegistrationEntryRequest{
		Entry: cEntry,
	})
	if err != nil {
		log.WithError(err).Error("Failed to create entry")
		return &entry.BatchCreateEntryResponse_Result{
			Status: api.CreateStatus(codes.Internal, "failed to create entry: %v", err),
		}
	}

	tEntry, err := api.RegistrationEntryToProto(resp.Entry)
	if err != nil {
		log.WithError(err).Error("Unable to convert registration entry")
		return &entry.BatchCreateEntryResponse_Result{
			Status: api.CreateStatus(codes.Internal, "unable to convert registration entry: %v", err),
		}
	}

	applyMask(tEntry, outputMask)

	return &entry.BatchCreateEntryResponse_Result{
		Status: api.OK(),
		Entry:  tEntry,
	}
}

func (s *Service) isEntryUnique(ctx context.Context, e *common.RegistrationEntry) *types.Status {
	resp, err := s.ds.ListRegistrationEntries(ctx, &datastore.ListRegistrationEntriesRequest{
		BySpiffeId: &wrappers.StringValue{
			Value: e.SpiffeId,
		},
		ByParentId: &wrappers.StringValue{
			Value: e.ParentId,
		},
		BySelectors: &datastore.BySelectors{
			Match:     datastore.BySelectors_MATCH_EXACT,
			Selectors: e.Selectors,
		},
	})
	if err != nil {
		return api.CreateStatus(codes.Internal, "failed to list entries: %v", err)
	}
	if len(resp.Entries) != 0 {
		return api.CreateStatus(codes.AlreadyExists, "entry already exists")
	}

	return nil
}

func (s *Service) BatchUpdateEntry(ctx context.Context, req *entry.BatchUpdateEntryRequest) (*entry.BatchUpdateEntryResponse, error) {
	return nil, status.Error(codes.Unimplemented, "method BatchUpdateEntry not implemented")
}

func (s *Service) BatchDeleteEntry(ctx context.Context, req *entry.BatchDeleteEntryRequest) (*entry.BatchDeleteEntryResponse, error) {
	var results []*entry.BatchDeleteEntryResponse_Result
	for _, id := range req.Ids {
		results = append(results, s.deleteEntry(ctx, id))
	}

	return &entry.BatchDeleteEntryResponse{
		Results: results,
	}, nil
}

func (s *Service) deleteEntry(ctx context.Context, id string) *entry.BatchDeleteEntryResponse_Result {
	log := rpccontext.Logger(ctx)

	if id == "" {
		log.Error("Invalid request: missing entry ID")
		return &entry.BatchDeleteEntryResponse_Result{
			Id:     id,
			Status: api.CreateStatus(codes.InvalidArgument, "missing entry ID"),
		}
	}

	log = log.WithField(telemetry.RegistrationID, id)

	_, err := s.ds.DeleteRegistrationEntry(ctx, &datastore.DeleteRegistrationEntryRequest{
		EntryId: id,
	})
	switch status.Code(err) {
	case codes.OK:
		return &entry.BatchDeleteEntryResponse_Result{
			Id:     id,
			Status: api.OK(),
		}
	case codes.NotFound:
		return &entry.BatchDeleteEntryResponse_Result{
			Id:     id,
			Status: api.StatusFromError(err),
		}
	default:
		log.WithError(err).Error("Failed to delete entry")
		return &entry.BatchDeleteEntryResponse_Result{
			Id:     id,
			Status: api.CreateStatus(codes.Internal, "failed to delete entry: %v", err),
		}
	}
}

func (s *Service) GetAuthorizedEntries(ctx context.Context, req *entry.GetAuthorizedEntriesRequest) (*entry.GetAuthorizedEntriesResponse, error) {
	return nil, status.Error(codes.Unimplemented, "method GetAuthorizedEntries not implemented")
}

func applyMask(e *types.Entry, mask *types.EntryMask) { //nolint: unused,deadcode
	if mask == nil {
		return
	}

	if !mask.SpiffeId {
		e.SpiffeId = nil
	}

	if !mask.ParentId {
		e.ParentId = nil
	}

	if !mask.Selectors {
		e.Selectors = nil
	}

	if !mask.Ttl {
		e.Ttl = 0
	}

	if !mask.FederatesWith {
		e.FederatesWith = nil
	}

	if !mask.Admin {
		e.Admin = false
	}

	if !mask.Downstream {
		e.Downstream = false
	}

	if !mask.ExpiresAt {
		e.ExpiresAt = 0
	}

	if !mask.DnsNames {
		e.DnsNames = nil
	}
}

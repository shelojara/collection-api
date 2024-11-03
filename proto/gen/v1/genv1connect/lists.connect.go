// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: v1/lists.proto

package genv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/shelojara/collection-api/proto/gen/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// ListsName is the fully-qualified name of the Lists service.
	ListsName = "v1.Lists"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// ListsGetListProcedure is the fully-qualified name of the Lists's GetList RPC.
	ListsGetListProcedure = "/v1.Lists/GetList"
	// ListsListListItemsProcedure is the fully-qualified name of the Lists's ListListItems RPC.
	ListsListListItemsProcedure = "/v1.Lists/ListListItems"
	// ListsCreateListProcedure is the fully-qualified name of the Lists's CreateList RPC.
	ListsCreateListProcedure = "/v1.Lists/CreateList"
	// ListsUpdateListProcedure is the fully-qualified name of the Lists's UpdateList RPC.
	ListsUpdateListProcedure = "/v1.Lists/UpdateList"
	// ListsCreateListStatusProcedure is the fully-qualified name of the Lists's CreateListStatus RPC.
	ListsCreateListStatusProcedure = "/v1.Lists/CreateListStatus"
	// ListsUpdateListStatusProcedure is the fully-qualified name of the Lists's UpdateListStatus RPC.
	ListsUpdateListStatusProcedure = "/v1.Lists/UpdateListStatus"
	// ListsCreateListItemProcedure is the fully-qualified name of the Lists's CreateListItem RPC.
	ListsCreateListItemProcedure = "/v1.Lists/CreateListItem"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	listsServiceDescriptor                = v1.File_v1_lists_proto.Services().ByName("Lists")
	listsGetListMethodDescriptor          = listsServiceDescriptor.Methods().ByName("GetList")
	listsListListItemsMethodDescriptor    = listsServiceDescriptor.Methods().ByName("ListListItems")
	listsCreateListMethodDescriptor       = listsServiceDescriptor.Methods().ByName("CreateList")
	listsUpdateListMethodDescriptor       = listsServiceDescriptor.Methods().ByName("UpdateList")
	listsCreateListStatusMethodDescriptor = listsServiceDescriptor.Methods().ByName("CreateListStatus")
	listsUpdateListStatusMethodDescriptor = listsServiceDescriptor.Methods().ByName("UpdateListStatus")
	listsCreateListItemMethodDescriptor   = listsServiceDescriptor.Methods().ByName("CreateListItem")
)

// ListsClient is a client for the v1.Lists service.
type ListsClient interface {
	GetList(context.Context, *connect.Request[v1.GetListRequest]) (*connect.Response[v1.List], error)
	ListListItems(context.Context, *connect.Request[v1.ListListItemsRequest]) (*connect.Response[v1.ListListItemsResponse], error)
	CreateList(context.Context, *connect.Request[v1.CreateListRequest]) (*connect.Response[v1.CreateListResponse], error)
	UpdateList(context.Context, *connect.Request[v1.UpdateListRequest]) (*connect.Response[v1.UpdateListResponse], error)
	CreateListStatus(context.Context, *connect.Request[v1.CreateListStatusRequest]) (*connect.Response[v1.CreateListStatusResponse], error)
	UpdateListStatus(context.Context, *connect.Request[v1.UpdateListStatusRequest]) (*connect.Response[v1.UpdateListStatusResponse], error)
	CreateListItem(context.Context, *connect.Request[v1.CreateListItemRequest]) (*connect.Response[v1.CreateListItemResponse], error)
}

// NewListsClient constructs a client for the v1.Lists service. By default, it uses the Connect
// protocol with the binary Protobuf Codec, asks for gzipped responses, and sends uncompressed
// requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewListsClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) ListsClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &listsClient{
		getList: connect.NewClient[v1.GetListRequest, v1.List](
			httpClient,
			baseURL+ListsGetListProcedure,
			connect.WithSchema(listsGetListMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		listListItems: connect.NewClient[v1.ListListItemsRequest, v1.ListListItemsResponse](
			httpClient,
			baseURL+ListsListListItemsProcedure,
			connect.WithSchema(listsListListItemsMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		createList: connect.NewClient[v1.CreateListRequest, v1.CreateListResponse](
			httpClient,
			baseURL+ListsCreateListProcedure,
			connect.WithSchema(listsCreateListMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		updateList: connect.NewClient[v1.UpdateListRequest, v1.UpdateListResponse](
			httpClient,
			baseURL+ListsUpdateListProcedure,
			connect.WithSchema(listsUpdateListMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		createListStatus: connect.NewClient[v1.CreateListStatusRequest, v1.CreateListStatusResponse](
			httpClient,
			baseURL+ListsCreateListStatusProcedure,
			connect.WithSchema(listsCreateListStatusMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		updateListStatus: connect.NewClient[v1.UpdateListStatusRequest, v1.UpdateListStatusResponse](
			httpClient,
			baseURL+ListsUpdateListStatusProcedure,
			connect.WithSchema(listsUpdateListStatusMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		createListItem: connect.NewClient[v1.CreateListItemRequest, v1.CreateListItemResponse](
			httpClient,
			baseURL+ListsCreateListItemProcedure,
			connect.WithSchema(listsCreateListItemMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// listsClient implements ListsClient.
type listsClient struct {
	getList          *connect.Client[v1.GetListRequest, v1.List]
	listListItems    *connect.Client[v1.ListListItemsRequest, v1.ListListItemsResponse]
	createList       *connect.Client[v1.CreateListRequest, v1.CreateListResponse]
	updateList       *connect.Client[v1.UpdateListRequest, v1.UpdateListResponse]
	createListStatus *connect.Client[v1.CreateListStatusRequest, v1.CreateListStatusResponse]
	updateListStatus *connect.Client[v1.UpdateListStatusRequest, v1.UpdateListStatusResponse]
	createListItem   *connect.Client[v1.CreateListItemRequest, v1.CreateListItemResponse]
}

// GetList calls v1.Lists.GetList.
func (c *listsClient) GetList(ctx context.Context, req *connect.Request[v1.GetListRequest]) (*connect.Response[v1.List], error) {
	return c.getList.CallUnary(ctx, req)
}

// ListListItems calls v1.Lists.ListListItems.
func (c *listsClient) ListListItems(ctx context.Context, req *connect.Request[v1.ListListItemsRequest]) (*connect.Response[v1.ListListItemsResponse], error) {
	return c.listListItems.CallUnary(ctx, req)
}

// CreateList calls v1.Lists.CreateList.
func (c *listsClient) CreateList(ctx context.Context, req *connect.Request[v1.CreateListRequest]) (*connect.Response[v1.CreateListResponse], error) {
	return c.createList.CallUnary(ctx, req)
}

// UpdateList calls v1.Lists.UpdateList.
func (c *listsClient) UpdateList(ctx context.Context, req *connect.Request[v1.UpdateListRequest]) (*connect.Response[v1.UpdateListResponse], error) {
	return c.updateList.CallUnary(ctx, req)
}

// CreateListStatus calls v1.Lists.CreateListStatus.
func (c *listsClient) CreateListStatus(ctx context.Context, req *connect.Request[v1.CreateListStatusRequest]) (*connect.Response[v1.CreateListStatusResponse], error) {
	return c.createListStatus.CallUnary(ctx, req)
}

// UpdateListStatus calls v1.Lists.UpdateListStatus.
func (c *listsClient) UpdateListStatus(ctx context.Context, req *connect.Request[v1.UpdateListStatusRequest]) (*connect.Response[v1.UpdateListStatusResponse], error) {
	return c.updateListStatus.CallUnary(ctx, req)
}

// CreateListItem calls v1.Lists.CreateListItem.
func (c *listsClient) CreateListItem(ctx context.Context, req *connect.Request[v1.CreateListItemRequest]) (*connect.Response[v1.CreateListItemResponse], error) {
	return c.createListItem.CallUnary(ctx, req)
}

// ListsHandler is an implementation of the v1.Lists service.
type ListsHandler interface {
	GetList(context.Context, *connect.Request[v1.GetListRequest]) (*connect.Response[v1.List], error)
	ListListItems(context.Context, *connect.Request[v1.ListListItemsRequest]) (*connect.Response[v1.ListListItemsResponse], error)
	CreateList(context.Context, *connect.Request[v1.CreateListRequest]) (*connect.Response[v1.CreateListResponse], error)
	UpdateList(context.Context, *connect.Request[v1.UpdateListRequest]) (*connect.Response[v1.UpdateListResponse], error)
	CreateListStatus(context.Context, *connect.Request[v1.CreateListStatusRequest]) (*connect.Response[v1.CreateListStatusResponse], error)
	UpdateListStatus(context.Context, *connect.Request[v1.UpdateListStatusRequest]) (*connect.Response[v1.UpdateListStatusResponse], error)
	CreateListItem(context.Context, *connect.Request[v1.CreateListItemRequest]) (*connect.Response[v1.CreateListItemResponse], error)
}

// NewListsHandler builds an HTTP handler from the service implementation. It returns the path on
// which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewListsHandler(svc ListsHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	listsGetListHandler := connect.NewUnaryHandler(
		ListsGetListProcedure,
		svc.GetList,
		connect.WithSchema(listsGetListMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	listsListListItemsHandler := connect.NewUnaryHandler(
		ListsListListItemsProcedure,
		svc.ListListItems,
		connect.WithSchema(listsListListItemsMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	listsCreateListHandler := connect.NewUnaryHandler(
		ListsCreateListProcedure,
		svc.CreateList,
		connect.WithSchema(listsCreateListMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	listsUpdateListHandler := connect.NewUnaryHandler(
		ListsUpdateListProcedure,
		svc.UpdateList,
		connect.WithSchema(listsUpdateListMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	listsCreateListStatusHandler := connect.NewUnaryHandler(
		ListsCreateListStatusProcedure,
		svc.CreateListStatus,
		connect.WithSchema(listsCreateListStatusMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	listsUpdateListStatusHandler := connect.NewUnaryHandler(
		ListsUpdateListStatusProcedure,
		svc.UpdateListStatus,
		connect.WithSchema(listsUpdateListStatusMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	listsCreateListItemHandler := connect.NewUnaryHandler(
		ListsCreateListItemProcedure,
		svc.CreateListItem,
		connect.WithSchema(listsCreateListItemMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/v1.Lists/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case ListsGetListProcedure:
			listsGetListHandler.ServeHTTP(w, r)
		case ListsListListItemsProcedure:
			listsListListItemsHandler.ServeHTTP(w, r)
		case ListsCreateListProcedure:
			listsCreateListHandler.ServeHTTP(w, r)
		case ListsUpdateListProcedure:
			listsUpdateListHandler.ServeHTTP(w, r)
		case ListsCreateListStatusProcedure:
			listsCreateListStatusHandler.ServeHTTP(w, r)
		case ListsUpdateListStatusProcedure:
			listsUpdateListStatusHandler.ServeHTTP(w, r)
		case ListsCreateListItemProcedure:
			listsCreateListItemHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedListsHandler returns CodeUnimplemented from all methods.
type UnimplementedListsHandler struct{}

func (UnimplementedListsHandler) GetList(context.Context, *connect.Request[v1.GetListRequest]) (*connect.Response[v1.List], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("v1.Lists.GetList is not implemented"))
}

func (UnimplementedListsHandler) ListListItems(context.Context, *connect.Request[v1.ListListItemsRequest]) (*connect.Response[v1.ListListItemsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("v1.Lists.ListListItems is not implemented"))
}

func (UnimplementedListsHandler) CreateList(context.Context, *connect.Request[v1.CreateListRequest]) (*connect.Response[v1.CreateListResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("v1.Lists.CreateList is not implemented"))
}

func (UnimplementedListsHandler) UpdateList(context.Context, *connect.Request[v1.UpdateListRequest]) (*connect.Response[v1.UpdateListResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("v1.Lists.UpdateList is not implemented"))
}

func (UnimplementedListsHandler) CreateListStatus(context.Context, *connect.Request[v1.CreateListStatusRequest]) (*connect.Response[v1.CreateListStatusResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("v1.Lists.CreateListStatus is not implemented"))
}

func (UnimplementedListsHandler) UpdateListStatus(context.Context, *connect.Request[v1.UpdateListStatusRequest]) (*connect.Response[v1.UpdateListStatusResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("v1.Lists.UpdateListStatus is not implemented"))
}

func (UnimplementedListsHandler) CreateListItem(context.Context, *connect.Request[v1.CreateListItemRequest]) (*connect.Response[v1.CreateListItemResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("v1.Lists.CreateListItem is not implemented"))
}

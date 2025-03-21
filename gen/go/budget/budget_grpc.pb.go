// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.25.6
// source: budget/budget.proto

package budgetv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	BudgetService_AddTransaction_FullMethodName       = "/budget.BudgetService/AddTransaction"
	BudgetService_GetTransactions_FullMethodName      = "/budget.BudgetService/GetTransactions"
	BudgetService_DeleteTransaction_FullMethodName    = "/budget.BudgetService/DeleteTransaction"
	BudgetService_GetBalance_FullMethodName           = "/budget.BudgetService/GetBalance"
	BudgetService_AddCategory_FullMethodName          = "/budget.BudgetService/AddCategory"
	BudgetService_GetCategories_FullMethodName        = "/budget.BudgetService/GetCategories"
	BudgetService_AddRecurringPayment_FullMethodName  = "/budget.BudgetService/AddRecurringPayment"
	BudgetService_GetRecurringPayments_FullMethodName = "/budget.BudgetService/GetRecurringPayments"
)

// BudgetServiceClient is the client API for BudgetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Определение сервиса BudgetService
type BudgetServiceClient interface {
	// Добавление транзакции (доход/расход)
	AddTransaction(ctx context.Context, in *TransactionRequest, opts ...grpc.CallOption) (*TransactionResponse, error)
	// Получение списка транзакций пользователя
	GetTransactions(ctx context.Context, in *GetTransactionsRequest, opts ...grpc.CallOption) (*GetTransactionsResponse, error)
	// Удаление транзакции
	DeleteTransaction(ctx context.Context, in *DeleteTransactionRequest, opts ...grpc.CallOption) (*DeleteTransactionResponse, error)
	// Получение текущего баланса
	GetBalance(ctx context.Context, in *GetBalanceRequest, opts ...grpc.CallOption) (*GetBalanceResponse, error)
	// Управление категориями расходов/доходов
	AddCategory(ctx context.Context, in *CategoryRequest, opts ...grpc.CallOption) (*CategoryResponse, error)
	GetCategories(ctx context.Context, in *GetCategoriesRequest, opts ...grpc.CallOption) (*GetCategoriesResponse, error)
	// Управление регулярными платежами
	AddRecurringPayment(ctx context.Context, in *RecurringPaymentRequest, opts ...grpc.CallOption) (*RecurringPaymentResponse, error)
	GetRecurringPayments(ctx context.Context, in *GetRecurringPaymentsRequest, opts ...grpc.CallOption) (*GetRecurringPaymentsResponse, error)
}

type budgetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBudgetServiceClient(cc grpc.ClientConnInterface) BudgetServiceClient {
	return &budgetServiceClient{cc}
}

func (c *budgetServiceClient) AddTransaction(ctx context.Context, in *TransactionRequest, opts ...grpc.CallOption) (*TransactionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, BudgetService_AddTransaction_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *budgetServiceClient) GetTransactions(ctx context.Context, in *GetTransactionsRequest, opts ...grpc.CallOption) (*GetTransactionsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTransactionsResponse)
	err := c.cc.Invoke(ctx, BudgetService_GetTransactions_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *budgetServiceClient) DeleteTransaction(ctx context.Context, in *DeleteTransactionRequest, opts ...grpc.CallOption) (*DeleteTransactionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteTransactionResponse)
	err := c.cc.Invoke(ctx, BudgetService_DeleteTransaction_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *budgetServiceClient) GetBalance(ctx context.Context, in *GetBalanceRequest, opts ...grpc.CallOption) (*GetBalanceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetBalanceResponse)
	err := c.cc.Invoke(ctx, BudgetService_GetBalance_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *budgetServiceClient) AddCategory(ctx context.Context, in *CategoryRequest, opts ...grpc.CallOption) (*CategoryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CategoryResponse)
	err := c.cc.Invoke(ctx, BudgetService_AddCategory_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *budgetServiceClient) GetCategories(ctx context.Context, in *GetCategoriesRequest, opts ...grpc.CallOption) (*GetCategoriesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCategoriesResponse)
	err := c.cc.Invoke(ctx, BudgetService_GetCategories_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *budgetServiceClient) AddRecurringPayment(ctx context.Context, in *RecurringPaymentRequest, opts ...grpc.CallOption) (*RecurringPaymentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RecurringPaymentResponse)
	err := c.cc.Invoke(ctx, BudgetService_AddRecurringPayment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *budgetServiceClient) GetRecurringPayments(ctx context.Context, in *GetRecurringPaymentsRequest, opts ...grpc.CallOption) (*GetRecurringPaymentsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetRecurringPaymentsResponse)
	err := c.cc.Invoke(ctx, BudgetService_GetRecurringPayments_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BudgetServiceServer is the server API for BudgetService service.
// All implementations must embed UnimplementedBudgetServiceServer
// for forward compatibility.
//
// Определение сервиса BudgetService
type BudgetServiceServer interface {
	// Добавление транзакции (доход/расход)
	AddTransaction(context.Context, *TransactionRequest) (*TransactionResponse, error)
	// Получение списка транзакций пользователя
	GetTransactions(context.Context, *GetTransactionsRequest) (*GetTransactionsResponse, error)
	// Удаление транзакции
	DeleteTransaction(context.Context, *DeleteTransactionRequest) (*DeleteTransactionResponse, error)
	// Получение текущего баланса
	GetBalance(context.Context, *GetBalanceRequest) (*GetBalanceResponse, error)
	// Управление категориями расходов/доходов
	AddCategory(context.Context, *CategoryRequest) (*CategoryResponse, error)
	GetCategories(context.Context, *GetCategoriesRequest) (*GetCategoriesResponse, error)
	// Управление регулярными платежами
	AddRecurringPayment(context.Context, *RecurringPaymentRequest) (*RecurringPaymentResponse, error)
	GetRecurringPayments(context.Context, *GetRecurringPaymentsRequest) (*GetRecurringPaymentsResponse, error)
	mustEmbedUnimplementedBudgetServiceServer()
}

// UnimplementedBudgetServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedBudgetServiceServer struct{}

func (UnimplementedBudgetServiceServer) AddTransaction(context.Context, *TransactionRequest) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTransaction not implemented")
}
func (UnimplementedBudgetServiceServer) GetTransactions(context.Context, *GetTransactionsRequest) (*GetTransactionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransactions not implemented")
}
func (UnimplementedBudgetServiceServer) DeleteTransaction(context.Context, *DeleteTransactionRequest) (*DeleteTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTransaction not implemented")
}
func (UnimplementedBudgetServiceServer) GetBalance(context.Context, *GetBalanceRequest) (*GetBalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBalance not implemented")
}
func (UnimplementedBudgetServiceServer) AddCategory(context.Context, *CategoryRequest) (*CategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCategory not implemented")
}
func (UnimplementedBudgetServiceServer) GetCategories(context.Context, *GetCategoriesRequest) (*GetCategoriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCategories not implemented")
}
func (UnimplementedBudgetServiceServer) AddRecurringPayment(context.Context, *RecurringPaymentRequest) (*RecurringPaymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddRecurringPayment not implemented")
}
func (UnimplementedBudgetServiceServer) GetRecurringPayments(context.Context, *GetRecurringPaymentsRequest) (*GetRecurringPaymentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRecurringPayments not implemented")
}
func (UnimplementedBudgetServiceServer) mustEmbedUnimplementedBudgetServiceServer() {}
func (UnimplementedBudgetServiceServer) testEmbeddedByValue()                       {}

// UnsafeBudgetServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BudgetServiceServer will
// result in compilation errors.
type UnsafeBudgetServiceServer interface {
	mustEmbedUnimplementedBudgetServiceServer()
}

func RegisterBudgetServiceServer(s grpc.ServiceRegistrar, srv BudgetServiceServer) {
	// If the following call pancis, it indicates UnimplementedBudgetServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&BudgetService_ServiceDesc, srv)
}

func _BudgetService_AddTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BudgetServiceServer).AddTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BudgetService_AddTransaction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BudgetServiceServer).AddTransaction(ctx, req.(*TransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BudgetService_GetTransactions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTransactionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BudgetServiceServer).GetTransactions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BudgetService_GetTransactions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BudgetServiceServer).GetTransactions(ctx, req.(*GetTransactionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BudgetService_DeleteTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BudgetServiceServer).DeleteTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BudgetService_DeleteTransaction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BudgetServiceServer).DeleteTransaction(ctx, req.(*DeleteTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BudgetService_GetBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BudgetServiceServer).GetBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BudgetService_GetBalance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BudgetServiceServer).GetBalance(ctx, req.(*GetBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BudgetService_AddCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BudgetServiceServer).AddCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BudgetService_AddCategory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BudgetServiceServer).AddCategory(ctx, req.(*CategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BudgetService_GetCategories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCategoriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BudgetServiceServer).GetCategories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BudgetService_GetCategories_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BudgetServiceServer).GetCategories(ctx, req.(*GetCategoriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BudgetService_AddRecurringPayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecurringPaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BudgetServiceServer).AddRecurringPayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BudgetService_AddRecurringPayment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BudgetServiceServer).AddRecurringPayment(ctx, req.(*RecurringPaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BudgetService_GetRecurringPayments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRecurringPaymentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BudgetServiceServer).GetRecurringPayments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BudgetService_GetRecurringPayments_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BudgetServiceServer).GetRecurringPayments(ctx, req.(*GetRecurringPaymentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BudgetService_ServiceDesc is the grpc.ServiceDesc for BudgetService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BudgetService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "budget.BudgetService",
	HandlerType: (*BudgetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddTransaction",
			Handler:    _BudgetService_AddTransaction_Handler,
		},
		{
			MethodName: "GetTransactions",
			Handler:    _BudgetService_GetTransactions_Handler,
		},
		{
			MethodName: "DeleteTransaction",
			Handler:    _BudgetService_DeleteTransaction_Handler,
		},
		{
			MethodName: "GetBalance",
			Handler:    _BudgetService_GetBalance_Handler,
		},
		{
			MethodName: "AddCategory",
			Handler:    _BudgetService_AddCategory_Handler,
		},
		{
			MethodName: "GetCategories",
			Handler:    _BudgetService_GetCategories_Handler,
		},
		{
			MethodName: "AddRecurringPayment",
			Handler:    _BudgetService_AddRecurringPayment_Handler,
		},
		{
			MethodName: "GetRecurringPayments",
			Handler:    _BudgetService_GetRecurringPayments_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "budget/budget.proto",
}

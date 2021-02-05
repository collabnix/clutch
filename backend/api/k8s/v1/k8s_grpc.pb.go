// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package k8sv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// K8SAPIClient is the client API for K8SAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type K8SAPIClient interface {
	DescribePod(ctx context.Context, in *DescribePodRequest, opts ...grpc.CallOption) (*DescribePodResponse, error)
	ListPods(ctx context.Context, in *ListPodsRequest, opts ...grpc.CallOption) (*ListPodsResponse, error)
	DeletePod(ctx context.Context, in *DeletePodRequest, opts ...grpc.CallOption) (*DeletePodResponse, error)
	UpdatePod(ctx context.Context, in *UpdatePodRequest, opts ...grpc.CallOption) (*UpdatePodResponse, error)
	ResizeHPA(ctx context.Context, in *ResizeHPARequest, opts ...grpc.CallOption) (*ResizeHPAResponse, error)
	DeleteHPA(ctx context.Context, in *DeleteHPARequest, opts ...grpc.CallOption) (*DeleteHPAResponse, error)
	UpdateDeployment(ctx context.Context, in *UpdateDeploymentRequest, opts ...grpc.CallOption) (*UpdateDeploymentResponse, error)
	DeleteDeployment(ctx context.Context, in *DeleteDeploymentRequest, opts ...grpc.CallOption) (*DeleteDeploymentResponse, error)
	DescribeService(ctx context.Context, in *DescribeServiceRequest, opts ...grpc.CallOption) (*DescribeServiceResponse, error)
	DeleteService(ctx context.Context, in *DeleteServiceRequest, opts ...grpc.CallOption) (*DeleteServiceResponse, error)
	DescribeStatefulSet(ctx context.Context, in *DescribeStatefulSetRequest, opts ...grpc.CallOption) (*DescribeStatefulSetResponse, error)
	UpdateStatefulSet(ctx context.Context, in *UpdateStatefulSetRequest, opts ...grpc.CallOption) (*UpdateStatefulSetResponse, error)
	DeleteStatefulSet(ctx context.Context, in *DeleteStatefulSetRequest, opts ...grpc.CallOption) (*DeleteStatefulSetResponse, error)
	DescribeCronJob(ctx context.Context, in *DescribeCronJobRequest, opts ...grpc.CallOption) (*DescribeCronJobResponse, error)
	DeleteCronJob(ctx context.Context, in *DeleteCronJobRequest, opts ...grpc.CallOption) (*DeleteCronJobResponse, error)
	ListConfigMaps(ctx context.Context, in *ListConfigMapsRequest, opts ...grpc.CallOption) (*ListConfigMapsResponse, error)
	DescribeConfigMap(ctx context.Context, in *DescribeConfigMapRequest, opts ...grpc.CallOption) (*DescribeConfigMapResponse, error)
	DeleteConfigMap(ctx context.Context, in *DeleteConfigMapRequest, opts ...grpc.CallOption) (*DeleteConfigMapResponse, error)
}

type k8SAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewK8SAPIClient(cc grpc.ClientConnInterface) K8SAPIClient {
	return &k8SAPIClient{cc}
}

func (c *k8SAPIClient) DescribePod(ctx context.Context, in *DescribePodRequest, opts ...grpc.CallOption) (*DescribePodResponse, error) {
	out := new(DescribePodResponse)
	err := c.cc.Invoke(ctx, "/clutch.k8s.v1.K8sAPI/DescribePod", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *k8SAPIClient) ListPods(ctx context.Context, in *ListPodsRequest, opts ...grpc.CallOption) (*ListPodsResponse, error) {
	out := new(ListPodsResponse)
	err := c.cc.Invoke(ctx, "/clutch.k8s.v1.K8sAPI/ListPods", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *k8SAPIClient) DeletePod(ctx context.Context, in *DeletePodRequest, opts ...grpc.CallOption) (*DeletePodResponse, error) {
	out := new(DeletePodResponse)
	err := c.cc.Invoke(ctx, "/clutch.k8s.v1.K8sAPI/DeletePod", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *k8SAPIClient) UpdatePod(ctx context.Context, in *UpdatePodRequest, opts ...grpc.CallOption) (*UpdatePodResponse, error) {
	out := new(UpdatePodResponse)
	err := c.cc.Invoke(ctx, "/clutch.k8s.v1.K8sAPI/UpdatePod", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *k8SAPIClient) ResizeHPA(ctx context.Context, in *ResizeHPARequest, opts ...grpc.CallOption) (*ResizeHPAResponse, error) {
	out := new(ResizeHPAResponse)
	err := c.cc.Invoke(ctx, "/clutch.k8s.v1.K8sAPI/ResizeHPA", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *k8SAPIClient) DeleteHPA(ctx context.Context, in *DeleteHPARequest, opts ...grpc.CallOption) (*DeleteHPAResponse, error) {
	out := new(DeleteHPAResponse)
	err := c.cc.Invoke(ctx, "/clutch.k8s.v1.K8sAPI/DeleteHPA", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *k8SAPIClient) UpdateDeployment(ctx context.Context, in *UpdateDeploymentRequest, opts ...grpc.CallOption) (*UpdateDeploymentResponse, error) {
	out := new(UpdateDeploymentResponse)
	err := c.cc.Invoke(ctx, "/clutch.k8s.v1.K8sAPI/UpdateDeployment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *k8SAPIClient) DeleteDeployment(ctx context.Context, in *DeleteDeploymentRequest, opts ...grpc.CallOption) (*DeleteDeploymentResponse, error) {
	out := new(DeleteDeploymentResponse)
	err := c.cc.Invoke(ctx, "/clutch.k8s.v1.K8sAPI/DeleteDeployment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *k8SAPIClient) DescribeService(ctx context.Context, in *DescribeServiceRequest, opts ...grpc.CallOption) (*DescribeServiceResponse, error) {
	out := new(DescribeServiceResponse)
	err := c.cc.Invoke(ctx, "/clutch.k8s.v1.K8sAPI/DescribeService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *k8SAPIClient) DeleteService(ctx context.Context, in *DeleteServiceRequest, opts ...grpc.CallOption) (*DeleteServiceResponse, error) {
	out := new(DeleteServiceResponse)
	err := c.cc.Invoke(ctx, "/clutch.k8s.v1.K8sAPI/DeleteService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *k8SAPIClient) DescribeStatefulSet(ctx context.Context, in *DescribeStatefulSetRequest, opts ...grpc.CallOption) (*DescribeStatefulSetResponse, error) {
	out := new(DescribeStatefulSetResponse)
	err := c.cc.Invoke(ctx, "/clutch.k8s.v1.K8sAPI/DescribeStatefulSet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *k8SAPIClient) UpdateStatefulSet(ctx context.Context, in *UpdateStatefulSetRequest, opts ...grpc.CallOption) (*UpdateStatefulSetResponse, error) {
	out := new(UpdateStatefulSetResponse)
	err := c.cc.Invoke(ctx, "/clutch.k8s.v1.K8sAPI/UpdateStatefulSet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *k8SAPIClient) DeleteStatefulSet(ctx context.Context, in *DeleteStatefulSetRequest, opts ...grpc.CallOption) (*DeleteStatefulSetResponse, error) {
	out := new(DeleteStatefulSetResponse)
	err := c.cc.Invoke(ctx, "/clutch.k8s.v1.K8sAPI/DeleteStatefulSet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *k8SAPIClient) DescribeCronJob(ctx context.Context, in *DescribeCronJobRequest, opts ...grpc.CallOption) (*DescribeCronJobResponse, error) {
	out := new(DescribeCronJobResponse)
	err := c.cc.Invoke(ctx, "/clutch.k8s.v1.K8sAPI/DescribeCronJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *k8SAPIClient) DeleteCronJob(ctx context.Context, in *DeleteCronJobRequest, opts ...grpc.CallOption) (*DeleteCronJobResponse, error) {
	out := new(DeleteCronJobResponse)
	err := c.cc.Invoke(ctx, "/clutch.k8s.v1.K8sAPI/DeleteCronJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *k8SAPIClient) ListConfigMaps(ctx context.Context, in *ListConfigMapsRequest, opts ...grpc.CallOption) (*ListConfigMapsResponse, error) {
	out := new(ListConfigMapsResponse)
	err := c.cc.Invoke(ctx, "/clutch.k8s.v1.K8sAPI/ListConfigMaps", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *k8SAPIClient) DescribeConfigMap(ctx context.Context, in *DescribeConfigMapRequest, opts ...grpc.CallOption) (*DescribeConfigMapResponse, error) {
	out := new(DescribeConfigMapResponse)
	err := c.cc.Invoke(ctx, "/clutch.k8s.v1.K8sAPI/DescribeConfigMap", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *k8SAPIClient) DeleteConfigMap(ctx context.Context, in *DeleteConfigMapRequest, opts ...grpc.CallOption) (*DeleteConfigMapResponse, error) {
	out := new(DeleteConfigMapResponse)
	err := c.cc.Invoke(ctx, "/clutch.k8s.v1.K8sAPI/DeleteConfigMap", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// K8SAPIServer is the server API for K8SAPI service.
// All implementations should embed UnimplementedK8SAPIServer
// for forward compatibility
type K8SAPIServer interface {
	DescribePod(context.Context, *DescribePodRequest) (*DescribePodResponse, error)
	ListPods(context.Context, *ListPodsRequest) (*ListPodsResponse, error)
	DeletePod(context.Context, *DeletePodRequest) (*DeletePodResponse, error)
	UpdatePod(context.Context, *UpdatePodRequest) (*UpdatePodResponse, error)
	ResizeHPA(context.Context, *ResizeHPARequest) (*ResizeHPAResponse, error)
	DeleteHPA(context.Context, *DeleteHPARequest) (*DeleteHPAResponse, error)
	UpdateDeployment(context.Context, *UpdateDeploymentRequest) (*UpdateDeploymentResponse, error)
	DeleteDeployment(context.Context, *DeleteDeploymentRequest) (*DeleteDeploymentResponse, error)
	DescribeService(context.Context, *DescribeServiceRequest) (*DescribeServiceResponse, error)
	DeleteService(context.Context, *DeleteServiceRequest) (*DeleteServiceResponse, error)
	DescribeStatefulSet(context.Context, *DescribeStatefulSetRequest) (*DescribeStatefulSetResponse, error)
	UpdateStatefulSet(context.Context, *UpdateStatefulSetRequest) (*UpdateStatefulSetResponse, error)
	DeleteStatefulSet(context.Context, *DeleteStatefulSetRequest) (*DeleteStatefulSetResponse, error)
	DescribeCronJob(context.Context, *DescribeCronJobRequest) (*DescribeCronJobResponse, error)
	DeleteCronJob(context.Context, *DeleteCronJobRequest) (*DeleteCronJobResponse, error)
	ListConfigMaps(context.Context, *ListConfigMapsRequest) (*ListConfigMapsResponse, error)
	DescribeConfigMap(context.Context, *DescribeConfigMapRequest) (*DescribeConfigMapResponse, error)
	DeleteConfigMap(context.Context, *DeleteConfigMapRequest) (*DeleteConfigMapResponse, error)
}

// UnimplementedK8SAPIServer should be embedded to have forward compatible implementations.
type UnimplementedK8SAPIServer struct {
}

func (UnimplementedK8SAPIServer) DescribePod(context.Context, *DescribePodRequest) (*DescribePodResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribePod not implemented")
}
func (UnimplementedK8SAPIServer) ListPods(context.Context, *ListPodsRequest) (*ListPodsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPods not implemented")
}
func (UnimplementedK8SAPIServer) DeletePod(context.Context, *DeletePodRequest) (*DeletePodResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePod not implemented")
}
func (UnimplementedK8SAPIServer) UpdatePod(context.Context, *UpdatePodRequest) (*UpdatePodResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePod not implemented")
}
func (UnimplementedK8SAPIServer) ResizeHPA(context.Context, *ResizeHPARequest) (*ResizeHPAResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResizeHPA not implemented")
}
func (UnimplementedK8SAPIServer) DeleteHPA(context.Context, *DeleteHPARequest) (*DeleteHPAResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteHPA not implemented")
}
func (UnimplementedK8SAPIServer) UpdateDeployment(context.Context, *UpdateDeploymentRequest) (*UpdateDeploymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDeployment not implemented")
}
func (UnimplementedK8SAPIServer) DeleteDeployment(context.Context, *DeleteDeploymentRequest) (*DeleteDeploymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDeployment not implemented")
}
func (UnimplementedK8SAPIServer) DescribeService(context.Context, *DescribeServiceRequest) (*DescribeServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeService not implemented")
}
func (UnimplementedK8SAPIServer) DeleteService(context.Context, *DeleteServiceRequest) (*DeleteServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteService not implemented")
}
func (UnimplementedK8SAPIServer) DescribeStatefulSet(context.Context, *DescribeStatefulSetRequest) (*DescribeStatefulSetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeStatefulSet not implemented")
}
func (UnimplementedK8SAPIServer) UpdateStatefulSet(context.Context, *UpdateStatefulSetRequest) (*UpdateStatefulSetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateStatefulSet not implemented")
}
func (UnimplementedK8SAPIServer) DeleteStatefulSet(context.Context, *DeleteStatefulSetRequest) (*DeleteStatefulSetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteStatefulSet not implemented")
}
func (UnimplementedK8SAPIServer) DescribeCronJob(context.Context, *DescribeCronJobRequest) (*DescribeCronJobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeCronJob not implemented")
}
func (UnimplementedK8SAPIServer) DeleteCronJob(context.Context, *DeleteCronJobRequest) (*DeleteCronJobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCronJob not implemented")
}
func (UnimplementedK8SAPIServer) ListConfigMaps(context.Context, *ListConfigMapsRequest) (*ListConfigMapsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListConfigMaps not implemented")
}
func (UnimplementedK8SAPIServer) DescribeConfigMap(context.Context, *DescribeConfigMapRequest) (*DescribeConfigMapResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeConfigMap not implemented")
}
func (UnimplementedK8SAPIServer) DeleteConfigMap(context.Context, *DeleteConfigMapRequest) (*DeleteConfigMapResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteConfigMap not implemented")
}

// UnsafeK8SAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to K8SAPIServer will
// result in compilation errors.
type UnsafeK8SAPIServer interface {
	mustEmbedUnimplementedK8SAPIServer()
}

func RegisterK8SAPIServer(s grpc.ServiceRegistrar, srv K8SAPIServer) {
	s.RegisterService(&K8SAPI_ServiceDesc, srv)
}

func _K8SAPI_DescribePod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribePodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(K8SAPIServer).DescribePod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clutch.k8s.v1.K8sAPI/DescribePod",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(K8SAPIServer).DescribePod(ctx, req.(*DescribePodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _K8SAPI_ListPods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPodsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(K8SAPIServer).ListPods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clutch.k8s.v1.K8sAPI/ListPods",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(K8SAPIServer).ListPods(ctx, req.(*ListPodsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _K8SAPI_DeletePod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(K8SAPIServer).DeletePod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clutch.k8s.v1.K8sAPI/DeletePod",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(K8SAPIServer).DeletePod(ctx, req.(*DeletePodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _K8SAPI_UpdatePod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(K8SAPIServer).UpdatePod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clutch.k8s.v1.K8sAPI/UpdatePod",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(K8SAPIServer).UpdatePod(ctx, req.(*UpdatePodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _K8SAPI_ResizeHPA_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResizeHPARequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(K8SAPIServer).ResizeHPA(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clutch.k8s.v1.K8sAPI/ResizeHPA",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(K8SAPIServer).ResizeHPA(ctx, req.(*ResizeHPARequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _K8SAPI_DeleteHPA_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteHPARequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(K8SAPIServer).DeleteHPA(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clutch.k8s.v1.K8sAPI/DeleteHPA",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(K8SAPIServer).DeleteHPA(ctx, req.(*DeleteHPARequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _K8SAPI_UpdateDeployment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDeploymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(K8SAPIServer).UpdateDeployment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clutch.k8s.v1.K8sAPI/UpdateDeployment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(K8SAPIServer).UpdateDeployment(ctx, req.(*UpdateDeploymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _K8SAPI_DeleteDeployment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDeploymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(K8SAPIServer).DeleteDeployment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clutch.k8s.v1.K8sAPI/DeleteDeployment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(K8SAPIServer).DeleteDeployment(ctx, req.(*DeleteDeploymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _K8SAPI_DescribeService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(K8SAPIServer).DescribeService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clutch.k8s.v1.K8sAPI/DescribeService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(K8SAPIServer).DescribeService(ctx, req.(*DescribeServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _K8SAPI_DeleteService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(K8SAPIServer).DeleteService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clutch.k8s.v1.K8sAPI/DeleteService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(K8SAPIServer).DeleteService(ctx, req.(*DeleteServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _K8SAPI_DescribeStatefulSet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeStatefulSetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(K8SAPIServer).DescribeStatefulSet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clutch.k8s.v1.K8sAPI/DescribeStatefulSet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(K8SAPIServer).DescribeStatefulSet(ctx, req.(*DescribeStatefulSetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _K8SAPI_UpdateStatefulSet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateStatefulSetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(K8SAPIServer).UpdateStatefulSet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clutch.k8s.v1.K8sAPI/UpdateStatefulSet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(K8SAPIServer).UpdateStatefulSet(ctx, req.(*UpdateStatefulSetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _K8SAPI_DeleteStatefulSet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteStatefulSetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(K8SAPIServer).DeleteStatefulSet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clutch.k8s.v1.K8sAPI/DeleteStatefulSet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(K8SAPIServer).DeleteStatefulSet(ctx, req.(*DeleteStatefulSetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _K8SAPI_DescribeCronJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeCronJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(K8SAPIServer).DescribeCronJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clutch.k8s.v1.K8sAPI/DescribeCronJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(K8SAPIServer).DescribeCronJob(ctx, req.(*DescribeCronJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _K8SAPI_DeleteCronJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCronJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(K8SAPIServer).DeleteCronJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clutch.k8s.v1.K8sAPI/DeleteCronJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(K8SAPIServer).DeleteCronJob(ctx, req.(*DeleteCronJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _K8SAPI_ListConfigMaps_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListConfigMapsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(K8SAPIServer).ListConfigMaps(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clutch.k8s.v1.K8sAPI/ListConfigMaps",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(K8SAPIServer).ListConfigMaps(ctx, req.(*ListConfigMapsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _K8SAPI_DescribeConfigMap_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeConfigMapRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(K8SAPIServer).DescribeConfigMap(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clutch.k8s.v1.K8sAPI/DescribeConfigMap",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(K8SAPIServer).DescribeConfigMap(ctx, req.(*DescribeConfigMapRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _K8SAPI_DeleteConfigMap_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteConfigMapRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(K8SAPIServer).DeleteConfigMap(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clutch.k8s.v1.K8sAPI/DeleteConfigMap",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(K8SAPIServer).DeleteConfigMap(ctx, req.(*DeleteConfigMapRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// K8SAPI_ServiceDesc is the grpc.ServiceDesc for K8SAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var K8SAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "clutch.k8s.v1.K8sAPI",
	HandlerType: (*K8SAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DescribePod",
			Handler:    _K8SAPI_DescribePod_Handler,
		},
		{
			MethodName: "ListPods",
			Handler:    _K8SAPI_ListPods_Handler,
		},
		{
			MethodName: "DeletePod",
			Handler:    _K8SAPI_DeletePod_Handler,
		},
		{
			MethodName: "UpdatePod",
			Handler:    _K8SAPI_UpdatePod_Handler,
		},
		{
			MethodName: "ResizeHPA",
			Handler:    _K8SAPI_ResizeHPA_Handler,
		},
		{
			MethodName: "DeleteHPA",
			Handler:    _K8SAPI_DeleteHPA_Handler,
		},
		{
			MethodName: "UpdateDeployment",
			Handler:    _K8SAPI_UpdateDeployment_Handler,
		},
		{
			MethodName: "DeleteDeployment",
			Handler:    _K8SAPI_DeleteDeployment_Handler,
		},
		{
			MethodName: "DescribeService",
			Handler:    _K8SAPI_DescribeService_Handler,
		},
		{
			MethodName: "DeleteService",
			Handler:    _K8SAPI_DeleteService_Handler,
		},
		{
			MethodName: "DescribeStatefulSet",
			Handler:    _K8SAPI_DescribeStatefulSet_Handler,
		},
		{
			MethodName: "UpdateStatefulSet",
			Handler:    _K8SAPI_UpdateStatefulSet_Handler,
		},
		{
			MethodName: "DeleteStatefulSet",
			Handler:    _K8SAPI_DeleteStatefulSet_Handler,
		},
		{
			MethodName: "DescribeCronJob",
			Handler:    _K8SAPI_DescribeCronJob_Handler,
		},
		{
			MethodName: "DeleteCronJob",
			Handler:    _K8SAPI_DeleteCronJob_Handler,
		},
		{
			MethodName: "ListConfigMaps",
			Handler:    _K8SAPI_ListConfigMaps_Handler,
		},
		{
			MethodName: "DescribeConfigMap",
			Handler:    _K8SAPI_DescribeConfigMap_Handler,
		},
		{
			MethodName: "DeleteConfigMap",
			Handler:    _K8SAPI_DeleteConfigMap_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "k8s/v1/k8s.proto",
}

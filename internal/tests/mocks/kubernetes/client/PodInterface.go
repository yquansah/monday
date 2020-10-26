// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	rest "k8s.io/client-go/rest"

	types "k8s.io/apimachinery/pkg/types"

	v1 "k8s.io/api/core/v1"

	v1beta1 "k8s.io/api/policy/v1beta1"

	watch "k8s.io/apimachinery/pkg/watch"
)

// PodInterface is an autogenerated mock type for the PodInterface type
type PodInterface struct {
	mock.Mock
}

// Bind provides a mock function with given fields: binding
func (_m *PodInterface) Bind(binding *v1.Binding) error {
	ret := _m.Called(binding)

	var r0 error
	if rf, ok := ret.Get(0).(func(*v1.Binding) error); ok {
		r0 = rf(binding)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Create provides a mock function with given fields: _a0
func (_m *PodInterface) Create(_a0 *v1.Pod) (*v1.Pod, error) {
	ret := _m.Called(_a0)

	var r0 *v1.Pod
	if rf, ok := ret.Get(0).(func(*v1.Pod) *v1.Pod); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.Pod)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*v1.Pod) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: name, options
func (_m *PodInterface) Delete(name string, options *metav1.DeleteOptions) error {
	ret := _m.Called(name, options)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, *metav1.DeleteOptions) error); ok {
		r0 = rf(name, options)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteCollection provides a mock function with given fields: options, listOptions
func (_m *PodInterface) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	ret := _m.Called(options, listOptions)

	var r0 error
	if rf, ok := ret.Get(0).(func(*metav1.DeleteOptions, metav1.ListOptions) error); ok {
		r0 = rf(options, listOptions)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Evict provides a mock function with given fields: eviction
func (_m *PodInterface) Evict(eviction *v1beta1.Eviction) error {
	ret := _m.Called(eviction)

	var r0 error
	if rf, ok := ret.Get(0).(func(*v1beta1.Eviction) error); ok {
		r0 = rf(eviction)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: name, options
func (_m *PodInterface) Get(name string, options metav1.GetOptions) (*v1.Pod, error) {
	ret := _m.Called(name, options)

	var r0 *v1.Pod
	if rf, ok := ret.Get(0).(func(string, metav1.GetOptions) *v1.Pod); ok {
		r0 = rf(name, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.Pod)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, metav1.GetOptions) error); ok {
		r1 = rf(name, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLogs provides a mock function with given fields: name, opts
func (_m *PodInterface) GetLogs(name string, opts *v1.PodLogOptions) *rest.Request {
	ret := _m.Called(name, opts)

	var r0 *rest.Request
	if rf, ok := ret.Get(0).(func(string, *v1.PodLogOptions) *rest.Request); ok {
		r0 = rf(name, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rest.Request)
		}
	}

	return r0
}

// List provides a mock function with given fields: opts
func (_m *PodInterface) List(opts metav1.ListOptions) (*v1.PodList, error) {
	ret := _m.Called(opts)

	var r0 *v1.PodList
	if rf, ok := ret.Get(0).(func(metav1.ListOptions) *v1.PodList); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.PodList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(metav1.ListOptions) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Patch provides a mock function with given fields: name, pt, data, subresources
func (_m *PodInterface) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (*v1.Pod, error) {
	_va := make([]interface{}, len(subresources))
	for _i := range subresources {
		_va[_i] = subresources[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, name, pt, data)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *v1.Pod
	if rf, ok := ret.Get(0).(func(string, types.PatchType, []byte, ...string) *v1.Pod); ok {
		r0 = rf(name, pt, data, subresources...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.Pod)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, types.PatchType, []byte, ...string) error); ok {
		r1 = rf(name, pt, data, subresources...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: _a0
func (_m *PodInterface) Update(_a0 *v1.Pod) (*v1.Pod, error) {
	ret := _m.Called(_a0)

	var r0 *v1.Pod
	if rf, ok := ret.Get(0).(func(*v1.Pod) *v1.Pod); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.Pod)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*v1.Pod) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateStatus provides a mock function with given fields: _a0
func (_m *PodInterface) UpdateStatus(_a0 *v1.Pod) (*v1.Pod, error) {
	ret := _m.Called(_a0)

	var r0 *v1.Pod
	if rf, ok := ret.Get(0).(func(*v1.Pod) *v1.Pod); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.Pod)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*v1.Pod) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Watch provides a mock function with given fields: opts
func (_m *PodInterface) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	ret := _m.Called(opts)

	var r0 watch.Interface
	if rf, ok := ret.Get(0).(func(metav1.ListOptions) watch.Interface); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(watch.Interface)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(metav1.ListOptions) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

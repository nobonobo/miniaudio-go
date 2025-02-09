package miniaudio

import "unsafe"

// ma_backend_callbacks
type BackendCallbacks struct {
	OnContextInit             func(context *Context, config *ContextConfig, callbacks *BackendCallbacks) Result                                            // ma_result (*)(ma_context*, const ma_context_config*, ma_backend_callbacks*)
	OnContextUninit           func(context *Context) Result                                                                                                // ma_result (*)(ma_context*)
	OnContextEnumerateDevices func(context *Context, callback Proc, userData unsafe.Pointer) Result                                                        // ma_result (*)(ma_context*, ma_enum_devices_callback_proc, void*)
	OnContextGetDeviceInfo    func(context *Context, deviceType DeviceType, deviceID *DeviceID, deviceInfo *DeviceInfo) Result                             // ma_result (*)(ma_context*, ma_device_type, const ma_device_id*, ma_device_info*)
	OnDeviceInit              func(device *Device, config *DeviceConfig, descriptorPlayback *DeviceDescriptor, descriptorCapture *DeviceDescriptor) Result // ma_result (*)(ma_device*, const ma_device_config*, ma_device_descriptor*, ma_device_descriptor*)
	OnDeviceUninit            func(device *Device) Result                                                                                                  // ma_result (*)(ma_device*)
	OnDeviceStart             func(device *Device) Result                                                                                                  // ma_result (*)(ma_device*)
	OnDeviceStop              func(device *Device) Result                                                                                                  // ma_result (*)(ma_device*)
	OnDeviceRead              func(device *Device, frames unsafe.Pointer, frameCount uint32, framesRead *uint32) Result                                    // ma_result (*)(ma_device*, void*, ma_uint32, ma_uint32*)
	OnDeviceWrite             func(device *Device, frames unsafe.Pointer, frameCount uint32, framesWritten *uint32) Result                                 // ma_result (*)(ma_device*, const void*, ma_uint32, ma_uint32*)
	OnDeviceDataLoop          func(device *Device) Result                                                                                                  // ma_result (*)(ma_device*)
	OnDeviceDataLoopWakeup    func(device *Device) Result                                                                                                  // ma_result (*)(ma_device*)
	OnDeviceGetInfo           func(device *Device, deviceType DeviceType, deviceInfo *DeviceInfo) Result                                                   // ma_result (*)(ma_device*, ma_device_type, ma_device_info*)
}

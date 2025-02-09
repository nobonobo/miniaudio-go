//go:build linux

package miniaudio

import "unsafe"

// ma_context
type Context struct {
	Callbacks               BackendCallbacks
	Backend                 Backend             // ma_backend
	Log                     *Log                // ma_log*
	OwnedLog                Log                 // Embedded log if owned by the context
	ThreadPriority          ThreadPriority      // ma_thread_priority
	ThreadStackSize         uintptr             // size_t
	UserData                unsafe.Pointer      // void*
	AllocationCallbacks     AllocationCallbacks // ma_allocation_callbacks
	DeviceEnumLock          Mutex               // ma_mutex
	DeviceInfoLock          Mutex               // ma_mutex
	DeviceInfoCapacity      uint32              // ma_uint32
	PlaybackDeviceInfoCount uint32              // ma_uint32
	CaptureDeviceInfoCount  uint32              // ma_uint32
	DeviceInfos             *DeviceInfo         // ma_device_info*

	PlatformUnion unsafe.Pointer // Union field to be cast to one of ALSA, Pulse, or Jack structs

	Posix struct {
		Unused int
	}
}

// ALSA backend (if supported)
type ContextALSA struct {
	AsoundSO    uintptr // ma_handle
	SndPcmOpen  Proc    // ma_proc
	SndPcmClose Proc    // ma_proc
	// Add other ALSA fields here
	InternalDeviceEnumLock Mutex  // ma_mutex
	UseVerboseDeviceEnum   Bool32 // ma_bool32
}

// PulseAudio backend (if supported)
type ContextPulse struct {
	PulseSO       uintptr // ma_handle
	PAMainLoopNew Proc    // ma_proc
	// Add other Pulse fields here
	PMainLoop       uintptr // pa_mainloop*
	PPulseContext   uintptr // pa_context*
	ApplicationName *byte   // *char
	ServerName      *byte   // *char
}

// JACK backend (if supported)
type ContextJack struct {
	JackSO         uintptr // ma_handle
	JackClientOpen Proc    // ma_proc
	// Add other JACK fields here
	PClientName    *byte  // *char
	TryStartServer Bool32 // ma_bool32
}

// struct ma_context
// {
//     ma_backend_callbacks callbacks;
//     ma_backend backend;                 /* DirectSound, ALSA, etc. */
//     ma_log* pLog;
//     ma_log log; /* Only used if the log is owned by the context. The pLog member will be set to &log in this case. */
//     ma_thread_priority threadPriority;
//     size_t threadStackSize;
//     void* pUserData;
//     ma_allocation_callbacks allocationCallbacks;
//     ma_mutex deviceEnumLock;            /* Used to make ma_context_get_devices() thread safe. */
//     ma_mutex deviceInfoLock;            /* Used to make ma_context_get_device_info() thread safe. */
//     ma_uint32 deviceInfoCapacity;       /* Total capacity of pDeviceInfos. */
//     ma_uint32 playbackDeviceInfoCount;
//     ma_uint32 captureDeviceInfoCount;
//     ma_device_info* pDeviceInfos;       /* Playback devices first, then capture. */

//     union
//     {
// #ifdef MA_SUPPORT_ALSA
//         struct
//         {
//             ma_handle asoundSO;
//             ma_proc snd_pcm_open;
//             ma_proc snd_pcm_close;
//             ma_proc snd_pcm_hw_params_sizeof;
//             ma_proc snd_pcm_hw_params_any;
//             ma_proc snd_pcm_hw_params_set_format;
//             ma_proc snd_pcm_hw_params_set_format_first;
//             ma_proc snd_pcm_hw_params_get_format_mask;
//             ma_proc snd_pcm_hw_params_set_channels;
//             ma_proc snd_pcm_hw_params_set_channels_near;
//             ma_proc snd_pcm_hw_params_set_channels_minmax;
//             ma_proc snd_pcm_hw_params_set_rate_resample;
//             ma_proc snd_pcm_hw_params_set_rate;
//             ma_proc snd_pcm_hw_params_set_rate_near;
//             ma_proc snd_pcm_hw_params_set_buffer_size_near;
//             ma_proc snd_pcm_hw_params_set_periods_near;
//             ma_proc snd_pcm_hw_params_set_access;
//             ma_proc snd_pcm_hw_params_get_format;
//             ma_proc snd_pcm_hw_params_get_channels;
//             ma_proc snd_pcm_hw_params_get_channels_min;
//             ma_proc snd_pcm_hw_params_get_channels_max;
//             ma_proc snd_pcm_hw_params_get_rate;
//             ma_proc snd_pcm_hw_params_get_rate_min;
//             ma_proc snd_pcm_hw_params_get_rate_max;
//             ma_proc snd_pcm_hw_params_get_buffer_size;
//             ma_proc snd_pcm_hw_params_get_periods;
//             ma_proc snd_pcm_hw_params_get_access;
//             ma_proc snd_pcm_hw_params_test_format;
//             ma_proc snd_pcm_hw_params_test_channels;
//             ma_proc snd_pcm_hw_params_test_rate;
//             ma_proc snd_pcm_hw_params;
//             ma_proc snd_pcm_sw_params_sizeof;
//             ma_proc snd_pcm_sw_params_current;
//             ma_proc snd_pcm_sw_params_get_boundary;
//             ma_proc snd_pcm_sw_params_set_avail_min;
//             ma_proc snd_pcm_sw_params_set_start_threshold;
//             ma_proc snd_pcm_sw_params_set_stop_threshold;
//             ma_proc snd_pcm_sw_params;
//             ma_proc snd_pcm_format_mask_sizeof;
//             ma_proc snd_pcm_format_mask_test;
//             ma_proc snd_pcm_get_chmap;
//             ma_proc snd_pcm_state;
//             ma_proc snd_pcm_prepare;
//             ma_proc snd_pcm_start;
//             ma_proc snd_pcm_drop;
//             ma_proc snd_pcm_drain;
//             ma_proc snd_pcm_reset;
//             ma_proc snd_device_name_hint;
//             ma_proc snd_device_name_get_hint;
//             ma_proc snd_card_get_index;
//             ma_proc snd_device_name_free_hint;
//             ma_proc snd_pcm_mmap_begin;
//             ma_proc snd_pcm_mmap_commit;
//             ma_proc snd_pcm_recover;
//             ma_proc snd_pcm_readi;
//             ma_proc snd_pcm_writei;
//             ma_proc snd_pcm_avail;
//             ma_proc snd_pcm_avail_update;
//             ma_proc snd_pcm_wait;
//             ma_proc snd_pcm_nonblock;
//             ma_proc snd_pcm_info;
//             ma_proc snd_pcm_info_sizeof;
//             ma_proc snd_pcm_info_get_name;
//             ma_proc snd_pcm_poll_descriptors;
//             ma_proc snd_pcm_poll_descriptors_count;
//             ma_proc snd_pcm_poll_descriptors_revents;
//             ma_proc snd_config_update_free_global;

//             ma_mutex internalDeviceEnumLock;
//             ma_bool32 useVerboseDeviceEnumeration;
//         } alsa;
// #endif
// #ifdef MA_SUPPORT_PULSEAUDIO
//         struct
//         {
//             ma_handle pulseSO;
//             ma_proc pa_mainloop_new;
//             ma_proc pa_mainloop_free;
//             ma_proc pa_mainloop_quit;
//             ma_proc pa_mainloop_get_api;
//             ma_proc pa_mainloop_iterate;
//             ma_proc pa_mainloop_wakeup;
//             ma_proc pa_threaded_mainloop_new;
//             ma_proc pa_threaded_mainloop_free;
//             ma_proc pa_threaded_mainloop_start;
//             ma_proc pa_threaded_mainloop_stop;
//             ma_proc pa_threaded_mainloop_lock;
//             ma_proc pa_threaded_mainloop_unlock;
//             ma_proc pa_threaded_mainloop_wait;
//             ma_proc pa_threaded_mainloop_signal;
//             ma_proc pa_threaded_mainloop_accept;
//             ma_proc pa_threaded_mainloop_get_retval;
//             ma_proc pa_threaded_mainloop_get_api;
//             ma_proc pa_threaded_mainloop_in_thread;
//             ma_proc pa_threaded_mainloop_set_name;
//             ma_proc pa_context_new;
//             ma_proc pa_context_unref;
//             ma_proc pa_context_connect;
//             ma_proc pa_context_disconnect;
//             ma_proc pa_context_set_state_callback;
//             ma_proc pa_context_get_state;
//             ma_proc pa_context_get_sink_info_list;
//             ma_proc pa_context_get_source_info_list;
//             ma_proc pa_context_get_sink_info_by_name;
//             ma_proc pa_context_get_source_info_by_name;
//             ma_proc pa_operation_unref;
//             ma_proc pa_operation_get_state;
//             ma_proc pa_channel_map_init_extend;
//             ma_proc pa_channel_map_valid;
//             ma_proc pa_channel_map_compatible;
//             ma_proc pa_stream_new;
//             ma_proc pa_stream_unref;
//             ma_proc pa_stream_connect_playback;
//             ma_proc pa_stream_connect_record;
//             ma_proc pa_stream_disconnect;
//             ma_proc pa_stream_get_state;
//             ma_proc pa_stream_get_sample_spec;
//             ma_proc pa_stream_get_channel_map;
//             ma_proc pa_stream_get_buffer_attr;
//             ma_proc pa_stream_set_buffer_attr;
//             ma_proc pa_stream_get_device_name;
//             ma_proc pa_stream_set_write_callback;
//             ma_proc pa_stream_set_read_callback;
//             ma_proc pa_stream_set_suspended_callback;
//             ma_proc pa_stream_set_moved_callback;
//             ma_proc pa_stream_is_suspended;
//             ma_proc pa_stream_flush;
//             ma_proc pa_stream_drain;
//             ma_proc pa_stream_is_corked;
//             ma_proc pa_stream_cork;
//             ma_proc pa_stream_trigger;
//             ma_proc pa_stream_begin_write;
//             ma_proc pa_stream_write;
//             ma_proc pa_stream_peek;
//             ma_proc pa_stream_drop;
//             ma_proc pa_stream_writable_size;
//             ma_proc pa_stream_readable_size;

//             /*pa_mainloop**/ ma_ptr pMainLoop;
//             /*pa_context**/ ma_ptr pPulseContext;
//             char* pApplicationName; /* Set when the context is initialized. Used by devices for their local pa_context objects. */
//             char* pServerName;      /* Set when the context is initialized. Used by devices for their local pa_context objects. */
//         } pulse;
// #endif
// #ifdef MA_SUPPORT_JACK
//         struct
//         {
//             ma_handle jackSO;
//             ma_proc jack_client_open;
//             ma_proc jack_client_close;
//             ma_proc jack_client_name_size;
//             ma_proc jack_set_process_callback;
//             ma_proc jack_set_buffer_size_callback;
//             ma_proc jack_on_shutdown;
//             ma_proc jack_get_sample_rate;
//             ma_proc jack_get_buffer_size;
//             ma_proc jack_get_ports;
//             ma_proc jack_activate;
//             ma_proc jack_deactivate;
//             ma_proc jack_connect;
//             ma_proc jack_port_register;
//             ma_proc jack_port_name;
//             ma_proc jack_port_get_buffer;
//             ma_proc jack_free;

//             char* pClientName;
//             ma_bool32 tryStartServer;
//         } jack;
// #endif
//     };

// 	struct
// 	{
// 		int _unused;
// 	} posix;
// };

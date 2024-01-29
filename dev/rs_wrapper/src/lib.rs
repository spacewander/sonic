use flat::Document;
use sonic_rs::private::config::Config;
use sonic_rs::private::flat;

#[repr(C)]
pub struct Dom {
    dom: *mut flat::Document,
    str_buf: *const u8,
    node: *const flat::Value,

    // TODO: should export error msg in Golang
    error_offset: i64,
    error_msg: *mut u8,
    error_msg_len: u64,
    error_msg_cap: u64,
}

const F_USE_NUMBER: u64 = 1 << 2;
const F_VALIDATE_STRING: u64 = 1 << 2;

/// # Safety
/// FFI wrapper.
#[no_mangle]
pub unsafe extern "C" fn sonic_rs_ffi_parse(json: *const u8, len: usize, config: u64) -> Dom {
    let json = std::slice::from_raw_parts(json, len);

    let config = Config {
        use_number: config & F_USE_NUMBER != 0,
        validate_string: config & F_VALIDATE_STRING != 0,
        disable_surrogates_error: true,
    };

    match flat::dom_from_slice_config(json, config) {
        Ok(dom) => {
            let dom = Box::into_raw(Box::new(dom));
            let str_buf = (*dom).json_buffer.as_ptr();
            let node = (*dom).root();
            Dom {
                dom,
                str_buf,
                node: node as *const _,
                error_offset: -1,
                error_msg: std::ptr::null_mut(),
                error_msg_len: 0,
                error_msg_cap: 0,
            }
        }
        Err(e) => {
            let mut msg = e.to_string();
            let error_msg = msg.as_mut_ptr();
            let error_msg_len = msg.len() as u64;
            let error_msg_cap = msg.capacity() as u64;
            std::mem::forget(msg);
            Dom {
                dom: std::ptr::null_mut(),
                node: std::ptr::null(),
                str_buf: std::ptr::null(),
                error_offset: e.offset() as i64,
                error_msg,
                error_msg_len,
                error_msg_cap,
            }
        }
    }
}

/// # Safety
/// FFI wrapper.
#[no_mangle]
pub unsafe extern "C" fn sonic_rs_ffi_free(dom: *mut Document, msg: *mut u8, msg_cap: u64) {
    unsafe {
        if !dom.is_null() {
            drop(Box::from_raw(dom));
        }

        if !msg.is_null() {
            let s = String::from_raw_parts(msg, 0, msg_cap as usize);
            drop(s);
        }
    }
}

extern crate libc;

#[no_mangle]
pub extern "C" fn dot_product(v1: *const i32, v2: *const i32, len: isize, result: *mut i32) {
    for i in 0..len {
        unsafe {
            let dot = *v1.offset(i) * (*v2.offset(i));
            *result.offset(i) = dot;
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]

    fn dot_product_works_when_typical() {
        let m1: &[i32] = &[10, 11, 12];
        let m2: &[i32] = &[20, 21, 22];
        let result: &mut [i32] = &mut [0, 0, 0];

        dot_product(m1.as_ptr(), m2.as_ptr(), 3, result.as_mut_ptr());

        assert!(result[0] == 200);
        assert!(result[1] == 231);
        assert!(result[2] == 264);
    }
}

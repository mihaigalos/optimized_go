extern crate libc;

#[no_mangle]
pub extern "C" fn dot_product(
    s1: *const i32,
    s2: *const i32,
    start: isize,
    stop: isize,
    result: *mut i32,
) {
    for i in start..stop {
        unsafe {
            let dot = *s1.offset(i) * (*s2.offset(i));
            *result.offset(i) = dot;
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn dot_product_works_when_typical() {
        let s1: &[i32] = &[10, 11, 12];
        let s2: &[i32] = &[20, 21, 22];
        let result: &mut [i32] = &mut [0, 0, 0];

        dot_product(s1.as_ptr(), s2.as_ptr(), 0, 3, result.as_mut_ptr());

        assert!(result[0] == 200);
        assert!(result[1] == 231);
        assert!(result[2] == 264);
    }
}

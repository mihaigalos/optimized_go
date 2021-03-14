use std::collections::HashMap;

pub fn dot_product(m1: &HashMap<u32, i32>, m2: &HashMap<u32, i32>, result: &mut HashMap<u32, i32>) {
    for i in 0..m1.len() {
        println!(
            "{}, {}",
            m1.get(&(i as u32)).unwrap(),
            m2.get(&(i as u32)).unwrap()
        );

        let dot = m1.get(&(i as u32)).unwrap() * m2.get(&(i as u32)).unwrap();
        result.insert(i as u32, dot);
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]

    fn dot_product_works() {
        let (mut m1, mut m2, mut result) = (HashMap::new(), HashMap::new(), HashMap::new());
        m1.insert(0, 10);
        m1.insert(1, 11);
        m1.insert(2, 12);
        m2.insert(0, 20);
        m2.insert(1, 21);
        m2.insert(2, 22);

        dot_product(&m1, &m2, &mut result);

        assert!(result.get(&0).unwrap() == &(200 as i32));
        assert!(result.get(&1).unwrap() == &(231 as i32));
        assert!(result.get(&2).unwrap() == &(264 as i32));
    }
}

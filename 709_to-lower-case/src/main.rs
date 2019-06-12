pub struct Solution {}


impl Solution {
    pub fn to_lower_case(str: String) -> String {
        let mut ret = String::from("");
        for c in str.bytes() {
            let mut d = 0;
            if c <= 90 && c >= 65 {  // Z A
                d = c + (97 -65);
                println!("{}", d);
                let e = &[d];
                let s = match std::str::from_utf8(e) {
                    Ok(v) => v,
                    Err(_e) => panic!("p")
                };

                ret.push_str(s);
            } else {
                let e = &[c];
                let s = match std::str::from_utf8(e) {
                    Ok(v) => v,
                    Err(_e) => panic!("p")
                };

                ret.push_str(s);

            }

        }

        return ret;
    }
}



#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test1() {
        let ret = Solution::to_lower_case("AAA".to_string());

        assert_eq!("aaa", ret);
    }
    #[test]
    fn test2() {
        let ret = Solution::to_lower_case("Hello".to_string());

        assert_eq!("hello", ret);
    }
}

pub fn reverse(input: &str) -> String {
    let mut ret = String::from("");
    input.chars().rev().for_each(
      |char| ret.push(char)
    );
    ret
}

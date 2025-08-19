use unicode_segmentation::UnicodeSegmentation;

pub fn reverse(input: &str) -> String {
    let mut ret = String::from("");
    UnicodeSegmentation::graphemes(input, true).rev().for_each(
      |char| ret.push_str(char)
    );
    ret
}

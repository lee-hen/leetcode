use two_sum;

fn main() {
    println!("Hello, leetcode!");
    let nums = vec![2, 7, 11, 15];
    let target = 9;
    println!("calling two_sum() get {:?}", two_sum::two_sum(nums, target));
}

use std::fs::File;
use std::io::{BufReader, BufRead};
fn main() {
    let p1_solution = p1();
    println!("Part 1 Solution: {}", p1_solution);
    let p2_solution = p2();
    println!("Part 2 Solution: {}", p2_solution);
}


fn p1() -> i64{
    let num_list = read_input("3-input.txt");

    let mut end_val = 0;
    
    for i in 0..num_list.len() {
        end_val += determine_max_joltage(2, &num_list[i]);
    }
    return end_val;
}

fn p2() -> i64{
    let num_list = read_input("3-input.txt");

    let mut end_val = 0;
    
    for i in 0..num_list.len() {
        end_val += determine_max_joltage(12, &num_list[i]);
    }
    return end_val;
}

fn read_input(file_name: &str) -> Vec<Vec<i8>>{
    //open the file in read-only mode
    let file = File::open(&file_name).expect("unable to open file");

    //create buffered reader
    let reader = BufReader::new(file);

    let mut v= Vec::new();

    //Read the file line by line using the lines() iterator and save to vector after converting to number
    for line in reader.lines(){
        let temp: Vec<i8> = line.expect("unable to read line").chars().map(|s| s.to_digit(10).unwrap() as i8).collect();
        v.push(temp);
    }
    return v;
}

fn determine_max_joltage(num_digits: i16, bank: &Vec<i8>) -> i64 {
    let mut joltage_digits: Vec<i8> = vec!();
    let mut last_index: i32 = 0;
    for i in 0..num_digits {
        let mid = &bank[(last_index) as usize..bank.len()-(num_digits as usize - i as usize - 1)];
        if let Some((index, max_val)) = mid.iter().rev().enumerate().max_by_key(|&(_, val)| val){
            joltage_digits.push(*max_val);
            last_index = last_index + (mid.len()-index) as i32;
        }
        else {
            println!("Empty Value");
        }
    }
    let mut joltage_value = 0;
    for i in 0..joltage_digits.len() {
        joltage_value += joltage_digits[i] as i64 * 10_i64.pow(num_digits as u32 - 1 - i as u32);
    }
    return joltage_value;
}
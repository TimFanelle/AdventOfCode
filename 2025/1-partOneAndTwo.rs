use std::fs::File;
use std::io::{BufReader, BufRead};
fn main() {
    let p1_solution = p1();
    println!("Part 1 Solution: {}", p1_solution);
    let p2_solution = p2();
    println!("Part 2 Solution: {}", p2_solution);
}

fn p1() -> i32{
    let num_list = read_input("1-input.txt");
    let mut cur_val = 50;
    let mut num_zero = 0;
    for i in 0..num_list.len(){
        (cur_val, _) = turn_dial(cur_val, &num_list[i]);
        if cur_val == 0 {
            num_zero = num_zero + 1;
        }
    }
    return num_zero;
}

fn p2() -> i32{
    let num_list = read_input("1-input.txt");
    let mut cur_val = 50;
    let mut num_zero = 0;
    let mut count ;
    for i in 0..num_list.len(){
        (cur_val, count) = turn_dial(cur_val, &num_list[i]);
        if cur_val == 0 {
            num_zero = num_zero + 1;
        }
        num_zero = num_zero + count
    }
    return num_zero;
}

fn read_input(file_name: &str) -> Vec<String>{
    //open the file in read-only mode
    let file = File::open(&file_name).expect("unable to open file");

    //create buffered reader
    let reader = BufReader::new(file);

    let mut v= Vec::new();

    //Read the file line by line using the lines() iterator and save to vector after converting to number
    for line in reader.lines(){
        v.push(line.expect("unable to read line"));
    }

    return v;
}

fn turn_dial(cur_value: i32, dial_instruct: &String) -> (i32, i32) {
    let head = dial_instruct.chars().next().unwrap();
    let mut tail: i32 = dial_instruct[1..].parse().expect("Failure");

    let mut count = 0;
    if tail > 100 {
        count += tail / 100;
        tail = tail % 100;
    }
    let end_value: i32;
    match head {
        'L' => {
            end_value = cur_value + (100-tail);
            if tail > cur_value && cur_value != 0{
                count += 1;
            }
        },
        'R' => {
            end_value = cur_value + tail;
            if end_value > 100 {
                count += 1;
            }
        },
        _ => end_value = 0
    }
    return (end_value % 100, count)
}
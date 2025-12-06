use std::fs::File;
use std::io::{BufReader, BufRead};

fn main() {
    let (values, operations) = read_input("6-input.txt");\
    let p1_solution = p1(&values, operations);
    println!("Part 1 Solution: {}", p1_solution);
    let p2_solution = p2();
    println!("Part 2 Solution: {}", p2_solution);
}

fn p1(values: &Vec<Vec<i32>>, operations: Vec<char>) -> i64{
    let mut end_val = 0;
    for i in 0..values.len() {
        if operations[i] == '+'{
            end_val += values[i].clone().into_iter().sum::<i32>() as i64;
        } else if operations[i] == '*' {
            let mut product = 1;
            for val in values[i].clone(){
                product *= val as i64;
            }
            end_val += product;
        } else {
            panic!("Non Valid Operation");
        }
    }
    return end_val;
}

fn p2() -> i64{
    let file = File::open("6-input.txt").expect("unable to open file");

    //create buffered reader
    let reader = BufReader::new(file);

    let mut lines = vec!();

    for line in reader.lines(){
        lines.push(line.expect("unable to read line").clone());
    }
    let longest_line_length = lines.iter().map(|l| l.len()).max().unwrap_or(0);
    let mut end_val: i64 = 0;

    let mut nums: Vec<i32> = vec!();
    for i in (0..longest_line_length).rev() {
        let strin: String = lines.iter().map(|s| s.chars().nth(i).unwrap_or(' ')).filter(|c| !c.is_whitespace()).collect();
        if let Some(check_char) = strin.chars().last(){
            if check_char == '*' || check_char == '+' {
                nums.push(strin[..strin.chars().count()-1].parse::<i32>().unwrap_or(0));
                //empty out vec
                match check_char {
                    '*' => {
                        let mut cur_val = 1;
                        for last in nums.drain(..){
                            cur_val *= last as i64;
                        }
                        end_val += cur_val as i64
                    },
                    '+' => {
                        let mut cur_val = 0;
                        for last in nums.drain(..){
                            cur_val += last;
                        }
                        end_val += cur_val as i64;

                    },
                    _ => panic!("Invalid operator")
                }
            } else {
                nums.push(strin.parse::<i32>().unwrap_or(0));
            }
        }
        
    }
    
    return end_val;
}

fn read_input(file_name: &str) -> (Vec<Vec<i32>>, Vec<char>){
    //open the file in read-only mode
    let file = File::open(&file_name).expect("unable to open file");

    //create buffered reader
    let reader = BufReader::new(file);

    let mut numbers = Vec::new();
    let mut operations = Vec::new();

    //Read the file line by line using the lines() iterator and save to vector after converting to number
    for line in reader.lines(){
        let temp = line.expect("unable to read line");
        let check = &temp[0..1];
        if check == "+" || check == "*"{
            let mut op_chars = temp.clone();
            op_chars.retain(|c| !c.is_whitespace());
            for char in op_chars.chars(){
                operations.push(char);
            }
        } else {
            let num_vals: Vec<i32> = temp.clone().split_whitespace().map(|v| v.parse::<i32>().unwrap()).collect();
            if numbers.len() == 0 {
                for _ in 0..num_vals.len(){
                    numbers.push(vec!());
                }
            }
            for i in 0..num_vals.len() {
                numbers[i].push(num_vals[i]);
            }
        }
    }
    return (numbers, operations);
}

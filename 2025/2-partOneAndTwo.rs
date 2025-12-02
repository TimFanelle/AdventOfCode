use std::fs::File;
use std::io::{BufReader, BufRead};
use std::fmt::Debug;
fn main() {
    let p1_solution = p1();
    println!("Part 1 Solution: {}", p1_solution);
    let p2_solution = p2();
    println!("Part 2 Solution: {}", p2_solution);
}

fn p1() -> i64{
    let num_list = read_input("2-input.txt");

    let mut end_val = 0;
    
    for i in 0..num_list.len() {
        for j in num_list[i].top..num_list[i].bottom+1 {
            let num_string = j.to_string();

            if validate_similarity(2, &num_string) {
                end_val += j;
            }
        }
    }
    return end_val;
}

fn p2() -> i64{
    let num_list = read_input("2-input.txt");

    let mut end_val = 0;
    
    for i in 0..num_list.len() {
        for j in num_list[i].top..num_list[i].bottom+1 {
            let num_string = j.to_string();

            for k in 2..num_string.chars().count()+1{
                if validate_similarity(k, &num_string) {
                    end_val += j;
                    break;
                }
            }
        }
    }
    return end_val;
}


#[derive(Debug)]
struct TB {
    top: i64,
    bottom: i64,
}

fn read_input(file_name: &str) -> Vec<TB>{
    //open the file in read-only mode
    let file = File::open(&file_name).expect("unable to open file");

    //create buffered reader
    let reader = BufReader::new(file);

    let mut v= Vec::new();

    //Read the file line by line using the lines() iterator and save to vector after converting to number
    for line in reader.lines(){
        v.push(line.expect("unable to read line"));
    }

    let mut tbs = Vec::new();
    let vals: Vec<&str> = v[0].split(",").collect();
    for i in 0..vals.len() {
        let temp: Vec<i64> = vals[i].split("-").map(|s| s.parse::<i64>().unwrap()).collect();
        tbs.push(TB{top: temp[0], bottom: temp[1]});
    }
    return tbs;
}

fn validate_similarity(split_amount: usize, full_s: &String) -> bool {
    if full_s.chars().count() % split_amount == 0 {
        let gen_len = full_s.chars().count()/split_amount;
        let s_in: String = full_s.chars().take(gen_len).collect();
        let mut cur_start = s_in.chars().count();
        while cur_start < full_s.chars().count() {
            let next_bit: String = full_s.chars().skip(cur_start).take(gen_len).collect();
            if next_bit != s_in {
                return false;
            }
            cur_start += gen_len;
        }
        return true
    }
    return false
}
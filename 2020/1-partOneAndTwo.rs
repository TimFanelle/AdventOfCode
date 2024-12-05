use std::fs::File;
use std::io::{BufReader, BufRead};


fn main(){
    let p1_solution = p1();
    println!("Part 1 solution: {}", p1_solution);
    let p2_solution = p2();
    println!("Part 2 solution: {}", p2_solution);
}

fn p1() -> i32{
    let num_list = read_input("1-input.txt");
    for i in 0..num_list.len(){
        for j in i..num_list.len(){
            if num_list[i] + num_list[j] == 2020{
                let outnum = num_list[i]*num_list[j];
                return outnum;
            }
        }
    }
    return 0;
}

fn p2() -> i32{
    let num_list = read_input("1-input.txt");
    for i in 0..num_list.len(){
        for j in i..num_list.len(){
            for k in j..num_list.len(){
                if num_list[i] + num_list[j] + num_list[k] == 2020{
                    let outnum = num_list[i]*num_list[j]*num_list[k];
                    return outnum;
                }
            }
        }
    }
    return 0;
}

fn read_input(file_name: &str) -> Vec<i32>{
    //open the file in read-only mode
    let file = File::open(&file_name).expect("unable to open file");

    //create buffered reader
    let reader = BufReader::new(file);

    let mut v: Vec<i32> = Vec::new();

    //Read the file line by line using the lines() iterator and save to vector after converting to number
    let mut temp: i32;
    for line in reader.lines(){
        temp = line.unwrap().parse().expect("Failure");
        v.push(temp);
    }

    return v;
}
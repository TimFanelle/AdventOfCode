use std::fs::File;
use std::io::{BufReader, BufRead};
use regex::Regex;


fn main(){
    let p1_solution = p1();
    println!("Part 1 solution: {}", p1_solution);
    let p2_solution = p2();
    println!("Part 2 solution: {}", p2_solution);
}

fn p1() -> i32 {
    let in_vec = read_input("2-input.txt");
    let total = in_vec.into_iter().filter(|val| validate_line_p1(val)).count() as i32;
    return total;
}

fn p2() -> i32 {
    let in_vec = read_input("2-input.txt");
    let total = in_vec.into_iter().filter(|val| validate_line_p2(val)).count() as i32;
    return total;
}

struct LineBreakdown {
    min_max: Vec<i32>,
    look_char: char,
    full_string: String
}
fn build_line_breakdown (mm: Vec<i32>, lc: char, fs: String) -> LineBreakdown{
    LineBreakdown{
        min_max: mm,
        look_char: lc,
        full_string: fs
    }
}

fn read_input(file_name: &str) -> Vec<LineBreakdown>{
    //open the file in read-only mode
    let file = File::open(&file_name).expect("unable to open file");

    //create buffered reader
    let reader = BufReader::new(file);

    let mut v: Vec<LineBreakdown> = Vec::new();

    //Read the file line by line using the lines() iterator and save to vector after converting to number
    let mut temp: String;

    let re_1 = Regex::new(r"\d*-\d*").unwrap();
    let re_2 = Regex::new(r"[a-z]:").unwrap();
    let re_3 = Regex::new(r"[a-z]{2,}").unwrap();
    for line in reader.lines(){
        temp = line.unwrap();
        let second_temp : Vec<i32> = re_1.find(&temp).unwrap().as_str().split("-").map(|s| s.parse().unwrap()).collect();
        let thirt_temp = re_2.find(&temp).unwrap().as_str();
        let third_temp : char = thirt_temp.chars().next().expect("weird char");
        let fourth_temp = re_3.find(&temp).unwrap().as_str();
        
        v.push(build_line_breakdown(second_temp, third_temp, fourth_temp.to_string()));
    }

    return v;
}

fn validate_line_p1(line_in: &LineBreakdown) -> bool {
    let count = line_in.full_string.chars().filter(|&c| c == line_in.look_char).count() as i32;
    
    return count >= line_in.min_max[0] && count <= line_in.min_max[1];
}

fn validate_line_p2(line_in: &LineBreakdown) -> bool {
    let a = line_in.full_string.chars().nth(line_in.min_max[0] as usize - 1 as usize) == Some(line_in.look_char);
    let b = line_in.full_string.chars().nth(line_in.min_max[1] as usize - 1 as usize) == Some(line_in.look_char);
    return (a || b) && !(a && b);
}
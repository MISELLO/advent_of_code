# Solutions for the Advent of Code event

This is the code (written in [Go](https://golang.org/)) I used to get the solutions for the [Advent of Code](https://adventofcode.com/) events. It may not be the best code, an optimal code or the fastest code. But it is my code.

## Badges

[![MIT License](https://img.shields.io/badge/License-MIT-green.svg)](https://choosealicense.com/licenses/mit/)
[![Go Report Card](https://goreportcard.com/badge/github.com/misello/advent_of_code)](https://goreportcard.com/report/github.com/misello/advent_of_code)
![Static Badge](https://img.shields.io/badge/stars%20%E2%AD%90-50-green)
![Static Badge](https://img.shields.io/badge/2023-☑-white)


## What is Advent of Code?

Advent of Code (or AoC) is an annual set of Christmas-themed computer programming challenges that follow an Advent calendar. It has been running since 2015. The programming puzzles cover a variety of skill sets and skill levels and can be solved using any programming language ([Wikipedia](https://en.wikipedia.org/wiki/Advent_of_Code)).

The most interesting part of it is not the answer itself (actually you can get it [here](https://aoc-puzzle-solver.streamlit.app/)). It has much more value what you code in order to get the answer and what you learn in the process. If you get stuck you can always search for hints on [Reddit](https://www.reddit.com/r/adventofcode/).

## Note

Please note that the problems definition/explanation and the inputs needed to make the code of this project work belong to [Advent of Code](https://adventofcode.com/) and not me, so I can't share it.

## Usage

Iside this repository you will find a folder for each year I have done. Inside this folder you will find one folder for each of the 25 days of the event. And inside the proper day you will find two [Go](https://golang.org/)) files named part1.go and part2.go. Each of this files corresponds to it's part of the puzzle.

If you want to execute one of this files you first need an input data as a text file. So, go to the [Advent of Code](https://adventofcode.com/) page and save any input, example or test of your own into a text file.

For example, to execute the first part of day 15 for 2023 just type:

```
git clone https://github.com/MISELLO/advent_of_code/
cd advent_of_code/2023/Day_15/
go run part1 /route_to_input_file/input.txt
```

## Thanks

Many thanks to [Eric Wastl](http://was.tl/) and all his team. Without you this would not be possible.

And if the reader has the chance, you should [support](https://adventofcode.com/support) him directly.
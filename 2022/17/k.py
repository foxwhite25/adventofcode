#!/usr/bin/env python3

from collections import defaultdict, Counter
from functools import lru_cache
import itertools
import math
import string
import sys
import timeit
import unittest

class Tests(unittest.TestCase):
	def setUp(self):
		self.testinput = """>>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>"""
		self.test = parse_input(self.testinput.split("\n"))

	def test_part1(self):
		self.assertEqual(part1(self.test), 3068)
		pass

	def test_part2(self):
		self.assertEqual(part2(self.test), 1514285714288)
		pass

def parse_input(i):
	for line in i:
		return list(c for c in line.strip())
	return

def part1(i):
	situation = [[0, 0, 0, 0, 0, 0, 0]]
	for _ in range(3):
		situation.append([0, 0, 0, 0, 0, 0, 0])
	curr_highest = -1
	rocknames = ["hline", "cross", "L", "vline", "square"]
	rock_width = {"hline": 4, "cross": 3, "L": 3, "vline": 1, "square": 2}
	rock_height = {"hline": 1, "cross": 3, "L": 3, "vline": 4, "square": 2}
	rocks = [[[1, 1, 1, 1]], [[0,1,0],[1,1,1],[0,1,0]],[[0,0,1],[0,0,1],[1,1,1]], [[1],[1],[1],[1]], [[1,1],[1,1]]]
	move_idx = 0
	for rock_num in range(2022):
		curr_rock = rocknames[rock_num % 5]
		curr_posns = get_posns(curr_highest + 1, curr_rock)
		for _ in range(rock_height[curr_rock]):
			situation.append([0, 0, 0, 0, 0, 0, 0])
		while True:
			if move_idx >= len(i):
				move_idx -= len(i)
			curr_posns = attempt_move(i[move_idx], curr_posns, situation)
			move_idx += 1
			fall_down, can_fall = fall_rock(curr_posns, situation)
			if not can_fall:
				for p in curr_posns:
					situation[p[0]][p[1]] = 1
					if p[0] > curr_highest:
						curr_highest = p[0]
				break
			curr_posns = fall_down
	return curr_highest + 1

def get_posns(h, name):
	if name == "hline":
		return [(h + 3, 2), (h + 3, 3), (h + 3, 4), (h + 3, 5)]
	elif name == "cross":
		return [(h + 3, 3), (h + 4, 2), (h + 4, 3), (h + 4, 4), (h + 5, 3)]
	elif name == "L":
		return [(h + 3, 2), (h + 3, 3), (h + 3, 4), (h + 4, 4), (h + 5, 4)]
	elif name == "vline":
		return [(h + 3, 2), (h + 4, 2), (h + 5, 2), (h + 6, 2)]
	elif name == "square":
		return [(h + 3, 2), (h + 3, 3), (h + 4, 2), (h + 4, 3)]
	else:
		assert False

def fall_rock(rock_posns, situation):
	new_posns = []
	for p in rock_posns:
		if p[0] == 0 or situation[p[0] - 1][p[1]] != 0:
			return None, False
		else:
			new_posns.append((p[0] - 1, p[1]))
	return new_posns, True

def attempt_move(move, rock_posns, situation):
	new_posns = []
	if move == "<":
		for p in rock_posns:
			if p[1] == 0 or situation[p[0]][p[1] - 1] != 0:
				return rock_posns
			new_posns.append((p[0], p[1] - 1))
		return new_posns
	elif move == ">":
		for p in rock_posns:
			if p[1] == 6 or situation[p[0]][p[1] + 1] != 0:
				return rock_posns
			new_posns.append((p[0], p[1] + 1))
		return new_posns
	else:
		assert False

def pretty_print(situation):
	t = ""
	for i in range(len(situation) - 1, -1, -1):
		o = "|"
		for c in situation[i]:
			if c == 0:
				o += "."
			else:
				o += "#"
		t += (o + "|\n")
	print(t)

def part2(i):
	situation = [[0, 0, 0, 0, 0, 0, 0]]
	for _ in range(3):
		situation.append([0, 0, 0, 0, 0, 0, 0])
	curr_highest = -1
	rocknames = ["hline", "cross", "L", "vline", "square"]
	rock_width = {"hline": 4, "cross": 3, "L": 3, "vline": 1, "square": 2}
	rock_height = {"hline": 1, "cross": 3, "L": 3, "vline": 4, "square": 2}
	rocks = [[[1, 1, 1, 1]], [[0,1,0],[1,1,1],[0,1,0]],[[0,0,1],[0,0,1],[1,1,1]], [[1],[1],[1],[1]], [[1,1],[1,1]]]
	move_idx = 0
	seven_peaks = [0,0,0,0,0,0,0]
	prev_seen = {}
	h_dict = {}
	for rock_num in range(1000000000000):
		curr_rock = rocknames[rock_num % 5]
		curr_posns = get_posns(curr_highest + 1, curr_rock)
		for _ in range(rock_height[curr_rock]):
			situation.append([0, 0, 0, 0, 0, 0, 0])
		while True:
			if move_idx >= len(i):
				move_idx -= len(i)
			curr_posns = attempt_move(i[move_idx], curr_posns, situation)
			move_idx += 1
			fall_down, can_fall = fall_rock(curr_posns, situation)
			if not can_fall:
				old_max = curr_highest
				for p in curr_posns:
					situation[p[0]][p[1]] = 1
					if p[0] > curr_highest:
						curr_highest = p[0]
				max_change = curr_highest - old_max
				for ind in range(7):
					seven_peaks[ind] -= max_change
				for p in curr_posns:
					this_peak = p[0] - curr_highest
					seven_peaks[p[1]] = max(seven_peaks[p[1]], this_peak)
				break
			curr_posns = fall_down
		h_dict[rock_num] = curr_highest
		k = (tuple(seven_peaks), move_idx, rock_num % 5)
		if k in prev_seen.keys():
			if prev_seen[k] != 0:
				then_count = prev_seen[k]
				highest_then = h_dict[prev_seen[k]]
				highest_now = curr_highest
				height_change = highest_now - highest_then
				cycle_size = rock_num - prev_seen[k]
				goal = 1000000000000
				goal -= prev_seen[k] # how many more rocks have to drop after cycling starts
				num_cycles = goal // cycle_size
				left_over = goal % cycle_size
				leftover_height = h_dict[then_count + left_over] - highest_then
				return highest_then + leftover_height + (num_cycles * height_change)
		prev_seen[k] = rock_num
	return curr_highest + 1

def main():
	t0 = timeit.default_timer()
	with open(sys.argv[1], "r") as f:
		i = parse_input(f.readlines())
	print(part1(i))
	t1 = timeit.default_timer()
	with open(sys.argv[1], "r") as f:
		i = parse_input(f.readlines())
	print(part2(i))
	t2 = timeit.default_timer()
	if len(sys.argv) > 2:
		print(f"Part 1: {t1 - t0} seconds")
		print(f"Part 2: {t2 - t1} seconds")

if __name__ == '__main__':
	if len(sys.argv) < 2:
		unittest.main()
	else:
		main()

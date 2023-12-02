import sys, os
import string
import numpy as np


def read_data(file_path: str):
    abs_path = os.path.abspath(file_path)
    with open(abs_path, "r") as f:
        data_ = f.read().strip()
    return data_

DIGITS = {
    "one": "1",
    "two": "2",
    "three": "3",
    "four": "4",
    "five": "5",
    "six": "6",
    "seven": "7",
    "eight": "8",
    "nine": "9"
}


if __name__ == "__main__":
    input_file_path = sys.argv[1]
    data = read_data(input_file_path)
    lines = data.splitlines()
    nums = []
    for idx, x in enumerate(lines):
        num = ""
        tmp_str = x
        first, fidx = "", np.inf
        last, lidx = "", np.inf
        first_digit = [tmp_str.find(y) for y in string.digits]
        first_digit = min([y for y in first_digit if y >= 0])
        last_digit = [tmp_str[::-1].find(y) for y in string.digits]
        last_digit = min([y for y in last_digit if y >= 0])
        assert first_digit >= 0 and last_digit >= 0
        num = [tmp_str[first_digit], tmp_str[::-1][last_digit][::-1]]
        for old, new in DIGITS.items():
            find_front = tmp_str.find(old)
            if find_front >= 0 and -1 < find_front < fidx:
                fidx = find_front
                first = old, new
            find_back = tmp_str[::-1].find(old[::-1])
            if find_back >= 0 and -1 < find_back < lidx:
                lidx = find_back
                last = old, new
        if fidx < np.inf and fidx < first_digit:
            num[0] = first[1]
        if lidx < np.inf and lidx < last_digit:
            num[1] = last[1]
        nums.append(num)
    print(sum([int("".join(x)) for x in nums]))

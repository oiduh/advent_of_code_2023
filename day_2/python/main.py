import sys, os

def read_data(file_path: str):
    with open(file_path, "r") as f:
        data_ = f.read().strip()
    return data_.split("\n")


if __name__ == "__main__":
    input_file_path = sys.argv[1]
    abs_path = os.path.abspath(input_file_path)
    data = read_data(abs_path)
    target = {"red": 12, "green": 13, "blue": 14}
    game_dict = {}
    for x in data:
        game_id, sets = x.split(": ")
        id_ = game_id.split(" ")[-1]
        sets = sets.split("; ")
        game_dict.update({id_: {"red": 0, "green": 0, "blue": 0}})
        for y in sets:
            colors = y.split(", ")
            for z in colors:
                amount, color = z.split(" ")
                amount_int = int(amount)
                current_amount = game_dict.get(id_)
                assert current_amount is not None
                new_amount = max(current_amount.get(color), amount_int)
                current_amount.update({color: new_amount})
                game_dict.update({id_: current_amount})

    filter_1 = dict(filter(
        lambda x: (
            x[1].get("red") <= target.get("red") and
            x[1].get("green") <= target.get("green") and
            x[1].get("blue") <= target.get("blue")
        ),
        game_dict.items()
    ))
    result_1 = sum(int(x) for x in filter_1.keys())
    print(result_1)

    filter_2 = dict(map(
        lambda x: (
                x[0], x[1].get("red") * x[1].get("green") * x[1].get("blue")
        ),
        game_dict.items()
    ))
    result_2 = sum(x for x in filter_2.values())
    print(result_2)

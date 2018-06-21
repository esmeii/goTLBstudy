"""Ground Truth"""

import re
import subprocess

import numpy as np
import pandas as pd

def run_benchmark(cmd, cwd):
    process = subprocess.Popen(cmd, shell=True, cwd=cwd,
                               stdout=subprocess.PIPE)
    (stdout, _) = process.communicate()

    m = re.search(r'Kernel [0-9\.]+ - [0-9\.]+: ([0-9\.]+)', str(stdout))
    return float(m.group(1))


def run():
    """ run benchmark and retuns a data frame that represents its result """
    data_columns = ['benchmark', 'arg', 'time']
    data = pd.DataFrame(columns=data_columns)

    process = subprocess.Popen("make", shell=True, cwd='.',
                            stdout=subprocess.DEVNULL)
    process.wait()

    for size in range(10000, 100000, 10000):
        print(size)
        for i in range(1, 21):
            time = run_benchmark('./kernel ' + str(size), '.')
            data = data.append(
                pd.DataFrame([['l1v_read', size, time]], columns=data_columns),
                ignore_index=True,
            )

    return data

def main():
    data = run()
    data.to_csv('r9nano.csv')
    grouped = data.groupby('arg')
    print(data.describe())

if __name__ == '__main__':
    main()


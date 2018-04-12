# NOTE: this file assumes that it runs in the pebblesdb directory inside another directory. e.g. pebblesdb/INFO/thisfile.py

import subprocess

commonalities = set()
strip_text = "rocksdb::"
with open("../../rocksdb::_locs.txt") as f:
    for line in f:
        strip_line = line[len(strip_text):].strip()
        proc = subprocess.Popen(['grep -rInH {} ..'.format(strip_line)], stdout=subprocess.PIPE, shell=True)
        (out, err) = proc.communicate()
        # print('{}'.format(out))
        if len(out) > 0:
            commonalities.add(line.strip())

print('{}'.format('\n'.join(sorted(commonalities))))

# Figure out all the files that contain the word 'rocks'
# requires that athe file rocks_locs.txt exists in this directory that contains the results of a grep for rocks
def unique_files():
    files_with_rocks = set()
    with open('rocks_locs.txt') as f:
        for line in f:
            files_with_rocks.add(line[:line.find(':')])

    print('{}'.format('\n'.join(sorted(files_with_rocks))))

# Figure out all the instances of rocksdb::*
# requires that athe file rocks_locs.txt exists in this directory that contains the results of a grep for rocksdb::
def rocksdb_namespace():
    names = set()
    search_text = 'rocksdb::'
    with open('rocksdb_locs.txt') as f:
        for line in f:
            found_loc = line.find(search_text, line.find(':')+2)
            while(found_loc != -1):
                end = len(line)
                for end_char in [' ', '(', ')', '<', '>', '&', '*', ';', '`', ',', ':', '.', '}', '[', ']']:
                    end_cand = line.find(end_char, found_loc+len(search_text))
                    end = end if end_cand == -1 else min(end, end_cand)
                cand = line[found_loc:end].strip()
                if len(cand) > 0:
                    names.add(cand)
                found_loc = line.find(search_text, found_loc+1)
    print('{}'.format('\n'.join(sorted(names))))

rocksdb_namespace()

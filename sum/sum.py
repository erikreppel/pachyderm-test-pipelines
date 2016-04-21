import os
import sys


if len(sys.argv) is not 3:
    print "Usage: python square.py <repopath> <outputpath>"
    sys.exit(1)

print 'starting'

input_path = sys.argv[1]
output_path = sys.argv[2]

print input_path, output_path

if not os.path.isdir(input_path):
    print "Invalid input path"
    sys.exit(1)
if not os.path.isdir(output_path) and output_path != '/pfs/out':
    print "Invalid output path"
    sys.exit(1)

print [f for f in os.walk(input_path)]

files = [[os.path.join(f[0], f[2][0]), f[2][0]]
            for f in os.walk(input_path) if f[2]]

print 'Found %s files' % len(files)

total = 0

file_write_path = os.path.join(output_path, 'sum')

for item in files:
    print "Opening", item[0]
    file_read_path = item[0]

    with open(file_read_path, 'r') as f:
        for line in f:
            total += int(line)

print 'The final total was', total

with open(file_write_path, 'w') as f:
    f.write(str(total))
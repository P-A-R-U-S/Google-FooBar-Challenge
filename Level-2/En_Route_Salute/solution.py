def answer(s):
    # your code here
    cr = 0
    result = 0
    for i, c in enumerate(s):
        if c == '>':
            cr += 1
        elif c == '<':
            result += cr

    return result * 2

print answer("--->-><-><-->-")
print answer(">----<")
print answer("<<>><")
def answer(start,   length):
    v1 = start
    v2 = start

    result = 0

    i = 0
    while i < length:
        v1 = v2 + i
        v2 = v1 + length - i - 1
        result ^= xor(v2) ^ xor(v1-1)
        i += 1

    return result


def xor(n):
    if n == 0:
        return 0

    if n & 3 == 0:
        return n
    elif n & 3 == 1:
        return 1
    elif n & 3 == 2:
        return n + 1
    elif n & 3 == 3:
        return 0


print answer(0, 1)
print answer(1, 1)
print answer(0, 3)
print answer(17, 4)
print answer(30, 10)
print answer(100, 10)
print answer(0, 200000)
print answer(0, 200001)
print answer(1, 200000)
print answer(1, 200001)

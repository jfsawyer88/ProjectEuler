#!/usr/bin/python
# list out permutations of a list


# all permutations
def perms(s, l = None):
    if l == None:
        out = []
        for i in xrange(len(s)):
            out = out + perms(s[:i] + s[i+1:], s[i])
        return out
    else:
        if 0 == len(s):
            return [l]
        else:
            return [l + 10 * n for n in perms(s)]


# permutations not divisible by 2, 3, or 5
def perms235(s, l = None):
    # top level
    if l == None:
        # don't bother if the sum of digits is 3
        if 0 == sum(s) % 3:
            return []
        else:
            out = []
            for i in xrange(len(s)):
                # don't bother putting an even number in the end
                if (1 == s[i] % 2) and (5 != s[i]):
                    out = out + perms235(s[:i] + s[i+1:], s[i])
            return out
    # no need to worry about divisibility inside
    elif l == len(s):
        out = []
        for i in xrange(len(s)):
            out = out + perms235(s[:i] + s[i+1:], s[i])
        return out
    # otherwise, recurse
    else:
        # base ccase
        if 0 == len(s):
            return [l]
        else:
            return [l + 10 * n for n in perms235(s, len(s))]


# permutations not divisible by 2, 3, or 5
# includes memoization
perms235_memo = {(2,): (2,), (3,): (3,), (5,): (5,)}
def perms235_m(s, l = None):

    # top level
    if l == None:

        # check if argument has been called before
        if s in perms235_memo:
            return perms235_memo[s]

        # don't bother if the sum of digits is 3
        if 0 == sum(s) % 3:
            return tuple()
        else:
            out = tuple()
            for i in xrange(len(s)):
                # don't bother putting an even number in the end
                if (1 == s[i] % 2) and (5 != s[i]):
                    out = out + perms235_m(s[:i] + s[i+1:], s[i])
            perms235_memo[s] = out
            return out
    # no need to worry about divisibility inside
    elif -1 == l:
        out = tuple()
        for i in xrange(len(s)):
            out = out + perms235_m(s[:i] + s[i+1:], s[i])
        # cannot add these to the memo
        #perms235_memo[s] = out
        return out
    # otherwise, recurse
    else:
        # base case
        if 0 == len(s):
            return (l,)
        else:
            return tuple([l + 10 * n for n in perms235_m(s, -1)])
# checks if the input is prime
# does not bother checking
# for divisibility by 2, 3, or 5
def is_prime235(n):
    if 1 == n: return False
    gaps = [4,2,4,2,4,6,2,6]
    i = 0
    p = 7
    while p < 1 + n ** 0.5:
        if 0 == n % p:
            return False
        p += gaps[i]
        i += 1
        i %= 8
    return True


# checks if the input is prime
# does not bother checking
# for divisibility by 2, 3, or 5
# memoization
is_prime235_memo = {1: False, 2: True, 3: True, 5: True}
def is_prime235_m(n):

    if n in is_prime235_memo:
        return is_prime235_memo[n]

    gaps = [4,2,4,2,4,6,2,6]
    i = 0
    p = 7
    while p < 1 + n ** 0.5:
        if 0 == n % p:
            is_prime235_memo[n] = False
            return False
        p += gaps[i]
        i += 1
        i %= 8
    is_prime235_memo[n] = True
    return True

# number of primes that can be made
# by concatentating all elemnts of s
def np(s):
    cnt = 0
    for perm in perms235(s):
        if is_prime235(perm):
            cnt += 1
    return cnt


# number of primes that can be made
# by concatentating all elemnts of s
# uses memozation
np_memo = {(2,): 1}
def np_m(s):
    if s in np_memo:
        return np_memo[s]
    cnt = 0
    for perm in perms235_m(s):
        if is_prime235_m(perm):
            cnt += 1
    np_memo[s] = cnt
    return cnt


def subsets(s):
    out = [tuple()]
    for x in s:
        out.extend([subset + (x,) for subset in out])
    return out


# full recursive memoized function
nps_memo = {tuple(): 1}
def nps(s):
    if s in nps_memo:
        return nps_memo[s]
    out = 0
    subs = subsets(s)
    for i in xrange(2 ** (len(s) - 1)):
        out += nps(subs[i]) * np_m(subs[-(1 + i)])
    nps_memo[s] = out
    return out

#################################################
#################################################
#################################################



## project euler 009

def factorize(N):
    # first test factor and list of all factors
    f=2; fs=[]
    while f*f < N:
        while 0 == N % f:
            fs.append(f)
            N = N//f
        f+=1
    if 1==N:
        return fs
    else:
        return fs+[N]


def divisors(N):
    pf  = factorize(N)
    pfg = {x:pf.count(x) from x in pf}
    

    

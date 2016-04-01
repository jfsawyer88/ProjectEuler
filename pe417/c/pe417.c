#include <stdio.h>
#include <sys/time.h>
#include <malloc.h>

long square(long n)
{
  return n * n;
}

long mpow(long a, long b, long m)
{
  if (0 == b) return 1;
  else
    {
      if (1 == b % 2) return (a * mpow(a, b - 1, m)) %m;
      else return square(mpow(a, b/2, m)) % m;
    }
}

int my_pow(int a, int b)
{
  int res = 1;
  while (b > 0)
    {
      res *= a;
      b--;
    }
  return res;
}

// multiplicative order of 10 mod p
long mult_order(long p)
{
  long d;
  long res = p - 1; //assume so until proven otherwise
  for (d = 2; d * d < p; d++)
    {
      // if d is a divisor of p - 1
      if (0 == (p - 1) % d)
	{
	  if ((1 == mpow(10, d, p) && (res >  d)))
	    {
	      res = d;
	    }
	  if ((1 == mpow(10, (p - 1)/d, p)) && (res > (p - 1)/d))
	    {
	      res = (p - 1)/d;
	    }
	}
    }
  return res;
}

int isqrt(int n)
{
  int xk0 = n;
  int xk1 = (xk0 + n / xk0) / 2;
  while ( ((xk1 - xk0) < -1) || (1 < (xk1-xk0)) )
    {
      xk0 = xk1;
      xk1 = (xk0 + n / xk0) / 2;
    }
  return xk1;
}

int fulldiv(n, p)
{
  while (0 == n % p) n /= p;
  return n;
}

int cntfacts(n, p)
{
  int res = 0;
  while (0 == n % p)
    {
      n /= p;
      res++;
    }
  return res;
}

int gcd(int a, int b)
{
  int t;
  while (b)
    {
      t = b;
      b = a % b;
      a = t;
    }
  return a;
}

int lcm(int a, int b)
{
  return (a * b) / gcd(a, b);
}


main()
{
  struct timeval tv1, tv2;
  gettimeofday(&tv1, NULL);

  int  i, j, m, l;
  long n, p;
  int N = 1000000; //input
  int sqrtN = isqrt(N+1);

  int *is_prime;
  is_prime = (int *)malloc(sizeof(int)*(sqrtN));
  is_prime[0] = 0; is_prime[1] = 0; is_prime[2] = 1;
  for (i = 3; i < sqrtN; i += 2) is_prime[i] = 1;
  for (i = 4; i < sqrtN; i += 2) is_prime[i] = 0;

  // the sieve
  for (i = 3; i*i < sqrtN; i += 2)
    {
      if (is_prime[i])
	{
	  for (j = i*i; j < sqrtN; j += i)
	    {
	      is_prime[j] = 0;
	    }
	}
    }

  // smallest prime factor
  int *spf;
  spf = (int *)malloc(sizeof(int)*(N+1));
  for (i = 0; i < N+1; i++) spf[i] = i;
  for (p = sqrtN - 1; p > 1; p--)
    {
      if (is_prime[p])
	{
	  for (n = p; n < N+1; n += p)
	    {
	      spf[n] = p;
	    }
	}
    }

  int *lam;
  lam = (int *)malloc(sizeof(int)*(N+1));
  lam[0] = 0; lam[1] = 0; lam[2] = 0;
  lam[3] = 1; lam[4] = 0; lam[5] = 0;

  for (n = 3*3; n < N+1; n *= 3)
    {
      if (1 != mpow(10, lam[n/3], n)) lam[n] = lam[n/3] * 3;
      else lam[n] = lam[n/3];
    }
  for (p = 7; p < N+1; p++)
    {
      if (p == spf[p])
	{
	  lam[p] = mult_order(p);
	  for (n = p*p; n < N+1; n *= p)
	    {
	      if (1 != mpow(10, lam[n/p], n)) lam[n] = lam[n/p]*p;
	      else lam[n] = lam[n/p];
	    }
	}
    }


  for (n = 6; n < N+1; n++)
    {
      //printf("n = %ld\n", n);
      // if divisible by 2 or 5 then answer
      // has already been computed
      if ( (0 == n % 2) || (0 == n % 5) )
	{
	  lam[n] = lam[fulldiv(fulldiv(n, 2), 5)];
	}
      else
	{
	  l = 1;
	  for (m = n; m > 1; m = fulldiv(m, spf[m]))
	    {
	      l = lcm(l, lam[my_pow(spf[m], cntfacts(n, spf[m]))]);
	    }
	  lam[n] = l;
	}
      //printf("lam[%ld] = %d\n", n, lam[n]);
    }

  long output = 0;
  for (i = 3; i < N+1; i++)
    {
      //printf("lam[%d] = %d\n", i, lam[i]);
      output += lam[i];
    }

  gettimeofday(&tv2, NULL);

  printf("Execution time: %f seconds\n",
         (double) (tv2.tv_usec - tv1.tv_usec)/1000000 + 
         (double) (tv2.tv_sec - tv1.tv_sec));
  printf("Output: %ld\n", output);

  return 0;
}

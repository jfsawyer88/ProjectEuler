#include <stdio.h>
#include <sys/time.h>


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


// largest s sucht that s^2 < n
long isqrt(long n)
{
  long xk0 = n;
  long xk1 = (xk0 + n / xk0) / 2;
  while ( ((xk1 - xk0) < -1) || (1 < (xk1-xk0)) )
    {
      xk0 = xk1;
      xk1 = (xk0 + n / xk0) / 2;
    }
  return xk1;
}


main()
{

  struct timeval tv1, tv2;
  gettimeofday(&tv1, NULL);
  //printf("check -1\n");
  long N = 1000000000000;
  //printf("check 0\n");
  long n;
  long sqrtN = 1 + isqrt(N);
  long i, p;
  //printf("check 1/2: sqrtN = %ld\n", sqrtN);

  // can be too long
  int prime_bool[sqrtN]; // prime bool array
  //printf("check 2/4\n");
  //initialize it!
  prime_bool[0] = 0;
  prime_bool[1] = 0;
  for (i = 2; i < sqrtN; i++)
    {
      prime_bool[i] = 1;
    }
  //printf("check 1\n");
  // sieve it!
  for (p = 2; p*p < sqrtN; p++)
    {

      // if p is prime
      if (1 == prime_bool[p])
	{
	  for (n = p * p; n < sqrtN; n += p)
	    {
	      prime_bool[n] = 0;
	    }
	}
    }
  //printf("check 2\n");
  // number of primes
  long nprimes = 0;
  for (p = 0; p < sqrtN; p++)
    {
      if (1 == prime_bool[p]) nprimes++;
    }
  //printf("nprimes = %ld\n", nprimes);
  // prime array
  long primes[nprimes];
  i = 0;
  for (p = 0; p < sqrtN; p++)
    {
      if (1 == prime_bool[p])
	{
	  primes[i] = p;
	  i++;
	}
    }

  int ifprime;
  for (n = N - 1; n >= 0; n--)
    {
      // prime until proven otherwise
      ifprime = 1;

      // if n is composite then continue
      for (i = 0; i < nprimes; i++)
	{
	  if (0 == n % primes[i])
	    {
	      ifprime = 0;
	      break;
	    }
	  //printf("%d not divisible by %d\n", n, primes[i]);
	}

      //if (ifprime) printf("%ld is prime, mult_order(%ld) = %ld\n\n", n, n, mult_order(n));
      //else printf("%ld is not prime\n", n);
      if (ifprime && (n == 1 + mult_order(n))) break;

    }


  gettimeofday(&tv2, NULL);  

  printf("Execution time: %f seconds\n",
         (double) (tv2.tv_usec - tv1.tv_usec)/1000000 + 
         (double) (tv2.tv_sec - tv1.tv_sec));
  printf("Output: %ld\n", n);

  return 0;
}

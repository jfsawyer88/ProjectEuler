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


main()
{

  struct timeval tv1, tv2;
  gettimeofday(&tv1, NULL);

  long N = 1000;

  long n;
  for (n = N - 1; n >= 0; n--)
    {
      if ( (0 != n % 2) && (0 != n % 5) )
	{
	  printf("cycle_length(%ld) = %ld\n", n, mult_order(n));
	  if (n == 1 + mult_order(n)) break;
	}
    }


  gettimeofday(&tv2, NULL);  

  printf("Execution time: %f seconds\n",
         (double) (tv2.tv_usec - tv1.tv_usec)/1000000 + 
         (double) (tv2.tv_sec - tv1.tv_sec));
  printf("Output: %ld\n", n);

  return 0;
}

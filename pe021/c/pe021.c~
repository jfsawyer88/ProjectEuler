#include <stdio.h>
#include <sys/time.h>

int sum_of_divisors(int n, int mem[])
{
  if (n < 10000)
    {
      return mem[n];
    }
  int c = 1;
  int d;
  for (d = 2; d * d <= n; d++)
    {
      if (0 == n % d)
	{
	  if (d * d == n)
	    {
	      c += d;
	    }
	  else
	    {
	      c += d + (n / d);
	    }
	}
    }

  return c;
}


main ()
{
  struct timeval tv1, tv2;
  gettimeofday(&tv1, NULL);

  int mem[10000];
  int i, j;
  for (i = 0; i < 10000; i++)
    {
      mem[i] = 0;
    }
  for (i = 1; i < 10000; i++)
    {
      for (j = 2 * i; j < 10000; j += i)
	{
	  mem[j] += i;
	}
    }

  int res = 0;
  int n, m;
  for (n = 1; n < 10000; n++)
    {
      m = sum_of_divisors(n, mem);
      if (n == sum_of_divisors(m, mem) && n != m)
	{
	  res += n;
	}
    }

  gettimeofday(&tv2, NULL);

  printf("Execution time: %f seconds\n",
         (double) (tv2.tv_usec - tv1.tv_usec)/1000000 + 
         (double) (tv2.tv_sec - tv1.tv_sec));
  printf("Output: %d\n", res);

}

#include <stdio.h>
#include <sys/time.h>
#include <malloc.h>


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


int min(int a, int b)
{
  if (a < b) return a;
  else return b;
}


int find_z0(int x, int *spf)
{
  int z = 1;
  while (x > 1)
    {
      if (0 == z % spf[x]) z /= spf[x];
      else z *= spf[x];
      x /= spf[x];
    }
  return z;
}


main()
{
  struct timeval tv1, tv2;
  gettimeofday(&tv1, NULL);

  int i, p, n;
  long x, y, z, z0, z1, sqrtz1;
  int limit = 100000000;

  // prime sieve
  int *spf; //smallest prime factor
  spf = (int *)malloc(sizeof(int)*(limit + 1));
  spf[0] = 0; spf[1] = 1;
  for (i = 2;   i < limit+1; i += 2) spf[i] = 2;
  for (i = 3;   i < limit+1; i += 2) spf[i] = i;
  for (p = 3; p*p < limit+1; p += 2)
    {
      if (p == spf[p]) // if p is prime
	{
	  for (n = p*p; n < limit+1; n += p)
	    {
	      spf[n] = min(spf[n], p);
	    }
	}
    }


  long output = 0;
  for (x = 3; x < limit+1; x++)
    {
      // if x-1 is prime
      if (spf[x-1] == x-1)
	{
	  z0 = find_z0(x, spf);
	  sqrtz1 = z0;
	  z = z0;
	  while (z < limit + 1)
	    {
	      if (z > x)
		{
		  // if z-1 is prime
		  if ( (z-1) == spf[z-1] )
		    {
		      y = isqrt(x * z);
		      // if y-1 is prime
		      if ( (y-1) == spf[y-1] )
			{
			  //printf("(%ld, %ld, %ld)\n", x-1, y-1, z-1);
			  output += (x-1) + (y-1) + (z-1);
			}
		    }
		}
	      sqrtz1 += z0;
	      z = sqrtz1*sqrtz1 / z0;
	    }
	}
    }


  /*
  n = 4567898;
  printf("%d = ", n);
  while (n > 1)
    {
      printf("%d*", spf[n]);
      n /= spf[n];
    }
  printf("\n");
  */


  gettimeofday(&tv2, NULL);  

  printf("Execution time: %f seconds\n",
         (double) (tv2.tv_usec - tv1.tv_usec)/1000000 + 
         (double) (tv2.tv_sec - tv1.tv_sec));
  printf("Output: %ld\n", output);

}

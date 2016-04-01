#include <stdio.h>
#include <sys/time.h>
#include <malloc.h>

int npk(int n, int p)
{
  int out = 1;
  while ( 0 == n % p )
    {
      n /= p;
      out *= p;
    }
  return out;
}


main()
{
  struct timeval tv1, tv2;
  gettimeofday(&tv1, NULL);

  int i, n, p, pk;
  long a;
  long output = 0;
  int N = 10000000;

  //int lpf[N + 1];
  int *lpf;
  lpf = (int *) malloc(sizeof(int)*(N+1));
  for (i = 0; i < N + 1; i++)
    {
      lpf[i] = 1;
    }

  for (p = 2; p < N + 1; p++)
    {
      // if p is prime
      if ( 1 == lpf[p] )
	{
	  // for all multiples of p
	  for (n = p; n < N + 1; n += p)
	    {
	      pk = npk(n, p);     // find highest power of p in n
	      if ( lpf[n] < pk )  // replace in array if larger than current
		{
		  lpf[n] = pk;
		}
	    }
	}
    }

  // main loop
  for (n = 2; n < N + 1; n++)
    {
      // if n is a prime power then it 
      // only contributes 1 to tht sum
      if ( n == lpf[n] )
	{
	  output++;
	}
      else
	{
	  for ( a = n - lpf[n]; a > n/2 - 1; a -= lpf[n] )
	    {
	      if ( (a + 1) == ((a + 1) * (a + 1)) % n )
		{
		  output += (a + 1);
		  break;
		}
	      else if ( a == (a * a) % n )
		{
		  output += a;
		  break;
		}
	    }
	}
    }



  gettimeofday(&tv2, NULL);  

  printf("Execution time: %f seconds\n",
         (double) (tv2.tv_usec - tv1.tv_usec)/1000000 + 
         (double) (tv2.tv_sec - tv1.tv_sec));
  printf("Output: %ld\n", output);

  return 0;
}

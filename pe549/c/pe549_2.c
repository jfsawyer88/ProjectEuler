#include <stdio.h>
#include <sys/time.h>
#include <malloc.h>


int max(int a, int b)
{
  if (a < b) return b;
  else return a;
}

int nf(int p, int n)
{
  int res = 0;
  while (0 == n % p)
    {
      n /= p;
      res++;
    }
  return res;
}

int nfp(int p, int n)
{
  int res = 0;
  while (n > 0)
    {
      n /= p;
      res += n;
    }
  return res;
}

int s(int p, int k)
{
 int s0 = (((k * (p - 1)) / p) * p) + p;
 if (k <= nfp(p, s0))
   {
     return s0;
   }
 else
   {
     //printf("p = %d, k = %d, s0 = %d, nfp(P, s0) = %d\n", p, k, s0, nfp(p, s0));
     return s0 + p;
   }
}


main()
{
  struct timeval tv1, tv2;
  gettimeofday(&tv1, NULL);

  int i, j, k, n, p;
  int N = 100000000; //input

  int *s_arr;
  s_arr = (int *)malloc(sizeof(int)*(N+1));
  for (i = 0; i < N+1; i++) s_arr[i] = 0;

  int *is_prime;
  is_prime = (int *)malloc(sizeof(int)*(N+1));
  is_prime[0] = 0; is_prime[1] = 0;
  for (i = 2; i < N+1; i += 1) is_prime[i] = 1;

  // the sieve
  for (p = 2; p < N+1; p++)
    {
      if (is_prime[p])
	{
	  for (n = p; n < N+1; n += p)
	    {
	      k = nf(p, n);
	      s_arr[n] = max(s_arr[n], s(p, k));
	      is_prime[n] = 0;
	    }
	}
    }


  long output = 0;
  for (n = 2; n < N+1; n++)
    {
      output += s_arr[n];
    }


  gettimeofday(&tv2, NULL);

  printf("Execution time: %f seconds\n",
         (double) (tv2.tv_usec - tv1.tv_usec)/1000000 + 
         (double) (tv2.tv_sec - tv1.tv_sec));
  printf("Output: %ld\n", output);

  return 0;
}

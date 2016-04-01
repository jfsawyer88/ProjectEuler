#include <stdio.h>
#include <sys/time.h>


long phi(n)
{
  long res = 1;
  long p = 2;
  while (1 < n)
    {
      if (0 == n % p)
	{
	  res *= p - 1;
	  n /= p;
	  while (0 == n % p)
	    {
	      res *= p;
	      n /=p;
	    }
	}
      p++;
    }
  return res;
}


long modpow(long a, long b, long m)
{
  if (b == 1) return a;
  if (0 == b % 2)
    {
      return modpow(a * a % m, b/2, m);
    }
  else
    {
      return (a * modpow(a * a % m, b/2, m) % m);
    }
}


long * doubles_cycle(long m)
{
  static long res[2];

  res[0] = 0;
  res[1] = phi(m);

  long t = 1;
  long h = modpow(2, res[1], m);
  while (t != h)
    {
      t = (t << 1) % m;
      h = (h << 1) % m;
      res[0]++;
    }

  return res;
}


long modtower(long b, long h, long m)
{
  if (m == 1) return 0;
  if (h == 1) return b % m;

  long *mu_lam = doubles_cycle(m);
  long mu  = mu_lam[0];
  long lam = mu_lam[1];

  //return modpow(2, mu + ((modtower(b, h - 1, lam) - mu) % lam), m);
  long res = modpow(2, mu + ((modtower(b, h - 1, lam) - mu) % lam), m);
  printf("res = %ld, b = %ld, h = %ld, m = %ld\nmu = %ld, lam = %ld\n\n", res, b, h, m, mu, lam);
}


main()
{

  struct timeval tv1, tv2;
  gettimeofday(&tv1, NULL);

  long m = 1475789056;
  //printf("phi(m) = %ld\n", phi(m));
  //long out = (63 + modtower(2, 7, m) + 2 * modtower(2, 1000, m)) % m;
  //long out = modtower(2, 7, m);
  long out = modtower(2, 6, 632481024);

  gettimeofday(&tv2, NULL);

  printf("Execution time: %f seconds\n",
         (double) (tv2.tv_usec - tv1.tv_usec)/1000000 +
         (double) (tv2.tv_sec - tv1.tv_sec));
  printf("Output: %ld\n", out);

  return 0;
}

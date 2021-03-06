#include <stdio.h>
#include <sys/time.h>


long * doubles_cycle(long m)
{
  static long res[2];
  res[0] = 0;
  res[1] = 1;

  long t = (1 << 1) % m;
  long h = (1 << 2) % m;
  while (t != h)
    {
      t = (t << 1) % m;
      h = (h << 2) % m;
    }

  t = 1;
  while (t != h)
    {
      t = (t << 1) % m;
      h = (h << 1) % m;
      res[0]++;
    }

  h = (t << 1) % m;
  while (t != h)
    {
      h = (h << 1) % m;
      res[1]++;
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


long modtower(long b, long h, long m)
{
  if (m == 1) return 0;
  if (h == 1) return b % m;

  long *mu_lam = doubles_cycle(m);
  long mu  = mu_lam[0];
  long lam = mu_lam[1];

  return modpow(2, mu + ((modtower(b, h - 1, lam) - mu) % lam), m);
}


main()
{

  struct timeval tv1, tv2;
  gettimeofday(&tv1, NULL);

  long m = 1475789056;
  long out = (63 + modtower(2, 7, m) + 2 * modtower(2, 1000, m)) % m;

  gettimeofday(&tv2, NULL);

  printf("Execution time: %f seconds\n",
         (double) (tv2.tv_usec - tv1.tv_usec)/1000000 +
         (double) (tv2.tv_sec - tv1.tv_sec));
  printf("Output: %ld\n", out);

  return 0;
}

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
      return a * modpow(a * a % m, b/2, m);
    }
}

long modtower(long b, long h, long m)
{
  printf("b = %ld, h = %ld, m = %ld\n", b, h, m);
  if (m == 1) return 0;
  if (h == 1) return b % m;

  long *mu_lam = doubles_cycle(m);
  long mu  = mu_lam[0];
  long lam = mu_lam[1];

  return modpow(2, lam + modtower(b, h - 1, lam), m);

}


main()
{
  struct timeval tv1, tv2;
  gettimeofday(&tv1, NULL);


  long *res;
  long m = 1475789056;
  while (1 < m)
    {
      res = doubles_cycle(m);
      //printf("m = %ld: mu = %ld, lam = %ld\n", m, res[0], res[1]);
      printf("mu = %ld, lam = %ld\n", res[0], res[1]);
      m = res[1];
    }

  long out;
  m = 1475789056;
  out = modtower(2, 6, m);

  gettimeofday(&tv2, NULL);

  printf("Execution time: %f seconds\n",
         (double) (tv2.tv_usec - tv1.tv_usec)/1000000 +
         (double) (tv2.tv_sec - tv1.tv_sec));
  printf("Output: %ld\n", out);

  return 0;
}

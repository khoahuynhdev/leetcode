"""
The problem is that / in Python always returns a float, and floats have limited precision (64-bit IEEE 754, which gives you about 53 bits of mantissa). When your number exceeds 2^53, the
   float loses precision and you get a wrong result.

  n = int("1111011110000011100000110001011011110010111001010111110001", 2)
  print(n)  # 279270273653256177

  # Float division — LOSES PRECISION for large numbers
  print(n / 2)    # 1.3963513682662809e+17  (float, rounded)
  print(int(n/2)) # 139635136826628080      (wrong!)

  # Integer division — EXACT
  print(n // 2)   # 139635136826628088      (correct)

  # Bit shift — EXACT, same as //2
  print(n >> 1)   # 139635136826628088      (correct)

  n / 2 converts to float, which can only represent integers exactly up to 2^53 (about 9 * 10^15). Your number is ~2.8 * 10^17, well beyond that limit, so the float rounds it and you get
  garbage.

  n // 2 and n >> 1 both stay in Python's arbitrary-precision integer space and give the exact answer. They're functionally equivalent for positive numbers. For negative numbers there's a
  subtle difference — // rounds toward negative infinity while >> is an arithmetic shift — but for interview problems you'll almost always be working with non-negative values where they're
   interchangeable.

  The rule of thumb: never use / when you want integer division in Python. Always use // or >>. This is actually a common interview bug — in Go, 10 / 3 gives you 3 (integer division), but
  in Python 10 / 3 gives you 3.3333....

"""

class Solution:
    def numSteps(self, s: str) -> int:
        cnt = 0
        num = int(s,2)
        while True:
            if num == 1:
                break
            if num % 2 == 0:
                cnt+=1          
                num >>=1
            else:
                cnt+=1
                num +=1
            
        return cnt

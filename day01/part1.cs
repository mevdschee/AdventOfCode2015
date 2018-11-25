using System;
using System.IO;

namespace Advent.Of
{
    class Code
    {
        static void Main(string[] args)
        {
            var s = File.ReadAllText("input");
            int n = 0;

            for (int i = 0; i < s.Length; i++)
            {
                if (s[i] == ')') n--;
                if (s[i] == '(') n++;
            }

            Console.WriteLine("{0}", n);
        }
    }
}
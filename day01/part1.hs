walk :: String -> Int
walk = sum . map move

move :: Char -> Int
move '(' = 1
move ')' = -1

main :: IO()
main = main' walk
main' f = readFile "input" >>= print . f

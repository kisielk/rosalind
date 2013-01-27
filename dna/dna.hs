import Data.Char

countChars (a, c, g, t) 'A' = (a + 1, c, g, t)
countChars (a, c, g, t) 'C' = (a, c + 1, g, t)
countChars (a, c, g, t) 'G' = (a ,c, g + 1, t)
countChars (a, c, g, t) 'T' = (a ,c, g, t + 1)
countChars (a, c, g, t)  _  = (a ,c, g, t)

printTuple (a, c, g, t) = show a ++ " " ++ show c ++ " " ++ show g ++ " " ++ show t ++ "\n"

countAllChars = foldl countChars (0, 0, 0, 0) . map toUpper

main = do
    interact (printTuple . countAllChars)

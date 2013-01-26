import Data.Char

main = do
    contents <- getContents
    putStr (map (\x -> if x == 'T' then 'U' else x) contents)

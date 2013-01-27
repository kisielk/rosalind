import Data.Char

complement 'C' = 'G'
complement 'G' = 'C'
complement 'A' = 'T'
complement 'T' = 'A'
complement  x  =  x

main = do
    interact (map complement . reverse)

import Data.Char

rna = map (\x -> if x == 'T' then 'U' else x)

main = do
    interact rna

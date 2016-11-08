fun zip_with bop nil _ = []
    | zip_with bop _ nil = []
    | zip_with bop (x::xs) (y::ys) = (bop(x, y)) :: (zip_with bop xs ys)
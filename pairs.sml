fun pair_with n [] = nil
        | pair_with n (hd::tl) = (n, hd)::(pair_with n tl)

fun all_pairs 0 = []
	| all_pairs 1 = []
	| all_pairs n = 
	let 
		fun gen_range s n = 
			if s = n then [s]
			else (gen_range s (n-1)) @ [n]
		fun gen_pairs x n =
			if (x+1) = n then [(x, n)] 
			else pair_with x (gen_range (x+1) n) @ (gen_pairs (x+1) n)
		val x = 1
	in 
		gen_pairs 1 n
	end
	

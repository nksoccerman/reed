fun all_bools n = 
    let 
    	fun head_all n [] = [] 
	    | head_all n (x::xs) = (n::x) :: (head_all n xs)
	fun do_appends 0 data = [[]]
	    | do_appends 1 data = data
	    | do_appends n data = do_appends (n-1) ((head_all true data) @ (head_all false data)) 
    in 
       do_appends n [[true], [false]]
    end

    

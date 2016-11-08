datatype logic = 
	AND of logic * logic
	| OR of logic * logic
	| NOT of logic
	| CONST of bool
	| VAR of string

val rng = Random.rand (0,0) 

fun rand_logic vars = 
	let
		fun getVal 0 vars = (hd vars)
			| getVal n vars = getVal (n-1) (tl vars)
		val x = Random.randRange(1,10) rng	
	in 
		if x = 1 then AND(rand_logic vars, rand_logic vars)
		else if x = 2 then OR(rand_logic vars, rand_logic vars)
		else if x = 3 then NOT(rand_logic vars)
		else let 
			val x = Random.randRange(1,(length vars)) rng
			in VAR(getVal (x-1) vars)
			end
	end

fun get_val key ((k, v)::tl) = 
		if k = key then v
		else (get_val key tl)

fun check (AND(a,b)) vars = if (check a vars) = true andalso (check b vars) = true then true else false
	| check (OR(a,b)) vars = if (check a vars) = true orelse (check b vars) = true then true else false
	| check (NOT(a)) vars = not (check a vars)
	| check (CONST(a)) vars = a
	| check (VAR(s)) vars = (get_val s vars)

fun vars_of vals = 
	let
		fun check a (hd::tl) = if a = hd then true else (check a tl)
			| check a [] = false
		fun clear_dups (hd::tl) = if (check hd tl) then (clear_dups tl)
			else hd::(clear_dups tl)
			| clear_dups [] = nil
		fun dovals (AND(a, b)) = (dovals a)@ (dovals b)
		| dovals (OR(a, b)) = (dovals a) @ (dovals b)
		| dovals (NOT(a)) = dovals a
		| dovals (CONST(b)) = []
		| dovals (VAR(a)) = a::[]
	in 
		clear_dups (dovals vals)
	end

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

fun truth_table logic = 
	let 
		val vars = vars_of logic
		val table = all_bools (length vars)
		fun map_pair (x::y) (n::z) =  (x,n)::(map_pair y z)
			| map_pair nil nil = nil
		fun do_evals logic vars (hd::tl) = 
				let val input = map_pair vars hd
				in (hd, (check logic input)) :: (do_evals logic vars tl)
				end
			| do_evals _ _ [] = nil
	in
		(vars, (do_evals logic vars table))
	end

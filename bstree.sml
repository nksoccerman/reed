datatype bstree = Lf | Br of int * bstree * bstree

fun contains x Lf = false
  | contains x (Br (k,l,r)) =
      if x = k then true
      else if x < k then contains x l
                    else contains x r

fun withalso x Lf = (Br (x,Lf,Lf))
  | withalso x (Br (k,l,r)) =
        if x = k then (Br (k,l,r))
        else if x < k then (Br (k,withalso x l,r))
                      else (Br (k,l,withalso x r))

datatype dir = Left | Right 

fun path x Lf = nil
	| path x (Br (k, l, r)) = 
		if (contains x (Br(k,l,r))) = false then (path x (withalso x (Br(k,l,r)))) 
		else if x > k then Right :: path x r
		else if x < k then Left :: path x l
		else nil

fun min (Br (k, Lf, _)) = k
	| min (Br (k, l, r)) = min l
	| min Lf = ~1

fun without x (Br(k,l,r)) = 
	let	
		fun brtolf x (Br(k,l,r)) = 
			if x < k then (Br (k, brtolf x l, r))
			else if x > k then (Br (k, l, brtolf x r))
			else Lf
		| brtolf x Lf = Lf
		fun dorem x min (Br(_,Lf,Lf)) = Lf
			| dorem x min (Br(_,Lf,r)) = r
			| dorem x min (Br(_,l,Lf)) = l
			| dorem x min (Br(_,l,r)) = (Br (min,l,brtolf min r))  
			| dorem x min Lf = Lf 
	in
		if x < k then (Br (k,without x l, r)) 
		else if x > k then (Br (k,l, without x r))
		else (dorem k (min r) (Br (k,l,r)))
	end
	| without x Lf = Lf

fun flatten (Br(k,Lf,Lf)) = [k]
	| flatten (Br(k, Lf, r)) = k::(flatten r)
	| flatten (Br(k, l, Lf)) = (flatten l)@[k]
	| flatten (Br(k, l, r)) = (flatten l)@[k]@(flatten r)

datatype quadsoln = 
         NoSoln | OneSoln of int | TwoSolns of int*int | AllSolns

fun sqrt x s = 
	if s * s = x then s 
	else if s * s < x then sqrt x (s+1)
	else ~1 

fun quadsolve 0 0 0 = AllSolns
	| quadsolve 0 b c = if c mod b = 0 then OneSoln (~(c div b)) else NoSoln
	| quadsolve a b c = 
		let 
			val d = sqrt ((b*b) - (4*a*c)) 0
			val x = (~b + d) div (2*a)
			val y = (~b - d) div (2*a)
			val xrem = (~b + d) mod (2*a)
			val yrem = (~b + d) mod (2*a)
		in 
			if d = ~1 then NoSoln
			else if d = 0 andalso xrem = 0 then OneSoln (x)
			else if xrem = 0 andalso yrem = 0 then TwoSolns (x, y)
			else if xrem = 0 andalso yrem <> 0 then OneSoln (x)
			else if xrem <> 0 andalso yrem = 0 then OneSoln (y)
			else NoSoln
		end

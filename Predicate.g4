grammar Predicate;

start: exp EOF;

exp:
    exp OR exp1
    | exp1;

exp1:
    exp1 AND exp2
    | exp2;

exp2:
    exp2 EQ exp1
	| exp2 NEQ exp1
	| exp2 LT exp1
	| exp2 GT exp1
	| exp2 LTE exp1
	| exp2 GTE exp1
    | exp2 IN LB (exp (CM exp)*)? RB
	| exp3;

exp3:
    exp3 MOD exp4
    | exp3 DIV exp4
    | exp3 MUL exp4
    | exp3 ADD exp4
    | exp3 SUB exp4
    | exp4;

exp4:
    ADD exp4
    | SUB exp4
    | NOT exp4
    | exp5;

exp5:
    LP exp RP
	| BOOL_LIT
	| INT_LIT
	| STRING_LIT
	| FLOAT_LIT
	| VAR
	;

// Literals
BOOL_LIT: TRUE | FALSE;
INT_LIT: Digit+;
STRING_LIT: SD (~[\n'"])*? SD ;
FLOAT_LIT: (Digit+ DT Digit* | Digit* DT Digit+ | Digit+) (E [+-]? Digit+)?;

// Keywords
IN: I N;
AND: A N D;
OR: O R;
TRUE: T R U E;
FALSE: F A L S E;

// Variable
VAR: (Letter | '_') (Letter | '_' | Digit)*;

// Ignore white spaces
WHITESPACE: [ \r\n\t]+ -> skip;

// Digit
fragment Digit: [0-9];

// Letter
fragment Letter: (A | B | C | D | E | F | G | H | I | J | K | L | M | N | O | P | Q | R | S | T | U | V | W | X | Y | Z);

// Case-insensitive letter
fragment A: ('a' | 'A');
fragment B: ('b' | 'B');
fragment C: ('c' | 'C');
fragment D: ('d' | 'D');
fragment E: ('e' | 'E');
fragment F: ('f' | 'F');
fragment G: ('g' | 'G');
fragment H: ('h' | 'H');
fragment I: ('i' | 'I');
fragment J: ('j' | 'J');
fragment K: ('k' | 'K');
fragment L: ('l' | 'L');
fragment M: ('m' | 'M');
fragment N: ('n' | 'N');
fragment O: ('o' | 'O');
fragment P: ('p' | 'P');
fragment Q: ('q' | 'Q');
fragment R: ('r' | 'R');
fragment S: ('s' | 'S');
fragment T: ('t' | 'T');
fragment U: ('u' | 'U');
fragment V: ('v' | 'V');
fragment W: ('w' | 'W');
fragment X: ('x' | 'X');
fragment Y: ('y' | 'Y');
fragment Z: ('z' | 'Z');

// Arithmetic Operators
MUL: '*';
DIV: '/';
ADD: '+';
SUB: '-';
MOD: '%';

// Conditional operators
NEQ: '!=';
EQ: '==';
GTE: '>=';
LTE: '<=';
LT: '<';
GT: '>';
NOT: '!';

// Delimitors
LB: '[';
RB: ']';
LP: '(';
RP: ')';
CM: ',';
SQ: '\'';
DQ: '"';
DT: '.';
SD: (SQ | DQ); // string's delimitor

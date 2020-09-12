# automaton
Deterministic Finite Automata (DFA) implementation in Go.

Should be able to compute any DFA.

One DFA is provided as proof.

## DFA included
### formal definition | definição formal
```
Σ = {a, b, c}
V = {S, X, Y, Z, T, W, V, K, L, M}
S = S
P:
  S -> aS | K
  K ->  b | bX
  X -> aY | L
  L ->  c | cZ
  Y -> aY | K | L
  Z -> bT | L
  T -> bT | M
  M ->  c | cW
```
### explicação
1. palavras terminam apenas em 'b' ou 'c', nunca em 'a'
2. b não pode ser concatenado a outro 'b' antes que surja um 'c'
3. a aparição de qualquer 'c' permite concatenação de 'b's posteriores
4. nenhum 'a' é permitido após qualquer 'c'
5. qualquer palavra sempre termina com o simbolo de maior "valor" na palavra`

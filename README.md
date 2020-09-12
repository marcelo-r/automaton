# automaton
Deterministic Finite Automata (DFA) implementation in Go.

Should be able to compute any DFA.

One DFA is provided as proof.

## DFA included
### formal definition | definição formal
```
considering: 
DFA = (Q, Σ, δ, q0, F)

then:
Q = {q1, q2, q3, q4, q5, q6}
Σ = {a, b, c}
q0 = q1
F = {S, X, Y, Z, T, W, V, K, L, M}
δ:
  S -> aS | K
  K ->  b | bX
  X -> aY | L
  L ->  c | cZ
  Y -> aY | K | L
  Z -> bT | L
  T -> bT | M
  M ->  c | cW
```
### explicação (pt-br)
1. palavras terminam apenas em 'b' ou 'c', nunca em 'a'
2. b não pode ser concatenado a outro 'b' antes que surja um 'c'
3. a aparição de qualquer 'c' permite concatenação de 'b's posteriores
4. nenhum 'a' é permitido após qualquer 'c'
5. qualquer palavra sempre termina com o simbolo de maior "valor" na palavra`

#### Done as a project for Introduction to Formal Languages class (circa Jun-2019).

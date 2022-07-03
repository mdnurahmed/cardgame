# cardgame
A simple card game problem.

# problem description
Card Game

You are given the following information:
1.	Below is the list of cards in ASCENDING orders (each card has alphanumeric and symbol):
2@	2#	2^	2*	3@	3#	3^	3*	4@	4#	4^	4*
5@	5#	5^	5*	6@	6#	6^	6*	7@	7#	7^	7*
8@	8#	8^	8*	9@	9#	9^	9*	10@	10#	10^	10*
J@	J#	J^	J*	Q@	Q#	Q^	Q*	K@	K#	K^	K*
A@	A#	A^	A*								

2.	Shuffle (randomize) the cards and display the result.

3.	Distribute the shuffled cards evenly in sequence to 4 players (i.e. 1st card to player 1, 2nd card to player 2, 3rd card to player 3, 4th card to player 4, 5th card to player 1, 6th card to player 2 and so on) and display the result.

4.	Evaluate the winner based on the following conditions:

-	Player with the highest number of cards with same alphanumeric part (i.e., K@, K#, K^, K*).
-	If more than 1 player has the same number of winning cards, the alphanumeric part with higher value won. If tie, the symbol part with high value won. Example:
Sample 1:	Player 1: K@, K#, K*	
	Player 2: A@, A#, A^	Winner
Sample 2:	Player 1: A@, A*	Winner
	Player 2: A#, A^	

5.	Show the winning result and player. 

# My strategy 

After random shuffling find out the frequency of the most frequent card of any type a player has. Also, find out the most valuable card a player has. This would be required to break tie. 

For example if a player has "A@, A#, A^" then he has 3 card of type A , and his most valuable card is A^.

If a player has "A@, A#, 2@, 2#" then he has 2 card of type A  and 2 card of type 2, so the frequency of the most frequent card of any type he has is 2 and his most valuable card is A# cause the frequency of type A card he has is equal to the frequency of most frequent card he has of any type and A# has the highest rank.

If a player has "A@ , 2@, 2#" then the frequency of the most frequent card he has is 2 and his most valuable card is 2# . 

In normal cases the player with most frequency wins , but in cases of tie the player with most valuable card wins .

# How to run 

Using Golang , 

```
git clone https://github.com/mdnurahmed/cardgame
cd cardgame
go install .
cardgame
```

to run unit tests, 
```
go test ./...
```

Using Docker , 
```
git clone https://github.com/mdnurahmed/cardgame
cd cardgame
docker build -t cardgame .
docker run -it cardgame /bin/sh
cardgame
go test ./... 
exit
```
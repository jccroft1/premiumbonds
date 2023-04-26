# Premium Bonds 

The UK Premium Bonds system is basically a lottery system that's a bit like an investment. 

You can see how they work here: 
https://www.nsandi.com/products/premium-bonds

This project aims to answer the question of whether it's actually worth using! 

If you're just interested in that, first follow the requirements, then head [here](./cmd/premiumbonds/README.md). 

## How? 

### Prize Allocation 

First, [premiumbonds.go](premiumbonds.go) can calculate the prize allocation for a given prize fund value. We can also tweak the odds and interest rate which are used to calculate the total fund size. These are currently set to the below: 

* March Fund Value: £330,527,200
* Odds: 24,000
* Interest: 3.3% 

### Simulation 

Next, we need some ground truth that we can check our calculations against. So [simulation.go](simulation/simulation.go) runs a mock of the actual lottery! This means generating a lot of random numbers so it can a while. This is because we want to simulate many years ahead so the values converge. 

### Prediction

However, rather than making compute intensive simulations we make use of probality distribution to estimate the overall results. 

Hypergeometric works well in this instance because it's a random draw WITHOUT replacement. 

https://en.wikipedia.org/wiki/Hypergeometric_distribution 

There's a number of central tendencies we can use but median is best because the data is ordinal (number of prizes you can win is discrete) and highly skewed (highly likely to win nothing). 

Applying this distribution to premium bonds is done by [hypergeometric.go](./hypergeometric/hypergeometric.go). 

### Results! 

To compare the results of the simulation and the hypergeometric distribution you can head to [compare](./cmd/compare/). 

As mentioned above to see your expected return for a given amount of money and time then see [this](./cmd/premiumbonds/README.md). 

## Requirements 

* Go! https://go.dev/

Yep, that's it! 

Just enter `go run .` from the command line in `./cmd/premiumbonds/` to get some results. 

## Credits 

To be honest, I was inspired to tackle this because the MSE Premium Bonds calculator was broken. 

https://www.moneysavingexpert.com/savings/premium-bonds/ 

## Disclaimer 

This is completely unaffiliated to the official premium bonds people NS&I. 

# Finance

- ## Bitcoin Stock-To-Flow Model and Why It Is A Bad Model

    The stock-to-flow model was popularized by a Dutch institutional investor who operates under the twitter account “PlanB”. In his paper “Modelling Bitcoin Value with Scarcity”, he asserts that certain precious metals  like gold and silver have maintained a monetary role throughout history because they have unforgeable scarcity and low rate of supply. For example, gold valuable because newly mined gold is insignificant to the current, aexisting stockpile because it is impossible to replicate the vast stores of gold around the globe. PlanB then goes on to explain that the same logic applies to bitcoin.

    In Bitcoin's case, it has unforgeable costliness because it costs a lot of electricity to produce new bitcoins. Production of bitcoin cannot be easily faked, as other miners have to verify the work of a miner before new bitcoin is mined. PlanB defined scarcity as low rate of supply and can be quantified with a metric called Stock-to-Flow (SF). SF is the ratio between current supply (stock) and new supply (flow).

    PlanB’s hypothesis is that “scarcity, as measured by SF, directly drives value.” He goes on to plot bitcoin’s SF against market value in US dollars as shown below.


    PlanB goes on to run a linear regression using natural logarithm of bitcoin’s SF metric as the independent variable and the natural logarithm of the market value as the dependent variable. This ends with a conclusion that there is a significant statistical relationship between stock-to-flow and market value. His model gave a prediction of bitcoin hitting a market value of $1 trillion after the previous bitcoin halving in May 2020 translating to a price of $55,000.

    Obviously, this prediction was way off, at the time of writing, bitcoin trades around $11,105. So why was the prediction so inaccurate? From a theoretical point of view, the model is based on a strong assertion that the USD market value of a commodity is derived directly from their rate of new supply. According to Nico Cordeiro, the CIO and fund manager at Strix Leviathan,  “The model is not supported by evidence other than the singular data points selected to chart gold and silver’s market value against bitcoin’s trajectory. Also, the naive application of a linear regression. Good statistical results such as high R-square, do not constitute a meaningful finding.“

    Theoretically, PlanB’s definition of “scarcity” is not scarcity by definition. He uses scarcity to describe an asset’s supply growth rate as measured by the SF metric with the assumption that increasing new supply depresses the price through increased selling pressure from producers. However, a high SF represents a dynamic where new supply is insignificant to the current supply. It can be argued that SF has no direct relationship with gold’s value over the last 100 years. Gold held market values between $60 billion to $9 trillion, all at the same SF of 60. 

    While a high SF value may be a necessary feature for a commodity to serve as a store of value, the metric says nothing about how market participants value said commodity.

# Computer Science
- ## Recursion and Fibonacci
    It is a bad idea to use recursion to find the fibonacci of a number because the implementation does a lot of repeated work, the time complexity: O(n) = O(n-1) + O(n-2) is exponetial.

# Maths
- ## Minimum value of a positive real number, y such that: y = sqrt((x+6)^2 + 25) + sqrt((x-6)^2 + 121)

    y is defined for all real values of x;

    y is a positive value for all real values of x.
    
    at x = 6:
    
    y = sqrt((0+6)^2 + 25) + sqrt((0-6)^2 + 121)

    y = sqrt(61) + sqrt(157)

    y = 20.34 
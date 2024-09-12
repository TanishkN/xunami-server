from backtesting import Backtest, Strategy
from backtesting.lib import crossover
from backtesting.test import SMA, GOOG

class MyStrategy(Strategy):
    def init(self):
        price = self.data.Close
        self.ma1 = self.I(SMA, price,10)#10 day ma
        self.ma2 = self.I(SMA, price,20)#10 day ma
    def next(self):
        if crossover(self.ma1,self.ma2):
            self.buy()
        elif crossover(self.ma2,self.ma1):
            self.sell
backtest = Backtest(GOOG, MyStrategy, commission=0.002,exclusive_orders=True)
stats=backtest.run()

print(stats)
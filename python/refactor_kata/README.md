# Agenda

You are tasked to improve existing accounting system of the prehistoric market.
You are given code of component that compute market's fee based on trade details.

Fee is usually computed as some percentage of difference between 
- price that was agreed between merchant
- `actual price` based on world-wide currencies exchange rate

In our setup fee is heavily affected by merchants characteristics so there is some complex business rules in 
place to take all possible options into considerations. 

There are some un-resolved challenges for computing `actual price`:
- Market is international so people from different places trade using different currencies.
- Some-times they even rely on different measurement systems.

Those part of functionality not yet fully supported, but YOU can change it and make disabled tests passed!

A previous programmer suddenly became victim during beta-testing by being un-able to persuade 
one of merchant that it is some tiny technicals issues that easy to fix 
(or it considered term `POC stage` as some kind of swearing?).

# Action plan

1. Re-factor code to simplify it and make more maintainable
2. Fix logical issues
3. Extend functionality to support non-direct currency conversions
4. Extend functionality to support computations of price in case of using different measurement systems
5. Reach 100% tests coverage 

# Learn your IDE!

We will use PyCharm as example.

### How to setup virtual environment
1. Using provided Makefile
```bash
make setup
```
2. Using PyCharm:

   [Detailed instructions with screenshots](https://www.jetbrains.com/help/pycharm/creating-virtual-environment.html).

### How to setup test runner in PyCharm
First you would need to [set up test runner](https://www.jetbrains.com/help/pycharm/testing-your-first-python-application.html#choose-test-runner) for project:
  - Open Project preferences (On Mac - ⌘,)
  - Navigate into `Tools -> Python integrated tools -> Testing -> Test runner`
  - Choose `pytest`

After that you can click on `tests` folder or individual test and either choose to
  - Run default runner using shortcut (On Mac - ⇧^R)
  - Manually choose test configuration under `More Run/Debug` menu - for example, in order to run tests with coverage

# Some useful shortcuts:
- Setup project interpreter: ⌘,
- Extract selected code to the method: ⌥⌘M
- Rename variables under cursor:  ⇧F6
- Change method's under cursor: ⌘F6
- Move selected class\method to another module: F6

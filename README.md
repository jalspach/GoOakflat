[![Gitpod Ready-to-Code](https://img.shields.io/badge/Gitpod-Ready--to--Code-blue?logo=gitpod)](https://gitpod.io/#https://github.com/jalspach/reactor)

# GoOakflat
## OakFlat simulation
VERY Simplified PWR Simulation in Go
https://www.ucsusa.org/sites/default/files/2019-10/pwr-intro.pdf

### Why?
This is all so that I can teach myself Go. It is based on the http://gamtech.com/oakflat.aspx that I used to enjoy playing in my youth (NERD lol).

### Dev setup
* Atom
  * go-debug
  * go-plus
  * go-signature-statusbar
  * platformio-ide-terminal

ToDo
* Everything lol
    * Basic mechanics
      * Move temp from device to device
      * Move radiation from device to device
      * Pressure and its relation to temp and flow
      * Speed and game time
      * Heat Exchanger formulas
        * Towers (with and without fan)
        * Heat Exchanger
      * Option - Do I build an "engine" function that cycles through each device in turn and updates the stats?
    * Interface
      * Show stats
      * Input Controls
      * Console or web?
      * Images would be great
      * Separate into client  server
        * This lets me focus on only doing one thing at a time
    * "Game"
      * Track
        * Spending
        * Profit
        * Max MWH (average and point in time)
        * Number of days w.o accident
        * High Score
        * Etc
      * Random Breakdowns
        * create failure scenarios

# Tasks
- [] Gather list of "parts"
    * need to refactor this into generic parts which get named when instantiated.
    * think about how to break these down (parts that move vs those that don't maybe)
      * parts with no volume need to have volume set to 0
    * rotating
    * static
- [] Clean up parts file and make it a package or module...learn to use package / module lol
- [] Build sample system
    * Tank
    * Feed
    * Pipe
    * Return1
    * Pump
    * Return2
- [] Test "flow" of water and Heat
    * Set initial temps
    * Request pump speed from user
    * Report temps out to user
    * Return to speed request

# Formulas / Conversions
This table will list some of the constants I plan to use.

Term | Simplified Value | Actual Value | Source
----------|----------|----------|----------
Specific dencity of h2o | 1000 kg/m³| 997 kg/m³ at specific temp and pressure | https://www.engineeringtoolbox.com/water-temperature-specific-gravity-d_1179.html
Heat transfer formula | use full formula | T(final) = (m1_T1 + m2_T2) / (m1 + m2) | https://sciencing.com/calculate-concentration-solution-different-concentrations-8680786.html
Volume of a tank | | |
Pressure to Temprature | 100psi / 1 deg F | | https://physics.stackexchange.com/questions/363328/water-pressure-vs-temperature,  https://opentextbc.ca/chemistry/chapter/9-2-relating-pressure-volume-amount-and-temperature-the-ideal-gas-law/
Steem pressure to RPM | this will need to be simplified as it requires nozzle angle calculations, etc.. | |
Steem flow rate | can be simplified as pressure to RPM is going to be also | | https://www.tlv.com/global/US/calculator/steam-flow-rate-through-piping.html
Rods to energy | will need to be simulated at first anyway | |
Water volume through a pipe at a specific pressure | | | https://sciencing.com/convert-psi-gpm-water-8174602.html

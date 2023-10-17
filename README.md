# lol-champselect-lookup

Welcome to lol lookup! This allows you to lookup a champion on the enemy team 
to see its counters and your mastery points on those counters.

# Startup Guide
1. Choose a location to download this project. Make it something that you can easily find the filepath for.
2. Clone the project with git clone into that file location.
3. Get the Ubuntu on Windows shell from the Windows app store if you don't already have it.
4. Open an Ubuntu terminal and type `source <filepath>` where filepath points to lol.sh. This might take some trial and error.
   Example: source ../../mnt/c/Users/myself/Desktop/lol.sh
5. To confirm that its working, type `lol`. You should see a "welcome to lol lookup!" message. 
6. Once you have the right filepath, copy the source command and run `nano ~/.bashrc`
7. At the bottom of this file, paste in your source command. This will make it so you never have to run the source command manually again.
You can now use lol lookup in the terminal anytime you like!

# Commands guide
## Set your name
`lol ign <summoner name>`

Use this command upon first downloading the project to set your summoner name for finding champion mastery. If your name has spaces, be sure to surround your name in quotes.

Examples:

`lol ign myName`

`lol ign "My name has spaces"`

## Look up a champion's counters

`lol <champion name>`

Use this command to return a list of the champion's counters along with information about how much mastery you have on each of those champions. You should not include apostrophes or spaces. Nunu & Willump is simply 'nunu'.

Examples:

`lol yasuo`

`lol reksai`

`lol leesin` 

Expected Output:
```
Visiting:  https://u.gg/lol/champions/lux/counter
Visiting:  https://championmastery.gg/summoner?summoner=summoner&region=NA&lang=en_US

Lux's best counters in Support are
Bard            55.09% WR in 1,218 games         mastery 7 with 86876 pts
Xerath          54.97% WR in 3,569 games         mastery 5 with 22171 pts
Janna           54.59% WR in 1,145 games         mastery 6 with 49652 pts
Vel'Koz         53.93% WR in 725 games           mastery 5 with 23845 pts
Pyke            53.86% WR in 2,549 games         mastery 4 with 19285 pts
Blitzcrank      53.12% WR in 3,200 games         mastery 5 with 40427 pts
Braum           52.9% WR in 690 games            mastery 4 with 14693 pts
Milio           52.82% WR in 2,448 games         mastery 5 with 21777 pts
Ashe            52.74% WR in 383 games           mastery 6 with 57703 pts
Taric           52.73% WR in 421 games           mastery 4 with 18397 pts
```

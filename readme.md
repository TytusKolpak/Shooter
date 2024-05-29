# Shooter game

This project is a 2d top-down game called Shooter since it mainly revolves around shooting enemies.
This game expects You to use a gamepad to control it.

## Playing the game

The simplest way to play the game is by downloading the whole project and then running the executable of this game (shooter.exe) which is placed in its root directory.

It can also be done if You go to the directory of this project in Your terminal and run the command `./shooter` if You are using Windows. Mac and Linux have slightly different pattern like just `shooter` or `shooter.exe`.

Keep in mind that the only elements really required to run this game is the shooter.exe file and sprites from `/sprites` directory for the images on the screen. All other files serve a purpose when developing the game.

### Rules

The goal of the game is to kill all the enemies that will be spawned. The enemies will come in 3 types. In the first of 3 stages it is the easiest enemy, then average and in the end the hardest. After the third stage passes no more enemies are spawned and You can shoot down the remaining enemies to win the game.

Should an enemy reach you before the game is won, then You loose it. In both cases Your statistics are displayed and appropriate message of type of the game end is displayed and You can play it again.

You begin the game with a fixed amount of projectiles to shoot. Any time You shoot it, it is removed from Your available projectiles. Once this number reaches 0 You can shoot no more. There will be more enemies than You have projectiles. To regain the ammunition You can pick it up from the ground where the enemy was shot down. If the projectile misses and leaves the screen then it is lost and You are left with that new decreased ammunition capacity.

### Controls

The movement of the player on the screen is controlled by the left stick of Your gamepad. Aiming is controlled by the right stick. Shooting is controlled by the right trigger.

Additionally You can pause the game using the left center button (like the small select or start button). If the game is over You can run it again immediately with the left center button again. If You just want to quit the game after it's over then You can do it with the right center button.

### Extra game information

While playing the game some game information will be displayed in the right upper corner of the window. These include time spent in the game, the amount of enemies You have shot down and the amount of bolts Your character in the game have left.

## Modifying the game to Your preferences

Should You be interested in modifying it's behavior, it can be done by changing values in file `entities/parameters.go`.
If You run the game by running it's executable file, to see changes introduced in the parameters file, You need to rebuild the executable by running command `go build` in the terminal while in the directory where the executable is placed. To do this You will need to have Go as a programming language installed on Your PC. If You don't have it installed then it can be done from https://go.dev/doc/install.

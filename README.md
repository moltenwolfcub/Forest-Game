# **Forest-Game**
### A game written in Go using Ebitengine

Currently very WIP. It's planned to be a strategy RPG style game set in a cozy forest.

# Installing
## Windows
Go to the Releases section on Github and select the version of the game you want to get (The latest is most adviseable). Download the `.exe` for the version of windows you are running.

## Mac & Linux
Go isn't capable of cross-compiling files with calls to C libraries (Which Ebitengine does) so I haven't managed to automate the build process of Mac or Linux. I might occasionally publish manual releases for these versions (as I run Linux myself) however if you can't find ther version for your OS then follow the steps below for compiling your own build. 

(If anyone knows how to get this working please let me know)

## Building from Source
1. Install a [Go compiler](https://go.dev/) on your computer
2. Obtain a local copy of the game's source. This can be done by cloning the repository `git clone git@github.com:moltenwolfcub/Forest-Game.git` or downloading the source off of a Github release. 
3. Run the command `go build` and a built executable file should appear in the directory named something along the lines of `Forest-Game`. (Make sure you have navigated to the downloaded source for this to work)
4. This is your final file that can be ran on any machine of that OS and architecture regardless of whether Go is installed

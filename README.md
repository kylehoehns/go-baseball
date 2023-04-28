# Go Baseball

I signed up to help coach my daughter's tball team and realized I needed an easy way to automatically create a batting order and assign fielding positions for each game. I also wanted each player to never play the same position in the same game and wanted each player to play in separate "groups" on defense. The field can be grouped into two by placing P, 1B, 2B, and RF into one group, while SS, 3B, LF, CF are another. Between batters, players rotate positions within that group.

This tool is a simple generator to help coaches assign players to a batting order and fielding positions during a game of baseball.

## Running

Use VS Code and open this project in a [Dev Container](https://code.visualstudio.com/docs/devcontainers/containers) without needing to install any tooling locally.


Create a file called `roster` in the root directory. Place each player's name on a separate line.

```text
Lily
Jack
Maya
Luke
Ava
Finn
Zoe
Owen
```

The `main` method takes in an argument for the desired number of innings to print. By default, that value is 9.

```shell
go run cmd/main.go 5
```

This will generate you an output that looks something like this.

```text
## Batting Order:
| Inning | 1          | 2          | 3          | 4          | 5         
|--------|------------|------------|------------|------------|------------|
|        | Maya       | Finn       | Owen       | Luke       | Lily      
|        | Jack       | Ava        | Zoe        | Maya       | Finn      
|        | Owen       | Luke       | Lily       | Jack       | Ava       
|        | Zoe        | Maya       | Finn       | Owen       | Luke      
|        | Lily       | Jack       | Ava        | Zoe        | Maya      
|        | Finn       | Owen       | Luke       | Lily       | Jack      
|        | Ava        | Zoe        | Maya       | Finn       | Owen      
|        | Luke       | Lily       | Jack       | Ava        | Zoe       

## Fielding Positions:
| Inning | 1          | 2          | 3          | 4          | 5         
|--------|------------|------------|------------|------------|------------|
| P      | Maya       | Zoe        | Ava        | Jack       | Lily      
| 1B     | Jack       | Lily       | Luke       | Owen       | Finn      
| 2B     | Owen       | Finn       | Maya       | Zoe        | Ava       
| 3B     | Zoe        | Ava        | Jack       | Lily       | Luke      
| SS     | Lily       | Luke       | Owen       | Finn       | Maya      
| LF     | Finn       | Maya       | Zoe        | Ava        | Jack      
| CF     | Ava        | Jack       | Lily       | Luke       | Owen      
| RF     | Luke       | Owen       | Finn       | Maya       | Zoe       
```

## Assumptions

* No player will be assigned the catcher position
* You will have exactly 7 or 8 players on your team
* This does not handle having players on the bench, or having a Rover, Left-Center or Right-Center additional fielding positions added.

## Implementation

This was a project where I utilized the GPT-4 LLM to generate a majority of the code based on prompts around what code I was wanting it to create.
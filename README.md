[![wercker status](https://app.wercker.com/status/459d395f95d69787d29861a22214dc4c/m "wercker status")](https://app.wercker.com/project/bykey/459d395f95d69787d29861a22214dc4c)

# Daemon

## Objective

Thought expirement to create a simple process which would scan live news feeds to hit on pre-conceived
criterion which would then trigger a series of events. The idea came from the recent re-read of
["Daemon" by Daniel Saurez](http://thedaemon.com)...

```
Matthew Sobol was a legendary computer game designer—the architect behind half a dozen popular online games.
His premature death from brain cancer depressed both gamers and his company’s stock price. But Sobol’s fans
weren’t the only ones to note his passing. He left behind something that was scanning Internet obituaries,
too — something that put in motion a whole series of programs upon his death. Programs that moved money.
Programs that recruited people. Programs that killed.

Confronted with a killer from beyond the grave, Detective Peter Sebeck comes face-to-face with the full
implications of our increasingly complex and interconnected world—one where the dead can read headlines,
steal identities, and carry out far-reaching plans without fear of retribution. Sebeck must find a way to
stop Sobol’s web of programs—his Daemon—before it achieves its ultimate purpose. And to do so, he must
uncover what that purpose is...
```

## The nasty bits

### daemon

The **daemon** is a simple process started by ```daemon run```. Essentially it is the scheduler for all
other processes internal to the application set to go off in regular intervals. The primary focus was
to create a reproducible environment manageming any environment variables and other management tasks.

### scour

***scour*** will take a newline delimited file of RSS 2.0 or Atom feeds and look for matching stories. The
matching logic is very basic but may be later expanded to give more flexibility. If it detects a match
then the story is piped to the ***scourge*** command.

### scourge

The last piece of the puzzle is ***scourge***. Its intent is to take a story which was a successfuly hit
and to execute the appropriate response. The response is in the form of chained commands which are executed
on the system.

## Stories

Story files were created to drive the match/hit processes as well as the associated responses. The format
for a story can be seen with the example of []():

```
{
  "included": ["gopher", "gala"],
  "excluded": ["game"],
  "triggers": [{
    "command": "touch",
    "arguments": ["gophers-unite"],
    "wait": false
  }]
}
```

To qualify as a successful match all of the terms in the **included** array must be contained in the title
or description of the article. Also, the terms in the **excluded** array must not be found within either of
those sections. Once it has been determined a match then the array of **triggers** will be ran.

## Compiling a list of feeds

The current list of feeds is actually a good starting place. These are compiled from the following sources:

  - Reuter
  - NPR
  - CNN
  - Fox News
  - NY Times
  - BBC
  - United Nations

However, you can change the news feeds at runtime by setting the environment variable for ```DAEMON_FEEDS```
to any file. The only requirement is to place each URI on its own line. For example, for some testing we
used a much smaller list and only included the following feeds:

```
http://golangweekly.com/rss/17nm799j
```

## Use cases

Besides being an experiment of what can be done in a weekend this could actually have some interesting uses.

  1. Time sensitive reactions to current events
  2. Dynamic story lines in gaming
  3. Interesting workflows in the work environment
  4. Taking over the world

## Example

[![daemon in action](https://img.youtube.com/vi/kMmzyPthHdU/0.jpg)](https://www.youtube.com/watch?v=kMmzyPthHdU)

With this arbitrary example we pull down the latest and greatest and then modify the story to and watch as
it hits on a news item and sends a notification to [dunst](http://knopwob.org/dunst/index.html) displaying
in the top-right corner.

# pigsty - Your analytics workflow helper

This project is using [Readme Driven Development](http://tom.preston-werner.com/2010/08/23/readme-driven-development.html).

None of the features described here is fully/partialy working, once it is, the README will reflect it. If you have contributions or feedbacks please feel welcome to do so.

## Why?

Working at [Gogobot](http://www.gogobot.com), we have multiple tasks running on map-reduce every day.

Those tasks are typically ran on EMR, for 45-120 minutes at a time.

Once a cluster is running, we want to queue up some tasks (steps) to it.

Since we don't want to spend money on idling clusters, we want to kill the cluster at the end.

This of course is covered by logs, monitoring and more.

## Features

1. Launch a cluster
2. Termination protection (optional)
3. Cutom bootstrapping steps configuration
4. Custom Steps configuration
5. Kill timers
6. Logs -> JSON log support for Logstash

## Contibutors

[@kensodev](http://twitter.com/kensodev) [Github](http://github.com/KensoDev)  

## License

See `LICENSE` file, but you already know the drill with my repos probably :)

## Contribute / Bug reports

Please feel free to open issues, add PRs, documentation, question or anything else.

If you want to contribute please make sure you add tests.







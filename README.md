# pigsty - Your analytics workflow helper

This project is using [Readme Driven Development](http://tom.preston-werner.com/2010/08/23/readme-driven-development.html).

None of the features described here is fully/partialy working, once it is, the README will reflect it. If you have contributions or feedbacks please feel welcome to do so.

## Why?

Working at [Gogobot](http://www.gogobot.com), we have multiple tasks running on map-reduce every day.

Those tasks are typically ran on EMR, for 45-120 minutes at a time.

Once a cluster is running, we want to queue up some tasks (steps) to it.

Since we don't want to spend money on idling clusters, we want to kill the cluster at the end.

This of course is covered by logs, monitoring and more.

## Features (IMPLEMENTED next to features that are already working)

1. Launch a cluster (IMPLEMENTED)
2. Termination protection (optional)
3. Cutom bootstrapping steps configuration (IMPLEMENTED)
4. Custom Steps configuration (IMPLEMENTED)
5. Kill timers
6. Logs -> JSON log support for Logstash

## How it works?

### Prerequisites

You need to define 2 environment variables for this command to work

`AWS_SECRET_ACCESS_KEY` and `AWS_ACCESS_KEY_ID`, this should be for a user that **has** acces to creating/terminating EMR clusters.

### Configuration

The command has multiple configuration options

#### Command line Arguments

1. `log-uri`: Where do you want to keep the logs of your steps/cluster run. Default: none
2. `cluster-name`: Name of the cluster that will be created. Default: `Pigsty`
3. `key-name`: AWS KeyPair that will be assigned to the cluster. Default none
4. `instance-count`: How many instances should your cluster include: Default: `3`
5. `master-instance-type`: Instance Type for the master node. Default `m3.xlarge`
6. `slave-instance-type`: Instance Type for the slave node. Default `m3.xlarge`
7. `bootstrap-actions-file`: Location of the bootstrap steps configuration file (See bootstrap steps configuration structure). Default none
8. `steps-configuration-file`: Location of the steps configuration file (See Steps json structure). Default non
9. `kill-when-idle`: true/false whether to kill the instance(s) when the steps are finished. Default true

#### JSON steps cnofiguration (for bootstrap actions)

If you want to have bootstrap actions attached to your cluster, the command needs to include a `bootstrap-actions-file`.  
This is just a simple JSON file that needs to look like this:

```javascript

  [
    {
		Args: [],
        "Name": "Install Hue",
        "Path": "s3://us-west-2.elasticmapreduce/libs/hue/install-hue"
    },
    {
        "Args": [],
        "Name": "Download mysql-connector jar",
        "Path": "s3://some-bucket-name/bootstrap.sh"
    }
]

  
```

#### Steps

Steps is a similar configuration file like bootstrap steps.

This is an example for installing some software. Pig, Hive, also running Hue.

```javascript

[
  {
    "Name": "Install Hive",
    "ActionOnFailure": "TERMINATE_CLUSTER",
    "MainClass": "None",
    "Jar": "s3://us-west-2.elasticmapreduce/libs/script-runner/script-runner.jar",
    "Args": [
      "s3://us-west-2.elasticmapreduce/libs/hive/hive-script",
      "--install-hive",
      "--base-path",
      "s3://us-west-2.elasticmapreduce/libs/hive",
      "--hive-versions",
      "latest"
    ]
  },
  {
    "Name": "Install Pig",
    "ActionOnFailure": "TERMINATE_CLUSTER",
    "MainClass": "None",
    "Jar": "s3://us-west-2.elasticmapreduce/libs/script-runner/script-runner.jar",
    "Args": [
      "s3://us-west-2.elasticmapreduce/libs/pig/pig-script",
      "--install-pig",
      "--base-path",
      "s3://us-west-2.elasticmapreduce/libs/pig",
      "--pig-versions",
      "latest"
    ]
  },
  {
    "Name": "Run Hue",
    "ActionOnFailure": "TERMINATE_CLUSTER",
    "MainClass": "None",
    "Jar": "s3://us-west-2.elasticmapreduce/libs/script-runner/script-runner.jar",
    "Args": [
      "s3://us-west-2.elasticmapreduce/libs/hue/run-hue"
    ]
  }
]
  
```



## Contibutors

[@kensodev](http://twitter.com/kensodev) [Github](http://github.com/KensoDev)  

## License

See `LICENSE` file, but you already know the drill with my repos probably :)

## Contribute / Bug reports

Please feel free to open issues, add PRs, documentation, question or anything else.

If you want to contribute please make sure you add tests.
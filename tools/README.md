# Tools
This sub-repo is for tools related to election, staking and voting.

# Dumper
Dumper dumps the information from staking contract (0x87c9dbff0016af23f5b1ab9b8e072124ab729193) and registration contract (0x95724986563028deb58f15c5fac19fa09304f32d) to a local csv file.

# Processor
Processor processes the csv produced by Dumper and generate another csv that breaks down the votes to each voters for each delegate. See the samples for more information.

# Generate breakdown of votes
```
go run tools/dumper/dumper.go > tools/processor/stats.csv
go run tools/processor/processor.go
```
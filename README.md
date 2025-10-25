# iptool

`iptool` is a lightweight command-line utility built in Go for developers, DevOps engineers, and network admins who work with IP ranges, subnets, and CIDRs on a daily basis. It provides quick, practical operations like viewing subnet info, converting CIDRs to ranges, splitting or merging networks, and validating IPs, etc.

## Features

### 1. Generate IP range within a CIDR

To generate the IP range within a CIDR use the `range` subcommand as shown below -

```
iptool range 192.168.1.0/24
```

This will return the starting ip address, ending ip address, and the total number of hosts within the CIDR.

### 2. Genrate CIDR for a given IP Range

To generate the CIDR for a given IP range (First IP Address and Last IP Address) use the `cidr` subcommand as shown below -

```
iptool cidr 192.168.1.0 192.168.1.255
```
This will return the quivalent CIDR block for IP range provided.

### 3. Check if IP is in CIDR/Subnet range

To check if a given IP address falls inside a CIDR/Subnet range use the `check` subcommand as shown below -

```
iptool check 192.168.1.48 192.168.1.0/24
```

This will check if the given IP address is inside the CIDR block or not.

### 4. Get information about a Subnet

To get a summarised information about a subnet range use the `info` subcommand as shown below -

```
iptool info 192.168.1.0/24
```
This will return a summary of the subnet.

### 5. Split a Subnet into smaller subnets

To split a bigger subnet into smaller subnets use the `split` subcommand with `range` flag as shown below -

```
iptool split 10.0.0.0/16 --range /24
```

To convert into specific number of subnets use the `count` flag like this -

```
iptool split 10.0.0.0/16 --count 4
```


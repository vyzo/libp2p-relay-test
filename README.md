# Libp2p Relay Test Programs

These are some simple programs and scripts to test libp2p relay performance.

# Measurements

We test by transferring a 1GB file between a server and a client,
first through a direct connection and then through every relay in the PL fleet.
The server is located in Paris, while the client is located in London.

At the time of this writing (2019/2/21) there are 15 relays operated by PL to support
IPFS and libp2p applications.
The relays are spread over 5 aws regions, with 9 in the US (N.Virgina, Ohio, Oregon) and 6 in Europe (Frankfurt, Dublin).

Results (Updated 2019/3/11 for large yamux window):

| Relay                                          | ping relay | ping server | transfer time    | region    |
|------------------------------------------------|------------|-------------|------------------|-----------|
| direct connection (baseline)                   |            |     7.4ms   |   7.187864928s   |           |
| QmZ5UovPa7H7Mqe31crwhmoEgAz8XZ6s1stzzWSNKFzuqU |  12.7ms    |    23.2ms   |  26.343105791s   | Dublin    |
| QmbLjvQtV1kkiBh9cWGeXaBoaoWkdnm6QDUFrf2LAkecTH |  12.7ms    |    23.2ms   |  21.850558434s   | Dublin    |
| QmQUPL5gvWp3VDBrouhpZx2L6un2ZBzMbLmcnjK45TMuVv |  12.7ms    |    23.1ms   |  21.481653414s   | Dublin    |
| QmTxKmg6bHwH3kvAwiuFpir9Ee3ogStqjkNeg8Lt7V4iWW |  11.8ms    |    26.8ms   |  35.163579596s   | Frankfurt |
| QmTt4EmqFDsYJiypRD5U4dUDzramwLVCgDQJS9GzkHWJGR |   9.9ms    |    26.8ms   |  58.232045306s   | Frankfurt |
| QmSqKqGZ5yr8adWK5txTn1BqG3Em6RhjL3LM2Q4pjUfYMP |  10.0ms    |    28.3ms   |  22.855141011s   | Frankfurt |
| QmWKiiMyB1qtBs6atLC7UXTk6kw8au2p6LGrET6AkmTCRr |  76.0ms    |   154.2ms   | 1m40.915977947s  | Ohio      |
| QmPZ3ZfLNSiiQXj75Ft5Qs2qnPCt1Csp1nZBUppfY997Yb |  75.3ms    |   154.3ms   | 1m40.047935383s  | Ohio      |
| QmYwiV8rxauHpPCbLukjwHAsy3AkLU4U7VF7NrMTxWsJS8 |  75.2ms    |   153.6ms   | 1m43.054258421s  | Ohio      |
| Qmbyw5e8pemFWb27Z6JJtR8gQ3hMRRkzynSYn66DmQ6yP4 |  85.0ms    |   173.7ms   | 1m55.582224015s  | Virginia  |
| QmaXbbcs7LRFuEoQcxfXqziZATzS68WT5DgFjYFgn3YYLX |  85.0ms    |   173.7ms   | 1m53.344062343s  | Virginia  |
| QmeaqKEz7NipyX5YJhtFb7NdAGjgp4wzGfHEpXESrtSL6Y |  85.0ms    |   173.7ms   | 2m37.460513334s  | Virginia  |
| QmbYSi1jdhnz5D4PiY5AB1CVRYK37A4K4jNamp655wTvWC | 143.8ms    |   299.8ms   | 3m22.792653437s  | Oregon    |
| QmaZK14yReY9Bio4Wxj8NLH7vmEBh91J4rS7HVktRhHiV8 | 143.5ms    |   303.2ms   | 3m18.798269751s  | Oregon    |
| Qmbn1WTN8WRPPhmayAe2eiMdcgKqdtm9MbWKga59bhFEoW | 143.4ms    |   302.2ms   | 3m21.895076093s  | Oregon    |

# Stress Testing

We stress test a relay by running multiple test clients and test servers.
We start with 1 test server and 100 clients, and increment by adding servers with a 100 clients each.
At 10 servers/1000 clients the relay becomes cpu-saturated.

1 server/100 clients:
![netdata-1-100](https://ipfs.io/ipfs/QmRx9rxB49v5trApfnAg5hi4t3eedWLG3yu6wQzd6o3sjY)

5 servers/500 clients:
![netdata-5-500](https://ipfs.io/ipfs/QmVbS46mxvJ6PTHaYgG6hUjT8vkFhTngpgTUfh7tkvojDH)

10 servers/1000 clients:
![netdata-10-100](https://ipfs.io/ipfs/Qmdi62ApVYxTtHXBLJpesRTt6xndG59mgmVPC1Rj9ibnTp)

A netdata snapshot from the saturated relay is available in [netdata-10-1000.snapshot](https://ipfs.io/ipfs/QmdCM7HkwiAYrKkdpZNAKuJrife9WsTwsUN9YREXtnrYoi) for interactive exploration.

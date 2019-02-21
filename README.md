# Libp2p Relay Test Programs

These are some simple programs and scripts to test libp2p relay performance.

# Measurements

We test by transferring a 1GB file between a server and a client,
first through a direct connection and then through every relay in the PL fleet.
The server is located in Paris, while the client is located in London.

At the time of this writing (2019/2/21) there are 15 relays operated by PL to support
IPFS and libp2p applications.
The relays are spread over 5 aws regions, with 9 in the US (N.Virgina, Ohio, Oregon) and 6 in Europe (Frankfurt, Dublin).

Results:

| Relay                                          | ping    | transfer time    |
|------------------------------------------------|---------|------------------|
| direct connection (baseline)                   | 7.7ms   | 33.964118735s    |
| QmPZ3ZfLNSiiQXj75Ft5Qs2qnPCt1Csp1nZBUppfY997Yb | 75.4ms  | 5m27.643753171s  |
| QmWKiiMyB1qtBs6atLC7UXTk6kw8au2p6LGrET6AkmTCRr | 75.2ms  | 4m15.151337098s  |
| QmYwiV8rxauHpPCbLukjwHAsy3AkLU4U7VF7NrMTxWsJS8 | 75.9ms  | 15m49.825359214s |
| Qmbyw5e8pemFWb27Z6JJtR8gQ3hMRRkzynSYn66DmQ6yP4 | 85.0ms  | 3m38.97912766s   |
| QmaXbbcs7LRFuEoQcxfXqziZATzS68WT5DgFjYFgn3YYLX | 85.0ms  | 4m2.533593947s   |
| QmeaqKEz7NipyX5YJhtFb7NdAGjgp4wzGfHEpXESrtSL6Y | 85.0ms  | 5m55.927897621s  |
| Qmbn1WTN8WRPPhmayAe2eiMdcgKqdtm9MbWKga59bhFEoW | 144.9ms | 7m11.629286569s  |
| QmaZK14yReY9Bio4Wxj8NLH7vmEBh91J4rS7HVktRhHiV8 | 144.2ms | 5m21.409484466s  |
| QmbYSi1jdhnz5D4PiY5AB1CVRYK37A4K4jNamp655wTvWC | 140.2ms | 18m28.022029431s |
| QmTxKmg6bHwH3kvAwiuFpir9Ee3ogStqjkNeg8Lt7V4iWW | 10.1ms  | 2m27.574585977s  |
| QmSqKqGZ5yr8adWK5txTn1BqG3Em6RhjL3LM2Q4pjUfYMP | 11.8ms  | 1m0.324375898s   |
| QmTt4EmqFDsYJiypRD5U4dUDzramwLVCgDQJS9GzkHWJGR | 11.5ms  | 1m1.178714762s   |
| QmbLjvQtV1kkiBh9cWGeXaBoaoWkdnm6QDUFrf2LAkecTH | 12.7ms  | 1m23.266873724s  |
| QmZ5UovPa7H7Mqe31crwhmoEgAz8XZ6s1stzzWSNKFzuqU | 12.6ms  | 1m25.655036374s  |
| QmQUPL5gvWp3VDBrouhpZx2L6un2ZBzMbLmcnjK45TMuVv | 12.7ms  | 53.261012677s    |

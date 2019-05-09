FOR /L %%A IN (1,1,30) DO (
  set IPFS_PATH=~/.ipfsSybil%%A
  START cmd /K sybil-ipfs log level dht info
)
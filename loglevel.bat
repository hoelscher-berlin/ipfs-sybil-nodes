FOR /L %%A IN (1,1,10) DO (
  set IPFS_PATH=~/.ipfsSybil%%A
  START cmd /K ipfs log level dht info
)